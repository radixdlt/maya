package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExecutedScenario struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Well-named addresses touched/created by the Scenario, keyed by their name.
	addresses ExecutedScenario_addressesable
	// Transactions successfully committed by the Scenario.
	committed_transactions []ExecutedScenarioTransactionable
	// The logical_name property
	logical_name *string
	// An index of the Scenario (reflecting its execution order).
	sequence_number *int32
}

// NewExecutedScenario instantiates a new ExecutedScenario and sets the default values.
func NewExecutedScenario() *ExecutedScenario {
	m := &ExecutedScenario{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateExecutedScenarioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExecutedScenarioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewExecutedScenario(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExecutedScenario) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddresses gets the addresses property value. Well-named addresses touched/created by the Scenario, keyed by their name.
// returns a ExecutedScenario_addressesable when successful
func (m *ExecutedScenario) GetAddresses() ExecutedScenario_addressesable {
	return m.addresses
}

// GetCommittedTransactions gets the committed_transactions property value. Transactions successfully committed by the Scenario.
// returns a []ExecutedScenarioTransactionable when successful
func (m *ExecutedScenario) GetCommittedTransactions() []ExecutedScenarioTransactionable {
	return m.committed_transactions
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExecutedScenario) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateExecutedScenario_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddresses(val.(ExecutedScenario_addressesable))
		}
		return nil
	}
	res["committed_transactions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateExecutedScenarioTransactionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ExecutedScenarioTransactionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ExecutedScenarioTransactionable)
				}
			}
			m.SetCommittedTransactions(res)
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
	res["sequence_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSequenceNumber(val)
		}
		return nil
	}
	return res
}

// GetLogicalName gets the logical_name property value. The logical_name property
// returns a *string when successful
func (m *ExecutedScenario) GetLogicalName() *string {
	return m.logical_name
}

// GetSequenceNumber gets the sequence_number property value. An index of the Scenario (reflecting its execution order).
// returns a *int32 when successful
func (m *ExecutedScenario) GetSequenceNumber() *int32 {
	return m.sequence_number
}

// Serialize serializes information the current object
func (m *ExecutedScenario) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("addresses", m.GetAddresses())
		if err != nil {
			return err
		}
	}
	if m.GetCommittedTransactions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommittedTransactions()))
		for i, v := range m.GetCommittedTransactions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("committed_transactions", cast)
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
		err := writer.WriteInt32Value("sequence_number", m.GetSequenceNumber())
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
func (m *ExecutedScenario) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddresses sets the addresses property value. Well-named addresses touched/created by the Scenario, keyed by their name.
func (m *ExecutedScenario) SetAddresses(value ExecutedScenario_addressesable) {
	m.addresses = value
}

// SetCommittedTransactions sets the committed_transactions property value. Transactions successfully committed by the Scenario.
func (m *ExecutedScenario) SetCommittedTransactions(value []ExecutedScenarioTransactionable) {
	m.committed_transactions = value
}

// SetLogicalName sets the logical_name property value. The logical_name property
func (m *ExecutedScenario) SetLogicalName(value *string) {
	m.logical_name = value
}

// SetSequenceNumber sets the sequence_number property value. An index of the Scenario (reflecting its execution order).
func (m *ExecutedScenario) SetSequenceNumber(value *int32) {
	m.sequence_number = value
}

type ExecutedScenarioable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddresses() ExecutedScenario_addressesable
	GetCommittedTransactions() []ExecutedScenarioTransactionable
	GetLogicalName() *string
	GetSequenceNumber() *int32
	SetAddresses(value ExecutedScenario_addressesable)
	SetCommittedTransactions(value []ExecutedScenarioTransactionable)
	SetLogicalName(value *string)
	SetSequenceNumber(value *int32)
}
