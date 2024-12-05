package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BinaryPlaintextMessageContent struct {
    PlaintextMessageContent
    // The hex-encoded value of a message that the author decided to provide as raw bytes.
    value_hex *string
}
// NewBinaryPlaintextMessageContent instantiates a new BinaryPlaintextMessageContent and sets the default values.
func NewBinaryPlaintextMessageContent()(*BinaryPlaintextMessageContent) {
    m := &BinaryPlaintextMessageContent{
        PlaintextMessageContent: *NewPlaintextMessageContent(),
    }
    return m
}
// CreateBinaryPlaintextMessageContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBinaryPlaintextMessageContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBinaryPlaintextMessageContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BinaryPlaintextMessageContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlaintextMessageContent.GetFieldDeserializers()
    res["value_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueHex(val)
        }
        return nil
    }
    return res
}
// GetValueHex gets the value_hex property value. The hex-encoded value of a message that the author decided to provide as raw bytes.
// returns a *string when successful
func (m *BinaryPlaintextMessageContent) GetValueHex()(*string) {
    return m.value_hex
}
// Serialize serializes information the current object
func (m *BinaryPlaintextMessageContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlaintextMessageContent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value_hex", m.GetValueHex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValueHex sets the value_hex property value. The hex-encoded value of a message that the author decided to provide as raw bytes.
func (m *BinaryPlaintextMessageContent) SetValueHex(value *string)() {
    m.value_hex = value
}
type BinaryPlaintextMessageContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlaintextMessageContentable
    GetValueHex()(*string)
    SetValueHex(value *string)()
}
