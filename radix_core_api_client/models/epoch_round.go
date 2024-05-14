package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EpochRound struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An integer between `0` and `10^10`, marking the epoch.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
	epoch *int64
	// An integer between `0` and `10^10`, marking the current round in an epoch
	round *int64
}

// NewEpochRound instantiates a new EpochRound and sets the default values.
func NewEpochRound() *EpochRound {
	m := &EpochRound{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEpochRoundFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEpochRoundFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEpochRound(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EpochRound) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEpoch gets the epoch property value. An integer between `0` and `10^10`, marking the epoch.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
// returns a *int64 when successful
func (m *EpochRound) GetEpoch() *int64 {
	return m.epoch
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EpochRound) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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

// GetRound gets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
// returns a *int64 when successful
func (m *EpochRound) GetRound() *int64 {
	return m.round
}

// Serialize serializes information the current object
func (m *EpochRound) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("epoch", m.GetEpoch())
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

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EpochRound) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEpoch sets the epoch property value. An integer between `0` and `10^10`, marking the epoch.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
func (m *EpochRound) SetEpoch(value *int64) {
	m.epoch = value
}

// SetRound sets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
func (m *EpochRound) SetRound(value *int64) {
	m.round = value
}

type EpochRoundable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEpoch() *int64
	GetRound() *int64
	SetEpoch(value *int64)
	SetRound(value *int64)
}
