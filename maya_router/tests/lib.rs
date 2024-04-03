use maya_router::{MayaRouterDepositEvent, MayaRouterTransferOutEvent};
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

struct MayaRouterSimulator {
    pub ledger: DefaultLedgerSimulator,
    pub component_address: ComponentAddress,
    pub _owner: User,
    pub signer: User,
    pub swapper: User,
    pub resources: IndexMap<String, ResourceAddress>,
}

impl MayaRouterSimulator {
    const BLUEPRINT_NAME: &'static str = "MayaRouter";

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

        let receipt = Self::create_component(&mut ledger, owner.badge.clone(), signer_rule);
        let component_address = receipt.expect_commit_success().new_component_addresses()[0];
        Self {
            ledger,
            component_address,
            _owner: owner,
            signer,
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
        asset: ResourceAddress,
        amount: Decimal,
        memo: String,
    ) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .withdraw_from_account(self.swapper.address, asset, amount)
            .take_all_from_worktop(asset, "bucket")
            .with_bucket("bucket", |builder, bucket| {
                builder.call_method(
                    self.component_address,
                    "deposit",
                    manifest_args!(self.swapper.address, bucket, memo),
                )
            })
            .build();
        self.ledger
            .execute_manifest(manifest, vec![self.swapper.badge.clone()])
    }

    // Only signer can call it
    pub fn transfer_out(
        &mut self,
        sender: ComponentAddress,
        badge: NonFungibleGlobalId,
        to: ComponentAddress,
        asset: ResourceAddress,
        amount: Decimal,
        memo: String,
    ) -> TransactionReceipt {
        let manifest = Self::manifest_builder()
            .call_method(
                self.component_address,
                "transfer_out",
                manifest_args!(sender, to, asset, amount, memo),
            )
            .build();
        self.ledger.execute_manifest(manifest, vec![badge])
    }

    pub fn update_signer(&mut self, new_signer: User) {
        let new_signer_rule = rule!(require(new_signer.badge.clone()));
        let manifest = Self::manifest_builder()
            .set_role(
                self.component_address,
                ModuleId::Main,
                RoleKey::new("signer"),
                new_signer_rule,
            )
            .build();
        let receipt = self
            .ledger
            .execute_manifest(manifest, vec![self.signer.badge.clone()]);

        receipt.expect_commit_success();

        self.signer = new_signer;
    }
}

#[test]
fn maya_router_swap_and_send_success() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO".to_string();
    let tx_out_memo = "OUT:".to_string();

    // Perform Swap
    // Act
    let receipt = maya_router.deposit(XRD, dec!(100), swap_memo.clone());

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
            asset: XRD,
            amount: dec!(100),
            memo: swap_memo,
        }
    );

    // Perform Send
    // Arrange
    let balance = maya_router.get_swapper_balance(XRD);

    // Act
    let receipt = maya_router.transfer_out(
        maya_router.signer.address,
        maya_router.signer.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo.clone(),
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
            sender: maya_router.signer.address,
            address: maya_router.swapper.address,
            asset: XRD,
            amount: dec!(100),
            memo: tx_out_memo,
        }
    );
    assert_eq!(balance + dec!(100), maya_router.get_swapper_balance(XRD));
}

#[test]
fn maya_router_swap_and_send_fail() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();

    // Act
    let swap_memo = "SWAP:MAYA.CACAO".to_string();
    let tx_out_memo = "OUT:".to_string();

    // Perform Swap
    // Act
    let receipt = maya_router.deposit(XRD, dec!(100), swap_memo.clone());

    // Assert
    receipt.expect_commit_success();

    // Perform Send of non-existing asset
    // Act
    let receipt = maya_router.transfer_out(
        maya_router.signer.address,
        maya_router.signer.badge.clone(),
        maya_router.swapper.address,
        *maya_router.resources.get("USDT").unwrap(),
        dec!(100),
        tx_out_memo.clone(),
    );

    // Assert
    receipt.expect_specific_failure(|e| match e {
        RuntimeError::ApplicationError(ApplicationError::PanicMessage(s)) => {
            s.contains("not available in the vault")
        }
        _ => false,
    });
}

#[test]
fn maya_router_update_signer_rule() {
    // Arrange
    let mut maya_router = MayaRouterSimulator::new();
    let swap_memo = "SWAP:MAYA.CACAO".to_string();
    let tx_out_memo = "OUT:".to_string();

    let old_signer_badge = maya_router.signer.badge.clone();

    // Act
    let receipt = maya_router.deposit(XRD, dec!(1000), swap_memo);
    receipt.expect_commit_success();

    // No error expected when using old signer badge
    let receipt = maya_router.transfer_out(
        maya_router.signer.address,
        old_signer_badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo.clone(),
    );
    receipt.expect_commit_success();

    // Update signer rule
    let new_signer = User::new(&mut maya_router.ledger);
    maya_router.update_signer(new_signer);

    // Expect AuthError when using old signer badge
    let receipt = maya_router.transfer_out(
        maya_router.signer.address,
        old_signer_badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo.clone(),
    );
    receipt.expect_specific_failure(|e| {
        matches!(
            e,
            RuntimeError::SystemModuleError(SystemModuleError::AuthError(AuthError::Unauthorized(
                ..
            )))
        )
    });

    // No error expected when using current signer badge
    let receipt = maya_router.transfer_out(
        maya_router.signer.address,
        maya_router.signer.badge.clone(),
        maya_router.swapper.address,
        XRD,
        dec!(100),
        tx_out_memo,
    );
    receipt.expect_commit_success();
}
