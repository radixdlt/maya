use radix_common::crypto::Ed25519PrivateKey;

pub fn clone_private_key(private_key: &Ed25519PrivateKey) -> Ed25519PrivateKey {
    Ed25519PrivateKey::from_bytes(&private_key.to_bytes()).expect("Cannot happen")
}

