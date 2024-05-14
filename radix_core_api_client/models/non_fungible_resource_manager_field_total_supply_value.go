package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldTotalSupplyValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The string-encoded decimal representing the total supply of this resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	total_supply *string
}

// NewNonFungibleResourceManagerFieldTotalSupplyValue instantiates a new NonFungibleResourceManagerFieldTotalSupplyValue and sets the default values.
func NewNonFungibleResourceManagerFieldTotalSupplyValue() *NonFungibleResourceManagerFieldTotalSupplyValue {
	m := &NonFungibleResourceManagerFieldTotalSupplyValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNonFungibleResourceManagerFieldTotalSupplyValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldTotalSupplyValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleResourceManagerFieldTotalSupplyValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["total_supply"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotalSupply(val)
		}
		return nil
	}
	return res
}

// GetTotalSupply gets the total_supply property value. The string-encoded decimal representing the total supply of this resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) GetTotalSupply() *string {
	return m.total_supply
}

// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("total_supply", m.GetTotalSupply())
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
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetTotalSupply sets the total_supply property value. The string-encoded decimal representing the total supply of this resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *NonFungibleResourceManagerFieldTotalSupplyValue) SetTotalSupply(value *string) {
	m.total_supply = value
}

type NonFungibleResourceManagerFieldTotalSupplyValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetTotalSupply() *string
	SetTotalSupply(value *string)
}
