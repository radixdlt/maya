package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessControllerFieldStateSubstate struct {
	Substate
	// The value property
	value AccessControllerFieldStateValueable
}

// NewAccessControllerFieldStateSubstate instantiates a new AccessControllerFieldStateSubstate and sets the default values.
func NewAccessControllerFieldStateSubstate() *AccessControllerFieldStateSubstate {
	m := &AccessControllerFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateAccessControllerFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessControllerFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccessControllerFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessControllerFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccessControllerFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(AccessControllerFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a AccessControllerFieldStateValueable when successful
func (m *AccessControllerFieldStateSubstate) GetValue() AccessControllerFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *AccessControllerFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AccessControllerFieldStateSubstate) SetValue(value AccessControllerFieldStateValueable) {
	m.value = value
}

type AccessControllerFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() AccessControllerFieldStateValueable
	SetValue(value AccessControllerFieldStateValueable)
}
