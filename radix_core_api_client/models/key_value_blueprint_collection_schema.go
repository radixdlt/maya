package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KeyValueBlueprintCollectionSchema struct {
	BlueprintCollectionSchema
	// Whether the entries of the key-value partition are allowed to own child nodes.
	allow_ownership *bool
	// The key_type_ref property
	key_type_ref BlueprintPayloadDefable
	// The value_type_ref property
	value_type_ref BlueprintPayloadDefable
}

// NewKeyValueBlueprintCollectionSchema instantiates a new KeyValueBlueprintCollectionSchema and sets the default values.
func NewKeyValueBlueprintCollectionSchema() *KeyValueBlueprintCollectionSchema {
	m := &KeyValueBlueprintCollectionSchema{
		BlueprintCollectionSchema: *NewBlueprintCollectionSchema(),
	}
	return m
}

// CreateKeyValueBlueprintCollectionSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeyValueBlueprintCollectionSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewKeyValueBlueprintCollectionSchema(), nil
}

// GetAllowOwnership gets the allow_ownership property value. Whether the entries of the key-value partition are allowed to own child nodes.
// returns a *bool when successful
func (m *KeyValueBlueprintCollectionSchema) GetAllowOwnership() *bool {
	return m.allow_ownership
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeyValueBlueprintCollectionSchema) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *KeyValueBlueprintCollectionSchema) GetKeyTypeRef() BlueprintPayloadDefable {
	return m.key_type_ref
}

// GetValueTypeRef gets the value_type_ref property value. The value_type_ref property
// returns a BlueprintPayloadDefable when successful
func (m *KeyValueBlueprintCollectionSchema) GetValueTypeRef() BlueprintPayloadDefable {
	return m.value_type_ref
}

// Serialize serializes information the current object
func (m *KeyValueBlueprintCollectionSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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

// SetAllowOwnership sets the allow_ownership property value. Whether the entries of the key-value partition are allowed to own child nodes.
func (m *KeyValueBlueprintCollectionSchema) SetAllowOwnership(value *bool) {
	m.allow_ownership = value
}

// SetKeyTypeRef sets the key_type_ref property value. The key_type_ref property
func (m *KeyValueBlueprintCollectionSchema) SetKeyTypeRef(value BlueprintPayloadDefable) {
	m.key_type_ref = value
}

// SetValueTypeRef sets the value_type_ref property value. The value_type_ref property
func (m *KeyValueBlueprintCollectionSchema) SetValueTypeRef(value BlueprintPayloadDefable) {
	m.value_type_ref = value
}

type KeyValueBlueprintCollectionSchemaable interface {
	BlueprintCollectionSchemaable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAllowOwnership() *bool
	GetKeyTypeRef() BlueprintPayloadDefable
	GetValueTypeRef() BlueprintPayloadDefable
	SetAllowOwnership(value *bool)
	SetKeyTypeRef(value BlueprintPayloadDefable)
	SetValueTypeRef(value BlueprintPayloadDefable)
}
