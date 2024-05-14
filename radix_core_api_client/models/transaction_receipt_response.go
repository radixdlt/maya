package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionReceiptResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The committed property
	committed CommittedTransactionable
}

// NewTransactionReceiptResponse instantiates a new TransactionReceiptResponse and sets the default values.
func NewTransactionReceiptResponse() *TransactionReceiptResponse {
	m := &TransactionReceiptResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionReceiptResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionReceiptResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionReceiptResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionReceiptResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCommitted gets the committed property value. The committed property
// returns a CommittedTransactionable when successful
func (m *TransactionReceiptResponse) GetCommitted() CommittedTransactionable {
	return m.committed
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionReceiptResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["committed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCommittedTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCommitted(val.(CommittedTransactionable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TransactionReceiptResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("committed", m.GetCommitted())
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
func (m *TransactionReceiptResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCommitted sets the committed property value. The committed property
func (m *TransactionReceiptResponse) SetCommitted(value CommittedTransactionable) {
	m.committed = value
}

type TransactionReceiptResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCommitted() CommittedTransactionable
	SetCommitted(value CommittedTransactionable)
}
