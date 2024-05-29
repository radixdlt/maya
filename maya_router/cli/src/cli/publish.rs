use crate::*;
use clap::Parser;
use maya_router_cli::manifest_execute::execute;
use maya_router_cli::network_selector::NetworkConfig;
use maya_router_cli::utils::*;
use radix_common::prelude::*;
use radix_engine_interface::prelude::*;
use radix_transactions::prelude::*;

#[derive(Parser, Debug)]
pub struct Publish {
    /// The hex-encoded private key of the package owner.
    owner_ed25519_private_key: String,

    /// Do not instantiate package after publishing
    #[clap(short, long)]
    dont_instantiate: bool,

    /// Network to Use [Stokenet | Mainnet]
    #[clap(short, long)]
    network: Option<String>,
}

impl Publish {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let network = match &self.network {
            Some(network) => NetworkConfig::new(network),
            None => NetworkConfig::default(),
        };

        let owner_private_key_bytes =
            hex::decode(self.owner_ed25519_private_key).map_err(|_| Error::PrivateKeyError)?;
        let owner_private_key = Ed25519PrivateKey::from_bytes(&owner_private_key_bytes)
            .map_err(|_| Error::PrivateKeyError)?;

        let owner_address =
            ComponentAddress::virtual_account_from_public_key(&owner_private_key.public_key());
        let owner_badge = NonFungibleGlobalId::from_public_key(&owner_private_key.public_key());

        let (code, definition) = package_get();
        let mut metadata = MetadataInit::new();
        metadata.set_metadata("name", "MayaRouter Package");
        metadata.set_metadata(
            "description",
            "MayaRouter package stores assets swappable with assets from other networks",
        );
        let owner_role = OwnerRole::Fixed(rule!(require(owner_badge)));
        let manifest = ManifestBuilder::new()
            .lock_fee(owner_address, dec!(100))
            .publish_package_advanced(None, code, definition, metadata, owner_role)
            .build();

        log::info!("Publishing package");
        let result = execute(
            network.clone(),
            clone_private_key_ed25519(&owner_private_key).into(),
            manifest,
        )?;

        writeln!(f, "TxId: {}", result.intent_hash).map_err(Error::IoError)?;

        let address_bech32_encoder =
            AddressBech32Encoder::new(&network.clone().network_definition());
        let package_address = result
            .new_entities
            .new_package_addresses
            .get_index(0)
            .unwrap();
        writeln!(
            f,
            "PackageAddress: {}",
            package_address.display(&address_bech32_encoder)
        )
        .map_err(Error::IoError)?;

        if !self.dont_instantiate {
            log::info!("Instantiating package");
            let manifest = ManifestBuilder::new()
                .lock_fee(owner_address, dec!(100))
                .call_function(
                    *package_address,
                    "MayaRouter",
                    "instantiate",
                    manifest_args!(),
                )
                .build();
            let result = execute(network.clone(), owner_private_key.into(), manifest)?;

            writeln!(f, "TxId: {}", result.intent_hash).map_err(Error::IoError)?;
            let component_address = result
                .new_entities
                .new_component_addresses
                .get_index(0)
                .unwrap();
            writeln!(
                f,
                "ComponentAddress: {}",
                component_address.display(&address_bech32_encoder)
            )
            .map_err(Error::IoError)
        } else {
            Ok(())
        }
    }
}
