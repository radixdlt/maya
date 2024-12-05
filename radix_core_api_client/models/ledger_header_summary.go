package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerHeaderSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The epoch_round property
    epoch_round EpochRoundable
    // The ledger_hashes property
    ledger_hashes LedgerHashesable
    // The proposer_timestamp property
    proposer_timestamp InstantMsable
}
// NewLedgerHeaderSummary instantiates a new LedgerHeaderSummary and sets the default values.
func NewLedgerHeaderSummary()(*LedgerHeaderSummary) {
    m := &LedgerHeaderSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLedgerHeaderSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerHeaderSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLedgerHeaderSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerHeaderSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEpochRound gets the epoch_round property value. The epoch_round property
// returns a EpochRoundable when successful
func (m *LedgerHeaderSummary) GetEpochRound()(EpochRoundable) {
    return m.epoch_round
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerHeaderSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["epoch_round"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEpochRoundFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpochRound(val.(EpochRoundable))
        }
        return nil
    }
    res["ledger_hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerHashesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLedgerHashes(val.(LedgerHashesable))
        }
        return nil
    }
    res["proposer_timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposerTimestamp(val.(InstantMsable))
        }
        return nil
    }
    return res
}
// GetLedgerHashes gets the ledger_hashes property value. The ledger_hashes property
// returns a LedgerHashesable when successful
func (m *LedgerHeaderSummary) GetLedgerHashes()(LedgerHashesable) {
    return m.ledger_hashes
}
// GetProposerTimestamp gets the proposer_timestamp property value. The proposer_timestamp property
// returns a InstantMsable when successful
func (m *LedgerHeaderSummary) GetProposerTimestamp()(InstantMsable) {
    return m.proposer_timestamp
}
// Serialize serializes information the current object
func (m *LedgerHeaderSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("epoch_round", m.GetEpochRound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ledger_hashes", m.GetLedgerHashes())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LedgerHeaderSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEpochRound sets the epoch_round property value. The epoch_round property
func (m *LedgerHeaderSummary) SetEpochRound(value EpochRoundable)() {
    m.epoch_round = value
}
// SetLedgerHashes sets the ledger_hashes property value. The ledger_hashes property
func (m *LedgerHeaderSummary) SetLedgerHashes(value LedgerHashesable)() {
    m.ledger_hashes = value
}
// SetProposerTimestamp sets the proposer_timestamp property value. The proposer_timestamp property
func (m *LedgerHeaderSummary) SetProposerTimestamp(value InstantMsable)() {
    m.proposer_timestamp = value
}
type LedgerHeaderSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEpochRound()(EpochRoundable)
    GetLedgerHashes()(LedgerHashesable)
    GetProposerTimestamp()(InstantMsable)
    SetEpochRound(value EpochRoundable)()
    SetLedgerHashes(value LedgerHashesable)()
    SetProposerTimestamp(value InstantMsable)()
}
