package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedLedgerTransaction struct {
	ParsedTransaction
	// The identifiers property
	identifiers ParsedLedgerTransactionIdentifiersable
	// The ledger_transaction property
	ledger_transaction LedgerTransactionable
}

// NewParsedLedgerTransaction instantiates a new ParsedLedgerTransaction and sets the default values.
func NewParsedLedgerTransaction() *ParsedLedgerTransaction {
	m := &ParsedLedgerTransaction{
		ParsedTransaction: *NewParsedTransaction(),
	}
	return m
}

// CreateParsedLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewParsedLedgerTransaction(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedLedgerTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ParsedTransaction.GetFieldDeserializers()
	res["identifiers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateParsedLedgerTransactionIdentifiersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIdentifiers(val.(ParsedLedgerTransactionIdentifiersable))
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
	return res
}

// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a ParsedLedgerTransactionIdentifiersable when successful
func (m *ParsedLedgerTransaction) GetIdentifiers() ParsedLedgerTransactionIdentifiersable {
	return m.identifiers
}

// GetLedgerTransaction gets the ledger_transaction property value. The ledger_transaction property
// returns a LedgerTransactionable when successful
func (m *ParsedLedgerTransaction) GetLedgerTransaction() LedgerTransactionable {
	return m.ledger_transaction
}

// Serialize serializes information the current object
func (m *ParsedLedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ParsedTransaction.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("identifiers", m.GetIdentifiers())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("ledger_transaction", m.GetLedgerTransaction())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *ParsedLedgerTransaction) SetIdentifiers(value ParsedLedgerTransactionIdentifiersable) {
	m.identifiers = value
}

// SetLedgerTransaction sets the ledger_transaction property value. The ledger_transaction property
func (m *ParsedLedgerTransaction) SetLedgerTransaction(value LedgerTransactionable) {
	m.ledger_transaction = value
}

type ParsedLedgerTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ParsedTransactionable
	GetIdentifiers() ParsedLedgerTransactionIdentifiersable
	GetLedgerTransaction() LedgerTransactionable
	SetIdentifiers(value ParsedLedgerTransactionIdentifiersable)
	SetLedgerTransaction(value LedgerTransactionable)
}
