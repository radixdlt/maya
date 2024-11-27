package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SystemSchemaStructure struct {
    SubstateSystemStructure
}
// NewSystemSchemaStructure instantiates a new SystemSchemaStructure and sets the default values.
func NewSystemSchemaStructure()(*SystemSchemaStructure) {
    m := &SystemSchemaStructure{
        SubstateSystemStructure: *NewSubstateSystemStructure(),
    }
    return m
}
// CreateSystemSchemaStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSystemSchemaStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSystemSchemaStructure(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SystemSchemaStructure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubstateSystemStructure.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SystemSchemaStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubstateSystemStructure.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SystemSchemaStructureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubstateSystemStructureable
}
