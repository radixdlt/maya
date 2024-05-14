package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountFieldStateSubstate struct {
	Substate
	// The value property
	value AccountFieldStateValueable
}

// NewAccountFieldStateSubstate instantiates a new AccountFieldStateSubstate and sets the default values.
func NewAccountFieldStateSubstate() *AccountFieldStateSubstate {
	m := &AccountFieldStateSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateAccountFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountFieldStateSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountFieldStateSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccountFieldStateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(AccountFieldStateValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a AccountFieldStateValueable when successful
func (m *AccountFieldStateSubstate) GetValue() AccountFieldStateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *AccountFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AccountFieldStateSubstate) SetValue(value AccountFieldStateValueable) {
	m.value = value
}

type AccountFieldStateSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() AccountFieldStateValueable
	SetValue(value AccountFieldStateValueable)
}
