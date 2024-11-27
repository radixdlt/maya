package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectRoleKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The object_module_id property
    object_module_id *ModuleId
    // The role_key property
    role_key *string
}
// NewObjectRoleKey instantiates a new ObjectRoleKey and sets the default values.
func NewObjectRoleKey()(*ObjectRoleKey) {
    m := &ObjectRoleKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateObjectRoleKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectRoleKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectRoleKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ObjectRoleKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectRoleKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["object_module_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseModuleId)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectModuleId(val.(*ModuleId))
        }
        return nil
    }
    res["role_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleKey(val)
        }
        return nil
    }
    return res
}
// GetObjectModuleId gets the object_module_id property value. The object_module_id property
// returns a *ModuleId when successful
func (m *ObjectRoleKey) GetObjectModuleId()(*ModuleId) {
    return m.object_module_id
}
// GetRoleKey gets the role_key property value. The role_key property
// returns a *string when successful
func (m *ObjectRoleKey) GetRoleKey()(*string) {
    return m.role_key
}
// Serialize serializes information the current object
func (m *ObjectRoleKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetObjectModuleId() != nil {
        cast := (*m.GetObjectModuleId()).String()
        err := writer.WriteStringValue("object_module_id", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("role_key", m.GetRoleKey())
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
func (m *ObjectRoleKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetObjectModuleId sets the object_module_id property value. The object_module_id property
func (m *ObjectRoleKey) SetObjectModuleId(value *ModuleId)() {
    m.object_module_id = value
}
// SetRoleKey sets the role_key property value. The role_key property
func (m *ObjectRoleKey) SetRoleKey(value *string)() {
    m.role_key = value
}
type ObjectRoleKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetObjectModuleId()(*ModuleId)
    GetRoleKey()(*string)
    SetObjectModuleId(value *ModuleId)()
    SetRoleKey(value *string)()
}
