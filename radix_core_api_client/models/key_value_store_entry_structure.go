package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KeyValueStoreEntryStructure struct {
    SubstateSystemStructure
    // An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
    key_full_type_id FullyScopedTypeIdable
    // An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
    value_full_type_id FullyScopedTypeIdable
}
// NewKeyValueStoreEntryStructure instantiates a new KeyValueStoreEntryStructure and sets the default values.
func NewKeyValueStoreEntryStructure()(*KeyValueStoreEntryStructure) {
    m := &KeyValueStoreEntryStructure{
        SubstateSystemStructure: *NewSubstateSystemStructure(),
    }
    return m
}
// CreateKeyValueStoreEntryStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeyValueStoreEntryStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyValueStoreEntryStructure(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeyValueStoreEntryStructure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubstateSystemStructure.GetFieldDeserializers()
    res["key_full_type_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFullyScopedTypeIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyFullTypeId(val.(FullyScopedTypeIdable))
        }
        return nil
    }
    res["value_full_type_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFullyScopedTypeIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueFullTypeId(val.(FullyScopedTypeIdable))
        }
        return nil
    }
    return res
}
// GetKeyFullTypeId gets the key_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
// returns a FullyScopedTypeIdable when successful
func (m *KeyValueStoreEntryStructure) GetKeyFullTypeId()(FullyScopedTypeIdable) {
    return m.key_full_type_id
}
// GetValueFullTypeId gets the value_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
// returns a FullyScopedTypeIdable when successful
func (m *KeyValueStoreEntryStructure) GetValueFullTypeId()(FullyScopedTypeIdable) {
    return m.value_full_type_id
}
// Serialize serializes information the current object
func (m *KeyValueStoreEntryStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubstateSystemStructure.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("key_full_type_id", m.GetKeyFullTypeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value_full_type_id", m.GetValueFullTypeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyFullTypeId sets the key_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
func (m *KeyValueStoreEntryStructure) SetKeyFullTypeId(value FullyScopedTypeIdable)() {
    m.key_full_type_id = value
}
// SetValueFullTypeId sets the value_full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
func (m *KeyValueStoreEntryStructure) SetValueFullTypeId(value FullyScopedTypeIdable)() {
    m.value_full_type_id = value
}
type KeyValueStoreEntryStructureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubstateSystemStructureable
    GetKeyFullTypeId()(FullyScopedTypeIdable)
    GetValueFullTypeId()(FullyScopedTypeIdable)
    SetKeyFullTypeId(value FullyScopedTypeIdable)()
    SetValueFullTypeId(value FullyScopedTypeIdable)()
}
