package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewV2Response_logs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The level property
    level *string
    // The message property
    message *string
}
// NewTransactionPreviewV2Response_logs instantiates a new TransactionPreviewV2Response_logs and sets the default values.
func NewTransactionPreviewV2Response_logs()(*TransactionPreviewV2Response_logs) {
    m := &TransactionPreviewV2Response_logs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2Response_logsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2Response_logsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewV2Response_logs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2Response_logs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2Response_logs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The level property
// returns a *string when successful
func (m *TransactionPreviewV2Response_logs) GetLevel()(*string) {
    return m.level
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *TransactionPreviewV2Response_logs) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2Response_logs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *TransactionPreviewV2Response_logs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLevel sets the level property value. The level property
func (m *TransactionPreviewV2Response_logs) SetLevel(value *string)() {
    m.level = value
}
// SetMessage sets the message property value. The message property
func (m *TransactionPreviewV2Response_logs) SetMessage(value *string)() {
    m.message = value
}
type TransactionPreviewV2Response_logsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLevel()(*string)
    GetMessage()(*string)
    SetLevel(value *string)()
    SetMessage(value *string)()
}
