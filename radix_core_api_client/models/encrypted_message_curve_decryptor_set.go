package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptedMessageCurveDecryptorSet a decryptor set for a particular ECDSA curve type.The (128-bit) AES-GCM symmetric key is encrypted separately for each decryptor public key via (256-bit) AES-KeyWrap.AES-KeyWrap uses a key derived via a KDF (Key Derivation Function) using a shared secret.For each decryptor public key, we create a shared curve point `G` via static Diffie-Helman between the decryptor public key, and a per-transaction ephemeral public key for that curve type.We then use that shared secret with a key derivation function to create the (256-bit) KEK (Key Encrypting Key):`KEK = HKDF(hash: Blake2b, secret: x co-ord of G, salt: [], length: 256 bits)`.
type EncryptedMessageCurveDecryptorSet struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The decryptors property
	decryptors []EncryptedMessageDecryptorable
	// The dh_ephemeral_public_key property
	dh_ephemeral_public_key PublicKeyable
}

// NewEncryptedMessageCurveDecryptorSet instantiates a new EncryptedMessageCurveDecryptorSet and sets the default values.
func NewEncryptedMessageCurveDecryptorSet() *EncryptedMessageCurveDecryptorSet {
	m := &EncryptedMessageCurveDecryptorSet{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEncryptedMessageCurveDecryptorSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEncryptedMessageCurveDecryptorSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEncryptedMessageCurveDecryptorSet(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EncryptedMessageCurveDecryptorSet) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDecryptors gets the decryptors property value. The decryptors property
// returns a []EncryptedMessageDecryptorable when successful
func (m *EncryptedMessageCurveDecryptorSet) GetDecryptors() []EncryptedMessageDecryptorable {
	return m.decryptors
}

// GetDhEphemeralPublicKey gets the dh_ephemeral_public_key property value. The dh_ephemeral_public_key property
// returns a PublicKeyable when successful
func (m *EncryptedMessageCurveDecryptorSet) GetDhEphemeralPublicKey() PublicKeyable {
	return m.dh_ephemeral_public_key
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EncryptedMessageCurveDecryptorSet) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["decryptors"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEncryptedMessageDecryptorFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EncryptedMessageDecryptorable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EncryptedMessageDecryptorable)
				}
			}
			m.SetDecryptors(res)
		}
		return nil
	}
	res["dh_ephemeral_public_key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePublicKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDhEphemeralPublicKey(val.(PublicKeyable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EncryptedMessageCurveDecryptorSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetDecryptors() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecryptors()))
		for i, v := range m.GetDecryptors() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("decryptors", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("dh_ephemeral_public_key", m.GetDhEphemeralPublicKey())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptedMessageCurveDecryptorSet) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDecryptors sets the decryptors property value. The decryptors property
func (m *EncryptedMessageCurveDecryptorSet) SetDecryptors(value []EncryptedMessageDecryptorable) {
	m.decryptors = value
}

// SetDhEphemeralPublicKey sets the dh_ephemeral_public_key property value. The dh_ephemeral_public_key property
func (m *EncryptedMessageCurveDecryptorSet) SetDhEphemeralPublicKey(value PublicKeyable) {
	m.dh_ephemeral_public_key = value
}

type EncryptedMessageCurveDecryptorSetable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDecryptors() []EncryptedMessageDecryptorable
	GetDhEphemeralPublicKey() PublicKeyable
	SetDecryptors(value []EncryptedMessageDecryptorable)
	SetDhEphemeralPublicKey(value PublicKeyable)
}
