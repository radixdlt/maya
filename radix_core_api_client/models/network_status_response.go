package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NetworkStatusResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The current_epoch_round property
	current_epoch_round EpochRoundable
	// A descriptor for the current protocol version that the node is running.
	current_protocol_version *string
	// The current_state_identifier property
	current_state_identifier CommittedStateIdentifierable
	// The genesis_epoch_round property
	genesis_epoch_round EpochRoundable
	// The post_genesis_epoch_round property
	post_genesis_epoch_round EpochRoundable
	// The post_genesis_state_identifier property
	post_genesis_state_identifier CommittedStateIdentifierable
	// The pre_genesis_state_identifier property
	pre_genesis_state_identifier CommittedStateIdentifierable
}

// NewNetworkStatusResponse instantiates a new NetworkStatusResponse and sets the default values.
func NewNetworkStatusResponse() *NetworkStatusResponse {
	m := &NetworkStatusResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNetworkStatusResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkStatusResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNetworkStatusResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkStatusResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCurrentEpochRound gets the current_epoch_round property value. The current_epoch_round property
// returns a EpochRoundable when successful
func (m *NetworkStatusResponse) GetCurrentEpochRound() EpochRoundable {
	return m.current_epoch_round
}

// GetCurrentProtocolVersion gets the current_protocol_version property value. A descriptor for the current protocol version that the node is running.
// returns a *string when successful
func (m *NetworkStatusResponse) GetCurrentProtocolVersion() *string {
	return m.current_protocol_version
}

// GetCurrentStateIdentifier gets the current_state_identifier property value. The current_state_identifier property
// returns a CommittedStateIdentifierable when successful
func (m *NetworkStatusResponse) GetCurrentStateIdentifier() CommittedStateIdentifierable {
	return m.current_state_identifier
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkStatusResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["current_epoch_round"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEpochRoundFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrentEpochRound(val.(EpochRoundable))
		}
		return nil
	}
	res["current_protocol_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrentProtocolVersion(val)
		}
		return nil
	}
	res["current_state_identifier"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedStateIdentifierFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrentStateIdentifier(val.(CommittedStateIdentifierable))
		}
		return nil
	}
	res["genesis_epoch_round"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEpochRoundFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGenesisEpochRound(val.(EpochRoundable))
		}
		return nil
	}
	res["post_genesis_epoch_round"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEpochRoundFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPostGenesisEpochRound(val.(EpochRoundable))
		}
		return nil
	}
	res["post_genesis_state_identifier"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedStateIdentifierFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPostGenesisStateIdentifier(val.(CommittedStateIdentifierable))
		}
		return nil
	}
	res["pre_genesis_state_identifier"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedStateIdentifierFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreGenesisStateIdentifier(val.(CommittedStateIdentifierable))
		}
		return nil
	}
	return res
}

// GetGenesisEpochRound gets the genesis_epoch_round property value. The genesis_epoch_round property
// returns a EpochRoundable when successful
func (m *NetworkStatusResponse) GetGenesisEpochRound() EpochRoundable {
	return m.genesis_epoch_round
}

// GetPostGenesisEpochRound gets the post_genesis_epoch_round property value. The post_genesis_epoch_round property
// returns a EpochRoundable when successful
func (m *NetworkStatusResponse) GetPostGenesisEpochRound() EpochRoundable {
	return m.post_genesis_epoch_round
}

// GetPostGenesisStateIdentifier gets the post_genesis_state_identifier property value. The post_genesis_state_identifier property
// returns a CommittedStateIdentifierable when successful
func (m *NetworkStatusResponse) GetPostGenesisStateIdentifier() CommittedStateIdentifierable {
	return m.post_genesis_state_identifier
}

// GetPreGenesisStateIdentifier gets the pre_genesis_state_identifier property value. The pre_genesis_state_identifier property
// returns a CommittedStateIdentifierable when successful
func (m *NetworkStatusResponse) GetPreGenesisStateIdentifier() CommittedStateIdentifierable {
	return m.pre_genesis_state_identifier
}

// Serialize serializes information the current object
func (m *NetworkStatusResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("current_epoch_round", m.GetCurrentEpochRound())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("current_protocol_version", m.GetCurrentProtocolVersion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("current_state_identifier", m.GetCurrentStateIdentifier())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("genesis_epoch_round", m.GetGenesisEpochRound())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("post_genesis_epoch_round", m.GetPostGenesisEpochRound())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("post_genesis_state_identifier", m.GetPostGenesisStateIdentifier())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("pre_genesis_state_identifier", m.GetPreGenesisStateIdentifier())
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
func (m *NetworkStatusResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCurrentEpochRound sets the current_epoch_round property value. The current_epoch_round property
func (m *NetworkStatusResponse) SetCurrentEpochRound(value EpochRoundable) {
	m.current_epoch_round = value
}

// SetCurrentProtocolVersion sets the current_protocol_version property value. A descriptor for the current protocol version that the node is running.
func (m *NetworkStatusResponse) SetCurrentProtocolVersion(value *string) {
	m.current_protocol_version = value
}

// SetCurrentStateIdentifier sets the current_state_identifier property value. The current_state_identifier property
func (m *NetworkStatusResponse) SetCurrentStateIdentifier(value CommittedStateIdentifierable) {
	m.current_state_identifier = value
}

// SetGenesisEpochRound sets the genesis_epoch_round property value. The genesis_epoch_round property
func (m *NetworkStatusResponse) SetGenesisEpochRound(value EpochRoundable) {
	m.genesis_epoch_round = value
}

// SetPostGenesisEpochRound sets the post_genesis_epoch_round property value. The post_genesis_epoch_round property
func (m *NetworkStatusResponse) SetPostGenesisEpochRound(value EpochRoundable) {
	m.post_genesis_epoch_round = value
}

// SetPostGenesisStateIdentifier sets the post_genesis_state_identifier property value. The post_genesis_state_identifier property
func (m *NetworkStatusResponse) SetPostGenesisStateIdentifier(value CommittedStateIdentifierable) {
	m.post_genesis_state_identifier = value
}

// SetPreGenesisStateIdentifier sets the pre_genesis_state_identifier property value. The pre_genesis_state_identifier property
func (m *NetworkStatusResponse) SetPreGenesisStateIdentifier(value CommittedStateIdentifierable) {
	m.pre_genesis_state_identifier = value
}

type NetworkStatusResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCurrentEpochRound() EpochRoundable
	GetCurrentProtocolVersion() *string
	GetCurrentStateIdentifier() CommittedStateIdentifierable
	GetGenesisEpochRound() EpochRoundable
	GetPostGenesisEpochRound() EpochRoundable
	GetPostGenesisStateIdentifier() CommittedStateIdentifierable
	GetPreGenesisStateIdentifier() CommittedStateIdentifierable
	SetCurrentEpochRound(value EpochRoundable)
	SetCurrentProtocolVersion(value *string)
	SetCurrentStateIdentifier(value CommittedStateIdentifierable)
	SetGenesisEpochRound(value EpochRoundable)
	SetPostGenesisEpochRound(value EpochRoundable)
	SetPostGenesisStateIdentifier(value CommittedStateIdentifierable)
	SetPreGenesisStateIdentifier(value CommittedStateIdentifierable)
}
