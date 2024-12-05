package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectInstanceTypeReference struct {
    ObjectSubstateTypeReference
    // The index of the generic parameter in the instance definition.Was called the `instance_type_index`.
    generic_index *int32
    // An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
    resolved_full_type_id FullyScopedTypeIdable
}
// NewObjectInstanceTypeReference instantiates a new ObjectInstanceTypeReference and sets the default values.
func NewObjectInstanceTypeReference()(*ObjectInstanceTypeReference) {
    m := &ObjectInstanceTypeReference{
        ObjectSubstateTypeReference: *NewObjectSubstateTypeReference(),
    }
    return m
}
// CreateObjectInstanceTypeReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectInstanceTypeReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectInstanceTypeReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectInstanceTypeReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ObjectSubstateTypeReference.GetFieldDeserializers()
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
// GetGenericIndex gets the generic_index property value. The index of the generic parameter in the instance definition.Was called the `instance_type_index`.
// returns a *int32 when successful
func (m *ObjectInstanceTypeReference) GetGenericIndex()(*int32) {
    return m.generic_index
}
// GetResolvedFullTypeId gets the resolved_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
// returns a FullyScopedTypeIdable when successful
func (m *ObjectInstanceTypeReference) GetResolvedFullTypeId()(FullyScopedTypeIdable) {
    return m.resolved_full_type_id
}
// Serialize serializes information the current object
func (m *ObjectInstanceTypeReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ObjectSubstateTypeReference.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("generic_index", m.GetGenericIndex())
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
// SetGenericIndex sets the generic_index property value. The index of the generic parameter in the instance definition.Was called the `instance_type_index`.
func (m *ObjectInstanceTypeReference) SetGenericIndex(value *int32)() {
    m.generic_index = value
}
// SetResolvedFullTypeId sets the resolved_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
func (m *ObjectInstanceTypeReference) SetResolvedFullTypeId(value FullyScopedTypeIdable)() {
    m.resolved_full_type_id = value
}
type ObjectInstanceTypeReferenceable interface {
    ObjectSubstateTypeReferenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGenericIndex()(*int32)
    GetResolvedFullTypeId()(FullyScopedTypeIdable)
    SetGenericIndex(value *int32)()
    SetResolvedFullTypeId(value FullyScopedTypeIdable)()
}
