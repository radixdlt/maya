package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExecutedScenarioTransaction struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
	intent_hash *string
	// The Bech32m-encoded human readable `IntentHash`.
	intent_hash_bech32m *string
	// The logical_name property
	logical_name *string
	// The state_version property
	state_version *int64
}

// NewExecutedScenarioTransaction instantiates a new ExecutedScenarioTransaction and sets the default values.
func NewExecutedScenarioTransaction() *ExecutedScenarioTransaction {
	m := &ExecutedScenarioTransaction{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateExecutedScenarioTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExecutedScenarioTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewExecutedScenarioTransaction(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExecutedScenarioTransaction) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExecutedScenarioTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["intent_hash_bech32m"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIntentHashBech32m(val)
		}
		return nil
	}
	res["logical_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLogicalName(val)
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
	return res
}

// GetIntentHash gets the intent_hash property value. The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
// returns a *string when successful
func (m *ExecutedScenarioTransaction) GetIntentHash() *string {
	return m.intent_hash
}

// GetIntentHashBech32m gets the intent_hash_bech32m property value. The Bech32m-encoded human readable `IntentHash`.
// returns a *string when successful
func (m *ExecutedScenarioTransaction) GetIntentHashBech32m() *string {
	return m.intent_hash_bech32m
}

// GetLogicalName gets the logical_name property value. The logical_name property
// returns a *string when successful
func (m *ExecutedScenarioTransaction) GetLogicalName() *string {
	return m.logical_name
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *ExecutedScenarioTransaction) GetStateVersion() *int64 {
	return m.state_version
}

// Serialize serializes information the current object
func (m *ExecutedScenarioTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("intent_hash", m.GetIntentHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("intent_hash_bech32m", m.GetIntentHashBech32m())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("logical_name", m.GetLogicalName())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecutedScenarioTransaction) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIntentHash sets the intent_hash property value. The hex-encoded intent hash for a user transaction, also known as the transaction id.This hash identifies the core content "intent" of the transaction. Each intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *ExecutedScenarioTransaction) SetIntentHash(value *string) {
	m.intent_hash = value
}

// SetIntentHashBech32m sets the intent_hash_bech32m property value. The Bech32m-encoded human readable `IntentHash`.
func (m *ExecutedScenarioTransaction) SetIntentHashBech32m(value *string) {
	m.intent_hash_bech32m = value
}

// SetLogicalName sets the logical_name property value. The logical_name property
func (m *ExecutedScenarioTransaction) SetLogicalName(value *string) {
	m.logical_name = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *ExecutedScenarioTransaction) SetStateVersion(value *int64) {
	m.state_version = value
}

type ExecutedScenarioTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIntentHash() *string
	GetIntentHashBech32m() *string
	GetLogicalName() *string
	GetStateVersion() *int64
	SetIntentHash(value *string)
	SetIntentHashBech32m(value *string)
	SetLogicalName(value *string)
	SetStateVersion(value *int64)
}
