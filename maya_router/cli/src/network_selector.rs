use std::str::FromStr;

use radix_common::network::*;

#[derive(Clone, Debug)]
pub struct NetworkConfig(NetworkDefinition);

impl NetworkConfig {
    pub fn new(network: &str) -> Self {
        let network_definition =
            NetworkDefinition::from_str(network).expect("Invalid network specified");
        Self(network_definition)
    }

    pub fn default() -> Self {
        Self(NetworkDefinition::stokenet())
    }

    pub fn gateway_base_url(&self) -> String {
        match self.0.logical_name.as_str() {
            "mainnet" => "https://mainnet.radixdlt.com".to_owned(),
            "stokenet" => "https://stokenet.radixdlt.com".to_owned(),
            _ => panic!("Network not supported"),
        }
    }

    pub fn network_definition(&self) -> NetworkDefinition {
        self.0.clone()
    }
}
