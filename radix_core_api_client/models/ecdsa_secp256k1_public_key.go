package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EcdsaSecp256k1PublicKey struct {
	PublicKey
	// The hex-encoded compressed ECDSA Secp256k1 public key (33 bytes)
	key_hex *string
}

// NewEcdsaSecp256k1PublicKey instantiates a new EcdsaSecp256k1PublicKey and sets the default values.
func NewEcdsaSecp256k1PublicKey() *EcdsaSecp256k1PublicKey {
	m := &EcdsaSecp256k1PublicKey{
		PublicKey: *NewPublicKey(),
	}
	return m
}

// CreateEcdsaSecp256k1PublicKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEcdsaSecp256k1PublicKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEcdsaSecp256k1PublicKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EcdsaSecp256k1PublicKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PublicKey.GetFieldDeserializers()
	res["key_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyHex(val)
		}
		return nil
	}
	return res
}

// GetKeyHex gets the key_hex property value. The hex-encoded compressed ECDSA Secp256k1 public key (33 bytes)
// returns a *string when successful
func (m *EcdsaSecp256k1PublicKey) GetKeyHex() *string {
	return m.key_hex
}

// Serialize serializes information the current object
func (m *EcdsaSecp256k1PublicKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.PublicKey.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("key_hex", m.GetKeyHex())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKeyHex sets the key_hex property value. The hex-encoded compressed ECDSA Secp256k1 public key (33 bytes)
func (m *EcdsaSecp256k1PublicKey) SetKeyHex(value *string) {
	m.key_hex = value
}

type EcdsaSecp256k1PublicKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PublicKeyable
	GetKeyHex() *string
	SetKeyHex(value *string)
}
