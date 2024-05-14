package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerHashes struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded root hash of the receipt tree. This captures the consensus-agreed output of each transaction on the ledger.
	receipt_tree_hash *string
	// The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
	state_tree_hash *string
	// The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
	transaction_tree_hash *string
}

// NewLedgerHashes instantiates a new LedgerHashes and sets the default values.
func NewLedgerHashes() *LedgerHashes {
	m := &LedgerHashes{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLedgerHashesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerHashesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLedgerHashes(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerHashes) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerHashes) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["receipt_tree_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReceiptTreeHash(val)
		}
		return nil
	}
	res["state_tree_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStateTreeHash(val)
		}
		return nil
	}
	res["transaction_tree_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionTreeHash(val)
		}
		return nil
	}
	return res
}

// GetReceiptTreeHash gets the receipt_tree_hash property value. The hex-encoded root hash of the receipt tree. This captures the consensus-agreed output of each transaction on the ledger.
// returns a *string when successful
func (m *LedgerHashes) GetReceiptTreeHash() *string {
	return m.receipt_tree_hash
}

// GetStateTreeHash gets the state_tree_hash property value. The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
// returns a *string when successful
func (m *LedgerHashes) GetStateTreeHash() *string {
	return m.state_tree_hash
}

// GetTransactionTreeHash gets the transaction_tree_hash property value. The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
// returns a *string when successful
func (m *LedgerHashes) GetTransactionTreeHash() *string {
	return m.transaction_tree_hash
}

// Serialize serializes information the current object
func (m *LedgerHashes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("receipt_tree_hash", m.GetReceiptTreeHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("state_tree_hash", m.GetStateTreeHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("transaction_tree_hash", m.GetTransactionTreeHash())
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
func (m *LedgerHashes) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetReceiptTreeHash sets the receipt_tree_hash property value. The hex-encoded root hash of the receipt tree. This captures the consensus-agreed output of each transaction on the ledger.
func (m *LedgerHashes) SetReceiptTreeHash(value *string) {
	m.receipt_tree_hash = value
}

// SetStateTreeHash sets the state_tree_hash property value. The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
func (m *LedgerHashes) SetStateTreeHash(value *string) {
	m.state_tree_hash = value
}

// SetTransactionTreeHash sets the transaction_tree_hash property value. The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
func (m *LedgerHashes) SetTransactionTreeHash(value *string) {
	m.transaction_tree_hash = value
}

type LedgerHashesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetReceiptTreeHash() *string
	GetStateTreeHash() *string
	GetTransactionTreeHash() *string
	SetReceiptTreeHash(value *string)
	SetStateTreeHash(value *string)
	SetTransactionTreeHash(value *string)
}
