package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintVersionKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The first part of the substate key `(blueprint_name, blueprint_version)`.
    blueprint_name *string
    // The second part of the substate key `(blueprint_name, blueprint_version)`.
    blueprint_version *string
}
// NewBlueprintVersionKey instantiates a new BlueprintVersionKey and sets the default values.
func NewBlueprintVersionKey()(*BlueprintVersionKey) {
    m := &BlueprintVersionKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintVersionKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintVersionKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintVersionKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintVersionKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlueprintName gets the blueprint_name property value. The first part of the substate key `(blueprint_name, blueprint_version)`.
// returns a *string when successful
func (m *BlueprintVersionKey) GetBlueprintName()(*string) {
    return m.blueprint_name
}
// GetBlueprintVersion gets the blueprint_version property value. The second part of the substate key `(blueprint_name, blueprint_version)`.
// returns a *string when successful
func (m *BlueprintVersionKey) GetBlueprintVersion()(*string) {
    return m.blueprint_version
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintVersionKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blueprint_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueprintName(val)
        }
        return nil
    }
    res["blueprint_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueprintVersion(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BlueprintVersionKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("blueprint_name", m.GetBlueprintName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blueprint_version", m.GetBlueprintVersion())
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
func (m *BlueprintVersionKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlueprintName sets the blueprint_name property value. The first part of the substate key `(blueprint_name, blueprint_version)`.
func (m *BlueprintVersionKey) SetBlueprintName(value *string)() {
    m.blueprint_name = value
}
// SetBlueprintVersion sets the blueprint_version property value. The second part of the substate key `(blueprint_name, blueprint_version)`.
func (m *BlueprintVersionKey) SetBlueprintVersion(value *string)() {
    m.blueprint_version = value
}
type BlueprintVersionKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlueprintName()(*string)
    GetBlueprintVersion()(*string)
    SetBlueprintName(value *string)()
    SetBlueprintVersion(value *string)()
}
