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
        vaults: IndexMap<ResourceAddress, Vault>,
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
            owner_badge: NonFungibleGlobalId,
            admin_rule: AccessRule,
            admin: Global<AnyComponent>,
        ) -> Global<MayaRouter> {
            let owner_role = OwnerRole::Fixed(rule!(require(owner_badge)));

            Self {
                admin: admin.address(),
                vaults: IndexMap::new(),
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
        //   sender - Address where to return refund if required. Must be the address of the method caller.
        //   bucket - bucket of assets
        //   memo   - message to emit when emitting deposit event
        pub fn deposit(&mut self, sender: Global<AnyComponent>, bucket: Bucket, memo: String) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            let amount = bucket.amount();
            let asset = bucket.resource_address();

            match self.vaults.get_mut(&asset) {
                Some(vault) => vault.put(bucket),
                None => {
                    self.vaults.insert(asset, Vault::with_bucket(bucket));
                }
            }

            // Send deposit event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterDepositEvent {
                sender: sender.address(),
                receiver: self.admin,
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of given asset to given address (only Bifrost Signer is allowed to call it).
        //   sender   - Address of the account, which currently controls the vault (has "admin" role)
        //   receiver - Address where to send assets (must be a real account)
        //   asset    - Resource address of the asset to send
        //   amount   - amount of asset to send
        //   memo     - message to emit when emitting sending the assets
        pub fn transfer_out(
            &mut self,
            sender: Global<AnyComponent>,
            receiver: Global<AnyComponent>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
        ) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            // Make sure receiver is a real account, not component.
            // Malicious component could eg. implement 'try_deposit_or_abort' method
            // to consume all gas, eg. with busy loop
            if receiver.blueprint_id().package_address != ACCOUNT_PACKAGE {
                Runtime::panic(format!(
                    "address {:?} is not a real account",
                    ComponentAddress::try_from(receiver.handle().as_node_id().as_bytes()).unwrap()
                ));
            }

            if let Some(vault) = self.vaults.get_mut(&asset) {
                let bucket = vault.take(amount);

                let mut account = Account::new(*receiver.handle());

                account.try_deposit_or_abort(bucket, None);

                // Send transfer out event to notify Bifrost Observer
                Runtime::emit_event(MayaRouterTransferOutEvent {
                    sender: sender.address(),
                    receiver: receiver.address(),
                    asset,
                    amount,
                    memo,
                });
            } else {
                Runtime::panic(format!("asset {:?} not available in the vault", asset));
            }
        }
    }
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterDepositEvent {
    pub sender: ComponentAddress,   // Address of the deposit sender
    pub receiver: ComponentAddress, // Address of the account, which currently controls the vault (has "admin" role)
    pub asset: ResourceAddress,     // Resource address of the deposited assets
    pub amount: Decimal,            // Amount of the deposited assets
    pub memo: String,               // Transaction memo
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterTransferOutEvent {
    pub sender: ComponentAddress, // Address of the account, which currently controls the vault (has "admin" role)
    pub receiver: ComponentAddress, // Address where assets were transferred
    pub asset: ResourceAddress,   // Resource address of the transferred assets
    pub amount: Decimal,          // Amount of the transferred assets
    pub memo: String,             // Transaction memo
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterUpdateAdminEvent {
    pub previous_admin: ComponentAddress, // Address of the previous admin account
    pub current_admin: ComponentAddress,  // Address of the current admin account
}
