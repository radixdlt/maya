package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectFieldStructure struct {
    SubstateSystemStructure
    // The value_schema property
    value_schema ObjectSubstateTypeReferenceable
}
// NewObjectFieldStructure instantiates a new ObjectFieldStructure and sets the default values.
func NewObjectFieldStructure()(*ObjectFieldStructure) {
    m := &ObjectFieldStructure{
        SubstateSystemStructure: *NewSubstateSystemStructure(),
    }
    return m
}
// CreateObjectFieldStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectFieldStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectFieldStructure(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectFieldStructure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubstateSystemStructure.GetFieldDeserializers()
    res["value_schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectSubstateTypeReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueSchema(val.(ObjectSubstateTypeReferenceable))
        }
        return nil
    }
    return res
}
// GetValueSchema gets the value_schema property value. The value_schema property
// returns a ObjectSubstateTypeReferenceable when successful
func (m *ObjectFieldStructure) GetValueSchema()(ObjectSubstateTypeReferenceable) {
    return m.value_schema
}
// Serialize serializes information the current object
func (m *ObjectFieldStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubstateSystemStructure.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("value_schema", m.GetValueSchema())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValueSchema sets the value_schema property value. The value_schema property
func (m *ObjectFieldStructure) SetValueSchema(value ObjectSubstateTypeReferenceable)() {
    m.value_schema = value
}
type ObjectFieldStructureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubstateSystemStructureable
    GetValueSchema()(ObjectSubstateTypeReferenceable)
    SetValueSchema(value ObjectSubstateTypeReferenceable)()
}
