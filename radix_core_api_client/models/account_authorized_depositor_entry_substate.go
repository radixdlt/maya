package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountAuthorizedDepositorEntrySubstate struct {
    Substate
    // The key property
    key AuthorizedDepositorKeyable
    // Empty value. The existence of the key implies the depositor is authorized.
    value AccountAuthorizedDepositorEntryValueable
}
// NewAccountAuthorizedDepositorEntrySubstate instantiates a new AccountAuthorizedDepositorEntrySubstate and sets the default values.
func NewAccountAuthorizedDepositorEntrySubstate()(*AccountAuthorizedDepositorEntrySubstate) {
    m := &AccountAuthorizedDepositorEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateAccountAuthorizedDepositorEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountAuthorizedDepositorEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccountAuthorizedDepositorEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountAuthorizedDepositorEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizedDepositorKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(AuthorizedDepositorKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountAuthorizedDepositorEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(AccountAuthorizedDepositorEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a AuthorizedDepositorKeyable when successful
func (m *AccountAuthorizedDepositorEntrySubstate) GetKey()(AuthorizedDepositorKeyable) {
    return m.key
}
// GetValue gets the value property value. Empty value. The existence of the key implies the depositor is authorized.
// returns a AccountAuthorizedDepositorEntryValueable when successful
func (m *AccountAuthorizedDepositorEntrySubstate) GetValue()(AccountAuthorizedDepositorEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AccountAuthorizedDepositorEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AccountAuthorizedDepositorEntrySubstate) SetKey(value AuthorizedDepositorKeyable)() {
    m.key = value
}
// SetValue sets the value property value. Empty value. The existence of the key implies the depositor is authorized.
func (m *AccountAuthorizedDepositorEntrySubstate) SetValue(value AccountAuthorizedDepositorEntryValueable)() {
    m.value = value
}
type AccountAuthorizedDepositorEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(AuthorizedDepositorKeyable)
    GetValue()(AccountAuthorizedDepositorEntryValueable)
    SetKey(value AuthorizedDepositorKeyable)()
    SetValue(value AccountAuthorizedDepositorEntryValueable)()
}
