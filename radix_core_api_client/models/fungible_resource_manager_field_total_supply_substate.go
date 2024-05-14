package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleResourceManagerFieldTotalSupplySubstate struct {
	Substate
	// The value property
	value FungibleResourceManagerFieldTotalSupplyValueable
}

// NewFungibleResourceManagerFieldTotalSupplySubstate instantiates a new FungibleResourceManagerFieldTotalSupplySubstate and sets the default values.
func NewFungibleResourceManagerFieldTotalSupplySubstate() *FungibleResourceManagerFieldTotalSupplySubstate {
	m := &FungibleResourceManagerFieldTotalSupplySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateFungibleResourceManagerFieldTotalSupplySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleResourceManagerFieldTotalSupplySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFungibleResourceManagerFieldTotalSupplySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleResourceManagerFieldTotalSupplySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFungibleResourceManagerFieldTotalSupplyValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(FungibleResourceManagerFieldTotalSupplyValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a FungibleResourceManagerFieldTotalSupplyValueable when successful
func (m *FungibleResourceManagerFieldTotalSupplySubstate) GetValue() FungibleResourceManagerFieldTotalSupplyValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *FungibleResourceManagerFieldTotalSupplySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *FungibleResourceManagerFieldTotalSupplySubstate) SetValue(value FungibleResourceManagerFieldTotalSupplyValueable) {
	m.value = value
}

type FungibleResourceManagerFieldTotalSupplySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() FungibleResourceManagerFieldTotalSupplyValueable
	SetValue(value FungibleResourceManagerFieldTotalSupplyValueable)
}
