package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EcdsaSecp256k1Signature struct {
	Signature
	// A hex-encoded recoverable ECDSA Secp256k1 signature (65 bytes). The first byte is the recovery id, the remaining 64 bytes are the compact signature, ie `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
	signature_hex *string
}

// NewEcdsaSecp256k1Signature instantiates a new EcdsaSecp256k1Signature and sets the default values.
func NewEcdsaSecp256k1Signature() *EcdsaSecp256k1Signature {
	m := &EcdsaSecp256k1Signature{
		Signature: *NewSignature(),
	}
	return m
}

// CreateEcdsaSecp256k1SignatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEcdsaSecp256k1SignatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEcdsaSecp256k1Signature(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EcdsaSecp256k1Signature) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetSignatureHex gets the signature_hex property value. A hex-encoded recoverable ECDSA Secp256k1 signature (65 bytes). The first byte is the recovery id, the remaining 64 bytes are the compact signature, ie `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
// returns a *string when successful
func (m *EcdsaSecp256k1Signature) GetSignatureHex() *string {
	return m.signature_hex
}

// Serialize serializes information the current object
func (m *EcdsaSecp256k1Signature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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

// SetSignatureHex sets the signature_hex property value. A hex-encoded recoverable ECDSA Secp256k1 signature (65 bytes). The first byte is the recovery id, the remaining 64 bytes are the compact signature, ie `CONCAT(R, s)` where `R` and `s` are each 32-bytes in padded big-endian format.
func (m *EcdsaSecp256k1Signature) SetSignatureHex(value *string) {
	m.signature_hex = value
}

type EcdsaSecp256k1Signatureable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Signatureable
	GetSignatureHex() *string
	SetSignatureHex(value *string)
}
