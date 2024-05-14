package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleVaultFieldFrozenStatusSubstate struct {
	Substate
	// The value property
	value NonFungibleVaultFieldFrozenStatusValueable
}

// NewNonFungibleVaultFieldFrozenStatusSubstate instantiates a new NonFungibleVaultFieldFrozenStatusSubstate and sets the default values.
func NewNonFungibleVaultFieldFrozenStatusSubstate() *NonFungibleVaultFieldFrozenStatusSubstate {
	m := &NonFungibleVaultFieldFrozenStatusSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateNonFungibleVaultFieldFrozenStatusSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleVaultFieldFrozenStatusSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleVaultFieldFrozenStatusSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleVaultFieldFrozenStatusSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleVaultFieldFrozenStatusValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(NonFungibleVaultFieldFrozenStatusValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a NonFungibleVaultFieldFrozenStatusValueable when successful
func (m *NonFungibleVaultFieldFrozenStatusSubstate) GetValue() NonFungibleVaultFieldFrozenStatusValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *NonFungibleVaultFieldFrozenStatusSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *NonFungibleVaultFieldFrozenStatusSubstate) SetValue(value NonFungibleVaultFieldFrozenStatusValueable) {
	m.value = value
}

type NonFungibleVaultFieldFrozenStatusSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() NonFungibleVaultFieldFrozenStatusValueable
	SetValue(value NonFungibleVaultFieldFrozenStatusValueable)
}
