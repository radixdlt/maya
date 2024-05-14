package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentValidatorSetValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The validator_set property
	validator_set []ActiveValidatorable
}

// NewConsensusManagerFieldCurrentValidatorSetValue instantiates a new ConsensusManagerFieldCurrentValidatorSetValue and sets the default values.
func NewConsensusManagerFieldCurrentValidatorSetValue() *ConsensusManagerFieldCurrentValidatorSetValue {
	m := &ConsensusManagerFieldCurrentValidatorSetValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateConsensusManagerFieldCurrentValidatorSetValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentValidatorSetValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusManagerFieldCurrentValidatorSetValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldCurrentValidatorSetValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentValidatorSetValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["validator_set"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateActiveValidatorFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ActiveValidatorable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ActiveValidatorable)
				}
			}
			m.SetValidatorSet(res)
		}
		return nil
	}
	return res
}

// GetValidatorSet gets the validator_set property value. The validator_set property
// returns a []ActiveValidatorable when successful
func (m *ConsensusManagerFieldCurrentValidatorSetValue) GetValidatorSet() []ActiveValidatorable {
	return m.validator_set
}

// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentValidatorSetValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetValidatorSet() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValidatorSet()))
		for i, v := range m.GetValidatorSet() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("validator_set", cast)
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
func (m *ConsensusManagerFieldCurrentValidatorSetValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetValidatorSet sets the validator_set property value. The validator_set property
func (m *ConsensusManagerFieldCurrentValidatorSetValue) SetValidatorSet(value []ActiveValidatorable) {
	m.validator_set = value
}

type ConsensusManagerFieldCurrentValidatorSetValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValidatorSet() []ActiveValidatorable
	SetValidatorSet(value []ActiveValidatorable)
}
