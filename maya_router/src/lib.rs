use scrypto::prelude::*;

type AllowanceKey = (PublicKey, ResourceAddress);

#[blueprint]
#[types(ResourceAddress, FungibleVault)]
#[types(AllowanceKey, Decimal)]
#[events(
    MayaRouterDepositEvent,
    MayaRouterTransferOutEvent,
    MayaRouterTransferVaultEvent
)]
mod maya_router {
    enable_method_auth! {
        methods {
            deposit => PUBLIC;
            transfer_out => PUBLIC;
            transfer_between_vaults => PUBLIC;
        }
    }

    // MayaRouter owns vaults with fungible resources per Vault (eg. Asgard Vault) public key.
    struct MayaRouter {
        locker: Global<AccountLocker>,
        balances: KeyValueStore<ResourceAddress, FungibleVault>,
        allowances: KeyValueStore<AllowanceKey, Decimal>,
    }

    impl MayaRouter {
        pub fn instantiate() -> Global<MayaRouter> {
            let (maya_router_address_reservation, maya_router_component_address) =
                Runtime::allocate_component_address(MayaRouter::blueprint_id());
            let global_caller_badge_rule =
                rule!(require(global_caller(maya_router_component_address)));

            let locker = Blueprint::<AccountLocker>::instantiate(
                OwnerRole::None,                  // owner
                global_caller_badge_rule.clone(), // storer
                rule!(deny_all),                  // storer_updater
                rule!(deny_all),                  // recoverer
                rule!(deny_all),                  // recoverer_updater
                None,                             // address_reservation
            );

            Self {
                locker,
                balances: KeyValueStore::new_with_registered_type(),
                allowances: KeyValueStore::new_with_registered_type(),
            }
            .instantiate()
            .prepare_to_globalize(OwnerRole::None)
            .with_address(maya_router_address_reservation)
            .globalize()
        }

        // Take some amount of assets from the vault indicated by the given key.
        // If amount is None, then take all.
        fn vault_take(
            &mut self,
            vault_key: PublicKey,
            asset: ResourceAddress,
            amount: Option<Decimal>,
            fee_to_lock: Decimal,
        ) -> FungibleBucket {
            let (bucket, allowance_is_zero) = {
                let mut vault = if asset == XRD {
                    let mut vault = self
                        .balances
                        .get_mut(&XRD)
                        .expect("asset XRD not available in the vault");
                    if fee_to_lock > 0.into() {
                        vault.lock_fee(fee_to_lock);
                    }
                    vault
                } else {
                    if fee_to_lock > 0.into() {
                        self.balances
                            .get_mut(&XRD)
                            .expect("asset XRD not available in the vault")
                            .lock_fee(fee_to_lock);
                    }
                    self.balances
                        .get_mut(&asset)
                        .expect(&format!("asset {:?} not available in the vault", asset))
                };

                // Check if indicated vault owns given asset
                let mut allowance = match self.allowances.get_mut(&(vault_key, asset)) {
                    Some(entry) => entry,
                    None => Runtime::panic(format!(
                        "asset {:?} not available in the vault {:?}",
                        asset, vault_key
                    )),
                };

                let bucket = match amount {
                    Some(amount) => {
                        if amount > vault.amount() {
                            Runtime::panic(format!(
                                "vault {:?} balance {:?} lower than taken amount {:?}",
                                vault_key,
                                vault.amount(),
                                amount
                            ));
                        } else {
                            *allowance -= amount;
                            vault.take(amount)
                        }
                    }
                    None => {
                        *allowance -= vault.amount();
                        vault.take_all()
                    }
                };

                (bucket, allowance.is_zero())
            };

            // Cleanup
            // We don't clean up balances, since not possible to drop a Vault.
            // Also the chances that Vault is empty are really low and event if that's the case,
            // then they will be refilled soon anyway.
            if allowance_is_zero {
                self.allowances.remove(&(vault_key, asset));
            }
            bucket
        }

        // Put bucket of assets into the vault indicated by the given key.
        // If vault of assets does not exist, then create it.
        fn vault_put(&mut self, vault_key: PublicKey, bucket: FungibleBucket) {
            let asset = bucket.resource_address();

            if self.allowances.get(&(vault_key, asset)).is_some() {
                let mut allowance = self
                    .allowances
                    .get_mut(&(vault_key, asset))
                    .expect("vault allowance should be present");
                *allowance += bucket.amount();
            } else {
                self.allowances.insert((vault_key, asset), bucket.amount());
            }

            if self.balances.get(&asset).is_some() {
                let mut vault = self
                    .balances
                    .get_mut(&asset)
                    .expect("asset should be present");
                vault.put(bucket);
            } else {
                self.balances
                    .insert(asset, FungibleVault::with_bucket(bucket));
            }
        }

        // Deposit some assets
        //   sender       - Address where to return refund if required. Must be the address of the method caller
        //   vault_key    - Public key of the Vault, which shall control deposited resources
        //   bucket       - bucket of assets
        //   memo         - message to emit when emitting deposit event
        pub fn deposit(
            &mut self,
            sender: Global<Account>,
            vault_key: PublicKey,
            bucket: FungibleBucket,
            memo: String,
        ) {
            // Make sure the sender account's owner's proof is present when calling this method.
            // This will be present if the sender has just withdrawn from their account.
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            let amount = bucket.amount();
            let asset = bucket.resource_address();

            self.vault_put(vault_key, bucket);

            // Send deposit event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterDepositEvent {
                sender: sender.address(),
                vault_key,
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of given asset to given address.
        //   vault_key    - Public key of the Vault, which controls transferred assets
        //   receiver     - Address where to send assets (must be a real account)
        //   asset        - Resource address of the asset to send
        //   amount       - amount of assets to send
        //   memo         - message to emit when emitting sending the assets
        pub fn transfer_out(
            &mut self,
            vault_key: PublicKey,
            receiver: Global<Account>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
            fee_to_lock: Decimal,
        ) {
            // Make sure the vault account's owner's proof is present when calling this method.
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &vault_key
            ))));

            let bucket = self.vault_take(vault_key, asset, Some(amount), fee_to_lock);

            self.locker.store(receiver, bucket.into(), true);

            // Send transfer out event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterTransferOutEvent {
                vault_key,
                receiver: receiver.address(),
                asset,
                amount,
                memo,
            });
        }

        pub fn transfer_between_vaults(
            &mut self,
            from_vault_key: PublicKey,
            to_vault_key: PublicKey,
            asset: ResourceAddress,
            memo: String,
        ) {
            // Make sure the vault account's owner's proof is present when calling this method.
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &from_vault_key
            ))));

            let bucket = self.vault_take(from_vault_key, asset, None, dec!(0));
            let amount = bucket.amount();
            self.vault_put(to_vault_key, bucket);

            // Send transfer out event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterTransferVaultEvent {
                from_vault_key,
                to_vault_key,
                asset,
                amount,
                memo,
            });
        }
    }
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq)]
pub struct MayaRouterDepositEvent {
    pub sender: ComponentAddress, // Address of the deposit sender
    pub vault_key: PublicKey,     // Public key of the Vault, which controls deposited assets
    pub asset: ResourceAddress,   // Resource address of the deposited assets
    pub amount: Decimal,          // Amount of the deposited assets
    pub memo: String,             // Maya Transaction memo with user intent
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq)]
pub struct MayaRouterTransferOutEvent {
    pub vault_key: PublicKey, // Public key of the Vault, which sends assets
    pub receiver: ComponentAddress, // Address where assets were transferred
    pub asset: ResourceAddress, // Resource address of the transferred assets
    pub amount: Decimal,      // Amount of the transferred assets
    pub memo: String,         // Maya Transaction memo with user intent
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq)]
pub struct MayaRouterTransferVaultEvent {
    pub from_vault_key: PublicKey, // Public key of the Vault, which sends assets
    pub to_vault_key: PublicKey,   // Public key of the Vault, which receives assets
    pub asset: ResourceAddress,    // Resource address of the transferred assets
    pub amount: Decimal,           // Amount of the transferred assets
    pub memo: String,              // Maya Transaction memo with user intent
}
