package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldValidatorRewardsSubstate struct {
	Substate
	// The value property
	value ConsensusManagerFieldValidatorRewardsValueable
}

// NewConsensusManagerFieldValidatorRewardsSubstate instantiates a new ConsensusManagerFieldValidatorRewardsSubstate and sets the default values.
func NewConsensusManagerFieldValidatorRewardsSubstate() *ConsensusManagerFieldValidatorRewardsSubstate {
	m := &ConsensusManagerFieldValidatorRewardsSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateConsensusManagerFieldValidatorRewardsSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldValidatorRewardsSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldValidatorRewardsSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldValidatorRewardsSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateConsensusManagerFieldValidatorRewardsValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ConsensusManagerFieldValidatorRewardsValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldValidatorRewardsValueable when successful
func (m *ConsensusManagerFieldValidatorRewardsSubstate) GetValue() ConsensusManagerFieldValidatorRewardsValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldValidatorRewardsSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ConsensusManagerFieldValidatorRewardsSubstate) SetValue(value ConsensusManagerFieldValidatorRewardsValueable) {
	m.value = value
}

type ConsensusManagerFieldValidatorRewardsSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() ConsensusManagerFieldValidatorRewardsValueable
	SetValue(value ConsensusManagerFieldValidatorRewardsValueable)
}
