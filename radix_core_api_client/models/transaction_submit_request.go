package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionSubmitRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If true, the transaction validity is freshly recalculated without using any caches (defaults false)
    force_recalculate *bool
    // The logical name of the network
    network *string
    // A hex-encoded, compiled notarized transaction.
    notarized_transaction_hex *string
}
// NewTransactionSubmitRequest instantiates a new TransactionSubmitRequest and sets the default values.
func NewTransactionSubmitRequest()(*TransactionSubmitRequest) {
    m := &TransactionSubmitRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionSubmitRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionSubmitRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionSubmitRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionSubmitRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionSubmitRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["force_recalculate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceRecalculate(val)
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["notarized_transaction_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotarizedTransactionHex(val)
        }
        return nil
    }
    return res
}
// GetForceRecalculate gets the force_recalculate property value. If true, the transaction validity is freshly recalculated without using any caches (defaults false)
// returns a *bool when successful
func (m *TransactionSubmitRequest) GetForceRecalculate()(*bool) {
    return m.force_recalculate
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionSubmitRequest) GetNetwork()(*string) {
    return m.network
}
// GetNotarizedTransactionHex gets the notarized_transaction_hex property value. A hex-encoded, compiled notarized transaction.
// returns a *string when successful
func (m *TransactionSubmitRequest) GetNotarizedTransactionHex()(*string) {
    return m.notarized_transaction_hex
}
// Serialize serializes information the current object
func (m *TransactionSubmitRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("force_recalculate", m.GetForceRecalculate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notarized_transaction_hex", m.GetNotarizedTransactionHex())
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
func (m *TransactionSubmitRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetForceRecalculate sets the force_recalculate property value. If true, the transaction validity is freshly recalculated without using any caches (defaults false)
func (m *TransactionSubmitRequest) SetForceRecalculate(value *bool)() {
    m.force_recalculate = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionSubmitRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetNotarizedTransactionHex sets the notarized_transaction_hex property value. A hex-encoded, compiled notarized transaction.
func (m *TransactionSubmitRequest) SetNotarizedTransactionHex(value *string)() {
    m.notarized_transaction_hex = value
}
type TransactionSubmitRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForceRecalculate()(*bool)
    GetNetwork()(*string)
    GetNotarizedTransactionHex()(*string)
    SetForceRecalculate(value *bool)()
    SetNetwork(value *string)()
    SetNotarizedTransactionHex(value *string)()
}
