use asgard_vault::AsgardSwapEvent;
use scrypto_test::prelude::LedgerSimulatorBuilder;
use scrypto_test::prelude::*;

fn asgard_vault_instantiate(
    ledger: &mut DefaultLedgerSimulator,
    owner_badge: NonFungibleGlobalId,
    signer_rule: AccessRule,
) -> ComponentAddress {
    let package_address = ledger.compile_and_publish(this_package!());
    let receipt = ledger.execute_manifest(
        ManifestBuilder::new()
            .lock_fee(ledger.faucet_component(), 500u32)
            .call_function(
                package_address,
                "AsgardVault",
                "instantiate",
                manifest_args!(owner_badge, signer_rule),
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
    badge: NonFungibleGlobalId,
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
    ledger.execute_manifest(manifest, vec![badge])
}

fn asgard_vault_send(
    ledger: &mut DefaultLedgerSimulator,
    asgard_component: ComponentAddress,
    badge: NonFungibleGlobalId,
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
    ledger.execute_manifest(manifest, vec![badge])
}

#[test]
fn test_asgard_vault_swap_and_send() {
    // Arrange
    let mut ledger = LedgerSimulatorBuilder::new().build();
    // Owner account
    let (owner_public_key, _, _owner_account) = ledger.new_account(false);
    let owner_badge = NonFungibleGlobalId::from_public_key(&owner_public_key);

    // Signer account
    let (signer_public_key, _, _signer_account) = ledger.new_account(false);
    let signer_badge = NonFungibleGlobalId::from_public_key(&signer_public_key);
    let signer_rule = rule!(require(signer_badge.clone()));
    // Swapper account
    let (swapper_public_key, _, swapper_account) = ledger.new_account(false);
    let swapper_badge = NonFungibleGlobalId::from_public_key(&swapper_public_key);

    let asgard_component =
        asgard_vault_instantiate(&mut ledger, owner_badge.clone(), signer_rule.clone());

    let into = "BTC".to_string();
    let dest_address = "btc_address".to_string();

    // Perform Swap
    // Act
    let receipt = asgard_vault_swap(
        &mut ledger,
        asgard_component,
        swapper_account,
        swapper_badge,
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
    let balance = ledger.get_component_balance(swapper_account, XRD);

    // Act
    let receipt = asgard_vault_send(
        &mut ledger,
        asgard_component,
        signer_badge.clone(),
        dec!(100),
        swapper_account,
    );

    // Assert
    let _result = receipt.expect_commit_success();

    assert_eq!(
        balance + dec!(100),
        ledger.get_component_balance(swapper_account, XRD)
    );
}
