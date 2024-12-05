package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionTrackerCollectionEntrySubstate struct {
    Substate
    // The key property
    key TransactionIdKeyable
    // The value property
    value TransactionTrackerCollectionEntryValueable
}
// NewTransactionTrackerCollectionEntrySubstate instantiates a new TransactionTrackerCollectionEntrySubstate and sets the default values.
func NewTransactionTrackerCollectionEntrySubstate()(*TransactionTrackerCollectionEntrySubstate) {
    m := &TransactionTrackerCollectionEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateTransactionTrackerCollectionEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionTrackerCollectionEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionTrackerCollectionEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionTrackerCollectionEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionIdKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(TransactionIdKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionTrackerCollectionEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(TransactionTrackerCollectionEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a TransactionIdKeyable when successful
func (m *TransactionTrackerCollectionEntrySubstate) GetKey()(TransactionIdKeyable) {
    return m.key
}
// GetValue gets the value property value. The value property
// returns a TransactionTrackerCollectionEntryValueable when successful
func (m *TransactionTrackerCollectionEntrySubstate) GetValue()(TransactionTrackerCollectionEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *TransactionTrackerCollectionEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. The key property
func (m *TransactionTrackerCollectionEntrySubstate) SetKey(value TransactionIdKeyable)() {
    m.key = value
}
// SetValue sets the value property value. The value property
func (m *TransactionTrackerCollectionEntrySubstate) SetValue(value TransactionTrackerCollectionEntryValueable)() {
    m.value = value
}
type TransactionTrackerCollectionEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(TransactionIdKeyable)
    GetValue()(TransactionTrackerCollectionEntryValueable)
    SetKey(value TransactionIdKeyable)()
    SetValue(value TransactionTrackerCollectionEntryValueable)()
}
