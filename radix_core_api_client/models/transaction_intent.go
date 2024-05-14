package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionIntent struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
	blobs_hex TransactionIntent_blobs_hexable
	// The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
	hash *string
	// The Bech32m-encoded human readable `IntentHash`.
	hash_bech32m *string
	// The header property
	header TransactionHeaderable
	// The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
	instructions *string
	// The message property
	message TransactionMessageable
}

// NewTransactionIntent instantiates a new TransactionIntent and sets the default values.
func NewTransactionIntent() *TransactionIntent {
	m := &TransactionIntent{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionIntent(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionIntent) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBlobsHex gets the blobs_hex property value. A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
// returns a TransactionIntent_blobs_hexable when successful
func (m *TransactionIntent) GetBlobsHex() TransactionIntent_blobs_hexable {
	return m.blobs_hex
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionIntent) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["blobs_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionIntent_blobs_hexFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlobsHex(val.(TransactionIntent_blobs_hexable))
		}
		return nil
	}
	res["hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHash(val)
		}
		return nil
	}
	res["hash_bech32m"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHashBech32m(val)
		}
		return nil
	}
	res["header"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionHeaderFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHeader(val.(TransactionHeaderable))
		}
		return nil
	}
	res["instructions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInstructions(val)
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionMessageFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val.(TransactionMessageable))
		}
		return nil
	}
	return res
}

// GetHash gets the hash property value. The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
// returns a *string when successful
func (m *TransactionIntent) GetHash() *string {
	return m.hash
}

// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `IntentHash`.
// returns a *string when successful
func (m *TransactionIntent) GetHashBech32m() *string {
	return m.hash_bech32m
}

// GetHeader gets the header property value. The header property
// returns a TransactionHeaderable when successful
func (m *TransactionIntent) GetHeader() TransactionHeaderable {
	return m.header
}

// GetInstructions gets the instructions property value. The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
// returns a *string when successful
func (m *TransactionIntent) GetInstructions() *string {
	return m.instructions
}

// GetMessage gets the message property value. The message property
// returns a TransactionMessageable when successful
func (m *TransactionIntent) GetMessage() TransactionMessageable {
	return m.message
}

// Serialize serializes information the current object
func (m *TransactionIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("blobs_hex", m.GetBlobsHex())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("hash", m.GetHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("hash_bech32m", m.GetHashBech32m())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("header", m.GetHeader())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("instructions", m.GetInstructions())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *TransactionIntent) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBlobsHex sets the blobs_hex property value. A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
func (m *TransactionIntent) SetBlobsHex(value TransactionIntent_blobs_hexable) {
	m.blobs_hex = value
}

// SetHash sets the hash property value. The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *TransactionIntent) SetHash(value *string) {
	m.hash = value
}

// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `IntentHash`.
func (m *TransactionIntent) SetHashBech32m(value *string) {
	m.hash_bech32m = value
}

// SetHeader sets the header property value. The header property
func (m *TransactionIntent) SetHeader(value TransactionHeaderable) {
	m.header = value
}

// SetInstructions sets the instructions property value. The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
func (m *TransactionIntent) SetInstructions(value *string) {
	m.instructions = value
}

// SetMessage sets the message property value. The message property
func (m *TransactionIntent) SetMessage(value TransactionMessageable) {
	m.message = value
}

type TransactionIntentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBlobsHex() TransactionIntent_blobs_hexable
	GetHash() *string
	GetHashBech32m() *string
	GetHeader() TransactionHeaderable
	GetInstructions() *string
	GetMessage() TransactionMessageable
	SetBlobsHex(value TransactionIntent_blobs_hexable)
	SetHash(value *string)
	SetHashBech32m(value *string)
	SetHeader(value TransactionHeaderable)
	SetInstructions(value *string)
	SetMessage(value TransactionMessageable)
}
