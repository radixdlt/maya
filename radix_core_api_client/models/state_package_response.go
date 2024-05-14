package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StatePackageResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The at_ledger_state property
	at_ledger_state LedgerStateSummaryable
	// The owner_role property
	owner_role Substateable
	// The royalty property
	royalty Substateable
}

// NewStatePackageResponse instantiates a new StatePackageResponse and sets the default values.
func NewStatePackageResponse() *StatePackageResponse {
	m := &StatePackageResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStatePackageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatePackageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStatePackageResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StatePackageResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StatePackageResponse) GetAtLedgerState() LedgerStateSummaryable {
	return m.at_ledger_state
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StatePackageResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["owner_role"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOwnerRole(val.(Substateable))
		}
		return nil
	}
	res["royalty"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyalty(val.(Substateable))
		}
		return nil
	}
	return res
}

// GetOwnerRole gets the owner_role property value. The owner_role property
// returns a Substateable when successful
func (m *StatePackageResponse) GetOwnerRole() Substateable {
	return m.owner_role
}

// GetRoyalty gets the royalty property value. The royalty property
// returns a Substateable when successful
func (m *StatePackageResponse) GetRoyalty() Substateable {
	return m.royalty
}

// Serialize serializes information the current object
func (m *StatePackageResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("owner_role", m.GetOwnerRole())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("royalty", m.GetRoyalty())
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
func (m *StatePackageResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StatePackageResponse) SetAtLedgerState(value LedgerStateSummaryable) {
	m.at_ledger_state = value
}

// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *StatePackageResponse) SetOwnerRole(value Substateable) {
	m.owner_role = value
}

// SetRoyalty sets the royalty property value. The royalty property
func (m *StatePackageResponse) SetRoyalty(value Substateable) {
	m.royalty = value
}

type StatePackageResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAtLedgerState() LedgerStateSummaryable
	GetOwnerRole() Substateable
	GetRoyalty() Substateable
	SetAtLedgerState(value LedgerStateSummaryable)
	SetOwnerRole(value Substateable)
	SetRoyalty(value Substateable)
}
