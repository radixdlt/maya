package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsEntityNonFungibleBalanceChanges struct {
	// The added property
	added []string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The Bech32m-encoded human readable version of the entity's address
	entity_address *string
	// The removed property
	removed []string
	// The Bech32m-encoded human readable version of the non-fungible resource's address
	resource_address *string
}

// NewLtsEntityNonFungibleBalanceChanges instantiates a new LtsEntityNonFungibleBalanceChanges and sets the default values.
func NewLtsEntityNonFungibleBalanceChanges() *LtsEntityNonFungibleBalanceChanges {
	m := &LtsEntityNonFungibleBalanceChanges{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsEntityNonFungibleBalanceChangesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsEntityNonFungibleBalanceChangesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsEntityNonFungibleBalanceChanges(), nil
}

// GetAdded gets the added property value. The added property
// returns a []string when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetAdded() []string {
	return m.added
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEntityAddress gets the entity_address property value. The Bech32m-encoded human readable version of the entity's address
// returns a *string when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetEntityAddress() *string {
	return m.entity_address
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["added"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetAdded(res)
		}
		return nil
	}
	res["entity_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEntityAddress(val)
		}
		return nil
	}
	res["removed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetRemoved(res)
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

// GetRemoved gets the removed property value. The removed property
// returns a []string when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetRemoved() []string {
	return m.removed
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the non-fungible resource's address
// returns a *string when successful
func (m *LtsEntityNonFungibleBalanceChanges) GetResourceAddress() *string {
	return m.resource_address
}

// Serialize serializes information the current object
func (m *LtsEntityNonFungibleBalanceChanges) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAdded() != nil {
		err := writer.WriteCollectionOfStringValues("added", m.GetAdded())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("entity_address", m.GetEntityAddress())
		if err != nil {
			return err
		}
	}
	if m.GetRemoved() != nil {
		err := writer.WriteCollectionOfStringValues("removed", m.GetRemoved())
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

// SetAdded sets the added property value. The added property
func (m *LtsEntityNonFungibleBalanceChanges) SetAdded(value []string) {
	m.added = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsEntityNonFungibleBalanceChanges) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEntityAddress sets the entity_address property value. The Bech32m-encoded human readable version of the entity's address
func (m *LtsEntityNonFungibleBalanceChanges) SetEntityAddress(value *string) {
	m.entity_address = value
}

// SetRemoved sets the removed property value. The removed property
func (m *LtsEntityNonFungibleBalanceChanges) SetRemoved(value []string) {
	m.removed = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the non-fungible resource's address
func (m *LtsEntityNonFungibleBalanceChanges) SetResourceAddress(value *string) {
	m.resource_address = value
}

type LtsEntityNonFungibleBalanceChangesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAdded() []string
	GetEntityAddress() *string
	GetRemoved() []string
	GetResourceAddress() *string
	SetAdded(value []string)
	SetEntityAddress(value *string)
	SetRemoved(value []string)
	SetResourceAddress(value *string)
}
