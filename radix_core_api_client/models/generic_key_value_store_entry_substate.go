package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericKeyValueStoreEntrySubstate struct {
	Substate
	// The key property
	key GenericKeyable
	// If not present, the entry has been deleted.
	value GenericKeyValueStoreEntryValueable
}

// NewGenericKeyValueStoreEntrySubstate instantiates a new GenericKeyValueStoreEntrySubstate and sets the default values.
func NewGenericKeyValueStoreEntrySubstate() *GenericKeyValueStoreEntrySubstate {
	m := &GenericKeyValueStoreEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateGenericKeyValueStoreEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericKeyValueStoreEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenericKeyValueStoreEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericKeyValueStoreEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateGenericKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(GenericKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateGenericKeyValueStoreEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(GenericKeyValueStoreEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a GenericKeyable when successful
func (m *GenericKeyValueStoreEntrySubstate) GetKey() GenericKeyable {
	return m.key
}

// GetValue gets the value property value. If not present, the entry has been deleted.
// returns a GenericKeyValueStoreEntryValueable when successful
func (m *GenericKeyValueStoreEntrySubstate) GetValue() GenericKeyValueStoreEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *GenericKeyValueStoreEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *GenericKeyValueStoreEntrySubstate) SetKey(value GenericKeyable) {
	m.key = value
}

// SetValue sets the value property value. If not present, the entry has been deleted.
func (m *GenericKeyValueStoreEntrySubstate) SetValue(value GenericKeyValueStoreEntryValueable) {
	m.value = value
}

type GenericKeyValueStoreEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() GenericKeyable
	GetValue() GenericKeyValueStoreEntryValueable
	SetKey(value GenericKeyable)
	SetValue(value GenericKeyValueStoreEntryValueable)
}
