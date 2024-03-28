use scrypto::prelude::*;

#[blueprint]
#[events(AsgardDepositEvent, AsgardTransferOutEvent)]
mod asgard_vault {
    enable_method_auth! {
        roles {
            // Bifrost Signer is allowed to perform finalized transactions
            signer => updatable_by: [signer, OWNER];
        },
        methods {
            deposit => PUBLIC;
            signer_send => restrict_to: [signer];
        }
    }

    struct AsgardVault {
        xrd: Vault,
    }

    impl AsgardVault {
        pub fn instantiate(
            owner_badge: NonFungibleGlobalId,
            signer_rule: AccessRule,
        ) -> Global<AsgardVault> {
            let owner_role = OwnerRole::Fixed(rule!(require(owner_badge)));

            Self {
                xrd: Vault::new(XRD),
            }
            .instantiate()
            .prepare_to_globalize(owner_role)
            .roles(roles! {
                signer => signer_rule;
            })
            .globalize()
        }

        // Deposit XRD into vault
        pub fn deposit(&mut self, xrds: Bucket, memo: String) {
            let amount = xrds.amount();
            let asset = xrds.resource_address();

            // TODO: add multiple vaults
            self.xrd.put(xrds);

            // Send deposit event to notify Bifrost Observer
            Runtime::emit_event(AsgardDepositEvent {
                vault: Runtime::global_address(),
                asset,
                amount,
                memo,
            });
        }

        // Send some amount of XRD to given address
        // Only Bifrost Signer is allowed to call it when finalizing
        // - swap from some asset to XRD
        // - remove liquidity request
        pub fn signer_send(
            &mut self,
            address: Global<AnyComponent>,
            asset: ResourceAddress,
            amount: Decimal,
            memo: String,
        ) {
            let bucket = self.xrd.take(amount);

            let mut account = Account::new(*address.handle());

            account.try_deposit_or_abort(bucket, None);

            // Send transfer out event to notify Bifrost Observer
            Runtime::emit_event(AsgardTransferOutEvent {
                vault: Runtime::global_address(),
                to: address.address(),
                asset,
                amount,
                memo,
            });
        }
    }
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct AsgardDepositEvent {
    pub vault: ComponentAddress, // Asgard Vault address
    pub asset: ResourceAddress,  // Resource address of the deposited assets
    pub amount: Decimal,         // Amount of the deposited assets
    pub memo: String,            // Transaction memo
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct AsgardTransferOutEvent {
    pub vault: ComponentAddress, // Asgard Vault address
    pub to: ComponentAddress,    // Address where to transfer to
    pub asset: ResourceAddress,  // Resource address of the transferred assets
    pub amount: Decimal,         // Amount of the transferred assets
    pub memo: String,            // Transaction memo
}
