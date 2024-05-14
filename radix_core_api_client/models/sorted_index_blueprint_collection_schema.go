package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SortedIndexBlueprintCollectionSchema struct {
	BlueprintCollectionSchema
	// Whether the entries of the sorted index partition are allowed to own child nodes.
	allow_ownership *bool
	// The key_type_ref property
	key_type_ref BlueprintPayloadDefable
	// The value_type_ref property
	value_type_ref BlueprintPayloadDefable
}

// NewSortedIndexBlueprintCollectionSchema instantiates a new SortedIndexBlueprintCollectionSchema and sets the default values.
func NewSortedIndexBlueprintCollectionSchema() *SortedIndexBlueprintCollectionSchema {
	m := &SortedIndexBlueprintCollectionSchema{
		BlueprintCollectionSchema: *NewBlueprintCollectionSchema(),
	}
	return m
}

// CreateSortedIndexBlueprintCollectionSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSortedIndexBlueprintCollectionSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSortedIndexBlueprintCollectionSchema(), nil
}

// GetAllowOwnership gets the allow_ownership property value. Whether the entries of the sorted index partition are allowed to own child nodes.
// returns a *bool when successful
func (m *SortedIndexBlueprintCollectionSchema) GetAllowOwnership() *bool {
	return m.allow_ownership
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SortedIndexBlueprintCollectionSchema) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BlueprintCollectionSchema.GetFieldDeserializers()
	res["allow_ownership"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAllowOwnership(val)
		}
		return nil
	}
	res["key_type_ref"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintPayloadDefFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyTypeRef(val.(BlueprintPayloadDefable))
		}
		return nil
	}
	res["value_type_ref"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintPayloadDefFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValueTypeRef(val.(BlueprintPayloadDefable))
		}
		return nil
	}
	return res
}

// GetKeyTypeRef gets the key_type_ref property value. The key_type_ref property
// returns a BlueprintPayloadDefable when successful
func (m *SortedIndexBlueprintCollectionSchema) GetKeyTypeRef() BlueprintPayloadDefable {
	return m.key_type_ref
}

// GetValueTypeRef gets the value_type_ref property value. The value_type_ref property
// returns a BlueprintPayloadDefable when successful
func (m *SortedIndexBlueprintCollectionSchema) GetValueTypeRef() BlueprintPayloadDefable {
	return m.value_type_ref
}

// Serialize serializes information the current object
func (m *SortedIndexBlueprintCollectionSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.BlueprintCollectionSchema.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteBoolValue("allow_ownership", m.GetAllowOwnership())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("key_type_ref", m.GetKeyTypeRef())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("value_type_ref", m.GetValueTypeRef())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAllowOwnership sets the allow_ownership property value. Whether the entries of the sorted index partition are allowed to own child nodes.
func (m *SortedIndexBlueprintCollectionSchema) SetAllowOwnership(value *bool) {
	m.allow_ownership = value
}

// SetKeyTypeRef sets the key_type_ref property value. The key_type_ref property
func (m *SortedIndexBlueprintCollectionSchema) SetKeyTypeRef(value BlueprintPayloadDefable) {
	m.key_type_ref = value
}

// SetValueTypeRef sets the value_type_ref property value. The value_type_ref property
func (m *SortedIndexBlueprintCollectionSchema) SetValueTypeRef(value BlueprintPayloadDefable) {
	m.value_type_ref = value
}

type SortedIndexBlueprintCollectionSchemaable interface {
	BlueprintCollectionSchemaable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAllowOwnership() *bool
	GetKeyTypeRef() BlueprintPayloadDefable
	GetValueTypeRef() BlueprintPayloadDefable
	SetAllowOwnership(value *bool)
	SetKeyTypeRef(value BlueprintPayloadDefable)
	SetValueTypeRef(value BlueprintPayloadDefable)
}
