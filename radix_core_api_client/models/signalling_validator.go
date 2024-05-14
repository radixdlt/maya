package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SignallingValidator struct {
	// A proportion (between 0 and 1) of the total active stake of an entire `current_validator_set` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	active_stake_proportion *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The index property
	index ActiveValidatorIndexable
}

// NewSignallingValidator instantiates a new SignallingValidator and sets the default values.
func NewSignallingValidator() *SignallingValidator {
	m := &SignallingValidator{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSignallingValidatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignallingValidatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSignallingValidator(), nil
}

// GetActiveStakeProportion gets the active_stake_proportion property value. A proportion (between 0 and 1) of the total active stake of an entire `current_validator_set` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SignallingValidator) GetActiveStakeProportion() *string {
	return m.active_stake_proportion
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SignallingValidator) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SignallingValidator) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active_stake_proportion"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActiveStakeProportion(val)
		}
		return nil
	}
	res["index"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateActiveValidatorIndexFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIndex(val.(ActiveValidatorIndexable))
		}
		return nil
	}
	return res
}

// GetIndex gets the index property value. The index property
// returns a ActiveValidatorIndexable when successful
func (m *SignallingValidator) GetIndex() ActiveValidatorIndexable {
	return m.index
}

// Serialize serializes information the current object
func (m *SignallingValidator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("active_stake_proportion", m.GetActiveStakeProportion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("index", m.GetIndex())
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

// SetActiveStakeProportion sets the active_stake_proportion property value. A proportion (between 0 and 1) of the total active stake of an entire `current_validator_set` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SignallingValidator) SetActiveStakeProportion(value *string) {
	m.active_stake_proportion = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignallingValidator) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIndex sets the index property value. The index property
func (m *SignallingValidator) SetIndex(value ActiveValidatorIndexable) {
	m.index = value
}

type SignallingValidatorable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActiveStakeProportion() *string
	GetIndex() ActiveValidatorIndexable
	SetActiveStakeProportion(value *string)
	SetIndex(value ActiveValidatorIndexable)
}
