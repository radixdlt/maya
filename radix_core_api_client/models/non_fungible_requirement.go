package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleRequirement struct {
	Requirement
	// The non_fungible property
	non_fungible NonFungibleGlobalIdable
}

// NewNonFungibleRequirement instantiates a new NonFungibleRequirement and sets the default values.
func NewNonFungibleRequirement() *NonFungibleRequirement {
	m := &NonFungibleRequirement{
		Requirement: *NewRequirement(),
	}
	return m
}

// CreateNonFungibleRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleRequirement(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleRequirement) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Requirement.GetFieldDeserializers()
	res["non_fungible"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleGlobalIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNonFungible(val.(NonFungibleGlobalIdable))
		}
		return nil
	}
	return res
}

// GetNonFungible gets the non_fungible property value. The non_fungible property
// returns a NonFungibleGlobalIdable when successful
func (m *NonFungibleRequirement) GetNonFungible() NonFungibleGlobalIdable {
	return m.non_fungible
}

// Serialize serializes information the current object
func (m *NonFungibleRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Requirement.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("non_fungible", m.GetNonFungible())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetNonFungible sets the non_fungible property value. The non_fungible property
func (m *NonFungibleRequirement) SetNonFungible(value NonFungibleGlobalIdable) {
	m.non_fungible = value
}

type NonFungibleRequirementable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Requirementable
	GetNonFungible() NonFungibleGlobalIdable
	SetNonFungible(value NonFungibleGlobalIdable)
}
