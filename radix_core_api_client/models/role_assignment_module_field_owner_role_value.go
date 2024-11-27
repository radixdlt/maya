package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoleAssignmentModuleFieldOwnerRoleValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The owner_role property
    owner_role OwnerRoleable
}
// NewRoleAssignmentModuleFieldOwnerRoleValue instantiates a new RoleAssignmentModuleFieldOwnerRoleValue and sets the default values.
func NewRoleAssignmentModuleFieldOwnerRoleValue()(*RoleAssignmentModuleFieldOwnerRoleValue) {
    m := &RoleAssignmentModuleFieldOwnerRoleValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRoleAssignmentModuleFieldOwnerRoleValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleAssignmentModuleFieldOwnerRoleValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignmentModuleFieldOwnerRoleValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoleAssignmentModuleFieldOwnerRoleValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoleAssignmentModuleFieldOwnerRoleValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["owner_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOwnerRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerRole(val.(OwnerRoleable))
        }
        return nil
    }
    return res
}
// GetOwnerRole gets the owner_role property value. The owner_role property
// returns a OwnerRoleable when successful
func (m *RoleAssignmentModuleFieldOwnerRoleValue) GetOwnerRole()(OwnerRoleable) {
    return m.owner_role
}
// Serialize serializes information the current object
func (m *RoleAssignmentModuleFieldOwnerRoleValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("owner_role", m.GetOwnerRole())
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
func (m *RoleAssignmentModuleFieldOwnerRoleValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *RoleAssignmentModuleFieldOwnerRoleValue) SetOwnerRole(value OwnerRoleable)() {
    m.owner_role = value
}
type RoleAssignmentModuleFieldOwnerRoleValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOwnerRole()(OwnerRoleable)
    SetOwnerRole(value OwnerRoleable)()
}
