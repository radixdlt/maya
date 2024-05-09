use scrypto::prelude::*;

#[blueprint]
#[types(Vec<u8>, ResourceAddress, FungibleVault, KeyValueStore<ResourceAddress, FungibleVault>)]
#[events(MayaRouterDepositEvent, MayaRouterTransferOutEvent)]
mod maya_router {
    enable_method_auth! {
        methods {
            deposit => PUBLIC;
            transfer_out => PUBLIC;
            transfer_via_locker => PUBLIC;
            direct_deposit_to_vault => PUBLIC;
            get_vault_balance => PUBLIC;
        }
    }

    // The MayaRouter component is responsible for managing access to the funds controlled
    // by Maya vaults (e.g. Asgard).
    // We use the term "vault_key" and "vault_address" to refer to Maya vaults (e.g. Asgard) - specifically,
    // to their public keys and the derived virtual account addresses respectively.
    // Not to be confused with Scrypto's vaults (e.g. FungibleVault).
    struct MayaRouter {
        locker: Global<AccountLocker>,
        // This KV store holds all the funds.
        // Each "Maya vault" owns a collection of Scrypto `FungibleVault`s (one for each owned resource).
        // For technical reasons, Maya vaults are identified by the bytes of virtual account addresses derived
        // from their public keys (e.g. `ComponentAddress::virtual_account_from_public_key(vault_key).to_vec()`).
        vaults: KeyValueStore<Vec<u8>, KeyValueStore<ResourceAddress, FungibleVault>>,
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
                vaults: KeyValueStore::new_with_registered_type(),
            }
            .instantiate()
            .prepare_to_globalize(OwnerRole::None)
            .with_address(maya_router_address_reservation)
            .globalize()
        }

        // Take some amount of resource from the vault corresponding to the given `vault_address`.
        // Optionally also lock an XRD fee.
        fn vault_take(
            &mut self,
            vault_address: ComponentAddress,
            resource_address: ResourceAddress,
            amount: Decimal,
            fee_to_lock: Decimal,
        ) -> FungibleBucket {
            if fee_to_lock.is_negative() {
                Runtime::panic(format!("Negative fee to lock {:?} provided", fee_to_lock));
            }

            let mut vault_resources = match self.vaults.get_mut(&vault_address.to_vec()) {
                Some(vault_resources) => vault_resources,
                None => Runtime::panic(format!(
                    "No resource has been deposited to vault {:?}",
                    vault_address
                )),
            };

            let mut vault = if resource_address == XRD {
                let mut vault = vault_resources
                    .get_mut(&XRD)
                    .expect("XRD not available in the vault");
                if fee_to_lock.is_positive() {
                    vault.lock_fee(fee_to_lock);
                }
                vault
            } else {
                if fee_to_lock.is_positive() {
                    vault_resources
                        .get_mut(&XRD)
                        .expect("XRD not available in the vault")
                        .lock_fee(fee_to_lock);
                }
                vault_resources.get_mut(&resource_address).expect(&format!(
                    "Resource {:?} not available in the vault {:?}",
                    resource_address, vault_address
                ))
            };

            if amount > vault.amount() {
                Runtime::panic(format!(
                    "Vault {:?} balance {:?} lower than taken amount {:?}",
                    vault_address,
                    vault.amount(),
                    amount
                ));
            } else {
                vault.take(amount)
            }
        }

        // Put a bucket of resource into the vault corresponding to the given `vault_address`.
        // Creates the vault if it doesn't exist.
        fn vault_put(&mut self, vault_address: ComponentAddress, bucket: FungibleBucket) {
            let resource_address = bucket.resource_address();

            let (vault_exists, resource_vault_exists) = match self
                .vaults
                .get(&vault_address.to_vec())
            {
                Some(vault_resources) => (true, vault_resources.get(&resource_address).is_some()),
                None => (false, false),
            };

            if vault_exists {
                let mut vault_resources = self
                    .vaults
                    .get_mut(&vault_address.to_vec())
                    .expect("Vault should be present");

                if resource_vault_exists {
                    let mut vault = vault_resources
                        .get_mut(&resource_address)
                        .expect("Resource should be present");
                    vault.put(bucket);
                } else {
                    vault_resources.insert(resource_address, FungibleVault::with_bucket(bucket));
                }
            } else {
                let vault_resources = KeyValueStore::new_with_registered_type();
                vault_resources.insert(resource_address, FungibleVault::with_bucket(bucket));

                self.vaults.insert(vault_address.to_vec(), vault_resources);
            };
        }

        // Deposit a bucket of resources to the vault corresponding to the given `vault_address`.
        // It is required that the provided `sender` account is one of the signers of the transaction.
        // `memo` is opaque to this component and will be forwarded to the Maya network.
        pub fn deposit(
            &mut self,
            sender: Global<Account>,
            vault_address: ComponentAddress,
            bucket: FungibleBucket,
            memo: String,
        ) {
            Runtime::assert_access_rule(sender.get_owner_role().rule);

            let amount = bucket.amount();
            let resource_address = bucket.resource_address();

            self.vault_put(vault_address, bucket);

            Runtime::emit_event(MayaRouterDepositEvent {
                sender: Runtime::bech32_encode_address(sender.address()),
                vault_address: Runtime::bech32_encode_address(vault_address),
                resource_address,
                amount,
                memo,
            });
        }

        // Withdraws a specified `amount` of resource from the vault corresponding to the given `vault_key`
        // and returns it to the caller.
        // It's up to the caller to perform an actual transfer of the funds to the end receiver.
        // `intended_receiver` parameter is passed as-is to the emmitted event, without any further verification.
        // A badge corresponding to the `vault_key` must be present (or, in other words, the transaction
        // must be signed by a corresponding private key).
        // This methods allows to optionally lock a fee from the XRD vault owned by `vault_key`.
        pub fn transfer_out(
            &mut self,
            vault_key: PublicKey,
            intended_receiver: String,
            resource_address: ResourceAddress,
            amount: Decimal,
            memo: String,
            fee_to_lock: Decimal,
        ) -> FungibleBucket {
            Runtime::assert_access_rule(rule!(require(NonFungibleGlobalId::from_public_key(
                &vault_key
            ))));

            let vault_address = ComponentAddress::virtual_account_from_public_key(&vault_key);

            let bucket = self.vault_take(vault_address, resource_address, amount, fee_to_lock);

            Runtime::emit_event(MayaRouterTransferOutEvent {
                vault_address: Runtime::bech32_encode_address(vault_address),
                intended_receiver,
                resource_address,
                amount,
                memo,
            });

            bucket
        }

        // Transfers the funds from `bucket` to `receiver` using this component's account locker.
        pub fn transfer_via_locker(
            &mut self,
            receiver: Global<Account>,
            bucket: FungibleBucket,
        ) {
            self.locker.store(receiver, bucket.into(), true);
        }

        // Directly deposits the funds from `bucket` to the vault identified by `to_vault_address`.
        pub fn direct_deposit_to_vault(
            &mut self,
            to_vault_address: ComponentAddress,
            bucket: FungibleBucket,
        ) {
            self.vault_put(to_vault_address, bucket);
        }

        // Returns a balance of the specified resource owned by a Maya vault corresponding to the given `vault_address`.
        pub fn get_vault_balance(
            &self,
            vault_address: ComponentAddress,
            resource_address: ResourceAddress,
        ) -> Decimal {
            let vault_kv_store = match self.vaults.get(&vault_address.to_vec()) {
                Some(vault_kv_store) => vault_kv_store,
                None => return Decimal::zero(),
            };

            let res = match vault_kv_store.get(&resource_address) {
                Some(vault) => vault.amount(),
                None => Decimal::zero(),
            };

            res
        }
    }
}

// An event indicating that the specified `amount` of resource (identified by `resource_address`) was deposited
// to a Maya vault (e.g. Asgard) identified by `vault_address`.
// The `memo` (opaque to this component) specifies a Maya-specific intent that was provided by the caller.
#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq)]
pub struct MayaRouterDepositEvent {
    pub sender: String,
    pub vault_address: String,
    pub resource_address: ResourceAddress,
    pub amount: Decimal,
    pub memo: String,
}

// An event indicating that the specified `amount` of resource (identified by `resource_address`) was withdrawn
// from a Maya vault (e.g. Asgard) identified by `vault_address` with the intention of being subsequently transferred to `intended_receiver`.
// The receiver can also be an address corresponding to another Maya vault key (in case of vault->vault transfers, or "migration").
// The `memo` (opaque to this component) specifies a Maya-specific intent that was provided by the caller.
#[derive(ScryptoSbor, ScryptoEvent, Debug, PartialEq, Eq)]
pub struct MayaRouterTransferOutEvent {
    pub vault_address: String,
    pub intended_receiver: String,
    pub resource_address: ResourceAddress,
    pub amount: Decimal,
    pub memo: String,
}
