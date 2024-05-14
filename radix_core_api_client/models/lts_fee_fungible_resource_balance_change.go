package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsFeeFungibleResourceBalanceChange struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The string-encoded decimal representing the amount of change for the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	balance_change *string
	// The Bech32m-encoded human readable version of the fungible resource's address
	resource_address *string
	// The type property
	typeEscaped *LtsFeeFungibleResourceBalanceChangeType
}

// NewLtsFeeFungibleResourceBalanceChange instantiates a new LtsFeeFungibleResourceBalanceChange and sets the default values.
func NewLtsFeeFungibleResourceBalanceChange() *LtsFeeFungibleResourceBalanceChange {
	m := &LtsFeeFungibleResourceBalanceChange{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsFeeFungibleResourceBalanceChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsFeeFungibleResourceBalanceChangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsFeeFungibleResourceBalanceChange(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsFeeFungibleResourceBalanceChange) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBalanceChange gets the balance_change property value. The string-encoded decimal representing the amount of change for the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *LtsFeeFungibleResourceBalanceChange) GetBalanceChange() *string {
	return m.balance_change
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsFeeFungibleResourceBalanceChange) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["balance_change"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBalanceChange(val)
		}
		return nil
	}
	res["resource_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResourceAddress(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseLtsFeeFungibleResourceBalanceChangeType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*LtsFeeFungibleResourceBalanceChangeType))
		}
		return nil
	}
	return res
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the fungible resource's address
// returns a *string when successful
func (m *LtsFeeFungibleResourceBalanceChange) GetResourceAddress() *string {
	return m.resource_address
}

// GetTypeEscaped gets the type property value. The type property
// returns a *LtsFeeFungibleResourceBalanceChangeType when successful
func (m *LtsFeeFungibleResourceBalanceChange) GetTypeEscaped() *LtsFeeFungibleResourceBalanceChangeType {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *LtsFeeFungibleResourceBalanceChange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("balance_change", m.GetBalanceChange())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
		if err != nil {
			return err
		}
	}
	if m.GetTypeEscaped() != nil {
		cast := (*m.GetTypeEscaped()).String()
		err := writer.WriteStringValue("type", &cast)
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
func (m *LtsFeeFungibleResourceBalanceChange) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBalanceChange sets the balance_change property value. The string-encoded decimal representing the amount of change for the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *LtsFeeFungibleResourceBalanceChange) SetBalanceChange(value *string) {
	m.balance_change = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the fungible resource's address
func (m *LtsFeeFungibleResourceBalanceChange) SetResourceAddress(value *string) {
	m.resource_address = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *LtsFeeFungibleResourceBalanceChange) SetTypeEscaped(value *LtsFeeFungibleResourceBalanceChangeType) {
	m.typeEscaped = value
}

type LtsFeeFungibleResourceBalanceChangeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBalanceChange() *string
	GetResourceAddress() *string
	GetTypeEscaped() *LtsFeeFungibleResourceBalanceChangeType
	SetBalanceChange(value *string)
	SetResourceAddress(value *string)
	SetTypeEscaped(value *LtsFeeFungibleResourceBalanceChangeType)
}
