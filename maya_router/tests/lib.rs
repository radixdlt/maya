use maya_router::{MayaRouterDepositEvent, MayaRouterTransferOutEvent};
use radix_engine::system::system_type_checker::TypeCheckError;
use scrypto_test::prelude::LedgerSimulatorBuilder;
use scrypto_test::prelude::*;

struct User {
    public_key: Ed25519PublicKey,
    _private_key: Ed25519PrivateKey,
    address: ComponentAddress,
    badge: NonFungibleGlobalId,
}

impl User {
    fn new(ledger: &mut DefaultLedgerSimulator) -> Self {
        let (public_key, private_key, address) = ledger.new_ed25519_virtual_account();
        let badge = NonFungibleGlobalId::from_public_key(&public_key);

        User {
            public_key,
            _private_key: private_key,
            address,
            badge,
        }
    }
}

struct MayaRouterSimulator {
    pub ledger: DefaultLedgerSimulator,
    pub component_address: ComponentAddress,
    pub _owner: User,
    pub asgard_vault_1: User,
    pub asgard_vault_2: User,
    pub swapper: User,
    pub resources: IndexMap<String, ResourceAddress>,
}

impl MayaRouterSimulator {
    const BLUEPRINT_NAME: &'static str = "MayaRouter";

    pub fn manifest_builder() -> ManifestBuilder {
        ManifestBuilder::new().lock_fee_from_faucet()
    }

    pub fn create_component(ledger: &mut DefaultLedgerSimulator) -> TransactionReceipt {
        let package_address = ledger.compile_and_publish(this_package!());
        ledger.execute_manifest(
            Self::manifest_builder()
                .call_function(
                    package_address,
                    Self::BLUEPRINT_NAME,
                    "instantiate",
                    manifest_args!(),
                )
                .build(),
            vec![],
        )
    }

    pub fn new() -> Self {
        let mut ledger = LedgerSimulatorBuilder::new().build();
        // Owner account
        let owner = User::new(&mut ledger);
        let asgard_vault_1 = User::new(&mut ledger);
        let asgard_vault_2 = User::new(&mut ledger);
        let swapper = User::new(&mut ledger);

        let mut resources = indexmap!();
        resources.insert("XRD".to_string(), XRD);
        for symbol in ["USDT", "ETH"] {
            resources.insert(
                symbol.to_string(),
                ledger.create_fungible_resource(dec!(20000), 18, swapper.address),
            );
        }
        for symbol in ["BTC", "SOL"] {
            resources.insert(
                symbol.to_string(),
                ledger.create_fungible_resource(dec!(20000), 18, owner.address),
            );
        }

        let receipt = Self::create_component(&mut ledger);
        let component_address = receipt.expect_commit_success().new_component_addresses()[0];
        Self {
            ledger,
            component_address,
            _owner: owner,
            asgard_vault_1,
            asgard_vault_2,
            swapper,
            resources,
        }
    }

    pub fn get_swapper_balance(&mut self, asset: ResourceAddress) -> Decimal {
        self.ledger
            .get_component_balance(self.swapper.address, asset)
    }

    pub fn deposit(
        &mut self,
        asgard_vault: Ed25519PublicKey,
        from: ComponentAddress,
        badge: NonFungibleGlobalId,
        asset: ResourceAddress,
        amount: Decimal,
        memo: &str,
    ) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .withdraw_from_account(from, asset, amount)
            .take_all_from_worktop(asset, "bucket")
            .with_bucket("bucket", |builder, bucket| {
                builder.call_method(
                    self.component_address,
                    "deposit",
                    manifest_args!(from, asgard_vault, bucket, memo.to_string()),
                )
            })
            .build();
        self.ledger.execute_manifest(manifest, vec![badge])
    }

    pub fn transfer_out(
        &mut self,
        asgard_vault: Ed25519PublicKey,
        badge: NonFungibleGlobalId,
        to: ComponentAddress,
        asset: ResourceAddress,
        amount: Decimal,
        memo: &str,
    ) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .call_method(
                self.component_address,
                "transfer_out",
                manifest_args!(asgard_vault, to, asset, amount, memo.to_string()),
            )
            .build();
        self.ledger.execute_manifest(manifest, vec![badge])
    }
}

#[test]
fn maya_router_deposit_and_transfer_out_success() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO";
    let tx_out_memo = "OUT:";

    // Perform Swap
    // Act
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_1.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        XRD,
        dec!(100),
        swap_memo,
    );

    // Assert
    let result = receipt.expect_commit_success();
    let events = result.application_events.as_slice();

    let event_data = events
        .iter()
        .find(|(type_identifier, _)| {
            type_identifier.eq(&EventTypeIdentifier(
                Emitter::Method(maya_router.component_address.into_node_id(), ModuleId::Main),
                "MayaRouterDepositEvent".to_string(),
            ))
        })
        .map(|(_, data)| data)
        .expect("MayaRouterDepositEvent not found");

    assert_eq!(
        scrypto_decode::<MayaRouterDepositEvent>(&event_data).unwrap(),
        MayaRouterDepositEvent {
            sender: maya_router.swapper.address,
            asgard_vault: maya_router.asgard_vault_1.public_key,
            asset: XRD,
            amount: dec!(100),
            memo: swap_memo.to_string(),
        }
    );

    // Perform Send
    // Arrange
    let balance = maya_router.get_swapper_balance(XRD);

    // Act
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_1.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );

    // Assert
    let result = receipt.expect_commit_success();
    let events = result.application_events.as_slice();

    let event_data = events
        .iter()
        .find(|(type_identifier, _)| {
            type_identifier.eq(&EventTypeIdentifier(
                Emitter::Method(maya_router.component_address.into_node_id(), ModuleId::Main),
                "MayaRouterTransferOutEvent".to_string(),
            ))
        })
        .map(|(_, data)| data)
        .expect("MayaRouterTransferOutEvent not found");

    assert_eq!(
        scrypto_decode::<MayaRouterTransferOutEvent>(&event_data).unwrap(),
        MayaRouterTransferOutEvent {
            asgard_vault: maya_router.asgard_vault_1.public_key,
            receiver: maya_router.swapper.address,
            asset: XRD,
            amount: dec!(100),
            memo: tx_out_memo.to_string(),
        }
    );
    assert_eq!(balance + dec!(100), maya_router.get_swapper_balance(XRD));
}

#[test]
fn maya_router_deposit_sender_non_real_account() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO";

    // Act
    let manifest = MayaRouterSimulator::manifest_builder()
        .withdraw_from_account(maya_router.swapper.address, XRD, dec!(100))
        .take_all_from_worktop(XRD, "bucket")
        .with_bucket("bucket", |builder, bucket| {
            builder.call_method(
                maya_router.component_address,
                "deposit",
                manifest_args!(
                    maya_router.component_address, // non-account component address
                    maya_router.asgard_vault_1.public_key,
                    bucket,
                    swap_memo.to_string()
                ),
            )
        })
        .build();
    let receipt = maya_router
        .ledger
        .execute_manifest(manifest, vec![maya_router.swapper.badge]);

    // Assert
    receipt.expect_specific_failure(|e| {
        matches!(
            e,
            RuntimeError::SystemError(SystemError::TypeCheckError(
                TypeCheckError::BlueprintPayloadValidationError(..)
            ))
        )
    });
}

#[test]
fn maya_router_transfer_out_asset_not_available() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO";
    let tx_out_memo = "OUT:";

    // Perform Swap
    // Act
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_1.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        XRD,
        dec!(100),
        swap_memo,
    );

    // Assert
    receipt.expect_commit_success();

    // Perform Send of non-existing asset
    // Act
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_1.badge.clone(),
        maya_router.swapper.address,
        *maya_router.resources.get("USDT").unwrap(),
        dec!(100),
        tx_out_memo,
    );

    // Assert
    receipt.expect_specific_failure(|e| match e {
        RuntimeError::ApplicationError(ApplicationError::PanicMessage(s)) => {
            s.contains("not available in the asgard vault")
        }
        _ => false,
    });
}

#[test]
fn maya_router_transfer_out_asgard_vault_not_available() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let tx_out_memo = "OUT:";

    // Perform Send from non-existing Asgard Vault
    // Act
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_1.badge.clone(),
        maya_router.swapper.address,
        *maya_router.resources.get("USDT").unwrap(),
        dec!(100),
        tx_out_memo,
    );

    // Assert
    receipt.expect_specific_failure(|e| match e {
        RuntimeError::ApplicationError(ApplicationError::PanicMessage(s)) => {
            s.contains("asgard vault") && s.contains("not available")
        }
        _ => false,
    });
}

#[test]
fn maya_router_transfer_out_assert_access_rule_failed() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO";
    let tx_out_memo = "OUT:";

    // Perform Swap
    // Act
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_1.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        XRD,
        dec!(100),
        swap_memo,
    );

    // Assert
    receipt.expect_commit_success();

    // Act
    // Expect AssertAccessRuleFailed when using Asgard Vault 2 badge for Asgard Vault 1
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_2.badge.clone(),
        maya_router.swapper.address,
        *maya_router.resources.get("USDT").unwrap(),
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_specific_failure(|e| {
        matches!(
            e,
            RuntimeError::SystemError(SystemError::AssertAccessRuleFailed)
        )
    });
}

#[test]
fn maya_router_multiple_asgard_vaults() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();
    let swap_memo = "SWAP:MAYA.CACAO";
    let tx_out_memo = "OUT:";

    // Act
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_1.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        XRD,
        dec!(1000),
        swap_memo,
    );
    receipt.expect_commit_success();

    let receipt = maya_router.deposit(
        maya_router.asgard_vault_2.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        *maya_router.resources.get("USDT").unwrap(),
        dec!(1000),
        swap_memo,
    );
    receipt.expect_commit_success();

    // No error expected when transferring XRD from Asgard Vault 1 with Asgard Vault 1 badge
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_1.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();

    // No error expected when transferring USDT from Asgard Vault 2 with Asgard Vault 2 badge
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_2.public_key,
        maya_router.asgard_vault_2.badge.clone(),
        maya_router.swapper.address,
        *maya_router.resources.get("USDT").unwrap(),
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();

    // Expect resources not available when trying to transfer XRD from Asgard Vault 2
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_2.public_key,
        maya_router.asgard_vault_2.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_specific_failure(|e| match e {
        RuntimeError::ApplicationError(ApplicationError::PanicMessage(s)) => {
            s.contains("not available in the asgard vault")
        }
        _ => false,
    });
}

#[test]
fn maya_router_move_assets_from_asgard_vault_1_to_asgard_vault_2() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();
    let swap_memo = "SWAP:MAYA.CACAO";
    let tx_out_memo = "OUT:";

    // Act
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_1.public_key,
        maya_router.swapper.address,
        maya_router.swapper.badge.clone(),
        XRD,
        dec!(1000),
        swap_memo,
    );
    receipt.expect_commit_success();

    // Transfer resources to Asgard Vault 2 address
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_1.public_key,
        maya_router.asgard_vault_1.badge.clone(),
        maya_router.asgard_vault_2.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();

    // Deposit resources from Asgard Vault 2 address into Asgard Vault 2
    let receipt = maya_router.deposit(
        maya_router.asgard_vault_2.public_key,
        maya_router.asgard_vault_2.address,
        maya_router.asgard_vault_2.badge.clone(),
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();

    // No error expected when transferring XRD from Asgard Vault 2
    let receipt = maya_router.transfer_out(
        maya_router.asgard_vault_2.public_key,
        maya_router.asgard_vault_2.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();
}
