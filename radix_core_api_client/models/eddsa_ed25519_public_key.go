package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EddsaEd25519PublicKey struct {
	PublicKey
	// The hex-encoded compressed EdDSA Ed25519 public key (32 bytes)
	key_hex *string
}

// NewEddsaEd25519PublicKey instantiates a new EddsaEd25519PublicKey and sets the default values.
func NewEddsaEd25519PublicKey() *EddsaEd25519PublicKey {
	m := &EddsaEd25519PublicKey{
		PublicKey: *NewPublicKey(),
	}
	return m
}

// CreateEddsaEd25519PublicKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEddsaEd25519PublicKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEddsaEd25519PublicKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EddsaEd25519PublicKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetKeyHex gets the key_hex property value. The hex-encoded compressed EdDSA Ed25519 public key (32 bytes)
// returns a *string when successful
func (m *EddsaEd25519PublicKey) GetKeyHex() *string {
	return m.key_hex
}

// Serialize serializes information the current object
func (m *EddsaEd25519PublicKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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

// SetKeyHex sets the key_hex property value. The hex-encoded compressed EdDSA Ed25519 public key (32 bytes)
func (m *EddsaEd25519PublicKey) SetKeyHex(value *string) {
	m.key_hex = value
}

type EddsaEd25519PublicKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PublicKeyable
	GetKeyHex() *string
	SetKeyHex(value *string)
}
