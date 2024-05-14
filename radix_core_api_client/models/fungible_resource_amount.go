package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleResourceAmount struct {
	ResourceAmount
	// The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	amount *string
}

// NewFungibleResourceAmount instantiates a new FungibleResourceAmount and sets the default values.
func NewFungibleResourceAmount() *FungibleResourceAmount {
	m := &FungibleResourceAmount{
		ResourceAmount: *NewResourceAmount(),
	}
	return m
}

// CreateFungibleResourceAmountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleResourceAmountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFungibleResourceAmount(), nil
}

// GetAmount gets the amount property value. The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FungibleResourceAmount) GetAmount() *string {
	return m.amount
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleResourceAmount) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ResourceAmount.GetFieldDeserializers()
	res["amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAmount(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *FungibleResourceAmount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ResourceAmount.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("amount", m.GetAmount())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAmount sets the amount property value. The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FungibleResourceAmount) SetAmount(value *string) {
	m.amount = value
}

type FungibleResourceAmountable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ResourceAmountable
	GetAmount() *string
	SetAmount(value *string)
}
