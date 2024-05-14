package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidatorFieldProtocolUpdateReadinessSignalSubstate struct {
	Substate
	// The value property
	value ValidatorFieldProtocolUpdateReadinessSignalValueable
}

// NewValidatorFieldProtocolUpdateReadinessSignalSubstate instantiates a new ValidatorFieldProtocolUpdateReadinessSignalSubstate and sets the default values.
func NewValidatorFieldProtocolUpdateReadinessSignalSubstate() *ValidatorFieldProtocolUpdateReadinessSignalSubstate {
	m := &ValidatorFieldProtocolUpdateReadinessSignalSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateValidatorFieldProtocolUpdateReadinessSignalSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidatorFieldProtocolUpdateReadinessSignalSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewValidatorFieldProtocolUpdateReadinessSignalSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidatorFieldProtocolUpdateReadinessSignalSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateValidatorFieldProtocolUpdateReadinessSignalValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ValidatorFieldProtocolUpdateReadinessSignalValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a ValidatorFieldProtocolUpdateReadinessSignalValueable when successful
func (m *ValidatorFieldProtocolUpdateReadinessSignalSubstate) GetValue() ValidatorFieldProtocolUpdateReadinessSignalValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ValidatorFieldProtocolUpdateReadinessSignalSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ValidatorFieldProtocolUpdateReadinessSignalSubstate) SetValue(value ValidatorFieldProtocolUpdateReadinessSignalValueable) {
	m.value = value
}

type ValidatorFieldProtocolUpdateReadinessSignalSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() ValidatorFieldProtocolUpdateReadinessSignalValueable
	SetValue(value ValidatorFieldProtocolUpdateReadinessSignalValueable)
}
