package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateComponentResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The at_ledger_state property
	at_ledger_state LedgerStateSummaryable
	// Any descendent nodes owned directly or indirectly by the component
	descendent_nodes []StateComponentDescendentNodeable
	// The info property
	info Substateable
	// The owner_role property
	owner_role Substateable
	// The royalty_accumulator property
	royalty_accumulator Substateable
	// The state property
	state Substateable
	// Any vaults owned directly or indirectly by the component
	vaults []VaultBalanceable
}

// NewStateComponentResponse instantiates a new StateComponentResponse and sets the default values.
func NewStateComponentResponse() *StateComponentResponse {
	m := &StateComponentResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateComponentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateComponentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateComponentResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateComponentResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateComponentResponse) GetAtLedgerState() LedgerStateSummaryable {
	return m.at_ledger_state
}

// GetDescendentNodes gets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
// returns a []StateComponentDescendentNodeable when successful
func (m *StateComponentResponse) GetDescendentNodes() []StateComponentDescendentNodeable {
	return m.descendent_nodes
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateComponentResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["descendent_nodes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateStateComponentDescendentNodeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]StateComponentDescendentNodeable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(StateComponentDescendentNodeable)
				}
			}
			m.SetDescendentNodes(res)
		}
		return nil
	}
	res["info"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInfo(val.(Substateable))
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
	res["royalty_accumulator"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyaltyAccumulator(val.(Substateable))
		}
		return nil
	}
	res["state"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetState(val.(Substateable))
		}
		return nil
	}
	res["vaults"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateVaultBalanceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]VaultBalanceable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(VaultBalanceable)
				}
			}
			m.SetVaults(res)
		}
		return nil
	}
	return res
}

// GetInfo gets the info property value. The info property
// returns a Substateable when successful
func (m *StateComponentResponse) GetInfo() Substateable {
	return m.info
}

// GetOwnerRole gets the owner_role property value. The owner_role property
// returns a Substateable when successful
func (m *StateComponentResponse) GetOwnerRole() Substateable {
	return m.owner_role
}

// GetRoyaltyAccumulator gets the royalty_accumulator property value. The royalty_accumulator property
// returns a Substateable when successful
func (m *StateComponentResponse) GetRoyaltyAccumulator() Substateable {
	return m.royalty_accumulator
}

// GetState gets the state property value. The state property
// returns a Substateable when successful
func (m *StateComponentResponse) GetState() Substateable {
	return m.state
}

// GetVaults gets the vaults property value. Any vaults owned directly or indirectly by the component
// returns a []VaultBalanceable when successful
func (m *StateComponentResponse) GetVaults() []VaultBalanceable {
	return m.vaults
}

// Serialize serializes information the current object
func (m *StateComponentResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
		if err != nil {
			return err
		}
	}
	if m.GetDescendentNodes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDescendentNodes()))
		for i, v := range m.GetDescendentNodes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("descendent_nodes", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("info", m.GetInfo())
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
		err := writer.WriteObjectValue("royalty_accumulator", m.GetRoyaltyAccumulator())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("state", m.GetState())
		if err != nil {
			return err
		}
	}
	if m.GetVaults() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVaults()))
		for i, v := range m.GetVaults() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("vaults", cast)
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
func (m *StateComponentResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateComponentResponse) SetAtLedgerState(value LedgerStateSummaryable) {
	m.at_ledger_state = value
}

// SetDescendentNodes sets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
func (m *StateComponentResponse) SetDescendentNodes(value []StateComponentDescendentNodeable) {
	m.descendent_nodes = value
}

// SetInfo sets the info property value. The info property
func (m *StateComponentResponse) SetInfo(value Substateable) {
	m.info = value
}

// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *StateComponentResponse) SetOwnerRole(value Substateable) {
	m.owner_role = value
}

// SetRoyaltyAccumulator sets the royalty_accumulator property value. The royalty_accumulator property
func (m *StateComponentResponse) SetRoyaltyAccumulator(value Substateable) {
	m.royalty_accumulator = value
}

// SetState sets the state property value. The state property
func (m *StateComponentResponse) SetState(value Substateable) {
	m.state = value
}

// SetVaults sets the vaults property value. Any vaults owned directly or indirectly by the component
func (m *StateComponentResponse) SetVaults(value []VaultBalanceable) {
	m.vaults = value
}

type StateComponentResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAtLedgerState() LedgerStateSummaryable
	GetDescendentNodes() []StateComponentDescendentNodeable
	GetInfo() Substateable
	GetOwnerRole() Substateable
	GetRoyaltyAccumulator() Substateable
	GetState() Substateable
	GetVaults() []VaultBalanceable
	SetAtLedgerState(value LedgerStateSummaryable)
	SetDescendentNodes(value []StateComponentDescendentNodeable)
	SetInfo(value Substateable)
	SetOwnerRole(value Substateable)
	SetRoyaltyAccumulator(value Substateable)
	SetState(value Substateable)
	SetVaults(value []VaultBalanceable)
}
