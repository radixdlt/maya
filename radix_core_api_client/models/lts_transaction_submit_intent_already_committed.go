package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionSubmitIntentAlreadyCommitted struct {
	LtsTransactionSubmitErrorDetails
	// The committed_as property
	committed_as CommittedIntentMetadataable
}

// NewLtsTransactionSubmitIntentAlreadyCommitted instantiates a new LtsTransactionSubmitIntentAlreadyCommitted and sets the default values.
func NewLtsTransactionSubmitIntentAlreadyCommitted() *LtsTransactionSubmitIntentAlreadyCommitted {
	m := &LtsTransactionSubmitIntentAlreadyCommitted{
		LtsTransactionSubmitErrorDetails: *NewLtsTransactionSubmitErrorDetails(),
	}
	return m
}

// CreateLtsTransactionSubmitIntentAlreadyCommittedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionSubmitIntentAlreadyCommittedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionSubmitIntentAlreadyCommitted(), nil
}

// GetCommittedAs gets the committed_as property value. The committed_as property
// returns a CommittedIntentMetadataable when successful
func (m *LtsTransactionSubmitIntentAlreadyCommitted) GetCommittedAs() CommittedIntentMetadataable {
	return m.committed_as
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionSubmitIntentAlreadyCommitted) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LtsTransactionSubmitErrorDetails.GetFieldDeserializers()
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
func (m *LtsTransactionSubmitIntentAlreadyCommitted) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LtsTransactionSubmitErrorDetails.Serialize(writer)
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
func (m *LtsTransactionSubmitIntentAlreadyCommitted) SetCommittedAs(value CommittedIntentMetadataable) {
	m.committed_as = value
}

type LtsTransactionSubmitIntentAlreadyCommittedable interface {
	LtsTransactionSubmitErrorDetailsable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCommittedAs() CommittedIntentMetadataable
	SetCommittedAs(value CommittedIntentMetadataable)
}
