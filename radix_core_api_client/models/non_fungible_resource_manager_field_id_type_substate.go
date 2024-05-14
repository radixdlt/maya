package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldIdTypeSubstate struct {
	Substate
	// The value property
	value NonFungibleResourceManagerFieldIdTypeValueable
}

// NewNonFungibleResourceManagerFieldIdTypeSubstate instantiates a new NonFungibleResourceManagerFieldIdTypeSubstate and sets the default values.
func NewNonFungibleResourceManagerFieldIdTypeSubstate() *NonFungibleResourceManagerFieldIdTypeSubstate {
	m := &NonFungibleResourceManagerFieldIdTypeSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateNonFungibleResourceManagerFieldIdTypeSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldIdTypeSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleResourceManagerFieldIdTypeSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldIdTypeSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleResourceManagerFieldIdTypeValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(NonFungibleResourceManagerFieldIdTypeValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a NonFungibleResourceManagerFieldIdTypeValueable when successful
func (m *NonFungibleResourceManagerFieldIdTypeSubstate) GetValue() NonFungibleResourceManagerFieldIdTypeValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldIdTypeSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *NonFungibleResourceManagerFieldIdTypeSubstate) SetValue(value NonFungibleResourceManagerFieldIdTypeValueable) {
	m.value = value
}

type NonFungibleResourceManagerFieldIdTypeSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() NonFungibleResourceManagerFieldIdTypeValueable
	SetValue(value NonFungibleResourceManagerFieldIdTypeValueable)
}
