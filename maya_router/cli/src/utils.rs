use radix_common::crypto::{Ed25519PrivateKey, Secp256k1PrivateKey};
use radix_engine_interface::blueprints::package::PackageDefinition;
use radix_transactions::signing::PrivateKey;
use scrypto_test::ledger_simulator::CompileProfile;
use std::path::PathBuf;
use std::str::FromStr;

pub fn clone_private_key_ed25519(private_key: &Ed25519PrivateKey) -> Ed25519PrivateKey {
    Ed25519PrivateKey::from_bytes(&private_key.to_bytes()).expect("Cannot happen")
}

pub fn clone_private_key_secp2561k(private_key: &Secp256k1PrivateKey) -> Secp256k1PrivateKey {
    Secp256k1PrivateKey::from_bytes(&private_key.to_bytes()).expect("Cannot happen")
}

pub fn clone_private_key(private_key: &PrivateKey) -> PrivateKey {
    match private_key {
        PrivateKey::Ed25519(private_key) => PrivateKey::Ed25519(clone_private_key_ed25519(private_key)),
        PrivateKey::Secp256k1(private_key) => PrivateKey::Secp256k1(clone_private_key_secp2561k(private_key)),
    }
}

pub fn package_get() -> (Vec<u8>, PackageDefinition) {
    let package_dir = PathBuf::from_str(env!("CARGO_MANIFEST_DIR"))
        .unwrap()
        .parent()
        .unwrap()
        .join("package");
    scrypto_test::prelude::Compile::compile(package_dir, CompileProfile::Standard)
}

