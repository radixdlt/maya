use crate::*;
use clap::Parser;
use maya_router_cli::constants::*;
use maya_router_cli::maya_router::*;
use radix_common::prelude::*;

#[derive(Parser, Debug)]
pub struct GetBalance {
    /// The hex-encoded Ed25519 private key of the Asgard Vault
    #[clap(
        short,
        long,
        default_value = ASGARD_VAULT_1_PRIVATE_KEY
    )]
    asgard_vault_private_key: String,

    /// The hex-encoded private key of the package owner.
    #[clap(
        short,
        long,
        default_value = OWNER_PRIVATE_KEY,
    )]
    owner_ed25519_private_key: String,

    /// Maya Router address
    #[clap(
        short,
        long,
        default_value = MAYA_ROUTER_ADDRESS,
    )]
    component_address: String,

    /// Network to Use [Stokenet | Mainnet]
    #[clap(short, long, default_value = "stokenet")]
    network: String,
}

impl GetBalance {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let mut maya_router =
            MayaRouterTester::new(&self.network, Some(&self.owner_ed25519_private_key));

        let address_bech32_decoder =
            AddressBech32Decoder::new(&maya_router.gateway_network_connector.network_definition);
        let address_bech32_encoder =
            AddressBech32Encoder::new(&maya_router.gateway_network_connector.network_definition);
        let component_address =
            ComponentAddress::try_from_bech32(&address_bech32_decoder, &self.component_address)
                .unwrap();

        maya_router.set_component_address(component_address);
        let asgarg_vault_account =
            AccountData::new_from_private_key(KeyType::Ed25519, &self.asgard_vault_private_key);

        let balance = maya_router.get_vault_balance(asgarg_vault_account.clone(), XRD);
        writeln!(
            f,
            "Asgard Vault\n\t{}\nbalance\n\t{} XRD",
            asgarg_vault_account
                .address
                .display(&address_bech32_encoder),
            balance,
        )
        .map_err(Error::IoError)?;

        Ok(())
    }
}
