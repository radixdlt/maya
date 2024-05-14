package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionSubmitIntentAlreadyCommitted struct {
	TransactionSubmitErrorDetails
	// The committed_as property
	committed_as CommittedIntentMetadataable
}

// NewTransactionSubmitIntentAlreadyCommitted instantiates a new TransactionSubmitIntentAlreadyCommitted and sets the default values.
func NewTransactionSubmitIntentAlreadyCommitted() *TransactionSubmitIntentAlreadyCommitted {
	m := &TransactionSubmitIntentAlreadyCommitted{
		TransactionSubmitErrorDetails: *NewTransactionSubmitErrorDetails(),
	}
	return m
}

// CreateTransactionSubmitIntentAlreadyCommittedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionSubmitIntentAlreadyCommittedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionSubmitIntentAlreadyCommitted(), nil
}

// GetCommittedAs gets the committed_as property value. The committed_as property
// returns a CommittedIntentMetadataable when successful
func (m *TransactionSubmitIntentAlreadyCommitted) GetCommittedAs() CommittedIntentMetadataable {
	return m.committed_as
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionSubmitIntentAlreadyCommitted) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TransactionSubmitErrorDetails.GetFieldDeserializers()
	res["committed_as"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedIntentMetadataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCommittedAs(val.(CommittedIntentMetadataable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TransactionSubmitIntentAlreadyCommitted) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TransactionSubmitErrorDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("committed_as", m.GetCommittedAs())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetCommittedAs sets the committed_as property value. The committed_as property
func (m *TransactionSubmitIntentAlreadyCommitted) SetCommittedAs(value CommittedIntentMetadataable) {
	m.committed_as = value
}

type TransactionSubmitIntentAlreadyCommittedable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TransactionSubmitErrorDetailsable
	GetCommittedAs() CommittedIntentMetadataable
	SetCommittedAs(value CommittedIntentMetadataable)
}
