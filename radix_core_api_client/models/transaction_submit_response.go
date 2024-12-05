package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionSubmitResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Is true if the transaction is a duplicate of an existing transaction in the mempool.
    duplicate *bool
}
// NewTransactionSubmitResponse instantiates a new TransactionSubmitResponse and sets the default values.
func NewTransactionSubmitResponse()(*TransactionSubmitResponse) {
    m := &TransactionSubmitResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionSubmitResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionSubmitResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionSubmitResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionSubmitResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDuplicate gets the duplicate property value. Is true if the transaction is a duplicate of an existing transaction in the mempool.
// returns a *bool when successful
func (m *TransactionSubmitResponse) GetDuplicate()(*bool) {
    return m.duplicate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionSubmitResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["duplicate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplicate(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TransactionSubmitResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("duplicate", m.GetDuplicate())
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
func (m *TransactionSubmitResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDuplicate sets the duplicate property value. Is true if the transaction is a duplicate of an existing transaction in the mempool.
func (m *TransactionSubmitResponse) SetDuplicate(value *bool)() {
    m.duplicate = value
}
type TransactionSubmitResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDuplicate()(*bool)
    SetDuplicate(value *bool)()
}
