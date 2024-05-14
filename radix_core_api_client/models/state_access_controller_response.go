package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateAccessControllerResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The at_ledger_state property
	at_ledger_state LedgerStateSummaryable
	// Any descendent nodes owned directly or indirectly by the component
	descendent_nodes []StateComponentDescendentNodeable
	// The owner_role property
	owner_role Substateable
	// The state property
	state Substateable
	// Any vaults owned directly or indirectly by the component
	vaults []VaultBalanceable
}

// NewStateAccessControllerResponse instantiates a new StateAccessControllerResponse and sets the default values.
func NewStateAccessControllerResponse() *StateAccessControllerResponse {
	m := &StateAccessControllerResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateAccessControllerResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateAccessControllerResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateAccessControllerResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateAccessControllerResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateAccessControllerResponse) GetAtLedgerState() LedgerStateSummaryable {
	return m.at_ledger_state
}

// GetDescendentNodes gets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
// returns a []StateComponentDescendentNodeable when successful
func (m *StateAccessControllerResponse) GetDescendentNodes() []StateComponentDescendentNodeable {
	return m.descendent_nodes
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateAccessControllerResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetOwnerRole gets the owner_role property value. The owner_role property
// returns a Substateable when successful
func (m *StateAccessControllerResponse) GetOwnerRole() Substateable {
	return m.owner_role
}

// GetState gets the state property value. The state property
// returns a Substateable when successful
func (m *StateAccessControllerResponse) GetState() Substateable {
	return m.state
}

// GetVaults gets the vaults property value. Any vaults owned directly or indirectly by the component
// returns a []VaultBalanceable when successful
func (m *StateAccessControllerResponse) GetVaults() []VaultBalanceable {
	return m.vaults
}

// Serialize serializes information the current object
func (m *StateAccessControllerResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteObjectValue("owner_role", m.GetOwnerRole())
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
func (m *StateAccessControllerResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateAccessControllerResponse) SetAtLedgerState(value LedgerStateSummaryable) {
	m.at_ledger_state = value
}

// SetDescendentNodes sets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
func (m *StateAccessControllerResponse) SetDescendentNodes(value []StateComponentDescendentNodeable) {
	m.descendent_nodes = value
}

// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *StateAccessControllerResponse) SetOwnerRole(value Substateable) {
	m.owner_role = value
}

// SetState sets the state property value. The state property
func (m *StateAccessControllerResponse) SetState(value Substateable) {
	m.state = value
}

// SetVaults sets the vaults property value. Any vaults owned directly or indirectly by the component
func (m *StateAccessControllerResponse) SetVaults(value []VaultBalanceable) {
	m.vaults = value
}

type StateAccessControllerResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAtLedgerState() LedgerStateSummaryable
	GetDescendentNodes() []StateComponentDescendentNodeable
	GetOwnerRole() Substateable
	GetState() Substateable
	GetVaults() []VaultBalanceable
	SetAtLedgerState(value LedgerStateSummaryable)
	SetDescendentNodes(value []StateComponentDescendentNodeable)
	SetOwnerRole(value Substateable)
	SetState(value Substateable)
	SetVaults(value []VaultBalanceable)
}
