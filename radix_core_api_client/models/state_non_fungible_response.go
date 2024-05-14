package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateNonFungibleResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The at_ledger_state property
	at_ledger_state LedgerStateSummaryable
	// The non_fungible property
	non_fungible Substateable
}

// NewStateNonFungibleResponse instantiates a new StateNonFungibleResponse and sets the default values.
func NewStateNonFungibleResponse() *StateNonFungibleResponse {
	m := &StateNonFungibleResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateNonFungibleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateNonFungibleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateNonFungibleResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateNonFungibleResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateNonFungibleResponse) GetAtLedgerState() LedgerStateSummaryable {
	return m.at_ledger_state
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateNonFungibleResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["at_ledger_state"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerStateSummaryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAtLedgerState(val.(LedgerStateSummaryable))
		}
		return nil
	}
	res["non_fungible"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNonFungible(val.(Substateable))
		}
		return nil
	}
	return res
}

// GetNonFungible gets the non_fungible property value. The non_fungible property
// returns a Substateable when successful
func (m *StateNonFungibleResponse) GetNonFungible() Substateable {
	return m.non_fungible
}

// Serialize serializes information the current object
func (m *StateNonFungibleResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("non_fungible", m.GetNonFungible())
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
func (m *StateNonFungibleResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateNonFungibleResponse) SetAtLedgerState(value LedgerStateSummaryable) {
	m.at_ledger_state = value
}

// SetNonFungible sets the non_fungible property value. The non_fungible property
func (m *StateNonFungibleResponse) SetNonFungible(value Substateable) {
	m.non_fungible = value
}

type StateNonFungibleResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAtLedgerState() LedgerStateSummaryable
	GetNonFungible() Substateable
	SetAtLedgerState(value LedgerStateSummaryable)
	SetNonFungible(value Substateable)
}
