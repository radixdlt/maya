use crate::*;
use clap::Parser;
use maya_router_cli::manifest_execute::execute;
use maya_router_cli::network_selector::NetworkConfig;
use maya_router_cli::utils::*;
use radix_common::prelude::*;
use radix_engine_interface::prelude::*;
use radix_transactions::prelude::*;
use rand::prelude::*;

#[derive(Parser, Debug)]
pub struct CreateAccount {
    /// Private key input value
    #[clap(short, long)]
    key_init: Option<u64>,

    /// If true, then don't initialize the account
    #[clap(short, long)]
    unitialized: bool,

    /// Network to Use [Stokenet | Mainnet]
    #[clap(short, long, default_value = "stokenet")]
    network: Option<String>,
}

impl CreateAccount {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let mut std_rng = rand::rngs::StdRng::from_entropy();

        let network = match &self.network {
            Some(network) => NetworkConfig::new(network),
            None => NetworkConfig::default(),
        };
        let key_init = self.key_init.unwrap_or(std_rng.next_u64());

        let private_key =
            Ed25519PrivateKey::from_u64(key_init).map_err(|_| Error::PrivateKeyError)?;
        let public_key = private_key.public_key();
        let address = ComponentAddress::virtual_account_from_public_key(&public_key);
        let address_bech32_encoder = AddressBech32Encoder::new(&network.network_definition());

        if !self.unitialized {
            log::info!("Initializing account");
            let private_key = clone_private_key_ed25519(&private_key);
            let manifest = ManifestBuilder::new()
                .lock_fee(FAUCET, dec!(100))
                .get_free_xrd_from_faucet()
                .take_all_from_worktop(XRD, "bucket")
                .try_deposit_or_abort(address, None, "bucket")
                .build();
            let result = execute(network, private_key.into(), manifest)?;

            writeln!(f, "TxId: {}", result.intent_hash).map_err(Error::IoError)?;
        }

        writeln!(f, "PrivateKey: {}", hex::encode(private_key.to_bytes()))
            .map_err(Error::IoError)?;
        writeln!(f, "PublicKey: {}", hex::encode(public_key.to_vec())).map_err(Error::IoError)?;
        writeln!(f, "Address: {}", address.display(&address_bech32_encoder)).map_err(Error::IoError)
    }
}
