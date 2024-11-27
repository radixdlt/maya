package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SystemTransaction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded system transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
    payload_hex *string
}
// NewSystemTransaction instantiates a new SystemTransaction and sets the default values.
func NewSystemTransaction()(*SystemTransaction) {
    m := &SystemTransaction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSystemTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSystemTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSystemTransaction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SystemTransaction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SystemTransaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payload_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadHex(val)
        }
        return nil
    }
    return res
}
// GetPayloadHex gets the payload_hex property value. The hex-encoded system transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
// returns a *string when successful
func (m *SystemTransaction) GetPayloadHex()(*string) {
    return m.payload_hex
}
// Serialize serializes information the current object
func (m *SystemTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("payload_hex", m.GetPayloadHex())
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
func (m *SystemTransaction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPayloadHex sets the payload_hex property value. The hex-encoded system transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
func (m *SystemTransaction) SetPayloadHex(value *string)() {
    m.payload_hex = value
}
type SystemTransactionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPayloadHex()(*string)
    SetPayloadHex(value *string)()
}
