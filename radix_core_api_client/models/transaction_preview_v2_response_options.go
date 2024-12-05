package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionPreviewV2ResponseOptions a set of flags to configure the response of the transaction preview.
type TransactionPreviewV2ResponseOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This flag controls whether the preview response will include a Core API receipt or not.If not provided, this defaults to `false` and no core api receipt is provided inthe response.
    core_api_receipt *bool
    // This flag controls whether the preview response will include execution logs.If not provided, this defaults to `false` and no logs will be provided in the response.
    logs *bool
    // This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
    radix_engine_toolkit_receipt *bool
}
// NewTransactionPreviewV2ResponseOptions instantiates a new TransactionPreviewV2ResponseOptions and sets the default values.
func NewTransactionPreviewV2ResponseOptions()(*TransactionPreviewV2ResponseOptions) {
    m := &TransactionPreviewV2ResponseOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2ResponseOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2ResponseOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewV2ResponseOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2ResponseOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCoreApiReceipt gets the core_api_receipt property value. This flag controls whether the preview response will include a Core API receipt or not.If not provided, this defaults to `false` and no core api receipt is provided inthe response.
// returns a *bool when successful
func (m *TransactionPreviewV2ResponseOptions) GetCoreApiReceipt()(*bool) {
    return m.core_api_receipt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2ResponseOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["core_api_receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreApiReceipt(val)
        }
        return nil
    }
    res["logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogs(val)
        }
        return nil
    }
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
// GetLogs gets the logs property value. This flag controls whether the preview response will include execution logs.If not provided, this defaults to `false` and no logs will be provided in the response.
// returns a *bool when successful
func (m *TransactionPreviewV2ResponseOptions) GetLogs()(*bool) {
    return m.logs
}
// GetRadixEngineToolkitReceipt gets the radix_engine_toolkit_receipt property value. This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
// returns a *bool when successful
func (m *TransactionPreviewV2ResponseOptions) GetRadixEngineToolkitReceipt()(*bool) {
    return m.radix_engine_toolkit_receipt
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2ResponseOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("core_api_receipt", m.GetCoreApiReceipt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("logs", m.GetLogs())
        if err != nil {
            return err
        }
    }
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
func (m *TransactionPreviewV2ResponseOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCoreApiReceipt sets the core_api_receipt property value. This flag controls whether the preview response will include a Core API receipt or not.If not provided, this defaults to `false` and no core api receipt is provided inthe response.
func (m *TransactionPreviewV2ResponseOptions) SetCoreApiReceipt(value *bool)() {
    m.core_api_receipt = value
}
// SetLogs sets the logs property value. This flag controls whether the preview response will include execution logs.If not provided, this defaults to `false` and no logs will be provided in the response.
func (m *TransactionPreviewV2ResponseOptions) SetLogs(value *bool)() {
    m.logs = value
}
// SetRadixEngineToolkitReceipt sets the radix_engine_toolkit_receipt property value. This flag controls whether the preview response will include a Radix Engine Toolkit serializablereceipt or not. If not provided, this defaults to `false` and no toolkit receipt is provided inthe response.
func (m *TransactionPreviewV2ResponseOptions) SetRadixEngineToolkitReceipt(value *bool)() {
    m.radix_engine_toolkit_receipt = value
}
type TransactionPreviewV2ResponseOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCoreApiReceipt()(*bool)
    GetLogs()(*bool)
    GetRadixEngineToolkitReceipt()(*bool)
    SetCoreApiReceipt(value *bool)()
    SetLogs(value *bool)()
    SetRadixEngineToolkitReceipt(value *bool)()
}
