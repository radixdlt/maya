package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleVaultFieldBalanceSubstate struct {
    Substate
    // The value property
    value NonFungibleVaultFieldBalanceValueable
}
// NewNonFungibleVaultFieldBalanceSubstate instantiates a new NonFungibleVaultFieldBalanceSubstate and sets the default values.
func NewNonFungibleVaultFieldBalanceSubstate()(*NonFungibleVaultFieldBalanceSubstate) {
    m := &NonFungibleVaultFieldBalanceSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateNonFungibleVaultFieldBalanceSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleVaultFieldBalanceSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleVaultFieldBalanceSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleVaultFieldBalanceSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNonFungibleVaultFieldBalanceValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(NonFungibleVaultFieldBalanceValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a NonFungibleVaultFieldBalanceValueable when successful
func (m *NonFungibleVaultFieldBalanceSubstate) GetValue()(NonFungibleVaultFieldBalanceValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *NonFungibleVaultFieldBalanceSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *NonFungibleVaultFieldBalanceSubstate) SetValue(value NonFungibleVaultFieldBalanceValueable)() {
    m.value = value
}
type NonFungibleVaultFieldBalanceSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(NonFungibleVaultFieldBalanceValueable)
    SetValue(value NonFungibleVaultFieldBalanceValueable)()
}
