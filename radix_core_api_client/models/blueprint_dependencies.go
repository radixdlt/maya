package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintDependencies struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dependencies property
    dependencies []string
}
// NewBlueprintDependencies instantiates a new BlueprintDependencies and sets the default values.
func NewBlueprintDependencies()(*BlueprintDependencies) {
    m := &BlueprintDependencies{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintDependenciesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintDependenciesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintDependencies(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintDependencies) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDependencies gets the dependencies property value. The dependencies property
// returns a []string when successful
func (m *BlueprintDependencies) GetDependencies()([]string) {
    return m.dependencies
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintDependencies) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDependencies(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BlueprintDependencies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDependencies() != nil {
        err := writer.WriteCollectionOfStringValues("dependencies", m.GetDependencies())
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
func (m *BlueprintDependencies) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDependencies sets the dependencies property value. The dependencies property
func (m *BlueprintDependencies) SetDependencies(value []string)() {
    m.dependencies = value
}
type BlueprintDependenciesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependencies()([]string)
    SetDependencies(value []string)()
}
