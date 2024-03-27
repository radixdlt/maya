use asgard_vault::AsgardSwapEvent;
use scrypto_test::prelude::LedgerSimulatorBuilder;
use scrypto_test::prelude::*;

fn asgard_vault_instantiate(ledger: &mut DefaultLedgerSimulator) -> ComponentAddress {
    let package_address = ledger.compile_and_publish(this_package!());
    let receipt = ledger.execute_manifest(
        ManifestBuilder::new()
            .lock_fee(ledger.faucet_component(), 500u32)
            .call_function(
                package_address,
                "AsgardVault",
                "instantiate",
                manifest_args!(),
            )
            .build(),
        vec![],
    );
    receipt.expect_commit_success().new_component_addresses()[0]
}

fn asgard_vault_swap(
    ledger: &mut DefaultLedgerSimulator,
    asgard_component: ComponentAddress,
    from_account: ComponentAddress,
    public_key: Secp256k1PublicKey,
    amount: Decimal,
    into: String,
    dest_address: String,
) -> TransactionReceiptV1 {
    let manifest = ManifestBuilder::new()
        .lock_fee_from_faucet()
        .withdraw_from_account(from_account, XRD, amount)
        .take_all_from_worktop(XRD, "xrd_bucket")
        .with_name_lookup(|builder, lookup| {
            let xrd_bucket = lookup.bucket("xrd_bucket");
            builder.call_method(
                asgard_component,
                "swap",
                manifest_args!(xrd_bucket, into, dest_address),
            )
        })
        .build();
    ledger.execute_manifest(
        manifest,
        vec![NonFungibleGlobalId::from_public_key(&public_key)],
    )
}

fn asgard_vault_send(
    ledger: &mut DefaultLedgerSimulator,
    asgard_component: ComponentAddress,
    amount: Decimal,
    dest_address: ComponentAddress,
) -> TransactionReceiptV1 {
    let manifest = ManifestBuilder::new()
        .lock_fee_from_faucet()
        .call_method(
            asgard_component,
            "signer_send",
            manifest_args!(amount, dest_address),
        )
        .build();
    ledger.execute_manifest(manifest, vec![])
}

#[test]
fn test_asgard_vault_swap_and_send() {
    // Arrange
    let mut ledger = LedgerSimulatorBuilder::new().build();
    let asgard_component = asgard_vault_instantiate(&mut ledger);
    let (public_key, _, account) = ledger.new_account(false);

    let into = "BTC".to_string();
    let dest_address = "btc_address".to_string();

    // Perform Swap
    // Act
    let receipt = asgard_vault_swap(
        &mut ledger,
        asgard_component,
        account,
        public_key,
        dec!(100),
        into.clone(),
        dest_address.clone(),
    );

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
            Emitter::Method(asgard_component.into_node_id(), ModuleId::Main),
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
    let balance = ledger.get_component_balance(account, XRD);

    // Act
    let receipt = asgard_vault_send(&mut ledger, asgard_component, dec!(100), account);

    // Assert
    let _result = receipt.expect_commit_success();

    assert_eq!(
        balance + dec!(100),
        ledger.get_component_balance(account, XRD)
    );
}
