package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStateAccountFungibleResourceBalanceRequest struct {
	// The Bech32m-encoded human readable version of the account's address
	account_address *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The logical name of the network
	network *string
	// The Bech32m-encoded human readable version of the resource address
	resource_address *string
}

// NewLtsStateAccountFungibleResourceBalanceRequest instantiates a new LtsStateAccountFungibleResourceBalanceRequest and sets the default values.
func NewLtsStateAccountFungibleResourceBalanceRequest() *LtsStateAccountFungibleResourceBalanceRequest {
	m := &LtsStateAccountFungibleResourceBalanceRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStateAccountFungibleResourceBalanceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStateAccountFungibleResourceBalanceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStateAccountFungibleResourceBalanceRequest(), nil
}

// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address
// returns a *string when successful
func (m *LtsStateAccountFungibleResourceBalanceRequest) GetAccountAddress() *string {
	return m.account_address
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStateAccountFungibleResourceBalanceRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStateAccountFungibleResourceBalanceRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["account_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccountAddress(val)
		}
		return nil
	}
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
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
	return res
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *LtsStateAccountFungibleResourceBalanceRequest) GetNetwork() *string {
	return m.network
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *LtsStateAccountFungibleResourceBalanceRequest) GetResourceAddress() *string {
	return m.resource_address
}

// Serialize serializes information the current object
func (m *LtsStateAccountFungibleResourceBalanceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("account_address", m.GetAccountAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
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
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAccountAddress sets the account_address property value. The Bech32m-encoded human readable version of the account's address
func (m *LtsStateAccountFungibleResourceBalanceRequest) SetAccountAddress(value *string) {
	m.account_address = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStateAccountFungibleResourceBalanceRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *LtsStateAccountFungibleResourceBalanceRequest) SetNetwork(value *string) {
	m.network = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *LtsStateAccountFungibleResourceBalanceRequest) SetResourceAddress(value *string) {
	m.resource_address = value
}

type LtsStateAccountFungibleResourceBalanceRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountAddress() *string
	GetNetwork() *string
	GetResourceAddress() *string
	SetAccountAddress(value *string)
	SetNetwork(value *string)
	SetResourceAddress(value *string)
}
