package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TypeInfoModuleFieldTypeInfoSubstate struct {
    Substate
    // The value property
    value TypeInfoModuleFieldTypeInfoValueable
}
// NewTypeInfoModuleFieldTypeInfoSubstate instantiates a new TypeInfoModuleFieldTypeInfoSubstate and sets the default values.
func NewTypeInfoModuleFieldTypeInfoSubstate()(*TypeInfoModuleFieldTypeInfoSubstate) {
    m := &TypeInfoModuleFieldTypeInfoSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateTypeInfoModuleFieldTypeInfoSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTypeInfoModuleFieldTypeInfoSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTypeInfoModuleFieldTypeInfoSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TypeInfoModuleFieldTypeInfoSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTypeInfoModuleFieldTypeInfoValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(TypeInfoModuleFieldTypeInfoValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a TypeInfoModuleFieldTypeInfoValueable when successful
func (m *TypeInfoModuleFieldTypeInfoSubstate) GetValue()(TypeInfoModuleFieldTypeInfoValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *TypeInfoModuleFieldTypeInfoSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TypeInfoModuleFieldTypeInfoSubstate) SetValue(value TypeInfoModuleFieldTypeInfoValueable)() {
    m.value = value
}
type TypeInfoModuleFieldTypeInfoSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(TypeInfoModuleFieldTypeInfoValueable)
    SetValue(value TypeInfoModuleFieldTypeInfoValueable)()
}
