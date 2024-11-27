package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionPreviewResponseOptions a set of flags to configure the response of the transaction preview.
type TransactionPreviewResponseOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
    radix_engine_toolkit_receipt *bool
}
// NewTransactionPreviewResponseOptions instantiates a new TransactionPreviewResponseOptions and sets the default values.
func NewTransactionPreviewResponseOptions()(*TransactionPreviewResponseOptions) {
    m := &TransactionPreviewResponseOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewResponseOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewResponseOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewResponseOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewResponseOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewResponseOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["radix_engine_toolkit_receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRadixEngineToolkitReceipt(val)
        }
        return nil
    }
    return res
}
// GetRadixEngineToolkitReceipt gets the radix_engine_toolkit_receipt property value. This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
// returns a *bool when successful
func (m *TransactionPreviewResponseOptions) GetRadixEngineToolkitReceipt()(*bool) {
    return m.radix_engine_toolkit_receipt
}
// Serialize serializes information the current object
func (m *TransactionPreviewResponseOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("radix_engine_toolkit_receipt", m.GetRadixEngineToolkitReceipt())
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
func (m *TransactionPreviewResponseOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRadixEngineToolkitReceipt sets the radix_engine_toolkit_receipt property value. This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
func (m *TransactionPreviewResponseOptions) SetRadixEngineToolkitReceipt(value *bool)() {
    m.radix_engine_toolkit_receipt = value
}
type TransactionPreviewResponseOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRadixEngineToolkitReceipt()(*bool)
    SetRadixEngineToolkitReceipt(value *bool)()
}
