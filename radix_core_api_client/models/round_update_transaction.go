package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoundUpdateTransaction struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An integer between `0` and `10^10`, marking the epoch.
	epoch *int64
	// The leader_proposal_history property
	leader_proposal_history LeaderProposalHistoryable
	// The proposer_timestamp property
	proposer_timestamp InstantMsable
	// An integer between `0` and `10^10`, marking the consensus round in the epoch
	round_in_epoch *int64
}

// NewRoundUpdateTransaction instantiates a new RoundUpdateTransaction and sets the default values.
func NewRoundUpdateTransaction() *RoundUpdateTransaction {
	m := &RoundUpdateTransaction{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateRoundUpdateTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoundUpdateTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoundUpdateTransaction(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoundUpdateTransaction) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEpoch gets the epoch property value. An integer between `0` and `10^10`, marking the epoch.
// returns a *int64 when successful
func (m *RoundUpdateTransaction) GetEpoch() *int64 {
	return m.epoch
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoundUpdateTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["leader_proposal_history"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLeaderProposalHistoryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLeaderProposalHistory(val.(LeaderProposalHistoryable))
		}
		return nil
	}
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
	res["round_in_epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoundInEpoch(val)
		}
		return nil
	}
	return res
}

// GetLeaderProposalHistory gets the leader_proposal_history property value. The leader_proposal_history property
// returns a LeaderProposalHistoryable when successful
func (m *RoundUpdateTransaction) GetLeaderProposalHistory() LeaderProposalHistoryable {
	return m.leader_proposal_history
}

// GetProposerTimestamp gets the proposer_timestamp property value. The proposer_timestamp property
// returns a InstantMsable when successful
func (m *RoundUpdateTransaction) GetProposerTimestamp() InstantMsable {
	return m.proposer_timestamp
}

// GetRoundInEpoch gets the round_in_epoch property value. An integer between `0` and `10^10`, marking the consensus round in the epoch
// returns a *int64 when successful
func (m *RoundUpdateTransaction) GetRoundInEpoch() *int64 {
	return m.round_in_epoch
}

// Serialize serializes information the current object
func (m *RoundUpdateTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("epoch", m.GetEpoch())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("leader_proposal_history", m.GetLeaderProposalHistory())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("proposer_timestamp", m.GetProposerTimestamp())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("round_in_epoch", m.GetRoundInEpoch())
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
func (m *RoundUpdateTransaction) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEpoch sets the epoch property value. An integer between `0` and `10^10`, marking the epoch.
func (m *RoundUpdateTransaction) SetEpoch(value *int64) {
	m.epoch = value
}

// SetLeaderProposalHistory sets the leader_proposal_history property value. The leader_proposal_history property
func (m *RoundUpdateTransaction) SetLeaderProposalHistory(value LeaderProposalHistoryable) {
	m.leader_proposal_history = value
}

// SetProposerTimestamp sets the proposer_timestamp property value. The proposer_timestamp property
func (m *RoundUpdateTransaction) SetProposerTimestamp(value InstantMsable) {
	m.proposer_timestamp = value
}

// SetRoundInEpoch sets the round_in_epoch property value. An integer between `0` and `10^10`, marking the consensus round in the epoch
func (m *RoundUpdateTransaction) SetRoundInEpoch(value *int64) {
	m.round_in_epoch = value
}

type RoundUpdateTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEpoch() *int64
	GetLeaderProposalHistory() LeaderProposalHistoryable
	GetProposerTimestamp() InstantMsable
	GetRoundInEpoch() *int64
	SetEpoch(value *int64)
	SetLeaderProposalHistory(value LeaderProposalHistoryable)
	SetProposerTimestamp(value InstantMsable)
	SetRoundInEpoch(value *int64)
}
