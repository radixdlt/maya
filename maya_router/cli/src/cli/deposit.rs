use crate::*;
use clap::Parser;
use maya_router_cli::constants::*;
use maya_router_cli::maya_router::*;
use radix_common::prelude::*;
use radix_engine_interface::prelude::*;

#[derive(Parser, Debug)]
pub struct Deposit {
    /// The hex-encoded Ed25519 private key of the depositor
    #[clap(
        short,
        long,
        default_value = DEPOSITOR_PRIVATE_KEY,
    )]
    depositor_private_key: String,

    // Asgard Vault address
    #[clap(
        short,
        long,
        default_value = ASGARD_VAULT_1_ADDRESS
    )]
    asgard_vault_address: String,

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

impl Deposit {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let amount = dec!(200);
        let mut maya_router =
            MayaRouterTester::new(&self.network, Some(&self.owner_ed25519_private_key));

        let address_bech32_decoder =
            AddressBech32Decoder::new(&maya_router.gateway_network_connector.network_definition);
        let address_bech32_encoder =
            AddressBech32Encoder::new(&maya_router.gateway_network_connector.network_definition);
        let component_address =
            ComponentAddress::try_from_bech32(&address_bech32_decoder, &self.component_address)
                .unwrap();
        let asgard_vault_address =
            ComponentAddress::try_from_bech32(&address_bech32_decoder, &self.asgard_vault_address)
                .unwrap();
        let depositor_account =
            AccountData::new_from_private_key(KeyType::Ed25519, &self.depositor_private_key);

        maya_router.set_component_address(component_address);

        writeln!(
            f,
            "Depositing\n\t{} XRD\nfrom\n\t{}\nto Asgard Vault\n\t{}",
            amount,
            depositor_account.address.display(&address_bech32_encoder),
            self.asgard_vault_address
        )
        .map_err(Error::IoError)?;

        maya_router.deposit(
            asgard_vault_address,
            depositor_account,
            XRD,
            amount,
            "SWAP:MAYA.CACAO",
        );

        Ok(())
    }
}
