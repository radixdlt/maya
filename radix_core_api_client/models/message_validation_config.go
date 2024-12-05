package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MessageValidationConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The max_decryptors property
    max_decryptors *string
    // The max_encrypted_message_length property
    max_encrypted_message_length *string
    // The max_mime_type_length property
    max_mime_type_length *string
    // The max_plaintext_message_length property
    max_plaintext_message_length *string
}
// NewMessageValidationConfig instantiates a new MessageValidationConfig and sets the default values.
func NewMessageValidationConfig()(*MessageValidationConfig) {
    m := &MessageValidationConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessageValidationConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMessageValidationConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageValidationConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MessageValidationConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MessageValidationConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_decryptors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxDecryptors(val)
        }
        return nil
    }
    res["max_encrypted_message_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxEncryptedMessageLength(val)
        }
        return nil
    }
    res["max_mime_type_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxMimeTypeLength(val)
        }
        return nil
    }
    res["max_plaintext_message_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPlaintextMessageLength(val)
        }
        return nil
    }
    return res
}
// GetMaxDecryptors gets the max_decryptors property value. The max_decryptors property
// returns a *string when successful
func (m *MessageValidationConfig) GetMaxDecryptors()(*string) {
    return m.max_decryptors
}
// GetMaxEncryptedMessageLength gets the max_encrypted_message_length property value. The max_encrypted_message_length property
// returns a *string when successful
func (m *MessageValidationConfig) GetMaxEncryptedMessageLength()(*string) {
    return m.max_encrypted_message_length
}
// GetMaxMimeTypeLength gets the max_mime_type_length property value. The max_mime_type_length property
// returns a *string when successful
func (m *MessageValidationConfig) GetMaxMimeTypeLength()(*string) {
    return m.max_mime_type_length
}
// GetMaxPlaintextMessageLength gets the max_plaintext_message_length property value. The max_plaintext_message_length property
// returns a *string when successful
func (m *MessageValidationConfig) GetMaxPlaintextMessageLength()(*string) {
    return m.max_plaintext_message_length
}
// Serialize serializes information the current object
func (m *MessageValidationConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("max_decryptors", m.GetMaxDecryptors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_encrypted_message_length", m.GetMaxEncryptedMessageLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_mime_type_length", m.GetMaxMimeTypeLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_plaintext_message_length", m.GetMaxPlaintextMessageLength())
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
func (m *MessageValidationConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxDecryptors sets the max_decryptors property value. The max_decryptors property
func (m *MessageValidationConfig) SetMaxDecryptors(value *string)() {
    m.max_decryptors = value
}
// SetMaxEncryptedMessageLength sets the max_encrypted_message_length property value. The max_encrypted_message_length property
func (m *MessageValidationConfig) SetMaxEncryptedMessageLength(value *string)() {
    m.max_encrypted_message_length = value
}
// SetMaxMimeTypeLength sets the max_mime_type_length property value. The max_mime_type_length property
func (m *MessageValidationConfig) SetMaxMimeTypeLength(value *string)() {
    m.max_mime_type_length = value
}
// SetMaxPlaintextMessageLength sets the max_plaintext_message_length property value. The max_plaintext_message_length property
func (m *MessageValidationConfig) SetMaxPlaintextMessageLength(value *string)() {
    m.max_plaintext_message_length = value
}
type MessageValidationConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxDecryptors()(*string)
    GetMaxEncryptedMessageLength()(*string)
    GetMaxMimeTypeLength()(*string)
    GetMaxPlaintextMessageLength()(*string)
    SetMaxDecryptors(value *string)()
    SetMaxEncryptedMessageLength(value *string)()
    SetMaxMimeTypeLength(value *string)()
    SetMaxPlaintextMessageLength(value *string)()
}
