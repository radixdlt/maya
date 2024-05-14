package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoyaltyModuleMethodRoyaltyEntrySubstate struct {
	Substate
	// The key property
	key MainMethodKeyable
	// If missing, it represents a free method.
	value RoyaltyModuleMethodRoyaltyEntryValueable
}

// NewRoyaltyModuleMethodRoyaltyEntrySubstate instantiates a new RoyaltyModuleMethodRoyaltyEntrySubstate and sets the default values.
func NewRoyaltyModuleMethodRoyaltyEntrySubstate() *RoyaltyModuleMethodRoyaltyEntrySubstate {
	m := &RoyaltyModuleMethodRoyaltyEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateRoyaltyModuleMethodRoyaltyEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoyaltyModuleMethodRoyaltyEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoyaltyModuleMethodRoyaltyEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateMainMethodKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(MainMethodKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRoyaltyModuleMethodRoyaltyEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(RoyaltyModuleMethodRoyaltyEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a MainMethodKeyable when successful
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) GetKey() MainMethodKeyable {
	return m.key
}

// GetValue gets the value property value. If missing, it represents a free method.
// returns a RoyaltyModuleMethodRoyaltyEntryValueable when successful
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) GetValue() RoyaltyModuleMethodRoyaltyEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) SetKey(value MainMethodKeyable) {
	m.key = value
}

// SetValue sets the value property value. If missing, it represents a free method.
func (m *RoyaltyModuleMethodRoyaltyEntrySubstate) SetValue(value RoyaltyModuleMethodRoyaltyEntryValueable) {
	m.value = value
}

type RoyaltyModuleMethodRoyaltyEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() MainMethodKeyable
	GetValue() RoyaltyModuleMethodRoyaltyEntryValueable
	SetKey(value MainMethodKeyable)
	SetValue(value RoyaltyModuleMethodRoyaltyEntryValueable)
}
