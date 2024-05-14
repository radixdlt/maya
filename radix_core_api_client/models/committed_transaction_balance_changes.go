package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommittedTransactionBalanceChanges struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
	fungible_entity_balance_changes []LtsEntityFungibleBalanceChangesable
	// Non fungible changes per entity and resource
	non_fungible_entity_balance_changes []LtsEntityNonFungibleBalanceChangesable
}

// NewCommittedTransactionBalanceChanges instantiates a new CommittedTransactionBalanceChanges and sets the default values.
func NewCommittedTransactionBalanceChanges() *CommittedTransactionBalanceChanges {
	m := &CommittedTransactionBalanceChanges{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCommittedTransactionBalanceChangesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommittedTransactionBalanceChangesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCommittedTransactionBalanceChanges(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommittedTransactionBalanceChanges) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommittedTransactionBalanceChanges) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["fungible_entity_balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsEntityFungibleBalanceChangesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsEntityFungibleBalanceChangesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsEntityFungibleBalanceChangesable)
				}
			}
			m.SetFungibleEntityBalanceChanges(res)
		}
		return nil
	}
	res["non_fungible_entity_balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsEntityNonFungibleBalanceChangesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsEntityNonFungibleBalanceChangesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsEntityNonFungibleBalanceChangesable)
				}
			}
			m.SetNonFungibleEntityBalanceChanges(res)
		}
		return nil
	}
	return res
}

// GetFungibleEntityBalanceChanges gets the fungible_entity_balance_changes property value. A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
// returns a []LtsEntityFungibleBalanceChangesable when successful
func (m *CommittedTransactionBalanceChanges) GetFungibleEntityBalanceChanges() []LtsEntityFungibleBalanceChangesable {
	return m.fungible_entity_balance_changes
}

// GetNonFungibleEntityBalanceChanges gets the non_fungible_entity_balance_changes property value. Non fungible changes per entity and resource
// returns a []LtsEntityNonFungibleBalanceChangesable when successful
func (m *CommittedTransactionBalanceChanges) GetNonFungibleEntityBalanceChanges() []LtsEntityNonFungibleBalanceChangesable {
	return m.non_fungible_entity_balance_changes
}

// Serialize serializes information the current object
func (m *CommittedTransactionBalanceChanges) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetFungibleEntityBalanceChanges() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFungibleEntityBalanceChanges()))
		for i, v := range m.GetFungibleEntityBalanceChanges() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("fungible_entity_balance_changes", cast)
		if err != nil {
			return err
		}
	}
	if m.GetNonFungibleEntityBalanceChanges() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonFungibleEntityBalanceChanges()))
		for i, v := range m.GetNonFungibleEntityBalanceChanges() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("non_fungible_entity_balance_changes", cast)
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
func (m *CommittedTransactionBalanceChanges) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFungibleEntityBalanceChanges sets the fungible_entity_balance_changes property value. A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
func (m *CommittedTransactionBalanceChanges) SetFungibleEntityBalanceChanges(value []LtsEntityFungibleBalanceChangesable) {
	m.fungible_entity_balance_changes = value
}

// SetNonFungibleEntityBalanceChanges sets the non_fungible_entity_balance_changes property value. Non fungible changes per entity and resource
func (m *CommittedTransactionBalanceChanges) SetNonFungibleEntityBalanceChanges(value []LtsEntityNonFungibleBalanceChangesable) {
	m.non_fungible_entity_balance_changes = value
}

type CommittedTransactionBalanceChangesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFungibleEntityBalanceChanges() []LtsEntityFungibleBalanceChangesable
	GetNonFungibleEntityBalanceChanges() []LtsEntityNonFungibleBalanceChangesable
	SetFungibleEntityBalanceChanges(value []LtsEntityFungibleBalanceChangesable)
	SetNonFungibleEntityBalanceChanges(value []LtsEntityNonFungibleBalanceChangesable)
}
