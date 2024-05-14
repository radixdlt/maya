package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionFormatOptions requested transaction formats to include in the response
type TransactionFormatOptions struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Whether to return the transaction balance changes (default false)
	balance_changes *bool
	// Whether to return the hex-encoded blobs (default false)
	blobs *bool
	// Whether to return the raw manifest (default true)
	manifest *bool
	// Whether to return the transaction message (default true)
	message *bool
	// Whether to return the raw hex-encoded ledger transaction bytes (default false)
	raw_ledger_transaction *bool
	// Whether to return the raw hex-encoded notarized transaction bytes (default true)
	raw_notarized_transaction *bool
	// Whether to return the raw hex-encoded system transaction bytes (default false)
	raw_system_transaction *bool
}

// NewTransactionFormatOptions instantiates a new TransactionFormatOptions and sets the default values.
func NewTransactionFormatOptions() *TransactionFormatOptions {
	m := &TransactionFormatOptions{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionFormatOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionFormatOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionFormatOptions(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionFormatOptions) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBalanceChanges gets the balance_changes property value. Whether to return the transaction balance changes (default false)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetBalanceChanges() *bool {
	return m.balance_changes
}

// GetBlobs gets the blobs property value. Whether to return the hex-encoded blobs (default false)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetBlobs() *bool {
	return m.blobs
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionFormatOptions) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBalanceChanges(val)
		}
		return nil
	}
	res["blobs"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlobs(val)
		}
		return nil
	}
	res["manifest"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetManifest(val)
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val)
		}
		return nil
	}
	res["raw_ledger_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRawLedgerTransaction(val)
		}
		return nil
	}
	res["raw_notarized_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRawNotarizedTransaction(val)
		}
		return nil
	}
	res["raw_system_transaction"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRawSystemTransaction(val)
		}
		return nil
	}
	return res
}

// GetManifest gets the manifest property value. Whether to return the raw manifest (default true)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetManifest() *bool {
	return m.manifest
}

// GetMessage gets the message property value. Whether to return the transaction message (default true)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetMessage() *bool {
	return m.message
}

// GetRawLedgerTransaction gets the raw_ledger_transaction property value. Whether to return the raw hex-encoded ledger transaction bytes (default false)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetRawLedgerTransaction() *bool {
	return m.raw_ledger_transaction
}

// GetRawNotarizedTransaction gets the raw_notarized_transaction property value. Whether to return the raw hex-encoded notarized transaction bytes (default true)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetRawNotarizedTransaction() *bool {
	return m.raw_notarized_transaction
}

// GetRawSystemTransaction gets the raw_system_transaction property value. Whether to return the raw hex-encoded system transaction bytes (default false)
// returns a *bool when successful
func (m *TransactionFormatOptions) GetRawSystemTransaction() *bool {
	return m.raw_system_transaction
}

// Serialize serializes information the current object
func (m *TransactionFormatOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("balance_changes", m.GetBalanceChanges())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("blobs", m.GetBlobs())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("manifest", m.GetManifest())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("message", m.GetMessage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("raw_ledger_transaction", m.GetRawLedgerTransaction())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("raw_notarized_transaction", m.GetRawNotarizedTransaction())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("raw_system_transaction", m.GetRawSystemTransaction())
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
func (m *TransactionFormatOptions) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBalanceChanges sets the balance_changes property value. Whether to return the transaction balance changes (default false)
func (m *TransactionFormatOptions) SetBalanceChanges(value *bool) {
	m.balance_changes = value
}

// SetBlobs sets the blobs property value. Whether to return the hex-encoded blobs (default false)
func (m *TransactionFormatOptions) SetBlobs(value *bool) {
	m.blobs = value
}

// SetManifest sets the manifest property value. Whether to return the raw manifest (default true)
func (m *TransactionFormatOptions) SetManifest(value *bool) {
	m.manifest = value
}

// SetMessage sets the message property value. Whether to return the transaction message (default true)
func (m *TransactionFormatOptions) SetMessage(value *bool) {
	m.message = value
}

// SetRawLedgerTransaction sets the raw_ledger_transaction property value. Whether to return the raw hex-encoded ledger transaction bytes (default false)
func (m *TransactionFormatOptions) SetRawLedgerTransaction(value *bool) {
	m.raw_ledger_transaction = value
}

// SetRawNotarizedTransaction sets the raw_notarized_transaction property value. Whether to return the raw hex-encoded notarized transaction bytes (default true)
func (m *TransactionFormatOptions) SetRawNotarizedTransaction(value *bool) {
	m.raw_notarized_transaction = value
}

// SetRawSystemTransaction sets the raw_system_transaction property value. Whether to return the raw hex-encoded system transaction bytes (default false)
func (m *TransactionFormatOptions) SetRawSystemTransaction(value *bool) {
	m.raw_system_transaction = value
}

type TransactionFormatOptionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBalanceChanges() *bool
	GetBlobs() *bool
	GetManifest() *bool
	GetMessage() *bool
	GetRawLedgerTransaction() *bool
	GetRawNotarizedTransaction() *bool
	GetRawSystemTransaction() *bool
	SetBalanceChanges(value *bool)
	SetBlobs(value *bool)
	SetManifest(value *bool)
	SetMessage(value *bool)
	SetRawLedgerTransaction(value *bool)
	SetRawNotarizedTransaction(value *bool)
	SetRawSystemTransaction(value *bool)
}
