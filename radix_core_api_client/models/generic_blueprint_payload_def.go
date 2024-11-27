package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericBlueprintPayloadDef struct {
    BlueprintPayloadDef
    // An index within the list of generic type substitutions.
    generic_index *int32
}
// NewGenericBlueprintPayloadDef instantiates a new GenericBlueprintPayloadDef and sets the default values.
func NewGenericBlueprintPayloadDef()(*GenericBlueprintPayloadDef) {
    m := &GenericBlueprintPayloadDef{
        BlueprintPayloadDef: *NewBlueprintPayloadDef(),
    }
    return m
}
// CreateGenericBlueprintPayloadDefFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericBlueprintPayloadDefFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenericBlueprintPayloadDef(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericBlueprintPayloadDef) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BlueprintPayloadDef.GetFieldDeserializers()
    res["generic_index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGenericIndex(val)
        }
        return nil
    }
    return res
}
// GetGenericIndex gets the generic_index property value. An index within the list of generic type substitutions.
// returns a *int32 when successful
func (m *GenericBlueprintPayloadDef) GetGenericIndex()(*int32) {
    return m.generic_index
}
// Serialize serializes information the current object
func (m *GenericBlueprintPayloadDef) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BlueprintPayloadDef.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("generic_index", m.GetGenericIndex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGenericIndex sets the generic_index property value. An index within the list of generic type substitutions.
func (m *GenericBlueprintPayloadDef) SetGenericIndex(value *int32)() {
    m.generic_index = value
}
type GenericBlueprintPayloadDefable interface {
    BlueprintPayloadDefable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGenericIndex()(*int32)
    SetGenericIndex(value *int32)()
}
