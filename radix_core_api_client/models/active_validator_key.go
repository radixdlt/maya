package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActiveValidatorKey struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The divided stake, giving a weighting for the validator,used as part of the sort key.
	stake_weighting *int32
	// The Bech32m-encoded human readable version of the component address
	validator_address *string
}

// NewActiveValidatorKey instantiates a new ActiveValidatorKey and sets the default values.
func NewActiveValidatorKey() *ActiveValidatorKey {
	m := &ActiveValidatorKey{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateActiveValidatorKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActiveValidatorKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewActiveValidatorKey(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActiveValidatorKey) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActiveValidatorKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["stake_weighting"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStakeWeighting(val)
		}
		return nil
	}
	res["validator_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValidatorAddress(val)
		}
		return nil
	}
	return res
}

// GetStakeWeighting gets the stake_weighting property value. The divided stake, giving a weighting for the validator,used as part of the sort key.
// returns a *int32 when successful
func (m *ActiveValidatorKey) GetStakeWeighting() *int32 {
	return m.stake_weighting
}

// GetValidatorAddress gets the validator_address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *ActiveValidatorKey) GetValidatorAddress() *string {
	return m.validator_address
}

// Serialize serializes information the current object
func (m *ActiveValidatorKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("stake_weighting", m.GetStakeWeighting())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("validator_address", m.GetValidatorAddress())
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
func (m *ActiveValidatorKey) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetStakeWeighting sets the stake_weighting property value. The divided stake, giving a weighting for the validator,used as part of the sort key.
func (m *ActiveValidatorKey) SetStakeWeighting(value *int32) {
	m.stake_weighting = value
}

// SetValidatorAddress sets the validator_address property value. The Bech32m-encoded human readable version of the component address
func (m *ActiveValidatorKey) SetValidatorAddress(value *string) {
	m.validator_address = value
}

type ActiveValidatorKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetStakeWeighting() *int32
	GetValidatorAddress() *string
	SetStakeWeighting(value *int32)
	SetValidatorAddress(value *string)
}
