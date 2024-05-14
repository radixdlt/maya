package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldConfigSubstate struct {
	Substate
	// The value property
	value ConsensusManagerFieldConfigValueable
}

// NewConsensusManagerFieldConfigSubstate instantiates a new ConsensusManagerFieldConfigSubstate and sets the default values.
func NewConsensusManagerFieldConfigSubstate() *ConsensusManagerFieldConfigSubstate {
	m := &ConsensusManagerFieldConfigSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateConsensusManagerFieldConfigSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldConfigSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldConfigSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldConfigSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateConsensusManagerFieldConfigValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ConsensusManagerFieldConfigValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldConfigValueable when successful
func (m *ConsensusManagerFieldConfigSubstate) GetValue() ConsensusManagerFieldConfigValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldConfigSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ConsensusManagerFieldConfigSubstate) SetValue(value ConsensusManagerFieldConfigValueable) {
	m.value = value
}

type ConsensusManagerFieldConfigSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() ConsensusManagerFieldConfigValueable
	SetValue(value ConsensusManagerFieldConfigValueable)
}
