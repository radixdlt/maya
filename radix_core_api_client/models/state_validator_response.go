package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateValidatorResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The Bech32m-encoded human readable version of the component address
	address *string
	// The at_ledger_state property
	at_ledger_state LedgerStateSummaryable
	// Any descendent nodes owned directly or indirectly by the component
	descendent_nodes []StateComponentDescendentNodeable
	// The owner_role property
	owner_role Substateable
	// The protocol_update_readiness_signal property
	protocol_update_readiness_signal Substateable
	// The state property
	state Substateable
	// Any vaults owned directly or indirectly by the component
	vaults []VaultBalanceable
}

// NewStateValidatorResponse instantiates a new StateValidatorResponse and sets the default values.
func NewStateValidatorResponse() *StateValidatorResponse {
	m := &StateValidatorResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateValidatorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateValidatorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateValidatorResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateValidatorResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *StateValidatorResponse) GetAddress() *string {
	return m.address
}

// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateValidatorResponse) GetAtLedgerState() LedgerStateSummaryable {
	return m.at_ledger_state
}

// GetDescendentNodes gets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
// returns a []StateComponentDescendentNodeable when successful
func (m *StateValidatorResponse) GetDescendentNodes() []StateComponentDescendentNodeable {
	return m.descendent_nodes
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateValidatorResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val)
		}
		return nil
	}
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
	res["protocol_update_readiness_signal"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProtocolUpdateReadinessSignal(val.(Substateable))
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
func (m *StateValidatorResponse) GetOwnerRole() Substateable {
	return m.owner_role
}

// GetProtocolUpdateReadinessSignal gets the protocol_update_readiness_signal property value. The protocol_update_readiness_signal property
// returns a Substateable when successful
func (m *StateValidatorResponse) GetProtocolUpdateReadinessSignal() Substateable {
	return m.protocol_update_readiness_signal
}

// GetState gets the state property value. The state property
// returns a Substateable when successful
func (m *StateValidatorResponse) GetState() Substateable {
	return m.state
}

// GetVaults gets the vaults property value. Any vaults owned directly or indirectly by the component
// returns a []VaultBalanceable when successful
func (m *StateValidatorResponse) GetVaults() []VaultBalanceable {
	return m.vaults
}

// Serialize serializes information the current object
func (m *StateValidatorResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("address", m.GetAddress())
		if err != nil {
			return err
		}
	}
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
		err := writer.WriteObjectValue("protocol_update_readiness_signal", m.GetProtocolUpdateReadinessSignal())
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
func (m *StateValidatorResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The Bech32m-encoded human readable version of the component address
func (m *StateValidatorResponse) SetAddress(value *string) {
	m.address = value
}

// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateValidatorResponse) SetAtLedgerState(value LedgerStateSummaryable) {
	m.at_ledger_state = value
}

// SetDescendentNodes sets the descendent_nodes property value. Any descendent nodes owned directly or indirectly by the component
func (m *StateValidatorResponse) SetDescendentNodes(value []StateComponentDescendentNodeable) {
	m.descendent_nodes = value
}

// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *StateValidatorResponse) SetOwnerRole(value Substateable) {
	m.owner_role = value
}

// SetProtocolUpdateReadinessSignal sets the protocol_update_readiness_signal property value. The protocol_update_readiness_signal property
func (m *StateValidatorResponse) SetProtocolUpdateReadinessSignal(value Substateable) {
	m.protocol_update_readiness_signal = value
}

// SetState sets the state property value. The state property
func (m *StateValidatorResponse) SetState(value Substateable) {
	m.state = value
}

// SetVaults sets the vaults property value. Any vaults owned directly or indirectly by the component
func (m *StateValidatorResponse) SetVaults(value []VaultBalanceable) {
	m.vaults = value
}

type StateValidatorResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() *string
	GetAtLedgerState() LedgerStateSummaryable
	GetDescendentNodes() []StateComponentDescendentNodeable
	GetOwnerRole() Substateable
	GetProtocolUpdateReadinessSignal() Substateable
	GetState() Substateable
	GetVaults() []VaultBalanceable
	SetAddress(value *string)
	SetAtLedgerState(value LedgerStateSummaryable)
	SetDescendentNodes(value []StateComponentDescendentNodeable)
	SetOwnerRole(value Substateable)
	SetProtocolUpdateReadinessSignal(value Substateable)
	SetState(value Substateable)
	SetVaults(value []VaultBalanceable)
}
