package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EddsaEd25519Signature struct {
	Signature
	// A hex-encoded EdDSA Ed25519 signature (64 bytes). This is `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
	signature_hex *string
}

// NewEddsaEd25519Signature instantiates a new EddsaEd25519Signature and sets the default values.
func NewEddsaEd25519Signature() *EddsaEd25519Signature {
	m := &EddsaEd25519Signature{
		Signature: *NewSignature(),
	}
	return m
}

// CreateEddsaEd25519SignatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEddsaEd25519SignatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEddsaEd25519Signature(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EddsaEd25519Signature) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Signature.GetFieldDeserializers()
	res["signature_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSignatureHex(val)
		}
		return nil
	}
	return res
}

// GetSignatureHex gets the signature_hex property value. A hex-encoded EdDSA Ed25519 signature (64 bytes). This is `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
// returns a *string when successful
func (m *EddsaEd25519Signature) GetSignatureHex() *string {
	return m.signature_hex
}

// Serialize serializes information the current object
func (m *EddsaEd25519Signature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Signature.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("signature_hex", m.GetSignatureHex())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetSignatureHex sets the signature_hex property value. A hex-encoded EdDSA Ed25519 signature (64 bytes). This is `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
func (m *EddsaEd25519Signature) SetSignatureHex(value *string) {
	m.signature_hex = value
}

type EddsaEd25519Signatureable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Signatureable
	GetSignatureHex() *string
	SetSignatureHex(value *string)
}
