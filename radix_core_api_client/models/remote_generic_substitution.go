package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteGenericSubstitution the generic substitution is provided remotely by a blueprint type.The `resolved_full_type_id` is added by the node, and is always present in the model returned from the transaction stream API.Other APIs may not resolve the type from the blueprint definition.
type RemoteGenericSubstitution struct {
    GenericSubstitution
    // An identifier for a defined type in the v1 blueprint version under the given package blueprint.
    blueprint_type_identifier BlueprintTypeIdentifierable
    // An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
    resolved_full_type_id FullyScopedTypeIdable
}
// NewRemoteGenericSubstitution instantiates a new RemoteGenericSubstitution and sets the default values.
func NewRemoteGenericSubstitution()(*RemoteGenericSubstitution) {
    m := &RemoteGenericSubstitution{
        GenericSubstitution: *NewGenericSubstitution(),
    }
    return m
}
// CreateRemoteGenericSubstitutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemoteGenericSubstitutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteGenericSubstitution(), nil
}
// GetBlueprintTypeIdentifier gets the blueprint_type_identifier property value. An identifier for a defined type in the v1 blueprint version under the given package blueprint.
// returns a BlueprintTypeIdentifierable when successful
func (m *RemoteGenericSubstitution) GetBlueprintTypeIdentifier()(BlueprintTypeIdentifierable) {
    return m.blueprint_type_identifier
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RemoteGenericSubstitution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GenericSubstitution.GetFieldDeserializers()
    res["blueprint_type_identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlueprintTypeIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueprintTypeIdentifier(val.(BlueprintTypeIdentifierable))
        }
        return nil
    }
    res["resolved_full_type_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFullyScopedTypeIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedFullTypeId(val.(FullyScopedTypeIdable))
        }
        return nil
    }
    return res
}
// GetResolvedFullTypeId gets the resolved_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
// returns a FullyScopedTypeIdable when successful
func (m *RemoteGenericSubstitution) GetResolvedFullTypeId()(FullyScopedTypeIdable) {
    return m.resolved_full_type_id
}
// Serialize serializes information the current object
func (m *RemoteGenericSubstitution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GenericSubstitution.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("blueprint_type_identifier", m.GetBlueprintTypeIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resolved_full_type_id", m.GetResolvedFullTypeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlueprintTypeIdentifier sets the blueprint_type_identifier property value. An identifier for a defined type in the v1 blueprint version under the given package blueprint.
func (m *RemoteGenericSubstitution) SetBlueprintTypeIdentifier(value BlueprintTypeIdentifierable)() {
    m.blueprint_type_identifier = value
}
// SetResolvedFullTypeId sets the resolved_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
func (m *RemoteGenericSubstitution) SetResolvedFullTypeId(value FullyScopedTypeIdable)() {
    m.resolved_full_type_id = value
}
type RemoteGenericSubstitutionable interface {
    GenericSubstitutionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlueprintTypeIdentifier()(BlueprintTypeIdentifierable)
    GetResolvedFullTypeId()(FullyScopedTypeIdable)
    SetBlueprintTypeIdentifier(value BlueprintTypeIdentifierable)()
    SetResolvedFullTypeId(value FullyScopedTypeIdable)()
}
