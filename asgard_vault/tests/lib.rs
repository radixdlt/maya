use asgard_vault::AsgardSwapEvent;
use scrypto_test::prelude::LedgerSimulatorBuilder;
use scrypto_test::prelude::*;

fn instantiate_asgard_vault(ledger: &mut DefaultLedgerSimulator) -> ComponentAddress {
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

#[test]
fn test_asgard_vault_swap() {
    // Arrange
    let mut ledger = LedgerSimulatorBuilder::new().build();
    let asgard_component = instantiate_asgard_vault(&mut ledger);
    let (public_key, _, account) = ledger.new_account(false);

    // Act
    let manifest = ManifestBuilder::new()
        .lock_fee_from_faucet()
        .withdraw_from_account(account, XRD, dec!(100))
        .take_all_from_worktop(XRD, "xrd_bucket")
        .with_name_lookup(|builder, lookup| {
            let xrd_bucket = lookup.bucket("xrd_bucket");
            builder.call_method(
                asgard_component,
                "swap",
                manifest_args!(xrd_bucket, "BTC", "btc_address"),
            )
        })
        .build();
    let receipt = ledger.execute_manifest(
        manifest,
        vec![NonFungibleGlobalId::from_public_key(&public_key)],
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
            into: "BTC".to_string(),
            dest_address: "btc_address".to_string(),
            limit: None,
            affiliate: None,
            fee: None,
        }
    );
}
