package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LeaderProposalHistory struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The current_leader property
	current_leader ActiveValidatorIndexable
	// The validators which were leaders of the "gap" rounds (i.e. since the previous `RoundUpdateValidatorTransaction` - which means that this list will contain exactly `current.round - previous.round - 1` elements). The validators on this list should be penalized during emissions at the end of the epoch.
	gap_round_leaders []ActiveValidatorIndexable
	// Whether the concluded round was conducted in a "fallback" mode (i.e. indicating a fault of the current leader). When `true`, the `current_leader` should be penalized during emissions in the same way as `gap_round_leaders`. When `false`, the `current_leader` is considered to have made this round's proposal successfully.
	is_fallback *bool
}

// NewLeaderProposalHistory instantiates a new LeaderProposalHistory and sets the default values.
func NewLeaderProposalHistory() *LeaderProposalHistory {
	m := &LeaderProposalHistory{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLeaderProposalHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLeaderProposalHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLeaderProposalHistory(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LeaderProposalHistory) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCurrentLeader gets the current_leader property value. The current_leader property
// returns a ActiveValidatorIndexable when successful
func (m *LeaderProposalHistory) GetCurrentLeader() ActiveValidatorIndexable {
	return m.current_leader
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LeaderProposalHistory) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["gap_round_leaders"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateActiveValidatorIndexFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ActiveValidatorIndexable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ActiveValidatorIndexable)
				}
			}
			m.SetGapRoundLeaders(res)
		}
		return nil
	}
	res["is_fallback"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsFallback(val)
		}
		return nil
	}
	return res
}

// GetGapRoundLeaders gets the gap_round_leaders property value. The validators which were leaders of the "gap" rounds (i.e. since the previous `RoundUpdateValidatorTransaction` - which means that this list will contain exactly `current.round - previous.round - 1` elements). The validators on this list should be penalized during emissions at the end of the epoch.
// returns a []ActiveValidatorIndexable when successful
func (m *LeaderProposalHistory) GetGapRoundLeaders() []ActiveValidatorIndexable {
	return m.gap_round_leaders
}

// GetIsFallback gets the is_fallback property value. Whether the concluded round was conducted in a "fallback" mode (i.e. indicating a fault of the current leader). When `true`, the `current_leader` should be penalized during emissions in the same way as `gap_round_leaders`. When `false`, the `current_leader` is considered to have made this round's proposal successfully.
// returns a *bool when successful
func (m *LeaderProposalHistory) GetIsFallback() *bool {
	return m.is_fallback
}

// Serialize serializes information the current object
func (m *LeaderProposalHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("current_leader", m.GetCurrentLeader())
		if err != nil {
			return err
		}
	}
	if m.GetGapRoundLeaders() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGapRoundLeaders()))
		for i, v := range m.GetGapRoundLeaders() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("gap_round_leaders", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_fallback", m.GetIsFallback())
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
func (m *LeaderProposalHistory) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCurrentLeader sets the current_leader property value. The current_leader property
func (m *LeaderProposalHistory) SetCurrentLeader(value ActiveValidatorIndexable) {
	m.current_leader = value
}

// SetGapRoundLeaders sets the gap_round_leaders property value. The validators which were leaders of the "gap" rounds (i.e. since the previous `RoundUpdateValidatorTransaction` - which means that this list will contain exactly `current.round - previous.round - 1` elements). The validators on this list should be penalized during emissions at the end of the epoch.
func (m *LeaderProposalHistory) SetGapRoundLeaders(value []ActiveValidatorIndexable) {
	m.gap_round_leaders = value
}

// SetIsFallback sets the is_fallback property value. Whether the concluded round was conducted in a "fallback" mode (i.e. indicating a fault of the current leader). When `true`, the `current_leader` should be penalized during emissions in the same way as `gap_round_leaders`. When `false`, the `current_leader` is considered to have made this round's proposal successfully.
func (m *LeaderProposalHistory) SetIsFallback(value *bool) {
	m.is_fallback = value
}

type LeaderProposalHistoryable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCurrentLeader() ActiveValidatorIndexable
	GetGapRoundLeaders() []ActiveValidatorIndexable
	GetIsFallback() *bool
	SetCurrentLeader(value ActiveValidatorIndexable)
	SetGapRoundLeaders(value []ActiveValidatorIndexable)
	SetIsFallback(value *bool)
}
