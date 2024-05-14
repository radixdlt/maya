package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionPayloadDetails struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An explanation for the error, if failed or rejected
	error_message *string
	// The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
	payload_hash *string
	// The Bech32m-encoded human readable `NotarizedTransactionHash`.
	payload_hash_bech32m *string
	// The state_version property
	state_version *int64
	// The status of the transaction payload, as per this node.A NotInMempool status means that it wasn't rejected at last execution attempt, but it's not currently in the mempool either.
	status *LtsTransactionPayloadStatus
}

// NewLtsTransactionPayloadDetails instantiates a new LtsTransactionPayloadDetails and sets the default values.
func NewLtsTransactionPayloadDetails() *LtsTransactionPayloadDetails {
	m := &LtsTransactionPayloadDetails{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsTransactionPayloadDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionPayloadDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionPayloadDetails(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsTransactionPayloadDetails) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetErrorMessage gets the error_message property value. An explanation for the error, if failed or rejected
// returns a *string when successful
func (m *LtsTransactionPayloadDetails) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionPayloadDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["error_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrorMessage(val)
		}
		return nil
	}
	res["payload_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPayloadHash(val)
		}
		return nil
	}
	res["payload_hash_bech32m"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPayloadHashBech32m(val)
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
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseLtsTransactionPayloadStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val.(*LtsTransactionPayloadStatus))
		}
		return nil
	}
	return res
}

// GetPayloadHash gets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
// returns a *string when successful
func (m *LtsTransactionPayloadDetails) GetPayloadHash() *string {
	return m.payload_hash
}

// GetPayloadHashBech32m gets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
// returns a *string when successful
func (m *LtsTransactionPayloadDetails) GetPayloadHashBech32m() *string {
	return m.payload_hash_bech32m
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LtsTransactionPayloadDetails) GetStateVersion() *int64 {
	return m.state_version
}

// GetStatus gets the status property value. The status of the transaction payload, as per this node.A NotInMempool status means that it wasn't rejected at last execution attempt, but it's not currently in the mempool either.
// returns a *LtsTransactionPayloadStatus when successful
func (m *LtsTransactionPayloadDetails) GetStatus() *LtsTransactionPayloadStatus {
	return m.status
}

// Serialize serializes information the current object
func (m *LtsTransactionPayloadDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("error_message", m.GetErrorMessage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payload_hash", m.GetPayloadHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payload_hash_bech32m", m.GetPayloadHashBech32m())
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
func (m *LtsTransactionPayloadDetails) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetErrorMessage sets the error_message property value. An explanation for the error, if failed or rejected
func (m *LtsTransactionPayloadDetails) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetPayloadHash sets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
func (m *LtsTransactionPayloadDetails) SetPayloadHash(value *string) {
	m.payload_hash = value
}

// SetPayloadHashBech32m sets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
func (m *LtsTransactionPayloadDetails) SetPayloadHashBech32m(value *string) {
	m.payload_hash_bech32m = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LtsTransactionPayloadDetails) SetStateVersion(value *int64) {
	m.state_version = value
}

// SetStatus sets the status property value. The status of the transaction payload, as per this node.A NotInMempool status means that it wasn't rejected at last execution attempt, but it's not currently in the mempool either.
func (m *LtsTransactionPayloadDetails) SetStatus(value *LtsTransactionPayloadStatus) {
	m.status = value
}

type LtsTransactionPayloadDetailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetErrorMessage() *string
	GetPayloadHash() *string
	GetPayloadHashBech32m() *string
	GetStateVersion() *int64
	GetStatus() *LtsTransactionPayloadStatus
	SetErrorMessage(value *string)
	SetPayloadHash(value *string)
	SetPayloadHashBech32m(value *string)
	SetStateVersion(value *int64)
	SetStatus(value *LtsTransactionPayloadStatus)
}
