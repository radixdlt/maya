package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleGlobalId struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The local_id property
	local_id NonFungibleLocalIdable
	// The Bech32m-encoded human readable version of the resource address
	resource_address *string
}

// NewNonFungibleGlobalId instantiates a new NonFungibleGlobalId and sets the default values.
func NewNonFungibleGlobalId() *NonFungibleGlobalId {
	m := &NonFungibleGlobalId{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNonFungibleGlobalIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleGlobalIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleGlobalId(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleGlobalId) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleGlobalId) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["local_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleLocalIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocalId(val.(NonFungibleLocalIdable))
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

// GetLocalId gets the local_id property value. The local_id property
// returns a NonFungibleLocalIdable when successful
func (m *NonFungibleGlobalId) GetLocalId() NonFungibleLocalIdable {
	return m.local_id
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *NonFungibleGlobalId) GetResourceAddress() *string {
	return m.resource_address
}

// Serialize serializes information the current object
func (m *NonFungibleGlobalId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("local_id", m.GetLocalId())
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

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NonFungibleGlobalId) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLocalId sets the local_id property value. The local_id property
func (m *NonFungibleGlobalId) SetLocalId(value NonFungibleLocalIdable) {
	m.local_id = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *NonFungibleGlobalId) SetResourceAddress(value *string) {
	m.resource_address = value
}

type NonFungibleGlobalIdable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLocalId() NonFungibleLocalIdable
	GetResourceAddress() *string
	SetLocalId(value NonFungibleLocalIdable)
	SetResourceAddress(value *string)
}
