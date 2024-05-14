package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MetadataModuleEntrySubstate struct {
	Substate
	// The key property
	key MetadataKeyable
	// If missing, it represents a non-existing or deleted value.
	value MetadataModuleEntryValueable
}

// NewMetadataModuleEntrySubstate instantiates a new MetadataModuleEntrySubstate and sets the default values.
func NewMetadataModuleEntrySubstate() *MetadataModuleEntrySubstate {
	m := &MetadataModuleEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateMetadataModuleEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetadataModuleEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMetadataModuleEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MetadataModuleEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateMetadataKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(MetadataKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateMetadataModuleEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(MetadataModuleEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a MetadataKeyable when successful
func (m *MetadataModuleEntrySubstate) GetKey() MetadataKeyable {
	return m.key
}

// GetValue gets the value property value. If missing, it represents a non-existing or deleted value.
// returns a MetadataModuleEntryValueable when successful
func (m *MetadataModuleEntrySubstate) GetValue() MetadataModuleEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *MetadataModuleEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *MetadataModuleEntrySubstate) SetKey(value MetadataKeyable) {
	m.key = value
}

// SetValue sets the value property value. If missing, it represents a non-existing or deleted value.
func (m *MetadataModuleEntrySubstate) SetValue(value MetadataModuleEntryValueable) {
	m.value = value
}

type MetadataModuleEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() MetadataKeyable
	GetValue() MetadataModuleEntryValueable
	SetKey(value MetadataKeyable)
	SetValue(value MetadataModuleEntryValueable)
}
