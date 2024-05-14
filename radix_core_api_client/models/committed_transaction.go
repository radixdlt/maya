package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommittedTransaction struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The balance_changes property
	balance_changes CommittedTransactionBalanceChangesable
	// The ledger_transaction property
	ledger_transaction LedgerTransactionable
	// An integer between `0` and `10^14`, marking the proposer timestamp in ms.
	proposer_timestamp_ms *int64
	// The transaction execution receipt
	receipt TransactionReceiptable
	// The resultant_state_identifiers property
	resultant_state_identifiers CommittedStateIdentifierable
}

// NewCommittedTransaction instantiates a new CommittedTransaction and sets the default values.
func NewCommittedTransaction() *CommittedTransaction {
	m := &CommittedTransaction{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCommittedTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommittedTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCommittedTransaction(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommittedTransaction) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBalanceChanges gets the balance_changes property value. The balance_changes property
// returns a CommittedTransactionBalanceChangesable when successful
func (m *CommittedTransaction) GetBalanceChanges() CommittedTransactionBalanceChangesable {
	return m.balance_changes
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommittedTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedTransactionBalanceChangesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBalanceChanges(val.(CommittedTransactionBalanceChangesable))
		}
		return nil
	}
	res["ledger_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLedgerTransaction(val.(LedgerTransactionable))
		}
		return nil
	}
	res["proposer_timestamp_ms"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProposerTimestampMs(val)
		}
		return nil
	}
	res["receipt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionReceiptFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReceipt(val.(TransactionReceiptable))
		}
		return nil
	}
	res["resultant_state_identifiers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedStateIdentifierFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResultantStateIdentifiers(val.(CommittedStateIdentifierable))
		}
		return nil
	}
	return res
}

// GetLedgerTransaction gets the ledger_transaction property value. The ledger_transaction property
// returns a LedgerTransactionable when successful
func (m *CommittedTransaction) GetLedgerTransaction() LedgerTransactionable {
	return m.ledger_transaction
}

// GetProposerTimestampMs gets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
// returns a *int64 when successful
func (m *CommittedTransaction) GetProposerTimestampMs() *int64 {
	return m.proposer_timestamp_ms
}

// GetReceipt gets the receipt property value. The transaction execution receipt
// returns a TransactionReceiptable when successful
func (m *CommittedTransaction) GetReceipt() TransactionReceiptable {
	return m.receipt
}

// GetResultantStateIdentifiers gets the resultant_state_identifiers property value. The resultant_state_identifiers property
// returns a CommittedStateIdentifierable when successful
func (m *CommittedTransaction) GetResultantStateIdentifiers() CommittedStateIdentifierable {
	return m.resultant_state_identifiers
}

// Serialize serializes information the current object
func (m *CommittedTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("balance_changes", m.GetBalanceChanges())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("ledger_transaction", m.GetLedgerTransaction())
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
		err := writer.WriteObjectValue("receipt", m.GetReceipt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("resultant_state_identifiers", m.GetResultantStateIdentifiers())
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
func (m *CommittedTransaction) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBalanceChanges sets the balance_changes property value. The balance_changes property
func (m *CommittedTransaction) SetBalanceChanges(value CommittedTransactionBalanceChangesable) {
	m.balance_changes = value
}

// SetLedgerTransaction sets the ledger_transaction property value. The ledger_transaction property
func (m *CommittedTransaction) SetLedgerTransaction(value LedgerTransactionable) {
	m.ledger_transaction = value
}

// SetProposerTimestampMs sets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
func (m *CommittedTransaction) SetProposerTimestampMs(value *int64) {
	m.proposer_timestamp_ms = value
}

// SetReceipt sets the receipt property value. The transaction execution receipt
func (m *CommittedTransaction) SetReceipt(value TransactionReceiptable) {
	m.receipt = value
}

// SetResultantStateIdentifiers sets the resultant_state_identifiers property value. The resultant_state_identifiers property
func (m *CommittedTransaction) SetResultantStateIdentifiers(value CommittedStateIdentifierable) {
	m.resultant_state_identifiers = value
}

type CommittedTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBalanceChanges() CommittedTransactionBalanceChangesable
	GetLedgerTransaction() LedgerTransactionable
	GetProposerTimestampMs() *int64
	GetReceipt() TransactionReceiptable
	GetResultantStateIdentifiers() CommittedStateIdentifierable
	SetBalanceChanges(value CommittedTransactionBalanceChangesable)
	SetLedgerTransaction(value LedgerTransactionable)
	SetProposerTimestampMs(value *int64)
	SetReceipt(value TransactionReceiptable)
	SetResultantStateIdentifiers(value CommittedStateIdentifierable)
}
