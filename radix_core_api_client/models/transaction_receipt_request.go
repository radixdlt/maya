package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionReceiptRequest struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
	intent_hash *string
	// The logical name of the network
	network *string
	// Requested transaction formats to include in the response
	transaction_format_options TransactionFormatOptionsable
}

// NewTransactionReceiptRequest instantiates a new TransactionReceiptRequest and sets the default values.
func NewTransactionReceiptRequest() *TransactionReceiptRequest {
	m := &TransactionReceiptRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionReceiptRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionReceiptRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionReceiptRequest(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionReceiptRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionReceiptRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["intent_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIntentHash(val)
		}
		return nil
	}
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
		}
		return nil
	}
	res["transaction_format_options"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionFormatOptionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionFormatOptions(val.(TransactionFormatOptionsable))
		}
		return nil
	}
	return res
}

// GetIntentHash gets the intent_hash property value. The intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
// returns a *string when successful
func (m *TransactionReceiptRequest) GetIntentHash() *string {
	return m.intent_hash
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionReceiptRequest) GetNetwork() *string {
	return m.network
}

// GetTransactionFormatOptions gets the transaction_format_options property value. Requested transaction formats to include in the response
// returns a TransactionFormatOptionsable when successful
func (m *TransactionReceiptRequest) GetTransactionFormatOptions() TransactionFormatOptionsable {
	return m.transaction_format_options
}

// Serialize serializes information the current object
func (m *TransactionReceiptRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("intent_hash", m.GetIntentHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("transaction_format_options", m.GetTransactionFormatOptions())
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
func (m *TransactionReceiptRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIntentHash sets the intent_hash property value. The intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
func (m *TransactionReceiptRequest) SetIntentHash(value *string) {
	m.intent_hash = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionReceiptRequest) SetNetwork(value *string) {
	m.network = value
}

// SetTransactionFormatOptions sets the transaction_format_options property value. Requested transaction formats to include in the response
func (m *TransactionReceiptRequest) SetTransactionFormatOptions(value TransactionFormatOptionsable) {
	m.transaction_format_options = value
}

type TransactionReceiptRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIntentHash() *string
	GetNetwork() *string
	GetTransactionFormatOptions() TransactionFormatOptionsable
	SetIntentHash(value *string)
	SetNetwork(value *string)
	SetTransactionFormatOptions(value TransactionFormatOptionsable)
}
