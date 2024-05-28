use crate::*;
use clap::Parser;

#[derive(Parser, Debug)]
pub struct Publish {
    /// The configuration that the user wants to use when publishing.
    configuration_selector: String,
    // ConfigurationSelector,
    /// The hex-encoded private key of the notary.
    notary_ed25519_private_key_hex: String,
}

impl Publish {
    pub fn run<O: std::io::Write>(self, f: &mut O) -> Result<(), Error> {
        writeln!(f, "{}", "publishing").map_err(Error::IoError)
    }
}
