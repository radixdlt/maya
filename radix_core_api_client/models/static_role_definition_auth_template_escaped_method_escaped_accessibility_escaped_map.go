package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StaticRoleDefinitionAuthTemplate_method_accessibility_map a map from a method identifier to MethodAccessibility
type StaticRoleDefinitionAuthTemplate_method_accessibility_map struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewStaticRoleDefinitionAuthTemplate_method_accessibility_map instantiates a new StaticRoleDefinitionAuthTemplate_method_accessibility_map and sets the default values.
func NewStaticRoleDefinitionAuthTemplate_method_accessibility_map()(*StaticRoleDefinitionAuthTemplate_method_accessibility_map) {
    m := &StaticRoleDefinitionAuthTemplate_method_accessibility_map{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStaticRoleDefinitionAuthTemplate_method_accessibility_mapFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStaticRoleDefinitionAuthTemplate_method_accessibility_mapFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStaticRoleDefinitionAuthTemplate_method_accessibility_map(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StaticRoleDefinitionAuthTemplate_method_accessibility_map) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StaticRoleDefinitionAuthTemplate_method_accessibility_map) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *StaticRoleDefinitionAuthTemplate_method_accessibility_map) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StaticRoleDefinitionAuthTemplate_method_accessibility_map) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type StaticRoleDefinitionAuthTemplate_method_accessibility_mapable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
