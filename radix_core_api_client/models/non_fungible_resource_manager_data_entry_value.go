package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NonFungibleResourceManagerDataEntryValue if missing, it represents a burned Non-Fungible.A Non-Fungible with that local id cannot be minted again - the id is not re-usable.
type NonFungibleResourceManagerDataEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The data_struct property
	data_struct DataStructable
}

// NewNonFungibleResourceManagerDataEntryValue instantiates a new NonFungibleResourceManagerDataEntryValue and sets the default values.
func NewNonFungibleResourceManagerDataEntryValue() *NonFungibleResourceManagerDataEntryValue {
	m := &NonFungibleResourceManagerDataEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNonFungibleResourceManagerDataEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerDataEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleResourceManagerDataEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleResourceManagerDataEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDataStruct gets the data_struct property value. The data_struct property
// returns a DataStructable when successful
func (m *NonFungibleResourceManagerDataEntryValue) GetDataStruct() DataStructable {
	return m.data_struct
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerDataEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["data_struct"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateDataStructFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDataStruct(val.(DataStructable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *NonFungibleResourceManagerDataEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("data_struct", m.GetDataStruct())
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
func (m *NonFungibleResourceManagerDataEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDataStruct sets the data_struct property value. The data_struct property
func (m *NonFungibleResourceManagerDataEntryValue) SetDataStruct(value DataStructable) {
	m.data_struct = value
}

type NonFungibleResourceManagerDataEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDataStruct() DataStructable
	SetDataStruct(value DataStructable)
}
