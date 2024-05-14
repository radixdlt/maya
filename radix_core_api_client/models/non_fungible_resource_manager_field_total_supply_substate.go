package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldTotalSupplySubstate struct {
	Substate
	// The value property
	value NonFungibleResourceManagerFieldTotalSupplyValueable
}

// NewNonFungibleResourceManagerFieldTotalSupplySubstate instantiates a new NonFungibleResourceManagerFieldTotalSupplySubstate and sets the default values.
func NewNonFungibleResourceManagerFieldTotalSupplySubstate() *NonFungibleResourceManagerFieldTotalSupplySubstate {
	m := &NonFungibleResourceManagerFieldTotalSupplySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateNonFungibleResourceManagerFieldTotalSupplySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldTotalSupplySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleResourceManagerFieldTotalSupplySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldTotalSupplySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleResourceManagerFieldTotalSupplyValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(NonFungibleResourceManagerFieldTotalSupplyValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a NonFungibleResourceManagerFieldTotalSupplyValueable when successful
func (m *NonFungibleResourceManagerFieldTotalSupplySubstate) GetValue() NonFungibleResourceManagerFieldTotalSupplyValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldTotalSupplySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *NonFungibleResourceManagerFieldTotalSupplySubstate) SetValue(value NonFungibleResourceManagerFieldTotalSupplyValueable) {
	m.value = value
}

type NonFungibleResourceManagerFieldTotalSupplySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() NonFungibleResourceManagerFieldTotalSupplyValueable
	SetValue(value NonFungibleResourceManagerFieldTotalSupplyValueable)
}
