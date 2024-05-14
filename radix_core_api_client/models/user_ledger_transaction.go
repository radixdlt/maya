package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserLedgerTransaction struct {
	LedgerTransaction
	// The notarized_transaction property
	notarized_transaction NotarizedTransactionable
}

// NewUserLedgerTransaction instantiates a new UserLedgerTransaction and sets the default values.
func NewUserLedgerTransaction() *UserLedgerTransaction {
	m := &UserLedgerTransaction{
		LedgerTransaction: *NewLedgerTransaction(),
	}
	return m
}

// CreateUserLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUserLedgerTransaction(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserLedgerTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerTransaction.GetFieldDeserializers()
	res["notarized_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNotarizedTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotarizedTransaction(val.(NotarizedTransactionable))
		}
		return nil
	}
	return res
}

// GetNotarizedTransaction gets the notarized_transaction property value. The notarized_transaction property
// returns a NotarizedTransactionable when successful
func (m *UserLedgerTransaction) GetNotarizedTransaction() NotarizedTransactionable {
	return m.notarized_transaction
}

// Serialize serializes information the current object
func (m *UserLedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerTransaction.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("notarized_transaction", m.GetNotarizedTransaction())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetNotarizedTransaction sets the notarized_transaction property value. The notarized_transaction property
func (m *UserLedgerTransaction) SetNotarizedTransaction(value NotarizedTransactionable) {
	m.notarized_transaction = value
}

type UserLedgerTransactionable interface {
	LedgerTransactionable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNotarizedTransaction() NotarizedTransactionable
	SetNotarizedTransaction(value NotarizedTransactionable)
}
