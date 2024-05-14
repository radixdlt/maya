package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionTrackerCollectionEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The status property
	status *TransactionTrackerTransactionStatus
}

// NewTransactionTrackerCollectionEntryValue instantiates a new TransactionTrackerCollectionEntryValue and sets the default values.
func NewTransactionTrackerCollectionEntryValue() *TransactionTrackerCollectionEntryValue {
	m := &TransactionTrackerCollectionEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionTrackerCollectionEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionTrackerCollectionEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionTrackerCollectionEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionTrackerCollectionEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionTrackerCollectionEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseTransactionTrackerTransactionStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val.(*TransactionTrackerTransactionStatus))
		}
		return nil
	}
	return res
}

// GetStatus gets the status property value. The status property
// returns a *TransactionTrackerTransactionStatus when successful
func (m *TransactionTrackerCollectionEntryValue) GetStatus() *TransactionTrackerTransactionStatus {
	return m.status
}

// Serialize serializes information the current object
func (m *TransactionTrackerCollectionEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetStatus() != nil {
		cast := (*m.GetStatus()).String()
		err := writer.WriteStringValue("status", &cast)
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
func (m *TransactionTrackerCollectionEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetStatus sets the status property value. The status property
func (m *TransactionTrackerCollectionEntryValue) SetStatus(value *TransactionTrackerTransactionStatus) {
	m.status = value
}

type TransactionTrackerCollectionEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetStatus() *TransactionTrackerTransactionStatus
	SetStatus(value *TransactionTrackerTransactionStatus)
}
