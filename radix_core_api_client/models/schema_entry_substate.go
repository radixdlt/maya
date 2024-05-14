package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SchemaEntrySubstate struct {
	Substate
	// The key property
	key SchemaKeyable
	// The value property
	value SchemaEntryValueable
}

// NewSchemaEntrySubstate instantiates a new SchemaEntrySubstate and sets the default values.
func NewSchemaEntrySubstate() *SchemaEntrySubstate {
	m := &SchemaEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateSchemaEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSchemaEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSchemaKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(SchemaKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSchemaEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(SchemaEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a SchemaKeyable when successful
func (m *SchemaEntrySubstate) GetKey() SchemaKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a SchemaEntryValueable when successful
func (m *SchemaEntrySubstate) GetValue() SchemaEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *SchemaEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Substate.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("key", m.GetKey())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKey sets the key property value. The key property
func (m *SchemaEntrySubstate) SetKey(value SchemaKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *SchemaEntrySubstate) SetValue(value SchemaEntryValueable) {
	m.value = value
}

type SchemaEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() SchemaKeyable
	GetValue() SchemaEntryValueable
	SetKey(value SchemaKeyable)
	SetValue(value SchemaEntryValueable)
}
