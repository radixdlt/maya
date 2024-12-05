package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamTransactionsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10000`, giving the total count of transactions in the returned response
    count *int32
    // The from_state_version property
    from_state_version *int64
    // The max_ledger_state_version property
    max_ledger_state_version *int64
    // The previous_state_identifiers property
    previous_state_identifiers CommittedStateIdentifierable
    // A ledger proof list starting from `from_state_version` (inclusive) stored by this node.
    proofs []LedgerProofable
    // A committed transactions list starting from the `from_state_version` (inclusive).
    transactions []CommittedTransactionable
}
// NewStreamTransactionsResponse instantiates a new StreamTransactionsResponse and sets the default values.
func NewStreamTransactionsResponse()(*StreamTransactionsResponse) {
    m := &StreamTransactionsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStreamTransactionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamTransactionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamTransactionsResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StreamTransactionsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. An integer between `0` and `10000`, giving the total count of transactions in the returned response
// returns a *int32 when successful
func (m *StreamTransactionsResponse) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamTransactionsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["previous_state_identifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommittedStateIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousStateIdentifiers(val.(CommittedStateIdentifierable))
        }
        return nil
    }
    res["proofs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLedgerProofFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LedgerProofable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LedgerProofable)
                }
            }
            m.SetProofs(res)
        }
        return nil
    }
    res["transactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommittedTransactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommittedTransactionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CommittedTransactionable)
                }
            }
            m.SetTransactions(res)
        }
        return nil
    }
    return res
}
// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *StreamTransactionsResponse) GetFromStateVersion()(*int64) {
    return m.from_state_version
}
// GetMaxLedgerStateVersion gets the max_ledger_state_version property value. The max_ledger_state_version property
// returns a *int64 when successful
func (m *StreamTransactionsResponse) GetMaxLedgerStateVersion()(*int64) {
    return m.max_ledger_state_version
}
// GetPreviousStateIdentifiers gets the previous_state_identifiers property value. The previous_state_identifiers property
// returns a CommittedStateIdentifierable when successful
func (m *StreamTransactionsResponse) GetPreviousStateIdentifiers()(CommittedStateIdentifierable) {
    return m.previous_state_identifiers
}
// GetProofs gets the proofs property value. A ledger proof list starting from `from_state_version` (inclusive) stored by this node.
// returns a []LedgerProofable when successful
func (m *StreamTransactionsResponse) GetProofs()([]LedgerProofable) {
    return m.proofs
}
// GetTransactions gets the transactions property value. A committed transactions list starting from the `from_state_version` (inclusive).
// returns a []CommittedTransactionable when successful
func (m *StreamTransactionsResponse) GetTransactions()([]CommittedTransactionable) {
    return m.transactions
}
// Serialize serializes information the current object
func (m *StreamTransactionsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("previous_state_identifiers", m.GetPreviousStateIdentifiers())
        if err != nil {
            return err
        }
    }
    if m.GetProofs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProofs()))
        for i, v := range m.GetProofs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("proofs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransactions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransactions()))
        for i, v := range m.GetTransactions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("transactions", cast)
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
func (m *StreamTransactionsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. An integer between `0` and `10000`, giving the total count of transactions in the returned response
func (m *StreamTransactionsResponse) SetCount(value *int32)() {
    m.count = value
}
// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *StreamTransactionsResponse) SetFromStateVersion(value *int64)() {
    m.from_state_version = value
}
// SetMaxLedgerStateVersion sets the max_ledger_state_version property value. The max_ledger_state_version property
func (m *StreamTransactionsResponse) SetMaxLedgerStateVersion(value *int64)() {
    m.max_ledger_state_version = value
}
// SetPreviousStateIdentifiers sets the previous_state_identifiers property value. The previous_state_identifiers property
func (m *StreamTransactionsResponse) SetPreviousStateIdentifiers(value CommittedStateIdentifierable)() {
    m.previous_state_identifiers = value
}
// SetProofs sets the proofs property value. A ledger proof list starting from `from_state_version` (inclusive) stored by this node.
func (m *StreamTransactionsResponse) SetProofs(value []LedgerProofable)() {
    m.proofs = value
}
// SetTransactions sets the transactions property value. A committed transactions list starting from the `from_state_version` (inclusive).
func (m *StreamTransactionsResponse) SetTransactions(value []CommittedTransactionable)() {
    m.transactions = value
}
type StreamTransactionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetFromStateVersion()(*int64)
    GetMaxLedgerStateVersion()(*int64)
    GetPreviousStateIdentifiers()(CommittedStateIdentifierable)
    GetProofs()([]LedgerProofable)
    GetTransactions()([]CommittedTransactionable)
    SetCount(value *int32)()
    SetFromStateVersion(value *int64)()
    SetMaxLedgerStateVersion(value *int64)()
    SetPreviousStateIdentifiers(value CommittedStateIdentifierable)()
    SetProofs(value []LedgerProofable)()
    SetTransactions(value []CommittedTransactionable)()
}
