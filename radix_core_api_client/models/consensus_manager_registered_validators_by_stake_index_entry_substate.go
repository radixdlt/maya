package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate struct {
	Substate
	// The key property
	key ActiveValidatorKeyable
	// The value property
	value ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable
}

// NewConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate instantiates a new ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate and sets the default values.
func NewConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate() *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate {
	m := &ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateActiveValidatorKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(ActiveValidatorKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateConsensusManagerRegisteredValidatorsByStakeIndexEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a ActiveValidatorKeyable when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) GetKey() ActiveValidatorKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) GetValue() ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Substate.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("key", m.GetKey())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKey sets the key property value. The key property
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) SetKey(value ActiveValidatorKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate) SetValue(value ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable) {
	m.value = value
}

type ConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() ActiveValidatorKeyable
	GetValue() ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable
	SetKey(value ActiveValidatorKeyable)
	SetValue(value ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable)
}
