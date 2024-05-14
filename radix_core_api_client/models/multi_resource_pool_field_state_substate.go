package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiResourcePoolFieldStateSubstate struct {
	Substate
	// The value property
	value MultiResourcePoolFieldStateValueable
}

// NewMultiResourcePoolFieldStateSubstate instantiates a new MultiResourcePoolFieldStateSubstate and sets the default values.
func NewMultiResourcePoolFieldStateSubstate() *MultiResourcePoolFieldStateSubstate {
	m := &MultiResourcePoolFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateMultiResourcePoolFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiResourcePoolFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMultiResourcePoolFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiResourcePoolFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateMultiResourcePoolFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(MultiResourcePoolFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a MultiResourcePoolFieldStateValueable when successful
func (m *MultiResourcePoolFieldStateSubstate) GetValue() MultiResourcePoolFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *MultiResourcePoolFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *MultiResourcePoolFieldStateSubstate) SetValue(value MultiResourcePoolFieldStateValueable) {
	m.value = value
}

type MultiResourcePoolFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() MultiResourcePoolFieldStateValueable
	SetValue(value MultiResourcePoolFieldStateValueable)
}
