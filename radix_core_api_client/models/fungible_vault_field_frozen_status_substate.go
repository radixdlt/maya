package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleVaultFieldFrozenStatusSubstate struct {
	Substate
	// The value property
	value FungibleVaultFieldFrozenStatusValueable
}

// NewFungibleVaultFieldFrozenStatusSubstate instantiates a new FungibleVaultFieldFrozenStatusSubstate and sets the default values.
func NewFungibleVaultFieldFrozenStatusSubstate() *FungibleVaultFieldFrozenStatusSubstate {
	m := &FungibleVaultFieldFrozenStatusSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateFungibleVaultFieldFrozenStatusSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleVaultFieldFrozenStatusSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFungibleVaultFieldFrozenStatusSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleVaultFieldFrozenStatusSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFungibleVaultFieldFrozenStatusValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(FungibleVaultFieldFrozenStatusValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a FungibleVaultFieldFrozenStatusValueable when successful
func (m *FungibleVaultFieldFrozenStatusSubstate) GetValue() FungibleVaultFieldFrozenStatusValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *FungibleVaultFieldFrozenStatusSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *FungibleVaultFieldFrozenStatusSubstate) SetValue(value FungibleVaultFieldFrozenStatusValueable) {
	m.value = value
}

type FungibleVaultFieldFrozenStatusSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() FungibleVaultFieldFrozenStatusValueable
	SetValue(value FungibleVaultFieldFrozenStatusValueable)
}
