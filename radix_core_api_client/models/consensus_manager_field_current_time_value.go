package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentTimeValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The proposer_timestamp property
	proposer_timestamp InstantMsable
}

// NewConsensusManagerFieldCurrentTimeValue instantiates a new ConsensusManagerFieldCurrentTimeValue and sets the default values.
func NewConsensusManagerFieldCurrentTimeValue() *ConsensusManagerFieldCurrentTimeValue {
	m := &ConsensusManagerFieldCurrentTimeValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateConsensusManagerFieldCurrentTimeValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentTimeValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldCurrentTimeValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldCurrentTimeValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentTimeValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["proposer_timestamp"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProposerTimestamp(val.(InstantMsable))
		}
		return nil
	}
	return res
}

// GetProposerTimestamp gets the proposer_timestamp property value. The proposer_timestamp property
// returns a InstantMsable when successful
func (m *ConsensusManagerFieldCurrentTimeValue) GetProposerTimestamp() InstantMsable {
	return m.proposer_timestamp
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentTimeValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("proposer_timestamp", m.GetProposerTimestamp())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConsensusManagerFieldCurrentTimeValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProposerTimestamp sets the proposer_timestamp property value. The proposer_timestamp property
func (m *ConsensusManagerFieldCurrentTimeValue) SetProposerTimestamp(value InstantMsable) {
	m.proposer_timestamp = value
}

type ConsensusManagerFieldCurrentTimeValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProposerTimestamp() InstantMsable
	SetProposerTimestamp(value InstantMsable)
}
