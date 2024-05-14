package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateConsensusManagerRequest struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Whether to include protocol update readiness signals of active validator set (default false).
	include_readiness_signals *bool
	// The logical name of the network
	network *string
}

// NewStateConsensusManagerRequest instantiates a new StateConsensusManagerRequest and sets the default values.
func NewStateConsensusManagerRequest() *StateConsensusManagerRequest {
	m := &StateConsensusManagerRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateConsensusManagerRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateConsensusManagerRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateConsensusManagerRequest(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateConsensusManagerRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateConsensusManagerRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["include_readiness_signals"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIncludeReadinessSignals(val)
		}
		return nil
	}
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
		}
		return nil
	}
	return res
}

// GetIncludeReadinessSignals gets the include_readiness_signals property value. Whether to include protocol update readiness signals of active validator set (default false).
// returns a *bool when successful
func (m *StateConsensusManagerRequest) GetIncludeReadinessSignals() *bool {
	return m.include_readiness_signals
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StateConsensusManagerRequest) GetNetwork() *string {
	return m.network
}

// Serialize serializes information the current object
func (m *StateConsensusManagerRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("include_readiness_signals", m.GetIncludeReadinessSignals())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
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
func (m *StateConsensusManagerRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIncludeReadinessSignals sets the include_readiness_signals property value. Whether to include protocol update readiness signals of active validator set (default false).
func (m *StateConsensusManagerRequest) SetIncludeReadinessSignals(value *bool) {
	m.include_readiness_signals = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *StateConsensusManagerRequest) SetNetwork(value *string) {
	m.network = value
}

type StateConsensusManagerRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIncludeReadinessSignals() *bool
	GetNetwork() *string
	SetIncludeReadinessSignals(value *bool)
	SetNetwork(value *string)
}
