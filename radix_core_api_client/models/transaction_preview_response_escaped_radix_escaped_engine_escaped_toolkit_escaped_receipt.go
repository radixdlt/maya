package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionPreviewResponse_radix_engine_toolkit_receipt an optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true when requesting a transaction preview from the API.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
type TransactionPreviewResponse_radix_engine_toolkit_receipt struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewTransactionPreviewResponse_radix_engine_toolkit_receipt instantiates a new TransactionPreviewResponse_radix_engine_toolkit_receipt and sets the default values.
func NewTransactionPreviewResponse_radix_engine_toolkit_receipt()(*TransactionPreviewResponse_radix_engine_toolkit_receipt) {
    m := &TransactionPreviewResponse_radix_engine_toolkit_receipt{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewResponse_radix_engine_toolkit_receiptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewResponse_radix_engine_toolkit_receiptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewResponse_radix_engine_toolkit_receipt(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewResponse_radix_engine_toolkit_receipt) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewResponse_radix_engine_toolkit_receipt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *TransactionPreviewResponse_radix_engine_toolkit_receipt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransactionPreviewResponse_radix_engine_toolkit_receipt) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type TransactionPreviewResponse_radix_engine_toolkit_receiptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
