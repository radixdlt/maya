package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FlashLedgerTransaction struct {
    LedgerTransaction
    // Direct state updates performed by a Flash Transaction.
    flashed_state_updates FlashedStateUpdatesable
    // Human-readable identifier of the flash transaction.
    name *string
}
// NewFlashLedgerTransaction instantiates a new FlashLedgerTransaction and sets the default values.
func NewFlashLedgerTransaction()(*FlashLedgerTransaction) {
    m := &FlashLedgerTransaction{
        LedgerTransaction: *NewLedgerTransaction(),
    }
    return m
}
// CreateFlashLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFlashLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFlashLedgerTransaction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FlashLedgerTransaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LedgerTransaction.GetFieldDeserializers()
    res["flashed_state_updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFlashedStateUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlashedStateUpdates(val.(FlashedStateUpdatesable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetFlashedStateUpdates gets the flashed_state_updates property value. Direct state updates performed by a Flash Transaction.
// returns a FlashedStateUpdatesable when successful
func (m *FlashLedgerTransaction) GetFlashedStateUpdates()(FlashedStateUpdatesable) {
    return m.flashed_state_updates
}
// GetName gets the name property value. Human-readable identifier of the flash transaction.
// returns a *string when successful
func (m *FlashLedgerTransaction) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *FlashLedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LedgerTransaction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("flashed_state_updates", m.GetFlashedStateUpdates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFlashedStateUpdates sets the flashed_state_updates property value. Direct state updates performed by a Flash Transaction.
func (m *FlashLedgerTransaction) SetFlashedStateUpdates(value FlashedStateUpdatesable)() {
    m.flashed_state_updates = value
}
// SetName sets the name property value. Human-readable identifier of the flash transaction.
func (m *FlashLedgerTransaction) SetName(value *string)() {
    m.name = value
}
type FlashLedgerTransactionable interface {
    LedgerTransactionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFlashedStateUpdates()(FlashedStateUpdatesable)
    GetName()(*string)
    SetFlashedStateUpdates(value FlashedStateUpdatesable)()
    SetName(value *string)()
}
