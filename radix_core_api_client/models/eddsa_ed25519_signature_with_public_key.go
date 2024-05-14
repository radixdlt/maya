package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EddsaEd25519SignatureWithPublicKey the EdDSA public key and signature
type EddsaEd25519SignatureWithPublicKey struct {
	SignatureWithPublicKey
	// The public_key property
	public_key EddsaEd25519PublicKeyable
	// The signature property
	signature EddsaEd25519Signatureable
}

// NewEddsaEd25519SignatureWithPublicKey instantiates a new EddsaEd25519SignatureWithPublicKey and sets the default values.
func NewEddsaEd25519SignatureWithPublicKey() *EddsaEd25519SignatureWithPublicKey {
	m := &EddsaEd25519SignatureWithPublicKey{
		SignatureWithPublicKey: *NewSignatureWithPublicKey(),
	}
	return m
}

// CreateEddsaEd25519SignatureWithPublicKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEddsaEd25519SignatureWithPublicKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEddsaEd25519SignatureWithPublicKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EddsaEd25519SignatureWithPublicKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SignatureWithPublicKey.GetFieldDeserializers()
	res["public_key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEddsaEd25519PublicKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublicKey(val.(EddsaEd25519PublicKeyable))
		}
		return nil
	}
	res["signature"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEddsaEd25519SignatureFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSignature(val.(EddsaEd25519Signatureable))
		}
		return nil
	}
	return res
}

// GetPublicKey gets the public_key property value. The public_key property
// returns a EddsaEd25519PublicKeyable when successful
func (m *EddsaEd25519SignatureWithPublicKey) GetPublicKey() EddsaEd25519PublicKeyable {
	return m.public_key
}

// GetSignature gets the signature property value. The signature property
// returns a EddsaEd25519Signatureable when successful
func (m *EddsaEd25519SignatureWithPublicKey) GetSignature() EddsaEd25519Signatureable {
	return m.signature
}

// Serialize serializes information the current object
func (m *EddsaEd25519SignatureWithPublicKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SignatureWithPublicKey.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("public_key", m.GetPublicKey())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("signature", m.GetSignature())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetPublicKey sets the public_key property value. The public_key property
func (m *EddsaEd25519SignatureWithPublicKey) SetPublicKey(value EddsaEd25519PublicKeyable) {
	m.public_key = value
}

// SetSignature sets the signature property value. The signature property
func (m *EddsaEd25519SignatureWithPublicKey) SetSignature(value EddsaEd25519Signatureable) {
	m.signature = value
}

type EddsaEd25519SignatureWithPublicKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SignatureWithPublicKeyable
	GetPublicKey() EddsaEd25519PublicKeyable
	GetSignature() EddsaEd25519Signatureable
	SetPublicKey(value EddsaEd25519PublicKeyable)
	SetSignature(value EddsaEd25519Signatureable)
}
