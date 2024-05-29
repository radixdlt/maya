use crate::*;
use clap::Parser;
use maya_router_cli::maya_router::*;
use radix_common::prelude::*;
use radix_engine_interface::prelude::*;

#[derive(Parser, Debug)]
pub struct DepositAndWithdraw {
    /// The hex-encoded private key of the package owner.
    #[clap(default_value = "0000000000000000000000000000000000000000000000007d336518a87c3284")]
    owner_ed25519_private_key: String,

    #[clap(
        default_value = "component_tdx_2_1crv4mqytgh9gqnnrzkgpwx7kwptrqgezjpvcjym78vsp8s836gkc63"
    )]
    component_address: String,

    /// Network to Use [Stokenet | Mainnet]
    #[clap(short, long)]
    network: Option<String>,
}

impl DepositAndWithdraw {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        let mut maya_router = MayaRouterTester::new(
            &self.network.unwrap_or("stokenet".to_string()),
            Some(&self.owner_ed25519_private_key),
        );

        let address_bech32_decoder =
            AddressBech32Decoder::new(&maya_router.gateway_network_connector.network_definition);
        let component_address =
            ComponentAddress::try_from_bech32(&address_bech32_decoder, &self.component_address)
                .unwrap();

        maya_router.set_component_address(component_address);
        maya_router.deposit(
            maya_router.asgard_vault_1.address,
            maya_router.swapper.clone(),
            XRD,
            dec!(200),
            "SWAP:MAYA.CACAO",
        );

        maya_router.transfer_out(
            maya_router.asgard_vault_1.clone(),
            maya_router.swapper.clone(),
            XRD,
            dec!(100),
            "OUT:",
            dec!(10),
        );
        Ok(())
    }
}
