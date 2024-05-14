package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EcdsaSecp256k1SignatureWithPublicKey because ECDSA has recoverable signatures, this only includes a signature
type EcdsaSecp256k1SignatureWithPublicKey struct {
	SignatureWithPublicKey
	// The recoverable_signature property
	recoverable_signature EcdsaSecp256k1Signatureable
}

// NewEcdsaSecp256k1SignatureWithPublicKey instantiates a new EcdsaSecp256k1SignatureWithPublicKey and sets the default values.
func NewEcdsaSecp256k1SignatureWithPublicKey() *EcdsaSecp256k1SignatureWithPublicKey {
	m := &EcdsaSecp256k1SignatureWithPublicKey{
		SignatureWithPublicKey: *NewSignatureWithPublicKey(),
	}
	return m
}

// CreateEcdsaSecp256k1SignatureWithPublicKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEcdsaSecp256k1SignatureWithPublicKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEcdsaSecp256k1SignatureWithPublicKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EcdsaSecp256k1SignatureWithPublicKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SignatureWithPublicKey.GetFieldDeserializers()
	res["recoverable_signature"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEcdsaSecp256k1SignatureFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecoverableSignature(val.(EcdsaSecp256k1Signatureable))
		}
		return nil
	}
	return res
}

// GetRecoverableSignature gets the recoverable_signature property value. The recoverable_signature property
// returns a EcdsaSecp256k1Signatureable when successful
func (m *EcdsaSecp256k1SignatureWithPublicKey) GetRecoverableSignature() EcdsaSecp256k1Signatureable {
	return m.recoverable_signature
}

// Serialize serializes information the current object
func (m *EcdsaSecp256k1SignatureWithPublicKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SignatureWithPublicKey.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("recoverable_signature", m.GetRecoverableSignature())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRecoverableSignature sets the recoverable_signature property value. The recoverable_signature property
func (m *EcdsaSecp256k1SignatureWithPublicKey) SetRecoverableSignature(value EcdsaSecp256k1Signatureable) {
	m.recoverable_signature = value
}

type EcdsaSecp256k1SignatureWithPublicKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SignatureWithPublicKeyable
	GetRecoverableSignature() EcdsaSecp256k1Signatureable
	SetRecoverableSignature(value EcdsaSecp256k1Signatureable)
}
