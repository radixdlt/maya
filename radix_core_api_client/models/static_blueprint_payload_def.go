package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StaticBlueprintPayloadDef struct {
    BlueprintPayloadDef
    // An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
    type_id ScopedTypeIdable
}
// NewStaticBlueprintPayloadDef instantiates a new StaticBlueprintPayloadDef and sets the default values.
func NewStaticBlueprintPayloadDef()(*StaticBlueprintPayloadDef) {
    m := &StaticBlueprintPayloadDef{
        BlueprintPayloadDef: *NewBlueprintPayloadDef(),
    }
    return m
}
// CreateStaticBlueprintPayloadDefFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStaticBlueprintPayloadDefFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStaticBlueprintPayloadDef(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StaticBlueprintPayloadDef) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BlueprintPayloadDef.GetFieldDeserializers()
    res["type_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScopedTypeIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeId(val.(ScopedTypeIdable))
        }
        return nil
    }
    return res
}
// GetTypeId gets the type_id property value. An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
// returns a ScopedTypeIdable when successful
func (m *StaticBlueprintPayloadDef) GetTypeId()(ScopedTypeIdable) {
    return m.type_id
}
// Serialize serializes information the current object
func (m *StaticBlueprintPayloadDef) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BlueprintPayloadDef.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("type_id", m.GetTypeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeId sets the type_id property value. An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
func (m *StaticBlueprintPayloadDef) SetTypeId(value ScopedTypeIdable)() {
    m.type_id = value
}
type StaticBlueprintPayloadDefable interface {
    BlueprintPayloadDefable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeId()(ScopedTypeIdable)
    SetTypeId(value ScopedTypeIdable)()
}
