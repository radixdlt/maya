use crate::*;
use clap::Parser;
use maya_router_cli::maya_router::*;
use radix_common::prelude::*;

#[derive(Parser, Debug)]
pub struct PublishAndInstantiate {
    /// The hex-encoded private key of the package owner.
    #[clap(default_value = "0000000000000000000000000000000000000000000000007d336518a87c3284")]
    owner_ed25519_private_key: String,

    /// Do not instantiate package after publishing
    #[clap(short, long)]
    dont_instantiate: bool,

    /// Network to Use [Stokenet | Mainnet]
    #[clap(short, long)]
    network: Option<String>,
}

impl PublishAndInstantiate {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let mut maya_router = MayaRouterTester::new(
            &self.network.unwrap_or("stokenet".to_string()),
            Some(&self.owner_ed25519_private_key),
        );

        let package_address = maya_router.publish();

        let address_bech32_encoder =
            AddressBech32Encoder::new(&maya_router.gateway_network_connector.network_definition);
        writeln!(
            f,
            "PackageAddress: {}",
            package_address.display(&address_bech32_encoder)
        )
        .map_err(Error::IoError)?;

        if !self.dont_instantiate {
            let _ = maya_router.instantiate(package_address);

            writeln!(f, "MayaRouter info: \n{}", maya_router).map_err(Error::IoError)
        } else {
            Ok(())
        }
    }
}
