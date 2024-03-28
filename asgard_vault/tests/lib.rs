use asgard_vault::AsgardSwapEvent;
use scrypto_test::prelude::LedgerSimulatorBuilder;
use scrypto_test::prelude::*;

struct User {
    _public_key: Secp256k1PublicKey,
    _private_key: Secp256k1PrivateKey,
    address: ComponentAddress,
    badge: NonFungibleGlobalId,
}

impl User {
    fn new(ledger: &mut DefaultLedgerSimulator) -> Self {
        let (public_key, private_key, address) = ledger.new_account(false);
        let badge = NonFungibleGlobalId::from_public_key(&public_key);

        User {
            _public_key: public_key,
            _private_key: private_key,
            address,
            badge,
        }
    }
}

struct AsgardVaultSimulator {
    pub ledger: DefaultLedgerSimulator,
    pub component_address: ComponentAddress,
    pub _owner: User,
    pub signer: User,
    pub swapper: User,
}

impl AsgardVaultSimulator {
    const BLUEPRINT_NAME: &'static str = "AsgardVault";

    pub fn manifest_builder() -> ManifestBuilder {
        ManifestBuilder::new().lock_fee_from_faucet()
    }

    pub fn create_component(
        ledger: &mut DefaultLedgerSimulator,
        owner_badge: NonFungibleGlobalId,
        signer_rule: AccessRule,
    ) -> TransactionReceipt {
        let package_address = ledger.compile_and_publish(this_package!());
        ledger.execute_manifest(
            Self::manifest_builder()
                .call_function(
                    package_address,
                    Self::BLUEPRINT_NAME,
                    "instantiate",
                    manifest_args!(owner_badge, signer_rule),
                )
                .build(),
            vec![],
        )
    }

    pub fn new() -> Self {
        let mut ledger = LedgerSimulatorBuilder::new().build();
        // Owner account
        let owner = User::new(&mut ledger);
        let signer = User::new(&mut ledger);
        let swapper = User::new(&mut ledger);

        let signer_rule = rule!(require(signer.badge.clone()));

        let receipt = Self::create_component(&mut ledger, owner.badge.clone(), signer_rule);
        let component_address = receipt.expect_commit_success().new_component_addresses()[0];
        Self {
            ledger,
            component_address,
            _owner: owner,
            signer,
            swapper,
        }
    }

    pub fn get_swapper_balance(&mut self) -> Decimal {
        self.ledger.get_component_balance(self.swapper.address, XRD)
    }

    pub fn swap(
        &mut self,
        amount: Decimal,
        into: String,
        dest_address: String,
    ) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .withdraw_from_account(self.swapper.address, XRD, amount)
            .take_all_from_worktop(XRD, "xrd_bucket")
            .with_name_lookup(|builder, lookup| {
                let xrd_bucket = lookup.bucket("xrd_bucket");
                builder.call_method(
                    self.component_address,
                    "swap",
                    manifest_args!(xrd_bucket, into, dest_address),
                )
            })
            .build();
        self.ledger
            .execute_manifest(manifest, vec![self.swapper.badge.clone()])
    }

    // Only signer can call it
    pub fn send(&mut self, amount: Decimal, dest_address: ComponentAddress) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .call_method(
                self.component_address,
                "signer_send",
                manifest_args!(amount, dest_address),
            )
            .build();
        self.ledger
            .execute_manifest(manifest, vec![self.signer.badge.clone()])
    }
}

#[test]
fn asgard_vault_swap_and_send() {
    // Arrange
    let mut sim = AsgardVaultSimulator::new();

    // Act
    let into = "BTC".to_string();
    let dest_address = "btc_address".to_string();

    // Perform Swap
    // Act
    let receipt = sim.swap(dec!(100), into.clone(), dest_address.clone());

    // Assert
    let result = receipt.expect_commit_success();
    let events = result.application_events.as_slice();
    let [
        _, /* LockFeeEvent */
        _, /* WithdrawEvent */
        _, /* WithdrawEvent */
        _, /* DepositEvent */
        (event_type_identifier, event_data), /* AsgardSwapEvent */
        _, /* PayFeeEvent */
        _, /* DepositEvent */
        _, /* BurnFungibleResourceEvent */
    ] = events else {
        panic!("Incorrect number of events: {}", events.len())
    };

    assert_eq!(
        event_type_identifier.to_owned(),
        EventTypeIdentifier(
            Emitter::Method(sim.component_address.into_node_id(), ModuleId::Main),
            "AsgardSwapEvent".to_string()
        )
    );
    assert_eq!(
        scrypto_decode::<AsgardSwapEvent>(&event_data).unwrap(),
        AsgardSwapEvent {
            into,
            dest_address,
            limit: None,
            affiliate: None,
            fee: None,
        }
    );

    // Perform Send
    // Arrange
    let balance = sim.get_swapper_balance();

    // Act
    let receipt = sim.send(dec!(100), sim.swapper.address);

    // Assert
    let _result = receipt.expect_commit_success();

    assert_eq!(balance + dec!(100), sim.get_swapper_balance());
}
