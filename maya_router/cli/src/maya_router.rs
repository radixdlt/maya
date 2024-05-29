use crate::manifest_execute::*;
use crate::network_connection_provider::*;
use crate::network_selector::*;
use crate::utils::*;
use radix_common::prelude::*;
use radix_engine_interface::prelude::*;
use radix_transactions::prelude::*;
use radix_transactions::signing::PrivateKey;
use rand::prelude::*;

pub enum KeyType {
    Ed25519,
    Secp256k1,
}

pub struct AccountData {
    public_key: PublicKey,
    private_key: PrivateKey,
    address: ComponentAddress,
    badge: NonFungibleGlobalId,
}

impl Clone for AccountData {
    fn clone(&self) -> Self {
        Self {
            public_key: self.public_key,
            private_key: clone_private_key(&self.private_key),
            address: self.address,
            badge: self.badge.clone(),
        }
    }
}

impl AccountData {
    fn new(
        network_connector: &mut GatewayNetworkConnector,
        key_type: KeyType,
        std_rng: &mut StdRng,
    ) -> Self {
        let account = Self::new_not_initialized(key_type, std_rng);
        let private_key = clone_private_key(&account.private_key);
        let manifest = ManifestBuilder::new()
            .lock_fee(FAUCET, dec!(100))
            .get_free_xrd_from_faucet()
            .take_all_from_worktop(XRD, "bucket")
            .try_deposit_or_abort(account.address, None, "bucket")
            .build();
        let _ = execute_2(network_connector, private_key, manifest).unwrap();
        account
    }

    fn new_not_initialized(key_type: KeyType, std_rng: &mut StdRng) -> Self {
        let account: (PublicKey, PrivateKey) = match key_type {
            KeyType::Ed25519 => {
                let private_key = Ed25519PrivateKey::from_u64(std_rng.next_u64())
                    .expect("Cannot generate private key from int");
                let public_key = private_key.public_key();
                (public_key.into(), private_key.into())
            }
            KeyType::Secp256k1 => {
                let private_key = Secp256k1PrivateKey::from_u64(std_rng.next_u64())
                    .expect("Cannot generate private key from int");
                let public_key = private_key.public_key();
                (public_key.into(), private_key.into())
            }
        };
        let address = ComponentAddress::virtual_account_from_public_key(&account.0.clone());

        let badge = NonFungibleGlobalId::from_public_key(&account.0);

        AccountData {
            public_key: account.0,
            private_key: account.1,
            address,
            badge,
        }
    }
}

pub struct MayaRouterTester {
    pub gateway_network_connector: GatewayNetworkConnector,
    pub std_rng: StdRng,
    pub component_address: Option<ComponentAddress>,
    pub owner: AccountData,
    pub asgard_vault_1: AccountData,
    pub asgard_vault_2: AccountData,
    pub swapper: AccountData,
    pub resources: IndexMap<String, ResourceAddress>,
}

impl MayaRouterTester {
    pub fn new(network: &str, owner_private_key: Option<&str>) -> Self {
        let private_key = match owner_private_key {
            Some(owner_private_key) => {
                let bytes = hex::decode(owner_private_key).expect("Cannot decode private_key");
                Ed25519PrivateKey::from_bytes(&bytes)
                    .expect("Cannot generate private key from bytes")
            }
            None => Ed25519PrivateKey::from_u64(222).expect("Cannot generate private key from int"),
        };
        let public_key = private_key.public_key();
        let address = ComponentAddress::virtual_account_from_public_key(&public_key);
        let badge = NonFungibleGlobalId::from_public_key(&public_key);

        let owner = AccountData {
            public_key: public_key.into(),
            private_key: clone_private_key_ed25519(&private_key).into(),
            address,
            badge,
        };

        let network = NetworkConfig::new(network);

        let mut gateway_network_connector = GatewayNetworkConnector::new(
            network.gateway_base_url(),
            network.network_definition(),
            PollingConfiguration {
                interval_in_seconds: 10,
                retries: 10,
            },
        );
        let mut std_rng = StdRng::from_entropy();
        let asgard_vault_1 = AccountData::new_not_initialized(KeyType::Secp256k1, &mut std_rng);
        let asgard_vault_2 = AccountData::new_not_initialized(KeyType::Ed25519, &mut std_rng);
        let swapper = AccountData::new(
            &mut gateway_network_connector,
            KeyType::Secp256k1,
            &mut std_rng,
        );
        Self {
            gateway_network_connector,
            std_rng,
            component_address: None,
            owner,
            asgard_vault_1,
            asgard_vault_2,
            swapper,
            resources: indexmap!("XRD".to_string() => XRD),
        }
    }

    pub fn set_component_address(&mut self, component_address: ComponentAddress) {
        self.component_address = Some(component_address);

        // Generate repeatable Asgard Vaults for given component address
        let mut seed: [u8; 32] = [0; 32];
        seed[0..30].clone_from_slice(component_address.to_vec().as_slice());
        self.std_rng = StdRng::from_seed(seed);
        self.asgard_vault_1 =
            AccountData::new_not_initialized(KeyType::Secp256k1, &mut self.std_rng);
        self.asgard_vault_2 = AccountData::new_not_initialized(KeyType::Ed25519, &mut self.std_rng);
    }

    pub fn publish(&mut self) -> PackageAddress {
        let (code, definition) = package_get();
        let mut metadata = MetadataInit::new();
        metadata.set_metadata("name", "MayaRouter Package");
        metadata.set_metadata(
            "description",
            "MayaRouter package stores assets swappable with assets from other networks",
        );
        let owner_role = OwnerRole::Fixed(rule!(require(self.owner.badge.clone())));
        let manifest = ManifestBuilder::new()
            .lock_fee(self.owner.address, dec!(100))
            .publish_package_advanced(None, code, definition, metadata, owner_role)
            .build();

        println!("Publish package");
        let result = execute_2(
            &mut self.gateway_network_connector,
            clone_private_key(&self.owner.private_key),
            manifest,
        )
        .unwrap();
        println!("TxId: {}", result.intent_hash);

        let package_address = result
            .new_entities
            .new_package_addresses
            .get_index(0)
            .unwrap();
        *package_address
    }

    pub fn instantiate(&mut self, package_address: PackageAddress) -> ComponentAddress {
        let manifest = ManifestBuilder::new()
            .lock_fee(self.owner.address, dec!(100))
            .call_function(
                package_address,
                "MayaRouter",
                "instantiate",
                manifest_args!(),
            )
            .build();

        println!("Instantiate package");
        let result = execute_2(
            &mut self.gateway_network_connector,
            clone_private_key(&self.owner.private_key),
            manifest,
        )
        .unwrap();
        println!("TxId: {}", result.intent_hash);

        let component_address = result
            .new_entities
            .new_component_addresses
            .get_index(0)
            .unwrap();
        self.set_component_address(*component_address);
        *component_address
    }

    pub fn deposit(
        &mut self,
        vault_address: ComponentAddress,
        from: AccountData,
        asset: ResourceAddress,
        amount: Decimal,
        memo: &str,
    ) {
        let manifest = ManifestBuilder::new()
            .lock_fee(from.address, dec!(10))
            .withdraw_from_account(from.address, asset, amount)
            .take_all_from_worktop(asset, "bucket")
            .with_bucket("bucket", |builder, bucket| {
                builder.call_method(
                    self.component_address.unwrap(),
                    "user_deposit",
                    manifest_args!(from.address, vault_address, bucket, memo.to_string()),
                )
            })
            .build();
        println!("Deposit");
        let result = execute_2(
            &mut self.gateway_network_connector,
            clone_private_key(&from.private_key),
            manifest,
        )
        .unwrap();
        println!("TxId: {}", result.intent_hash);
    }

    pub fn transfer_out(
        &mut self,
        from_vault: AccountData,
        to: AccountData,
        asset: ResourceAddress,
        amount: Decimal,
        memo: &str,
        fee_to_lock: Decimal,
    ) {
        let manifest = ManifestBuilder::new()
            .call_method(
                self.component_address.unwrap(),
                "lock_fee",
                manifest_args!(from_vault.address, fee_to_lock),
            )
            .call_method(
                self.component_address.unwrap(),
                "withdraw",
                manifest_args!(
                    from_vault.address,
                    asset,
                    to.address,
                    Option::<()>::None,
                    amount,
                    memo.to_string()
                ),
            )
            .take_all_from_worktop(asset, "asset")
            .call_method_with_name_lookup(self.component_address.unwrap(), "transfer", |lookup| {
                (to.address, lookup.bucket("asset"))
            })
            .build();

        println!("Withdraw");
        let result = execute_2(
            &mut self.gateway_network_connector,
            clone_private_key(&from_vault.private_key),
            manifest,
        )
        .unwrap();
        println!("TxId: {}", result.intent_hash);
    }

    pub fn transfer_between_vaults(
        &mut self,
        from_vault: AccountData,
        to_vault: AccountData,
        asset: ResourceAddress,
        amount: Decimal,
        memo: &str,
        fee_to_lock: Decimal,
    ) {
        let manifest = ManifestBuilder::new()
            .call_method(
                self.component_address.unwrap(),
                "lock_fee",
                manifest_args!(from_vault.address, fee_to_lock),
            )
            .call_method(
                self.component_address.unwrap(),
                "withdraw",
                manifest_args!(
                    from_vault.address,
                    asset,
                    to_vault.address,
                    Option::<()>::None,
                    amount,
                    memo.to_string()
                ),
            )
            .take_all_from_worktop(asset, "asset")
            .call_method_with_name_lookup(
                self.component_address.unwrap(),
                "direct_deposit",
                |lookup| (to_vault.address, lookup.bucket("asset")),
            )
            .build();

        println!("Transfer between vaults");
        let result = execute_2(
            &mut self.gateway_network_connector,
            clone_private_key(&from_vault.private_key),
            manifest,
        )
        .unwrap();
        println!("TxId: {}", result.intent_hash);
    }
}

// PrivateKey: 0000000000000000000000000000000000000000000000000000000000000001
// PublicKey: 4cb5abf6ad79fbf5abbccafcc269d85cd2651ed4b885b5869f241aedf0a5ba29
// Address: account_tdx_2_12xsvygvltz4uhsht6tdrfxktzpmnl77r0d40j8agmujgdj02el3l9v

#[test]
fn maya_router_publish_and_instantiate() {
    // Arrange
    let mut maya_router = MayaRouterTester::new(
        "stokenet",
        Some("0000000000000000000000000000000000000000000000000000000000000001"),
    );
    let address_bech32_encoder = AddressBech32Encoder::new(&NetworkDefinition::stokenet());

    let package_address = maya_router.publish();
    println!(
        "PackageAddress: {}",
        package_address.display(&address_bech32_encoder)
    );
    let component_address = maya_router.instantiate(package_address);
    println!(
        "ComponentAddress: {}",
        component_address.display(&address_bech32_encoder)
    );

    maya_router.deposit(
        maya_router.asgard_vault_1.address,
        maya_router.swapper.clone(),
        XRD,
        dec!(200),
        "SWAP:MAYA.CACAO",
    );

    maya_router.transfer_out(
        maya_router.asgard_vault_1.clone(),
        maya_router.swapper.clone(),
        XRD,
        dec!(100),
        "OUT:",
        dec!(10),
    );
}

// Test assumes that package is published and instantiated
// Publish package
// TxId: txid_tdx_2_1mdn07rz57nxjuv0k0aqh4lls90z082a8ufq44qzqz8kdudex45dqaqd6nh
// PackageAddress: package_tdx_2_1p5sa4nnvchp95dy70f94h7prn5ss6gwhx0yq8evfns7tqumq35jaqr
// Instantiate package
// TxId: txid_tdx_2_1cgn0x5wk6f85ukd2ga8he5vth7h6w6mgqcyyc4wtw4n696r2aq4szwl66v
// ComponentAddress: component_tdx_2_1cqfz7c2t37dpmnaz9zpkhhaa6442j6dl0t4l49szeds7e8ryzp5ran

#[test]
fn maya_router_deposit_and_withdraw_success() {
    // Arrange
    let mut maya_router = MayaRouterTester::new(
        "stokenet",
        Some("0000000000000000000000000000000000000000000000000000000000000001"),
    );
    let component_address =
        "component_tdx_2_1cqfz7c2t37dpmnaz9zpkhhaa6442j6dl0t4l49szeds7e8ryzp5ran";
    let address_bech32_decoder = AddressBech32Decoder::new(&NetworkDefinition::stokenet());
    let component_address =
        ComponentAddress::try_from_bech32(&address_bech32_decoder, component_address).unwrap();

    maya_router.set_component_address(component_address);
    maya_router.deposit(
        maya_router.asgard_vault_1.address,
        maya_router.swapper.clone(),
        XRD,
        dec!(200),
        "SWAP:MAYA.CACAO",
    );

    maya_router.transfer_out(
        maya_router.asgard_vault_1.clone(),
        maya_router.swapper.clone(),
        XRD,
        dec!(100),
        "OUT:",
        dec!(10),
    );
}
