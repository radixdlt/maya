package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoundUpdateLedgerTransaction struct {
	LedgerTransaction
	// The round_update_transaction property
	round_update_transaction RoundUpdateTransactionable
}

// NewRoundUpdateLedgerTransaction instantiates a new RoundUpdateLedgerTransaction and sets the default values.
func NewRoundUpdateLedgerTransaction() *RoundUpdateLedgerTransaction {
	m := &RoundUpdateLedgerTransaction{
		LedgerTransaction: *NewLedgerTransaction(),
	}
	return m
}

// CreateRoundUpdateLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoundUpdateLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoundUpdateLedgerTransaction(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoundUpdateLedgerTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerTransaction.GetFieldDeserializers()
	res["round_update_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRoundUpdateTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoundUpdateTransaction(val.(RoundUpdateTransactionable))
		}
		return nil
	}
	return res
}

// GetRoundUpdateTransaction gets the round_update_transaction property value. The round_update_transaction property
// returns a RoundUpdateTransactionable when successful
func (m *RoundUpdateLedgerTransaction) GetRoundUpdateTransaction() RoundUpdateTransactionable {
	return m.round_update_transaction
}

// Serialize serializes information the current object
func (m *RoundUpdateLedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerTransaction.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("round_update_transaction", m.GetRoundUpdateTransaction())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRoundUpdateTransaction sets the round_update_transaction property value. The round_update_transaction property
func (m *RoundUpdateLedgerTransaction) SetRoundUpdateTransaction(value RoundUpdateTransactionable) {
	m.round_update_transaction = value
}

type RoundUpdateLedgerTransactionable interface {
	LedgerTransactionable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRoundUpdateTransaction() RoundUpdateTransactionable
	SetRoundUpdateTransaction(value RoundUpdateTransactionable)
}
