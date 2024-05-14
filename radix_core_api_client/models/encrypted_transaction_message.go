package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptedTransactionMessage a `PlaintextTransactionMessage` encrypted with "Multi-party ECIES" for a number of decryptors (public keys).
type EncryptedTransactionMessage struct {
	TransactionMessage
	// The curve_decryptor_sets property
	curve_decryptor_sets []EncryptedMessageCurveDecryptorSetable
	// The hex-encoded (128-bit) AES-GCM encrypted bytes of an SBOR-encoded `PlaintextTransactionMessage`.The bytes are serialized as the concatenation `Nonce/IV (12 bytes) || Cipher (variable length) || Tag/MAC (16 bytes)`:
	encrypted_hex *string
}

// NewEncryptedTransactionMessage instantiates a new EncryptedTransactionMessage and sets the default values.
func NewEncryptedTransactionMessage() *EncryptedTransactionMessage {
	m := &EncryptedTransactionMessage{
		TransactionMessage: *NewTransactionMessage(),
	}
	return m
}

// CreateEncryptedTransactionMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEncryptedTransactionMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEncryptedTransactionMessage(), nil
}

// GetCurveDecryptorSets gets the curve_decryptor_sets property value. The curve_decryptor_sets property
// returns a []EncryptedMessageCurveDecryptorSetable when successful
func (m *EncryptedTransactionMessage) GetCurveDecryptorSets() []EncryptedMessageCurveDecryptorSetable {
	return m.curve_decryptor_sets
}

// GetEncryptedHex gets the encrypted_hex property value. The hex-encoded (128-bit) AES-GCM encrypted bytes of an SBOR-encoded `PlaintextTransactionMessage`.The bytes are serialized as the concatenation `Nonce/IV (12 bytes) || Cipher (variable length) || Tag/MAC (16 bytes)`:
// returns a *string when successful
func (m *EncryptedTransactionMessage) GetEncryptedHex() *string {
	return m.encrypted_hex
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EncryptedTransactionMessage) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TransactionMessage.GetFieldDeserializers()
	res["curve_decryptor_sets"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEncryptedMessageCurveDecryptorSetFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EncryptedMessageCurveDecryptorSetable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EncryptedMessageCurveDecryptorSetable)
				}
			}
			m.SetCurveDecryptorSets(res)
		}
		return nil
	}
	res["encrypted_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEncryptedHex(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EncryptedTransactionMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TransactionMessage.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetCurveDecryptorSets() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCurveDecryptorSets()))
		for i, v := range m.GetCurveDecryptorSets() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("curve_decryptor_sets", cast)
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("encrypted_hex", m.GetEncryptedHex())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetCurveDecryptorSets sets the curve_decryptor_sets property value. The curve_decryptor_sets property
func (m *EncryptedTransactionMessage) SetCurveDecryptorSets(value []EncryptedMessageCurveDecryptorSetable) {
	m.curve_decryptor_sets = value
}

// SetEncryptedHex sets the encrypted_hex property value. The hex-encoded (128-bit) AES-GCM encrypted bytes of an SBOR-encoded `PlaintextTransactionMessage`.The bytes are serialized as the concatenation `Nonce/IV (12 bytes) || Cipher (variable length) || Tag/MAC (16 bytes)`:
func (m *EncryptedTransactionMessage) SetEncryptedHex(value *string) {
	m.encrypted_hex = value
}

type EncryptedTransactionMessageable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TransactionMessageable
	GetCurveDecryptorSets() []EncryptedMessageCurveDecryptorSetable
	GetEncryptedHex() *string
	SetCurveDecryptorSets(value []EncryptedMessageCurveDecryptorSetable)
	SetEncryptedHex(value *string)
}
