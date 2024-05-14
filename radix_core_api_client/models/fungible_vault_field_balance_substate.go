package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleVaultFieldBalanceSubstate struct {
	Substate
	// The value property
	value FungibleVaultFieldBalanceValueable
}

// NewFungibleVaultFieldBalanceSubstate instantiates a new FungibleVaultFieldBalanceSubstate and sets the default values.
func NewFungibleVaultFieldBalanceSubstate() *FungibleVaultFieldBalanceSubstate {
	m := &FungibleVaultFieldBalanceSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateFungibleVaultFieldBalanceSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleVaultFieldBalanceSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFungibleVaultFieldBalanceSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleVaultFieldBalanceSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFungibleVaultFieldBalanceValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(FungibleVaultFieldBalanceValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a FungibleVaultFieldBalanceValueable when successful
func (m *FungibleVaultFieldBalanceSubstate) GetValue() FungibleVaultFieldBalanceValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *FungibleVaultFieldBalanceSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Substate.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. The value property
func (m *FungibleVaultFieldBalanceSubstate) SetValue(value FungibleVaultFieldBalanceValueable) {
	m.value = value
}

type FungibleVaultFieldBalanceSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() FungibleVaultFieldBalanceValueable
	SetValue(value FungibleVaultFieldBalanceValueable)
}
