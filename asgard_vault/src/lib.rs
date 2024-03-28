use scrypto::prelude::*;

#[blueprint]
#[events(AsgardSwapEvent)]
mod asgard_vault {
    enable_method_auth! {
        roles {
            // Bifrost Signer is allowed to perform finalized transactions
            signer => updatable_by: [signer, OWNER];
        },
        methods {
            swap => PUBLIC;
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

        // Swap XRD into some other assets
        pub fn swap(&mut self, xrds: Bucket, into: String, dest_address: String) {
            self.xrd.put(xrds);

            // Send swap event to notify Bifrost Observer
            Runtime::emit_event(AsgardSwapEvent {
                into,
                dest_address,
                limit: None,
                affiliate: None,
                fee: None,
            });
        }

        // Send some amount of XRD to given address
        // Only Bifrost Signer is allowed to call it when finalizing
        // - swap from some asset to XRD
        // - remove liquidity request
        pub fn signer_send(&mut self, amount: Decimal, address: Global<AnyComponent>) {
            let bucket = self.xrd.take(amount);

            let mut account = Account::new(*address.handle());

            account.try_deposit_or_abort(bucket, None);
        }
    }
}

#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq, PartialOrd, Ord)]
pub struct AsgardSwapEvent {
    pub into: String,
    pub dest_address: String,
    pub limit: Option<Decimal>,
    pub affiliate: Option<String>,
    pub fee: Option<Decimal>,
}
