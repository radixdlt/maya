package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EncryptedMessageDecryptor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded wrapped key bytes from applying RFC 3394 (256-bit) AES-KeyWrap to the message ephemeral public key, with the secret KEK provided by static Diffie-Helman between the decryptor public key, and the `dh_ephemeral_public_key` for that curve type.The bytes are serialized (according to RFC 3394) as the concatenation `IV (first 8 bytes) || Cipher`.The message ephemeral public key is 128-bit in V1 transactions, but 256-bit in V2 transactions.The Cipher is the wrapped ephemeral public key, and is encoded as two 64-bit blocks in V1, and four 64-bit blocks in V2.
    aes_wrapped_key_hex *string
    // The last 8 bytes of the Blake2b-256 hash of the public key bytes, in their standard Radix byte-serialization.
    public_key_fingerprint_hex *string
}
// NewEncryptedMessageDecryptor instantiates a new EncryptedMessageDecryptor and sets the default values.
func NewEncryptedMessageDecryptor()(*EncryptedMessageDecryptor) {
    m := &EncryptedMessageDecryptor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEncryptedMessageDecryptorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEncryptedMessageDecryptorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptedMessageDecryptor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EncryptedMessageDecryptor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAesWrappedKeyHex gets the aes_wrapped_key_hex property value. The hex-encoded wrapped key bytes from applying RFC 3394 (256-bit) AES-KeyWrap to the message ephemeral public key, with the secret KEK provided by static Diffie-Helman between the decryptor public key, and the `dh_ephemeral_public_key` for that curve type.The bytes are serialized (according to RFC 3394) as the concatenation `IV (first 8 bytes) || Cipher`.The message ephemeral public key is 128-bit in V1 transactions, but 256-bit in V2 transactions.The Cipher is the wrapped ephemeral public key, and is encoded as two 64-bit blocks in V1, and four 64-bit blocks in V2.
// returns a *string when successful
func (m *EncryptedMessageDecryptor) GetAesWrappedKeyHex()(*string) {
    return m.aes_wrapped_key_hex
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EncryptedMessageDecryptor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aes_wrapped_key_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAesWrappedKeyHex(val)
        }
        return nil
    }
    res["public_key_fingerprint_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKeyFingerprintHex(val)
        }
        return nil
    }
    return res
}
// GetPublicKeyFingerprintHex gets the public_key_fingerprint_hex property value. The last 8 bytes of the Blake2b-256 hash of the public key bytes, in their standard Radix byte-serialization.
// returns a *string when successful
func (m *EncryptedMessageDecryptor) GetPublicKeyFingerprintHex()(*string) {
    return m.public_key_fingerprint_hex
}
// Serialize serializes information the current object
func (m *EncryptedMessageDecryptor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aes_wrapped_key_hex", m.GetAesWrappedKeyHex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_key_fingerprint_hex", m.GetPublicKeyFingerprintHex())
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
func (m *EncryptedMessageDecryptor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAesWrappedKeyHex sets the aes_wrapped_key_hex property value. The hex-encoded wrapped key bytes from applying RFC 3394 (256-bit) AES-KeyWrap to the message ephemeral public key, with the secret KEK provided by static Diffie-Helman between the decryptor public key, and the `dh_ephemeral_public_key` for that curve type.The bytes are serialized (according to RFC 3394) as the concatenation `IV (first 8 bytes) || Cipher`.The message ephemeral public key is 128-bit in V1 transactions, but 256-bit in V2 transactions.The Cipher is the wrapped ephemeral public key, and is encoded as two 64-bit blocks in V1, and four 64-bit blocks in V2.
func (m *EncryptedMessageDecryptor) SetAesWrappedKeyHex(value *string)() {
    m.aes_wrapped_key_hex = value
}
// SetPublicKeyFingerprintHex sets the public_key_fingerprint_hex property value. The last 8 bytes of the Blake2b-256 hash of the public key bytes, in their standard Radix byte-serialization.
func (m *EncryptedMessageDecryptor) SetPublicKeyFingerprintHex(value *string)() {
    m.public_key_fingerprint_hex = value
}
type EncryptedMessageDecryptorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAesWrappedKeyHex()(*string)
    GetPublicKeyFingerprintHex()(*string)
    SetAesWrappedKeyHex(value *string)()
    SetPublicKeyFingerprintHex(value *string)()
}
