package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResourceAmount struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The Bech32m-encoded human readable version of the resource address
	resource_address *string
	// The resource_type property
	resource_type *ResourceType
}

// NewResourceAmount instantiates a new ResourceAmount and sets the default values.
func NewResourceAmount() *ResourceAmount {
	m := &ResourceAmount{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateResourceAmountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceAmountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("resource_type")
		if err != nil {
			return nil, err
		}
		if mappingValueNode != nil {
			mappingValue, err := mappingValueNode.GetStringValue()
			if err != nil {
				return nil, err
			}
			if mappingValue != nil {
				switch *mappingValue {
				case "Fungible":
					return NewFungibleResourceAmount(), nil
				case "NonFungible":
					return NewNonFungibleResourceAmount(), nil
				}
			}
		}
	}
	return NewResourceAmount(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ResourceAmount) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResourceAmount) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["resource_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseResourceType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResourceType(val.(*ResourceType))
		}
		return nil
	}
	return res
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ResourceAmount) GetResourceAddress() *string {
	return m.resource_address
}

// GetResourceType gets the resource_type property value. The resource_type property
// returns a *ResourceType when successful
func (m *ResourceAmount) GetResourceType() *ResourceType {
	return m.resource_type
}

// Serialize serializes information the current object
func (m *ResourceAmount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
		if err != nil {
			return err
		}
	}
	if m.GetResourceType() != nil {
		cast := (*m.GetResourceType()).String()
		err := writer.WriteStringValue("resource_type", &cast)
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
func (m *ResourceAmount) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *ResourceAmount) SetResourceAddress(value *string) {
	m.resource_address = value
}

// SetResourceType sets the resource_type property value. The resource_type property
func (m *ResourceAmount) SetResourceType(value *ResourceType) {
	m.resource_type = value
}

type ResourceAmountable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetResourceAddress() *string
	GetResourceType() *ResourceType
	SetResourceAddress(value *string)
	SetResourceType(value *ResourceType)
}
