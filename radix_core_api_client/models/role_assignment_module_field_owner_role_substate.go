package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoleAssignmentModuleFieldOwnerRoleSubstate struct {
    Substate
    // The value property
    value RoleAssignmentModuleFieldOwnerRoleValueable
}
// NewRoleAssignmentModuleFieldOwnerRoleSubstate instantiates a new RoleAssignmentModuleFieldOwnerRoleSubstate and sets the default values.
func NewRoleAssignmentModuleFieldOwnerRoleSubstate()(*RoleAssignmentModuleFieldOwnerRoleSubstate) {
    m := &RoleAssignmentModuleFieldOwnerRoleSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateRoleAssignmentModuleFieldOwnerRoleSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleAssignmentModuleFieldOwnerRoleSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignmentModuleFieldOwnerRoleSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoleAssignmentModuleFieldOwnerRoleSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentModuleFieldOwnerRoleValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(RoleAssignmentModuleFieldOwnerRoleValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a RoleAssignmentModuleFieldOwnerRoleValueable when successful
func (m *RoleAssignmentModuleFieldOwnerRoleSubstate) GetValue()(RoleAssignmentModuleFieldOwnerRoleValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *RoleAssignmentModuleFieldOwnerRoleSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RoleAssignmentModuleFieldOwnerRoleSubstate) SetValue(value RoleAssignmentModuleFieldOwnerRoleValueable)() {
    m.value = value
}
type RoleAssignmentModuleFieldOwnerRoleSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(RoleAssignmentModuleFieldOwnerRoleValueable)
    SetValue(value RoleAssignmentModuleFieldOwnerRoleValueable)()
}
