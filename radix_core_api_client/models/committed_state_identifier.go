package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommittedStateIdentifier struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded root hash of the receipt tree. This captures the consensus-agreed output of each transaction on the ledger.
	receipt_tree_hash *string
	// The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
	state_tree_hash *string
	// The state_version property
	state_version *int64
	// The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
	transaction_tree_hash *string
}

// NewCommittedStateIdentifier instantiates a new CommittedStateIdentifier and sets the default values.
func NewCommittedStateIdentifier() *CommittedStateIdentifier {
	m := &CommittedStateIdentifier{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCommittedStateIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommittedStateIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCommittedStateIdentifier(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommittedStateIdentifier) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommittedStateIdentifier) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["state_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStateVersion(val)
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
func (m *CommittedStateIdentifier) GetReceiptTreeHash() *string {
	return m.receipt_tree_hash
}

// GetStateTreeHash gets the state_tree_hash property value. The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
// returns a *string when successful
func (m *CommittedStateIdentifier) GetStateTreeHash() *string {
	return m.state_tree_hash
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *CommittedStateIdentifier) GetStateVersion() *int64 {
	return m.state_version
}

// GetTransactionTreeHash gets the transaction_tree_hash property value. The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
// returns a *string when successful
func (m *CommittedStateIdentifier) GetTransactionTreeHash() *string {
	return m.transaction_tree_hash
}

// Serialize serializes information the current object
func (m *CommittedStateIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteInt64Value("state_version", m.GetStateVersion())
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
func (m *CommittedStateIdentifier) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetReceiptTreeHash sets the receipt_tree_hash property value. The hex-encoded root hash of the receipt tree. This captures the consensus-agreed output of each transaction on the ledger.
func (m *CommittedStateIdentifier) SetReceiptTreeHash(value *string) {
	m.receipt_tree_hash = value
}

// SetStateTreeHash sets the state_tree_hash property value. The hex-encoded root hash of the state tree. This captures the current state of the state on the ledger.
func (m *CommittedStateIdentifier) SetStateTreeHash(value *string) {
	m.state_tree_hash = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *CommittedStateIdentifier) SetStateVersion(value *int64) {
	m.state_version = value
}

// SetTransactionTreeHash sets the transaction_tree_hash property value. The hex-encoded root hash of the transaction tree. This captures the ledger transactions committed to the ledger.
func (m *CommittedStateIdentifier) SetTransactionTreeHash(value *string) {
	m.transaction_tree_hash = value
}

type CommittedStateIdentifierable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetReceiptTreeHash() *string
	GetStateTreeHash() *string
	GetStateVersion() *int64
	GetTransactionTreeHash() *string
	SetReceiptTreeHash(value *string)
	SetStateTreeHash(value *string)
	SetStateVersion(value *int64)
	SetTransactionTreeHash(value *string)
}
