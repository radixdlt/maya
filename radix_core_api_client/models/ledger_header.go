package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerHeader struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^14`, marking the consensus parent round timestamp in ms.
    consensus_parent_round_timestamp_ms *int64
    // An integer between `0` and `10^10`, marking the epoch.
    epoch *int64
    // The hashes property
    hashes LedgerHashesable
    // The next_epoch property
    next_epoch NextEpochable
    // If present, indicates that this proof triggers the enactment of the given protocol version.
    next_protocol_version *string
    // An integer between `0` and `10^14`, marking the proposer timestamp in ms.
    proposer_timestamp_ms *int64
    // An integer between `0` and `10^10`, marking the current round in an epoch
    round *int64
    // The state_version property
    state_version *int64
}
// NewLedgerHeader instantiates a new LedgerHeader and sets the default values.
func NewLedgerHeader()(*LedgerHeader) {
    m := &LedgerHeader{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLedgerHeaderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerHeaderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLedgerHeader(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerHeader) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsensusParentRoundTimestampMs gets the consensus_parent_round_timestamp_ms property value. An integer between `0` and `10^14`, marking the consensus parent round timestamp in ms.
// returns a *int64 when successful
func (m *LedgerHeader) GetConsensusParentRoundTimestampMs()(*int64) {
    return m.consensus_parent_round_timestamp_ms
}
// GetEpoch gets the epoch property value. An integer between `0` and `10^10`, marking the epoch.
// returns a *int64 when successful
func (m *LedgerHeader) GetEpoch()(*int64) {
    return m.epoch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerHeader) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["consensus_parent_round_timestamp_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsensusParentRoundTimestampMs(val)
        }
        return nil
    }
    res["epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpoch(val)
        }
        return nil
    }
    res["hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerHashesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashes(val.(LedgerHashesable))
        }
        return nil
    }
    res["next_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNextEpochFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextEpoch(val.(NextEpochable))
        }
        return nil
    }
    res["next_protocol_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextProtocolVersion(val)
        }
        return nil
    }
    res["proposer_timestamp_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposerTimestampMs(val)
        }
        return nil
    }
    res["round"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRound(val)
        }
        return nil
    }
    res["state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateVersion(val)
        }
        return nil
    }
    return res
}
// GetHashes gets the hashes property value. The hashes property
// returns a LedgerHashesable when successful
func (m *LedgerHeader) GetHashes()(LedgerHashesable) {
    return m.hashes
}
// GetNextEpoch gets the next_epoch property value. The next_epoch property
// returns a NextEpochable when successful
func (m *LedgerHeader) GetNextEpoch()(NextEpochable) {
    return m.next_epoch
}
// GetNextProtocolVersion gets the next_protocol_version property value. If present, indicates that this proof triggers the enactment of the given protocol version.
// returns a *string when successful
func (m *LedgerHeader) GetNextProtocolVersion()(*string) {
    return m.next_protocol_version
}
// GetProposerTimestampMs gets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
// returns a *int64 when successful
func (m *LedgerHeader) GetProposerTimestampMs()(*int64) {
    return m.proposer_timestamp_ms
}
// GetRound gets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
// returns a *int64 when successful
func (m *LedgerHeader) GetRound()(*int64) {
    return m.round
}
// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LedgerHeader) GetStateVersion()(*int64) {
    return m.state_version
}
// Serialize serializes information the current object
func (m *LedgerHeader) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("consensus_parent_round_timestamp_ms", m.GetConsensusParentRoundTimestampMs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("epoch", m.GetEpoch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hashes", m.GetHashes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("next_epoch", m.GetNextEpoch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("next_protocol_version", m.GetNextProtocolVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("proposer_timestamp_ms", m.GetProposerTimestampMs())
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
        err := writer.WriteInt64Value("state_version", m.GetStateVersion())
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
func (m *LedgerHeader) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsensusParentRoundTimestampMs sets the consensus_parent_round_timestamp_ms property value. An integer between `0` and `10^14`, marking the consensus parent round timestamp in ms.
func (m *LedgerHeader) SetConsensusParentRoundTimestampMs(value *int64)() {
    m.consensus_parent_round_timestamp_ms = value
}
// SetEpoch sets the epoch property value. An integer between `0` and `10^10`, marking the epoch.
func (m *LedgerHeader) SetEpoch(value *int64)() {
    m.epoch = value
}
// SetHashes sets the hashes property value. The hashes property
func (m *LedgerHeader) SetHashes(value LedgerHashesable)() {
    m.hashes = value
}
// SetNextEpoch sets the next_epoch property value. The next_epoch property
func (m *LedgerHeader) SetNextEpoch(value NextEpochable)() {
    m.next_epoch = value
}
// SetNextProtocolVersion sets the next_protocol_version property value. If present, indicates that this proof triggers the enactment of the given protocol version.
func (m *LedgerHeader) SetNextProtocolVersion(value *string)() {
    m.next_protocol_version = value
}
// SetProposerTimestampMs sets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
func (m *LedgerHeader) SetProposerTimestampMs(value *int64)() {
    m.proposer_timestamp_ms = value
}
// SetRound sets the round property value. An integer between `0` and `10^10`, marking the current round in an epoch
func (m *LedgerHeader) SetRound(value *int64)() {
    m.round = value
}
// SetStateVersion sets the state_version property value. The state_version property
func (m *LedgerHeader) SetStateVersion(value *int64)() {
    m.state_version = value
}
type LedgerHeaderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsensusParentRoundTimestampMs()(*int64)
    GetEpoch()(*int64)
    GetHashes()(LedgerHashesable)
    GetNextEpoch()(NextEpochable)
    GetNextProtocolVersion()(*string)
    GetProposerTimestampMs()(*int64)
    GetRound()(*int64)
    GetStateVersion()(*int64)
    SetConsensusParentRoundTimestampMs(value *int64)()
    SetEpoch(value *int64)()
    SetHashes(value LedgerHashesable)()
    SetNextEpoch(value NextEpochable)()
    SetNextProtocolVersion(value *string)()
    SetProposerTimestampMs(value *int64)()
    SetRound(value *int64)()
    SetStateVersion(value *int64)()
}
