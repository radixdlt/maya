use scrypto::prelude::*;

#[blueprint]
#[events(
    MayaRouterDepositEvent,
    MayaRouterTransferOutEvent,
    MayaRouterTransferAsgardVaultEvent
)]
mod maya_router {
    enable_method_auth! {
        methods {
            deposit => PUBLIC;
            transfer_out => PUBLIC;
            transfer_out_to_asgard_vault => PUBLIC;
        }
    }

    // MayaRouter owns vaults with resources per Asgard Vault public key.
    struct MayaRouter {
        vaults: KeyValueStore<Ed25519PublicKey, KeyValueStore<ResourceAddress, Vault>>,
    }

    impl MayaRouter {
        pub fn instantiate() -> Global<MayaRouter> {
            let owner_role = OwnerRole::None;

            Self {
                vaults: KeyValueStore::new(),
            }
            .instantiate()
            .prepare_to_globalize(owner_role)
            .globalize()
        }

        fn asgard_vault_take(
            &mut self,
            asgard_vault: Ed25519PublicKey,
            asset: ResourceAddress,
            amount: Option<Decimal>,
        ) -> Bucket {
            // Check if Asgard Vault and given asset exists in Asgard Vault.
            // If yes then get its balance.
            let balance = match self.vaults.get(&asgard_vault) {
                Some(asset_vault_kv_store) => match asset_vault_kv_store.get(&asset) {
                    Some(vault) => vault.amount(),
                    None => Runtime::panic(format!(
                        "asset {:?} not available in the asgard vault {:?}",
                        asset, asgard_vault
                    )),
                },
                None => Runtime::panic(format!("asgard vault {:?} not available", asgard_vault)),
            };

            // Check the amount to be taken.
            // If greater than balance then return error.
            let amount = if let Some(amount) = amount {
                if amount > balance {
                    Runtime::panic(format!(
                        "asgard vault {:?} balance {:?} lower than taken amount {:?}",
                        asgard_vault, balance, amount
                    ));
                } else {
                    amount
                }
            } else {
                balance
            };

            let mut asset_vault_kv_store = self.vaults.get_mut(&asgard_vault).unwrap();
            // Would be nice to remove empty vault from the store, but currently
            // is not possible to remove objects persisted in substore store.
            let mut vault = asset_vault_kv_store.get_mut(&asset).unwrap();
            vault.take(amount)
        }

        fn asgard_vault_put(&mut self, asgard_vault: Ed25519PublicKey, bucket: Bucket) {
            let asset = bucket.resource_address();

            let (asgard_vault_exists, asset_exists) = match self.vaults.get(&asgard_vault) {
                Some(asset_vault_kv_store) => match asset_vault_kv_store.get(&asset) {
                    Some(_) => (true, true),
                    None => (true, false),
                },
                None => (false, false),
            };

            if asgard_vault_exists {
                let mut asset_vault_kv_store = self.vaults.get_mut(&asgard_vault).unwrap();

                if asset_exists {
                    let mut vault = asset_vault_kv_store.get_mut(&asset).unwrap();
                    vault.put(bucket);
                } else {
                    asset_vault_kv_store.insert(asset, Vault::with_bucket(bucket));
                }
            } else {
                let asset_vault_kv_store = KeyValueStore::new();
                asset_vault_kv_store.insert(asset, Vault::with_bucket(bucket));

                self.vaults.insert(asgard_vault, asset_vault_kv_store);
            };
        }

        // Deposit some assets
        //   sender       - Address where to return refund if required. Must be the address of the method caller
        //   asgard_vault - Public key of the Asgard Vault, which shall control deposited resources
        //   bucket       - bucket of assets
        //   memo         - message to emit when emitting deposit event
        pub fn deposit(
            &mut self,
            sender: Global<Account>,
            asgard_vault: Ed25519PublicKey,
            bucket: Bucket,
            memo: String,
        ) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            let amount = bucket.amount();
            let asset = bucket.resource_address();

            self.asgard_vault_put(asgard_vault, bucket);

            // Send deposit event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterDepositEvent {
                sender: sender.address(),
                asgard_vault,
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of given asset to given address.
        //   asgard_vault - Public key of the Asgard Vault, which controls transferred assets
        //   receiver     - Address where to send assets (must be a real account)
        //   asset        - Resource address of the asset to send
        //   amount       - amount of assets to send
        //   memo         - message to emit when emitting sending the assets
        pub fn transfer_out(
            &mut self,
            asgard_vault: Ed25519PublicKey,
            receiver: Global<Account>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
        ) {
            // Make sure asgard vault is the one that calls this method
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &asgard_vault
            ))));

            let bucket = self.asgard_vault_take(asgard_vault, asset, Some(amount));

            let mut account = Account::new(*receiver.handle());

            // TODO: Use Account locker in case deposit fails
            account.try_deposit_or_abort(bucket, None);

            // Send transfer out event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterTransferOutEvent {
                asgard_vault,
                receiver: receiver.address(),
                asset,
                amount,
                memo,
            });
        }

        pub fn transfer_out_to_asgard_vault(
            &mut self,
            from_asgard_vault: Ed25519PublicKey,
            to_asgard_vault: Ed25519PublicKey,
            asset: ResourceAddress,
            memo: String,
        ) {
            // Make sure asgard vault is the one that calls this method
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &from_asgard_vault
            ))));

            let bucket = self.asgard_vault_take(from_asgard_vault, asset, None);
            let amount = bucket.amount();
            self.asgard_vault_put(to_asgard_vault, bucket);

            // Send transfer out event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterTransferAsgardVaultEvent {
                from_asgard_vault,
                to_asgard_vault,
                asset,
                amount,
                memo,
            });
        }
    }
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterDepositEvent {
    pub sender: ComponentAddress,       // Address of the deposit sender
    pub asgard_vault: Ed25519PublicKey, // Public key of the Asgard Vault, which controls deposited assets
    pub asset: ResourceAddress,         // Resource address of the deposited assets
    pub amount: Decimal,                // Amount of the deposited assets
    pub memo: String,                   // Maya Transaction memo with user intent
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterTransferOutEvent {
    pub asgard_vault: Ed25519PublicKey, // Public key of the Asgard Vault, which sends assets
    pub receiver: ComponentAddress,     // Address where assets were transferred
    pub asset: ResourceAddress,         // Resource address of the transferred assets
    pub amount: Decimal,                // Amount of the transferred assets
    pub memo: String,                   // Maya Transaction memo with user intent
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterTransferAsgardVaultEvent {
    pub from_asgard_vault: Ed25519PublicKey, // Public key of the Asgard Vault, which sends assets
    pub to_asgard_vault: Ed25519PublicKey, // Public key of the Asgard Vault, which receives assets
    pub asset: ResourceAddress,            // Resource address of the transferred assets
    pub amount: Decimal,                   // Amount of the transferred assets
    pub memo: String,                      // Maya Transaction memo with user intent
}
