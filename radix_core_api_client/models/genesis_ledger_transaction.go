package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenesisLedgerTransaction struct {
	LedgerTransaction
	// The first genesis "transaction" flashes state into the database to prepare for the bootstrap transaction.Such a transaction does not have an associated `system_transaction`
	is_flash *bool
	// The system_transaction property
	system_transaction SystemTransactionable
}

// NewGenesisLedgerTransaction instantiates a new GenesisLedgerTransaction and sets the default values.
func NewGenesisLedgerTransaction() *GenesisLedgerTransaction {
	m := &GenesisLedgerTransaction{
		LedgerTransaction: *NewLedgerTransaction(),
	}
	return m
}

// CreateGenesisLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenesisLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenesisLedgerTransaction(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenesisLedgerTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerTransaction.GetFieldDeserializers()
	res["is_flash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsFlash(val)
		}
		return nil
	}
	res["system_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSystemTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSystemTransaction(val.(SystemTransactionable))
		}
		return nil
	}
	return res
}

// GetIsFlash gets the is_flash property value. The first genesis "transaction" flashes state into the database to prepare for the bootstrap transaction.Such a transaction does not have an associated `system_transaction`
// returns a *bool when successful
func (m *GenesisLedgerTransaction) GetIsFlash() *bool {
	return m.is_flash
}

// GetSystemTransaction gets the system_transaction property value. The system_transaction property
// returns a SystemTransactionable when successful
func (m *GenesisLedgerTransaction) GetSystemTransaction() SystemTransactionable {
	return m.system_transaction
}

// Serialize serializes information the current object
func (m *GenesisLedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerTransaction.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteBoolValue("is_flash", m.GetIsFlash())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("system_transaction", m.GetSystemTransaction())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetIsFlash sets the is_flash property value. The first genesis "transaction" flashes state into the database to prepare for the bootstrap transaction.Such a transaction does not have an associated `system_transaction`
func (m *GenesisLedgerTransaction) SetIsFlash(value *bool) {
	m.is_flash = value
}

// SetSystemTransaction sets the system_transaction property value. The system_transaction property
func (m *GenesisLedgerTransaction) SetSystemTransaction(value SystemTransactionable) {
	m.system_transaction = value
}

type GenesisLedgerTransactionable interface {
	LedgerTransactionable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIsFlash() *bool
	GetSystemTransaction() SystemTransactionable
	SetIsFlash(value *bool)
	SetSystemTransaction(value SystemTransactionable)
}
