package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintDefinitionEntryValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The definition property
    definition BlueprintDefinitionable
}
// NewPackageBlueprintDefinitionEntryValue instantiates a new PackageBlueprintDefinitionEntryValue and sets the default values.
func NewPackageBlueprintDefinitionEntryValue()(*PackageBlueprintDefinitionEntryValue) {
    m := &PackageBlueprintDefinitionEntryValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageBlueprintDefinitionEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintDefinitionEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageBlueprintDefinitionEntryValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageBlueprintDefinitionEntryValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefinition gets the definition property value. The definition property
// returns a BlueprintDefinitionable when successful
func (m *PackageBlueprintDefinitionEntryValue) GetDefinition()(BlueprintDefinitionable) {
    return m.definition
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintDefinitionEntryValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlueprintDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(BlueprintDefinitionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PackageBlueprintDefinitionEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("definition", m.GetDefinition())
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
func (m *PackageBlueprintDefinitionEntryValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefinition sets the definition property value. The definition property
func (m *PackageBlueprintDefinitionEntryValue) SetDefinition(value BlueprintDefinitionable)() {
    m.definition = value
}
type PackageBlueprintDefinitionEntryValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefinition()(BlueprintDefinitionable)
    SetDefinition(value BlueprintDefinitionable)()
}
