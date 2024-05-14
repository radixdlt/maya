package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EpochChangeCondition struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An integer between `0` and `10^10`, specifying the maximum number of rounds per epoch
	max_round_count *int64
	// An integer between `0` and `10^10`, specifying the minimum number of rounds per epoch
	min_round_count *int64
	// An integer between `0` and `10^10`, specifying the target number of milliseconds per epoch,assuming the round number is within the min and max range.
	target_duration_millis *int64
}

// NewEpochChangeCondition instantiates a new EpochChangeCondition and sets the default values.
func NewEpochChangeCondition() *EpochChangeCondition {
	m := &EpochChangeCondition{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEpochChangeConditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEpochChangeConditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEpochChangeCondition(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EpochChangeCondition) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EpochChangeCondition) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["max_round_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxRoundCount(val)
		}
		return nil
	}
	res["min_round_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMinRoundCount(val)
		}
		return nil
	}
	res["target_duration_millis"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTargetDurationMillis(val)
		}
		return nil
	}
	return res
}

// GetMaxRoundCount gets the max_round_count property value. An integer between `0` and `10^10`, specifying the maximum number of rounds per epoch
// returns a *int64 when successful
func (m *EpochChangeCondition) GetMaxRoundCount() *int64 {
	return m.max_round_count
}

// GetMinRoundCount gets the min_round_count property value. An integer between `0` and `10^10`, specifying the minimum number of rounds per epoch
// returns a *int64 when successful
func (m *EpochChangeCondition) GetMinRoundCount() *int64 {
	return m.min_round_count
}

// GetTargetDurationMillis gets the target_duration_millis property value. An integer between `0` and `10^10`, specifying the target number of milliseconds per epoch,assuming the round number is within the min and max range.
// returns a *int64 when successful
func (m *EpochChangeCondition) GetTargetDurationMillis() *int64 {
	return m.target_duration_millis
}

// Serialize serializes information the current object
func (m *EpochChangeCondition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("max_round_count", m.GetMaxRoundCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("min_round_count", m.GetMinRoundCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("target_duration_millis", m.GetTargetDurationMillis())
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
func (m *EpochChangeCondition) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetMaxRoundCount sets the max_round_count property value. An integer between `0` and `10^10`, specifying the maximum number of rounds per epoch
func (m *EpochChangeCondition) SetMaxRoundCount(value *int64) {
	m.max_round_count = value
}

// SetMinRoundCount sets the min_round_count property value. An integer between `0` and `10^10`, specifying the minimum number of rounds per epoch
func (m *EpochChangeCondition) SetMinRoundCount(value *int64) {
	m.min_round_count = value
}

// SetTargetDurationMillis sets the target_duration_millis property value. An integer between `0` and `10^10`, specifying the target number of milliseconds per epoch,assuming the round number is within the min and max range.
func (m *EpochChangeCondition) SetTargetDurationMillis(value *int64) {
	m.target_duration_millis = value
}

type EpochChangeConditionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMaxRoundCount() *int64
	GetMinRoundCount() *int64
	GetTargetDurationMillis() *int64
	SetMaxRoundCount(value *int64)
	SetMinRoundCount(value *int64)
	SetTargetDurationMillis(value *int64)
}
