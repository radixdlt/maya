use radix_engine_interface::blueprints::account::ACCOUNT_BLUEPRINT;
use scrypto::prelude::*;

#[blueprint]
#[events(
    MayaRouterDepositEvent,
    MayaRouterTransferOutEvent,
    MayaRouterUpdateAdminEvent
)]
mod maya_router {
    enable_method_auth! {
        roles {
            // Bifrost Signer is allowed to perform finalized transactions
            admin => updatable_by: [SELF];
        },
        methods {
            deposit => PUBLIC;
            transfer_out => restrict_to: [admin];
            update_admin => restrict_to: [admin];
        }
    }

    struct MayaRouter {
        admin: ComponentAddress,
        vaults: KeyValueStore<Ed25519PublicKey, KeyValueStore<ResourceAddress, Vault>>,
    }

    // TODO:
    // - consider a potential need to support old and new badge for "admin" role
    //   for a given period of time
    //   eg. old vault which is being retired still accepts
    //       deposits, and those deposits shall be possible to transfer out.
    // - consider adding a method that transfers all resources to another MayaRouter
    //   (old one could be compromised or lacks some features)
    impl MayaRouter {
        pub fn instantiate(
            admin_rule: AccessRule,
            admin: Global<AnyComponent>,
        ) -> Global<MayaRouter> {
            let owner_role = OwnerRole::None;

            Self {
                admin: admin.address(),
                vaults: KeyValueStore::new(),
            }
            .instantiate()
            .prepare_to_globalize(owner_role)
            .roles(roles! {
                admin => admin_rule;
            })
            .globalize()
        }

        // Update "admin" role
        //   admin_rule - new access rule
        //   admin      - new admin account address
        pub fn update_admin(&mut self, admin_rule: AccessRule, admin: Global<AnyComponent>) {
            Runtime::global_component().set_role("admin", admin_rule);

            // Send update admin event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterUpdateAdminEvent {
                previous_admin: self.admin,
                current_admin: admin.address(),
            });
            self.admin = admin.address();
        }

        // Deposit some assets
        //   sender       - Address where to return refund if required. Must be the address of the method caller
        //   asgard_vault - Public key of the Asgard Vault, which shall control deposited resources
        //   bucket       - bucket of assets
        //   memo         - message to emit when emitting deposit event
        pub fn deposit(
            &mut self,
            sender: Global<AnyComponent>,
            asgard_vault: Ed25519PublicKey,
            bucket: Bucket,
            memo: String,
        ) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            let amount = bucket.amount();
            let asset = bucket.resource_address();

            let asgard_vault_exists = if let Some(_) = self.vaults.get(&asgard_vault) {
                true
            } else {
                false
            };

            if asgard_vault_exists {
                let mut asset_vault_kv_store = self.vaults.get_mut(&asgard_vault).unwrap();
                let asset_exists = if let Some(_) = asset_vault_kv_store.get(&asset) {
                    true
                } else {
                    false
                };

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

            // Send deposit event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterDepositEvent {
                sender: sender.address(),
                asgard_vault,
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of given asset to given address (only Bifrost Signer is allowed to call it).
        //   asgard_vault - Public key of the Asgard Vault, which controls transferred assets
        //   receiver     - Address where to send assets (must be a real account)
        //   asset        - Resource address of the asset to send
        //   amount       - amount of assets to send
        //   memo         - message to emit when emitting sending the assets
        pub fn transfer_out(
            &mut self,
            asgard_vault: Ed25519PublicKey,
            receiver: Global<AnyComponent>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
        ) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &asgard_vault
            ))));

            // Make sure receiver is a real account, not component.
            // Malicious component could eg. implement 'try_deposit_or_abort' method
            // to consume all gas, eg. with busy loop
            if !receiver.blueprint_id().eq(&BlueprintId {
                package_address: ACCOUNT_PACKAGE,
                blueprint_name: ACCOUNT_BLUEPRINT.to_string(),
            }) {
                Runtime::panic(format!(
                    "address {:?} is not a real account",
                    ComponentAddress::try_from(receiver.handle().as_node_id().as_bytes()).unwrap()
                ));
            }

            if let Some(mut asset_vault_kv_store) = self.vaults.get_mut(&asgard_vault) {
                if let Some(mut vault) = asset_vault_kv_store.get_mut(&asset) {
                    let bucket = vault.take(amount);

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
                } else {
                    Runtime::panic(format!(
                        "asset {:?} not available in the asgard vault {:?}",
                        asset, asgard_vault
                    ));
                }
            } else {
                Runtime::panic(format!("asgard vault {:?} not available", asgard_vault));
            }
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
pub struct MayaRouterUpdateAdminEvent {
    pub previous_admin: ComponentAddress, // Address of the previous admin account
    pub current_admin: ComponentAddress,  // Address of the current admin account
}
