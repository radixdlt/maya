package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountVaultEntrySubstate struct {
    Substate
    // The key property
    key ResourceKeyable
    // The value property
    value AccountVaultEntryValueable
}
// NewAccountVaultEntrySubstate instantiates a new AccountVaultEntrySubstate and sets the default values.
func NewAccountVaultEntrySubstate()(*AccountVaultEntrySubstate) {
    m := &AccountVaultEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateAccountVaultEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountVaultEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccountVaultEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountVaultEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(ResourceKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountVaultEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(AccountVaultEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a ResourceKeyable when successful
func (m *AccountVaultEntrySubstate) GetKey()(ResourceKeyable) {
    return m.key
}
// GetValue gets the value property value. The value property
// returns a AccountVaultEntryValueable when successful
func (m *AccountVaultEntrySubstate) GetValue()(AccountVaultEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AccountVaultEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AccountVaultEntrySubstate) SetKey(value ResourceKeyable)() {
    m.key = value
}
// SetValue sets the value property value. The value property
func (m *AccountVaultEntrySubstate) SetValue(value AccountVaultEntryValueable)() {
    m.value = value
}
type AccountVaultEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(ResourceKeyable)
    GetValue()(AccountVaultEntryValueable)
    SetKey(value ResourceKeyable)()
    SetValue(value AccountVaultEntryValueable)()
}
