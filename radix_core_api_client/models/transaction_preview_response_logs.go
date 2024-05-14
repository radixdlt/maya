package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewResponse_logs struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The level property
	level *string
	// The message property
	message *string
}

// NewTransactionPreviewResponse_logs instantiates a new TransactionPreviewResponse_logs and sets the default values.
func NewTransactionPreviewResponse_logs() *TransactionPreviewResponse_logs {
	m := &TransactionPreviewResponse_logs{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionPreviewResponse_logsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewResponse_logsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionPreviewResponse_logs(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewResponse_logs) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewResponse_logs) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["level"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLevel(val)
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val)
		}
		return nil
	}
	return res
}

// GetLevel gets the level property value. The level property
// returns a *string when successful
func (m *TransactionPreviewResponse_logs) GetLevel() *string {
	return m.level
}

// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *TransactionPreviewResponse_logs) GetMessage() *string {
	return m.message
}

// Serialize serializes information the current object
func (m *TransactionPreviewResponse_logs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("level", m.GetLevel())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *TransactionPreviewResponse_logs) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLevel sets the level property value. The level property
func (m *TransactionPreviewResponse_logs) SetLevel(value *string) {
	m.level = value
}

// SetMessage sets the message property value. The message property
func (m *TransactionPreviewResponse_logs) SetMessage(value *string) {
	m.message = value
}

type TransactionPreviewResponse_logsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLevel() *string
	GetMessage() *string
	SetLevel(value *string)
	SetMessage(value *string)
}
