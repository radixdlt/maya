package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldStateValue struct {
	// The actual_epoch_start property
	actual_epoch_start InstantMsable
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The current_leader property
	current_leader ActiveValidatorIndexable
	// The effective_epoch_start property
	effective_epoch_start InstantMsable
	// An integer between `0` and `10^10`, marking the current epoch
	epoch *int64
	// The is_started property
	is_started *bool
	// An integer between `0` and `10^10`, marking the current round in an epoch
	round *int64
}

// NewConsensusManagerFieldStateValue instantiates a new ConsensusManagerFieldStateValue and sets the default values.
func NewConsensusManagerFieldStateValue() *ConsensusManagerFieldStateValue {
	m := &ConsensusManagerFieldStateValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateConsensusManagerFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldStateValue(), nil
}

// GetActualEpochStart gets the actual_epoch_start property value. The actual_epoch_start property
// returns a InstantMsable when successful
func (m *ConsensusManagerFieldStateValue) GetActualEpochStart() InstantMsable {
	return m.actual_epoch_start
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldStateValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCurrentLeader gets the current_leader property value. The current_leader property
// returns a ActiveValidatorIndexable when successful
func (m *ConsensusManagerFieldStateValue) GetCurrentLeader() ActiveValidatorIndexable {
	return m.current_leader
}

// GetEffectiveEpochStart gets the effective_epoch_start property value. The effective_epoch_start property
// returns a InstantMsable when successful
func (m *ConsensusManagerFieldStateValue) GetEffectiveEpochStart() InstantMsable {
	return m.effective_epoch_start
}

// GetEpoch gets the epoch property value. An integer between `0` and `10^10`, marking the current epoch
// returns a *int64 when successful
func (m *ConsensusManagerFieldStateValue) GetEpoch() *int64 {
	return m.epoch
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldStateValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["actual_epoch_start"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActualEpochStart(val.(InstantMsable))
		}
		return nil
	}
	res["current_leader"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateActiveValidatorIndexFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrentLeader(val.(ActiveValidatorIndexable))
		}
		return nil
	}
	res["effective_epoch_start"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEffectiveEpochStart(val.(InstantMsable))
		}
		return nil
	}
	res["epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEpoch(val)
		}
		return nil
	}
	res["is_started"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsStarted(val)
		}
		return nil
	}
	res["round"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRound(val)
		}
		return nil
	}
	return res
}

// GetIsStarted gets the is_started property value. The is_started property
// returns a *bool when successful
func (m *ConsensusManagerFieldStateValue) GetIsStarted() *bool {
	return m.is_started
}

// GetRound gets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
// returns a *int64 when successful
func (m *ConsensusManagerFieldStateValue) GetRound() *int64 {
	return m.round
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("actual_epoch_start", m.GetActualEpochStart())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("current_leader", m.GetCurrentLeader())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("effective_epoch_start", m.GetEffectiveEpochStart())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("epoch", m.GetEpoch())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_started", m.GetIsStarted())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("round", m.GetRound())
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

// SetActualEpochStart sets the actual_epoch_start property value. The actual_epoch_start property
func (m *ConsensusManagerFieldStateValue) SetActualEpochStart(value InstantMsable) {
	m.actual_epoch_start = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConsensusManagerFieldStateValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCurrentLeader sets the current_leader property value. The current_leader property
func (m *ConsensusManagerFieldStateValue) SetCurrentLeader(value ActiveValidatorIndexable) {
	m.current_leader = value
}

// SetEffectiveEpochStart sets the effective_epoch_start property value. The effective_epoch_start property
func (m *ConsensusManagerFieldStateValue) SetEffectiveEpochStart(value InstantMsable) {
	m.effective_epoch_start = value
}

// SetEpoch sets the epoch property value. An integer between `0` and `10^10`, marking the current epoch
func (m *ConsensusManagerFieldStateValue) SetEpoch(value *int64) {
	m.epoch = value
}

// SetIsStarted sets the is_started property value. The is_started property
func (m *ConsensusManagerFieldStateValue) SetIsStarted(value *bool) {
	m.is_started = value
}

// SetRound sets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
func (m *ConsensusManagerFieldStateValue) SetRound(value *int64) {
	m.round = value
}

type ConsensusManagerFieldStateValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActualEpochStart() InstantMsable
	GetCurrentLeader() ActiveValidatorIndexable
	GetEffectiveEpochStart() InstantMsable
	GetEpoch() *int64
	GetIsStarted() *bool
	GetRound() *int64
	SetActualEpochStart(value InstantMsable)
	SetCurrentLeader(value ActiveValidatorIndexable)
	SetEffectiveEpochStart(value InstantMsable)
	SetEpoch(value *int64)
	SetIsStarted(value *bool)
	SetRound(value *int64)
}
