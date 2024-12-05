package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NonFungibleResourceManagerDataEntrySubstate if the NF has been burned, the value is deleted and empty.
type NonFungibleResourceManagerDataEntrySubstate struct {
    Substate
    // The key property
    key LocalNonFungibleKeyable
    // If missing, it represents a burned Non-Fungible.A Non-Fungible with that local id cannot be minted again - the id is not re-usable.
    value NonFungibleResourceManagerDataEntryValueable
}
// NewNonFungibleResourceManagerDataEntrySubstate instantiates a new NonFungibleResourceManagerDataEntrySubstate and sets the default values.
func NewNonFungibleResourceManagerDataEntrySubstate()(*NonFungibleResourceManagerDataEntrySubstate) {
    m := &NonFungibleResourceManagerDataEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateNonFungibleResourceManagerDataEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerDataEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleResourceManagerDataEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerDataEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocalNonFungibleKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(LocalNonFungibleKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNonFungibleResourceManagerDataEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(NonFungibleResourceManagerDataEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a LocalNonFungibleKeyable when successful
func (m *NonFungibleResourceManagerDataEntrySubstate) GetKey()(LocalNonFungibleKeyable) {
    return m.key
}
// GetValue gets the value property value. If missing, it represents a burned Non-Fungible.A Non-Fungible with that local id cannot be minted again - the id is not re-usable.
// returns a NonFungibleResourceManagerDataEntryValueable when successful
func (m *NonFungibleResourceManagerDataEntrySubstate) GetValue()(NonFungibleResourceManagerDataEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *NonFungibleResourceManagerDataEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *NonFungibleResourceManagerDataEntrySubstate) SetKey(value LocalNonFungibleKeyable)() {
    m.key = value
}
// SetValue sets the value property value. If missing, it represents a burned Non-Fungible.A Non-Fungible with that local id cannot be minted again - the id is not re-usable.
func (m *NonFungibleResourceManagerDataEntrySubstate) SetValue(value NonFungibleResourceManagerDataEntryValueable)() {
    m.value = value
}
type NonFungibleResourceManagerDataEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(LocalNonFungibleKeyable)
    GetValue()(NonFungibleResourceManagerDataEntryValueable)
    SetKey(value LocalNonFungibleKeyable)()
    SetValue(value NonFungibleResourceManagerDataEntryValueable)()
}
