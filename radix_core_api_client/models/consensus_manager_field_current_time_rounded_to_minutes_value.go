package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentTimeRoundedToMinutesValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The proposer_timestamp_rounded_down_to_minute property
	proposer_timestamp_rounded_down_to_minute InstantMsable
}

// NewConsensusManagerFieldCurrentTimeRoundedToMinutesValue instantiates a new ConsensusManagerFieldCurrentTimeRoundedToMinutesValue and sets the default values.
func NewConsensusManagerFieldCurrentTimeRoundedToMinutesValue() *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue {
	m := &ConsensusManagerFieldCurrentTimeRoundedToMinutesValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateConsensusManagerFieldCurrentTimeRoundedToMinutesValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentTimeRoundedToMinutesValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldCurrentTimeRoundedToMinutesValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["proposer_timestamp_rounded_down_to_minute"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProposerTimestampRoundedDownToMinute(val.(InstantMsable))
		}
		return nil
	}
	return res
}

// GetProposerTimestampRoundedDownToMinute gets the proposer_timestamp_rounded_down_to_minute property value. The proposer_timestamp_rounded_down_to_minute property
// returns a InstantMsable when successful
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) GetProposerTimestampRoundedDownToMinute() InstantMsable {
	return m.proposer_timestamp_rounded_down_to_minute
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("proposer_timestamp_rounded_down_to_minute", m.GetProposerTimestampRoundedDownToMinute())
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
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProposerTimestampRoundedDownToMinute sets the proposer_timestamp_rounded_down_to_minute property value. The proposer_timestamp_rounded_down_to_minute property
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesValue) SetProposerTimestampRoundedDownToMinute(value InstantMsable) {
	m.proposer_timestamp_rounded_down_to_minute = value
}

type ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProposerTimestampRoundedDownToMinute() InstantMsable
	SetProposerTimestampRoundedDownToMinute(value InstantMsable)
}
