package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StaticRoleDefinitionAuthTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A map from a method identifier to MethodAccessibility
    method_accessibility_map StaticRoleDefinitionAuthTemplate_method_accessibility_mapable
    // A map from role name to role details
    roles StaticRoleDefinitionAuthTemplate_rolesable
    // The role_specification property
    role_specification *RoleSpecification
}
// NewStaticRoleDefinitionAuthTemplate instantiates a new StaticRoleDefinitionAuthTemplate and sets the default values.
func NewStaticRoleDefinitionAuthTemplate()(*StaticRoleDefinitionAuthTemplate) {
    m := &StaticRoleDefinitionAuthTemplate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStaticRoleDefinitionAuthTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStaticRoleDefinitionAuthTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStaticRoleDefinitionAuthTemplate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StaticRoleDefinitionAuthTemplate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StaticRoleDefinitionAuthTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["method_accessibility_map"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStaticRoleDefinitionAuthTemplate_method_accessibility_mapFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodAccessibilityMap(val.(StaticRoleDefinitionAuthTemplate_method_accessibility_mapable))
        }
        return nil
    }
    res["role_specification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleSpecification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleSpecification(val.(*RoleSpecification))
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStaticRoleDefinitionAuthTemplate_rolesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoles(val.(StaticRoleDefinitionAuthTemplate_rolesable))
        }
        return nil
    }
    return res
}
// GetMethodAccessibilityMap gets the method_accessibility_map property value. A map from a method identifier to MethodAccessibility
// returns a StaticRoleDefinitionAuthTemplate_method_accessibility_mapable when successful
func (m *StaticRoleDefinitionAuthTemplate) GetMethodAccessibilityMap()(StaticRoleDefinitionAuthTemplate_method_accessibility_mapable) {
    return m.method_accessibility_map
}
// GetRoles gets the roles property value. A map from role name to role details
// returns a StaticRoleDefinitionAuthTemplate_rolesable when successful
func (m *StaticRoleDefinitionAuthTemplate) GetRoles()(StaticRoleDefinitionAuthTemplate_rolesable) {
    return m.roles
}
// GetRoleSpecification gets the role_specification property value. The role_specification property
// returns a *RoleSpecification when successful
func (m *StaticRoleDefinitionAuthTemplate) GetRoleSpecification()(*RoleSpecification) {
    return m.role_specification
}
// Serialize serializes information the current object
func (m *StaticRoleDefinitionAuthTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("method_accessibility_map", m.GetMethodAccessibilityMap())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    if m.GetRoleSpecification() != nil {
        cast := (*m.GetRoleSpecification()).String()
        err := writer.WriteStringValue("role_specification", &cast)
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
func (m *StaticRoleDefinitionAuthTemplate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMethodAccessibilityMap sets the method_accessibility_map property value. A map from a method identifier to MethodAccessibility
func (m *StaticRoleDefinitionAuthTemplate) SetMethodAccessibilityMap(value StaticRoleDefinitionAuthTemplate_method_accessibility_mapable)() {
    m.method_accessibility_map = value
}
// SetRoles sets the roles property value. A map from role name to role details
func (m *StaticRoleDefinitionAuthTemplate) SetRoles(value StaticRoleDefinitionAuthTemplate_rolesable)() {
    m.roles = value
}
// SetRoleSpecification sets the role_specification property value. The role_specification property
func (m *StaticRoleDefinitionAuthTemplate) SetRoleSpecification(value *RoleSpecification)() {
    m.role_specification = value
}
type StaticRoleDefinitionAuthTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMethodAccessibilityMap()(StaticRoleDefinitionAuthTemplate_method_accessibility_mapable)
    GetRoles()(StaticRoleDefinitionAuthTemplate_rolesable)
    GetRoleSpecification()(*RoleSpecification)
    SetMethodAccessibilityMap(value StaticRoleDefinitionAuthTemplate_method_accessibility_mapable)()
    SetRoles(value StaticRoleDefinitionAuthTemplate_rolesable)()
    SetRoleSpecification(value *RoleSpecification)()
}
