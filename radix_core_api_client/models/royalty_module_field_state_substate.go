package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoyaltyModuleFieldStateSubstate struct {
	Substate
	// The value property
	value RoyaltyModuleFieldStateValueable
}

// NewRoyaltyModuleFieldStateSubstate instantiates a new RoyaltyModuleFieldStateSubstate and sets the default values.
func NewRoyaltyModuleFieldStateSubstate() *RoyaltyModuleFieldStateSubstate {
	m := &RoyaltyModuleFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateRoyaltyModuleFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoyaltyModuleFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoyaltyModuleFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoyaltyModuleFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRoyaltyModuleFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(RoyaltyModuleFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a RoyaltyModuleFieldStateValueable when successful
func (m *RoyaltyModuleFieldStateSubstate) GetValue() RoyaltyModuleFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *RoyaltyModuleFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *RoyaltyModuleFieldStateSubstate) SetValue(value RoyaltyModuleFieldStateValueable) {
	m.value = value
}

type RoyaltyModuleFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() RoyaltyModuleFieldStateValueable
	SetValue(value RoyaltyModuleFieldStateValueable)
}
