package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionParseResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parsed property
    parsed ParsedTransactionable
}
// NewTransactionParseResponse instantiates a new TransactionParseResponse and sets the default values.
func NewTransactionParseResponse()(*TransactionParseResponse) {
    m := &TransactionParseResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionParseResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionParseResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionParseResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionParseResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionParseResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParsedTransactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParsed(val.(ParsedTransactionable))
        }
        return nil
    }
    return res
}
// GetParsed gets the parsed property value. The parsed property
// returns a ParsedTransactionable when successful
func (m *TransactionParseResponse) GetParsed()(ParsedTransactionable) {
    return m.parsed
}
// Serialize serializes information the current object
func (m *TransactionParseResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parsed", m.GetParsed())
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
func (m *TransactionParseResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParsed sets the parsed property value. The parsed property
func (m *TransactionParseResponse) SetParsed(value ParsedTransactionable)() {
    m.parsed = value
}
type TransactionParseResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParsed()(ParsedTransactionable)
    SetParsed(value ParsedTransactionable)()
}
