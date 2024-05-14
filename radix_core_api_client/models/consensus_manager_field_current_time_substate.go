package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentTimeSubstate struct {
	Substate
	// The value property
	value ConsensusManagerFieldCurrentTimeValueable
}

// NewConsensusManagerFieldCurrentTimeSubstate instantiates a new ConsensusManagerFieldCurrentTimeSubstate and sets the default values.
func NewConsensusManagerFieldCurrentTimeSubstate() *ConsensusManagerFieldCurrentTimeSubstate {
	m := &ConsensusManagerFieldCurrentTimeSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateConsensusManagerFieldCurrentTimeSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentTimeSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldCurrentTimeSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentTimeSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateConsensusManagerFieldCurrentTimeValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ConsensusManagerFieldCurrentTimeValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldCurrentTimeValueable when successful
func (m *ConsensusManagerFieldCurrentTimeSubstate) GetValue() ConsensusManagerFieldCurrentTimeValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentTimeSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ConsensusManagerFieldCurrentTimeSubstate) SetValue(value ConsensusManagerFieldCurrentTimeValueable) {
	m.value = value
}

type ConsensusManagerFieldCurrentTimeSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() ConsensusManagerFieldCurrentTimeValueable
	SetValue(value ConsensusManagerFieldCurrentTimeValueable)
}
