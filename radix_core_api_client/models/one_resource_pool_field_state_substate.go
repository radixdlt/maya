package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OneResourcePoolFieldStateSubstate struct {
	Substate
	// The value property
	value OneResourcePoolFieldStateValueable
}

// NewOneResourcePoolFieldStateSubstate instantiates a new OneResourcePoolFieldStateSubstate and sets the default values.
func NewOneResourcePoolFieldStateSubstate() *OneResourcePoolFieldStateSubstate {
	m := &OneResourcePoolFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateOneResourcePoolFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOneResourcePoolFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOneResourcePoolFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OneResourcePoolFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOneResourcePoolFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(OneResourcePoolFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a OneResourcePoolFieldStateValueable when successful
func (m *OneResourcePoolFieldStateSubstate) GetValue() OneResourcePoolFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *OneResourcePoolFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Substate.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. The value property
func (m *OneResourcePoolFieldStateSubstate) SetValue(value OneResourcePoolFieldStateValueable) {
	m.value = value
}

type OneResourcePoolFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() OneResourcePoolFieldStateValueable
	SetValue(value OneResourcePoolFieldStateValueable)
}
