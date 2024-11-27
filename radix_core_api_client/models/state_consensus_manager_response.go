package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateConsensusManagerResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The at_ledger_state property
    at_ledger_state LedgerStateSummaryable
    // The config property
    config Substateable
    // The current_proposal_statistic property
    current_proposal_statistic Substateable
    // The current_time property
    current_time Substateable
    // The current_time_rounded_to_minutes property
    current_time_rounded_to_minutes Substateable
    // Protocol versions signalled by the current validator set.Every validator from `current_validator_set` will be referenced by exactly one of the items here.Only returned if enabled by `include_readiness_signals` on your request.
    current_validator_readiness_signals []ProtocolVersionReadinessable
    // The current_validator_set property
    current_validator_set Substateable
    // The state property
    state Substateable
}
// NewStateConsensusManagerResponse instantiates a new StateConsensusManagerResponse and sets the default values.
func NewStateConsensusManagerResponse()(*StateConsensusManagerResponse) {
    m := &StateConsensusManagerResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateConsensusManagerResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateConsensusManagerResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateConsensusManagerResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateConsensusManagerResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateConsensusManagerResponse) GetAtLedgerState()(LedgerStateSummaryable) {
    return m.at_ledger_state
}
// GetConfig gets the config property value. The config property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetConfig()(Substateable) {
    return m.config
}
// GetCurrentProposalStatistic gets the current_proposal_statistic property value. The current_proposal_statistic property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetCurrentProposalStatistic()(Substateable) {
    return m.current_proposal_statistic
}
// GetCurrentTime gets the current_time property value. The current_time property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetCurrentTime()(Substateable) {
    return m.current_time
}
// GetCurrentTimeRoundedToMinutes gets the current_time_rounded_to_minutes property value. The current_time_rounded_to_minutes property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetCurrentTimeRoundedToMinutes()(Substateable) {
    return m.current_time_rounded_to_minutes
}
// GetCurrentValidatorReadinessSignals gets the current_validator_readiness_signals property value. Protocol versions signalled by the current validator set.Every validator from `current_validator_set` will be referenced by exactly one of the items here.Only returned if enabled by `include_readiness_signals` on your request.
// returns a []ProtocolVersionReadinessable when successful
func (m *StateConsensusManagerResponse) GetCurrentValidatorReadinessSignals()([]ProtocolVersionReadinessable) {
    return m.current_validator_readiness_signals
}
// GetCurrentValidatorSet gets the current_validator_set property value. The current_validator_set property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetCurrentValidatorSet()(Substateable) {
    return m.current_validator_set
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateConsensusManagerResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["at_ledger_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtLedgerState(val.(LedgerStateSummaryable))
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(Substateable))
        }
        return nil
    }
    res["current_proposal_statistic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentProposalStatistic(val.(Substateable))
        }
        return nil
    }
    res["current_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentTime(val.(Substateable))
        }
        return nil
    }
    res["current_time_rounded_to_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentTimeRoundedToMinutes(val.(Substateable))
        }
        return nil
    }
    res["current_validator_readiness_signals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProtocolVersionReadinessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProtocolVersionReadinessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProtocolVersionReadinessable)
                }
            }
            m.SetCurrentValidatorReadinessSignals(res)
        }
        return nil
    }
    res["current_validator_set"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValidatorSet(val.(Substateable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(Substateable))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
// returns a Substateable when successful
func (m *StateConsensusManagerResponse) GetState()(Substateable) {
    return m.state
}
// Serialize serializes information the current object
func (m *StateConsensusManagerResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_proposal_statistic", m.GetCurrentProposalStatistic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_time", m.GetCurrentTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_time_rounded_to_minutes", m.GetCurrentTimeRoundedToMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentValidatorReadinessSignals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCurrentValidatorReadinessSignals()))
        for i, v := range m.GetCurrentValidatorReadinessSignals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("current_validator_readiness_signals", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_validator_set", m.GetCurrentValidatorSet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("state", m.GetState())
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
func (m *StateConsensusManagerResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateConsensusManagerResponse) SetAtLedgerState(value LedgerStateSummaryable)() {
    m.at_ledger_state = value
}
// SetConfig sets the config property value. The config property
func (m *StateConsensusManagerResponse) SetConfig(value Substateable)() {
    m.config = value
}
// SetCurrentProposalStatistic sets the current_proposal_statistic property value. The current_proposal_statistic property
func (m *StateConsensusManagerResponse) SetCurrentProposalStatistic(value Substateable)() {
    m.current_proposal_statistic = value
}
// SetCurrentTime sets the current_time property value. The current_time property
func (m *StateConsensusManagerResponse) SetCurrentTime(value Substateable)() {
    m.current_time = value
}
// SetCurrentTimeRoundedToMinutes sets the current_time_rounded_to_minutes property value. The current_time_rounded_to_minutes property
func (m *StateConsensusManagerResponse) SetCurrentTimeRoundedToMinutes(value Substateable)() {
    m.current_time_rounded_to_minutes = value
}
// SetCurrentValidatorReadinessSignals sets the current_validator_readiness_signals property value. Protocol versions signalled by the current validator set.Every validator from `current_validator_set` will be referenced by exactly one of the items here.Only returned if enabled by `include_readiness_signals` on your request.
func (m *StateConsensusManagerResponse) SetCurrentValidatorReadinessSignals(value []ProtocolVersionReadinessable)() {
    m.current_validator_readiness_signals = value
}
// SetCurrentValidatorSet sets the current_validator_set property value. The current_validator_set property
func (m *StateConsensusManagerResponse) SetCurrentValidatorSet(value Substateable)() {
    m.current_validator_set = value
}
// SetState sets the state property value. The state property
func (m *StateConsensusManagerResponse) SetState(value Substateable)() {
    m.state = value
}
type StateConsensusManagerResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSummaryable)
    GetConfig()(Substateable)
    GetCurrentProposalStatistic()(Substateable)
    GetCurrentTime()(Substateable)
    GetCurrentTimeRoundedToMinutes()(Substateable)
    GetCurrentValidatorReadinessSignals()([]ProtocolVersionReadinessable)
    GetCurrentValidatorSet()(Substateable)
    GetState()(Substateable)
    SetAtLedgerState(value LedgerStateSummaryable)()
    SetConfig(value Substateable)()
    SetCurrentProposalStatistic(value Substateable)()
    SetCurrentTime(value Substateable)()
    SetCurrentTimeRoundedToMinutes(value Substateable)()
    SetCurrentValidatorReadinessSignals(value []ProtocolVersionReadinessable)()
    SetCurrentValidatorSet(value Substateable)()
    SetState(value Substateable)()
}
