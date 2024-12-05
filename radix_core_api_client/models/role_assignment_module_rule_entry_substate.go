package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoleAssignmentModuleRuleEntrySubstate struct {
    Substate
    // The key property
    key ObjectRoleKeyable
    // The value property
    value RoleAssignmentModuleRuleEntryValueable
}
// NewRoleAssignmentModuleRuleEntrySubstate instantiates a new RoleAssignmentModuleRuleEntrySubstate and sets the default values.
func NewRoleAssignmentModuleRuleEntrySubstate()(*RoleAssignmentModuleRuleEntrySubstate) {
    m := &RoleAssignmentModuleRuleEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateRoleAssignmentModuleRuleEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleAssignmentModuleRuleEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignmentModuleRuleEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoleAssignmentModuleRuleEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectRoleKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(ObjectRoleKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentModuleRuleEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(RoleAssignmentModuleRuleEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a ObjectRoleKeyable when successful
func (m *RoleAssignmentModuleRuleEntrySubstate) GetKey()(ObjectRoleKeyable) {
    return m.key
}
// GetValue gets the value property value. The value property
// returns a RoleAssignmentModuleRuleEntryValueable when successful
func (m *RoleAssignmentModuleRuleEntrySubstate) GetValue()(RoleAssignmentModuleRuleEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *RoleAssignmentModuleRuleEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RoleAssignmentModuleRuleEntrySubstate) SetKey(value ObjectRoleKeyable)() {
    m.key = value
}
// SetValue sets the value property value. The value property
func (m *RoleAssignmentModuleRuleEntrySubstate) SetValue(value RoleAssignmentModuleRuleEntryValueable)() {
    m.value = value
}
type RoleAssignmentModuleRuleEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(ObjectRoleKeyable)
    GetValue()(RoleAssignmentModuleRuleEntryValueable)
    SetKey(value ObjectRoleKeyable)()
    SetValue(value RoleAssignmentModuleRuleEntryValueable)()
}
