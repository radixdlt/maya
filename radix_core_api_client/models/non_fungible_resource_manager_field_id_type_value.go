package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldIdTypeValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The non_fungible_id_type property
	non_fungible_id_type *NonFungibleIdType
}

// NewNonFungibleResourceManagerFieldIdTypeValue instantiates a new NonFungibleResourceManagerFieldIdTypeValue and sets the default values.
func NewNonFungibleResourceManagerFieldIdTypeValue() *NonFungibleResourceManagerFieldIdTypeValue {
	m := &NonFungibleResourceManagerFieldIdTypeValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNonFungibleResourceManagerFieldIdTypeValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldIdTypeValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleResourceManagerFieldIdTypeValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleResourceManagerFieldIdTypeValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldIdTypeValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["non_fungible_id_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseNonFungibleIdType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNonFungibleIdType(val.(*NonFungibleIdType))
		}
		return nil
	}
	return res
}

// GetNonFungibleIdType gets the non_fungible_id_type property value. The non_fungible_id_type property
// returns a *NonFungibleIdType when successful
func (m *NonFungibleResourceManagerFieldIdTypeValue) GetNonFungibleIdType() *NonFungibleIdType {
	return m.non_fungible_id_type
}

// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldIdTypeValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetNonFungibleIdType() != nil {
		cast := (*m.GetNonFungibleIdType()).String()
		err := writer.WriteStringValue("non_fungible_id_type", &cast)
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
func (m *NonFungibleResourceManagerFieldIdTypeValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNonFungibleIdType sets the non_fungible_id_type property value. The non_fungible_id_type property
func (m *NonFungibleResourceManagerFieldIdTypeValue) SetNonFungibleIdType(value *NonFungibleIdType) {
	m.non_fungible_id_type = value
}

type NonFungibleResourceManagerFieldIdTypeValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNonFungibleIdType() *NonFungibleIdType
	SetNonFungibleIdType(value *NonFungibleIdType)
}
