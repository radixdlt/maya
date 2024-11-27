package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStreamAccountTransactionOutcomesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A committed transaction outcomes list starting from the `from_state_version` (inclusive).
    committed_transaction_outcomes []LtsCommittedTransactionOutcomeable
    // An integer between `0` and `10000`, giving the total count of transactions in the returned response
    count *int32
    // The from_state_version property
    from_state_version *int64
    // The max_ledger_state_version property
    max_ledger_state_version *int64
}
// NewLtsStreamAccountTransactionOutcomesResponse instantiates a new LtsStreamAccountTransactionOutcomesResponse and sets the default values.
func NewLtsStreamAccountTransactionOutcomesResponse()(*LtsStreamAccountTransactionOutcomesResponse) {
    m := &LtsStreamAccountTransactionOutcomesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsStreamAccountTransactionOutcomesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStreamAccountTransactionOutcomesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsStreamAccountTransactionOutcomesResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommittedTransactionOutcomes gets the committed_transaction_outcomes property value. A committed transaction outcomes list starting from the `from_state_version` (inclusive).
// returns a []LtsCommittedTransactionOutcomeable when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetCommittedTransactionOutcomes()([]LtsCommittedTransactionOutcomeable) {
    return m.committed_transaction_outcomes
}
// GetCount gets the count property value. An integer between `0` and `10000`, giving the total count of transactions in the returned response
// returns a *int32 when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["committed_transaction_outcomes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLtsCommittedTransactionOutcomeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LtsCommittedTransactionOutcomeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LtsCommittedTransactionOutcomeable)
                }
            }
            m.SetCommittedTransactionOutcomes(res)
        }
        return nil
    }
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["from_state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromStateVersion(val)
        }
        return nil
    }
    res["max_ledger_state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLedgerStateVersion(val)
        }
        return nil
    }
    return res
}
// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetFromStateVersion()(*int64) {
    return m.from_state_version
}
// GetMaxLedgerStateVersion gets the max_ledger_state_version property value. The max_ledger_state_version property
// returns a *int64 when successful
func (m *LtsStreamAccountTransactionOutcomesResponse) GetMaxLedgerStateVersion()(*int64) {
    return m.max_ledger_state_version
}
// Serialize serializes information the current object
func (m *LtsStreamAccountTransactionOutcomesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommittedTransactionOutcomes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommittedTransactionOutcomes()))
        for i, v := range m.GetCommittedTransactionOutcomes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("committed_transaction_outcomes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("from_state_version", m.GetFromStateVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("max_ledger_state_version", m.GetMaxLedgerStateVersion())
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
func (m *LtsStreamAccountTransactionOutcomesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommittedTransactionOutcomes sets the committed_transaction_outcomes property value. A committed transaction outcomes list starting from the `from_state_version` (inclusive).
func (m *LtsStreamAccountTransactionOutcomesResponse) SetCommittedTransactionOutcomes(value []LtsCommittedTransactionOutcomeable)() {
    m.committed_transaction_outcomes = value
}
// SetCount sets the count property value. An integer between `0` and `10000`, giving the total count of transactions in the returned response
func (m *LtsStreamAccountTransactionOutcomesResponse) SetCount(value *int32)() {
    m.count = value
}
// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *LtsStreamAccountTransactionOutcomesResponse) SetFromStateVersion(value *int64)() {
    m.from_state_version = value
}
// SetMaxLedgerStateVersion sets the max_ledger_state_version property value. The max_ledger_state_version property
func (m *LtsStreamAccountTransactionOutcomesResponse) SetMaxLedgerStateVersion(value *int64)() {
    m.max_ledger_state_version = value
}
type LtsStreamAccountTransactionOutcomesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommittedTransactionOutcomes()([]LtsCommittedTransactionOutcomeable)
    GetCount()(*int32)
    GetFromStateVersion()(*int64)
    GetMaxLedgerStateVersion()(*int64)
    SetCommittedTransactionOutcomes(value []LtsCommittedTransactionOutcomeable)()
    SetCount(value *int32)()
    SetFromStateVersion(value *int64)()
    SetMaxLedgerStateVersion(value *int64)()
}
