use scrypto::prelude::*;

#[blueprint]
#[events(
    MayaRouterDepositEvent,
    MayaRouterTransferOutEvent,
    MayaRouterUpdateSignerEvent
)]
mod maya_router {
    enable_method_auth! {
        roles {
            // Bifrost Signer is allowed to perform finalized transactions
            signer => updatable_by: [SELF];
        },
        methods {
            deposit => PUBLIC;
            transfer_out => restrict_to: [signer];
            update_signer => restrict_to: [signer];
        }
    }

    struct MayaRouter {
        signer: ComponentAddress,
        vaults: IndexMap<ResourceAddress, Vault>,
    }

    // TODO:
    // - consider a potential need to support old and new badge for signer role
    //   for a given period of time
    //   eg. old vault which is being retired still accepts
    //       deposits, and those deposits shall be possible to transfer out.
    // - consider adding a method that transfers all resources to another MayaRouter
    //   (old one could be compromised or lacks some features)
    impl MayaRouter {
        pub fn instantiate(
            owner_badge: NonFungibleGlobalId,
            signer_rule: AccessRule,
            signer: Global<AnyComponent>,
        ) -> Global<MayaRouter> {
            let owner_role = OwnerRole::Fixed(rule!(require(owner_badge)));

            Self {
                signer: signer.address(),
                vaults: IndexMap::new(),
            }
            .instantiate()
            .prepare_to_globalize(owner_role)
            .roles(roles! {
                signer => signer_rule;
            })
            .globalize()
        }

        pub fn update_signer(&mut self, signer_rule: AccessRule, signer: Global<AnyComponent>) {
            Runtime::global_component().set_role("signer", signer_rule);

            // Send update signer event to notify Bifrost Observer
            Runtime::emit_event(MayaRouterUpdateSignerEvent {
                previous_signer: self.signer,
                current_signer: signer.address(),
            });
            self.signer = signer.address();
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
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of given asset to given address (only Bifrost Signer is allowed to call it).
        //   sender  - Address of the account, which currently controls the vault (has "signer" role)
        //   address - Address where to send assets
        //   asset   - Resource address of the asset to send
        //   amount  - amount of asset to send
        //   memo    - message to emit when emitting sending the assets
        pub fn transfer_out(
            &mut self,
            sender: Global<AnyComponent>,
            address: Global<AnyComponent>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
        ) {
            // Make sure sender is the one that calls this method
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            // Make sure address is a real account, not component.
            // Malicious component could eg. implement 'try_deposit_or_abort' method
            // to consume all gas, eg. with busy loop
            if address.blueprint_id().package_address != ACCOUNT_PACKAGE {
                Runtime::panic(format!(
                    "address {:?} is not a real account",
                    ComponentAddress::try_from(address.handle().as_node_id().as_bytes()).unwrap()
                ));
            }

            if let Some(mut vault) = self.vaults.get_mut(&asset) {
                let bucket = vault.take(amount);

                let mut account = Account::new(*address.handle());

                account.try_deposit_or_abort(bucket, None);

                // Send transfer out event to notify Bifrost Observer
                Runtime::emit_event(MayaRouterTransferOutEvent {
                    sender: sender.address(),
                    address: address.address(),
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
    pub sender: ComponentAddress, // Address of the deposit sender
    pub asset: ResourceAddress,   // Resource address of the deposited assets
    pub amount: Decimal,          // Amount of the deposited assets
    pub memo: String,             // Transaction memo
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterTransferOutEvent {
    pub sender: ComponentAddress, // Address of the sender (must have the "signer" role)
    pub address: ComponentAddress, // Address where to transfer to
    pub asset: ResourceAddress,   // Resource address of the transferred assets
    pub amount: Decimal,          // Amount of the transferred assets
    pub memo: String,             // Transaction memo
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct MayaRouterUpdateSignerEvent {
    pub previous_signer: ComponentAddress, // Address of the previous signer
    pub current_signer: ComponentAddress,  // Address of the current signer
}
