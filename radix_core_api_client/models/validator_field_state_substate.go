package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidatorFieldStateSubstate struct {
	Substate
	// The value property
	value ValidatorFieldStateValueable
}

// NewValidatorFieldStateSubstate instantiates a new ValidatorFieldStateSubstate and sets the default values.
func NewValidatorFieldStateSubstate() *ValidatorFieldStateSubstate {
	m := &ValidatorFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateValidatorFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidatorFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewValidatorFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidatorFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateValidatorFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ValidatorFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a ValidatorFieldStateValueable when successful
func (m *ValidatorFieldStateSubstate) GetValue() ValidatorFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ValidatorFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ValidatorFieldStateSubstate) SetValue(value ValidatorFieldStateValueable) {
	m.value = value
}

type ValidatorFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() ValidatorFieldStateValueable
	SetValue(value ValidatorFieldStateValueable)
}
