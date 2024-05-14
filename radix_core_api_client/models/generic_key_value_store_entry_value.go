package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GenericKeyValueStoreEntryValue if not present, the entry has been deleted.
type GenericKeyValueStoreEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The data property
	data DataStructable
}

// NewGenericKeyValueStoreEntryValue instantiates a new GenericKeyValueStoreEntryValue and sets the default values.
func NewGenericKeyValueStoreEntryValue() *GenericKeyValueStoreEntryValue {
	m := &GenericKeyValueStoreEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateGenericKeyValueStoreEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericKeyValueStoreEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenericKeyValueStoreEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericKeyValueStoreEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetData gets the data property value. The data property
// returns a DataStructable when successful
func (m *GenericKeyValueStoreEntryValue) GetData() DataStructable {
	return m.data
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericKeyValueStoreEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateDataStructFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetData(val.(DataStructable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *GenericKeyValueStoreEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("data", m.GetData())
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
func (m *GenericKeyValueStoreEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetData sets the data property value. The data property
func (m *GenericKeyValueStoreEntryValue) SetData(value DataStructable) {
	m.data = value
}

type GenericKeyValueStoreEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetData() DataStructable
	SetData(value DataStructable)
}
