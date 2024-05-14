package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalNonFungibleKey struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The non_fungible_local_id property
	non_fungible_local_id NonFungibleLocalIdable
}

// NewLocalNonFungibleKey instantiates a new LocalNonFungibleKey and sets the default values.
func NewLocalNonFungibleKey() *LocalNonFungibleKey {
	m := &LocalNonFungibleKey{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalNonFungibleKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalNonFungibleKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalNonFungibleKey(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalNonFungibleKey) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalNonFungibleKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["non_fungible_local_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleLocalIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNonFungibleLocalId(val.(NonFungibleLocalIdable))
		}
		return nil
	}
	return res
}

// GetNonFungibleLocalId gets the non_fungible_local_id property value. The non_fungible_local_id property
// returns a NonFungibleLocalIdable when successful
func (m *LocalNonFungibleKey) GetNonFungibleLocalId() NonFungibleLocalIdable {
	return m.non_fungible_local_id
}

// Serialize serializes information the current object
func (m *LocalNonFungibleKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("non_fungible_local_id", m.GetNonFungibleLocalId())
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
func (m *LocalNonFungibleKey) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNonFungibleLocalId sets the non_fungible_local_id property value. The non_fungible_local_id property
func (m *LocalNonFungibleKey) SetNonFungibleLocalId(value NonFungibleLocalIdable) {
	m.non_fungible_local_id = value
}

type LocalNonFungibleKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNonFungibleLocalId() NonFungibleLocalIdable
	SetNonFungibleLocalId(value NonFungibleLocalIdable)
}
