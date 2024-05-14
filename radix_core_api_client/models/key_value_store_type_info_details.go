package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KeyValueStoreTypeInfoDetails struct {
	TypeInfoDetails
	// The key_value_store_info property
	key_value_store_info KeyValueStoreInfoable
}

// NewKeyValueStoreTypeInfoDetails instantiates a new KeyValueStoreTypeInfoDetails and sets the default values.
func NewKeyValueStoreTypeInfoDetails() *KeyValueStoreTypeInfoDetails {
	m := &KeyValueStoreTypeInfoDetails{
		TypeInfoDetails: *NewTypeInfoDetails(),
	}
	return m
}

// CreateKeyValueStoreTypeInfoDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeyValueStoreTypeInfoDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewKeyValueStoreTypeInfoDetails(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeyValueStoreTypeInfoDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TypeInfoDetails.GetFieldDeserializers()
	res["key_value_store_info"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateKeyValueStoreInfoFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyValueStoreInfo(val.(KeyValueStoreInfoable))
		}
		return nil
	}
	return res
}

// GetKeyValueStoreInfo gets the key_value_store_info property value. The key_value_store_info property
// returns a KeyValueStoreInfoable when successful
func (m *KeyValueStoreTypeInfoDetails) GetKeyValueStoreInfo() KeyValueStoreInfoable {
	return m.key_value_store_info
}

// Serialize serializes information the current object
func (m *KeyValueStoreTypeInfoDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TypeInfoDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("key_value_store_info", m.GetKeyValueStoreInfo())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKeyValueStoreInfo sets the key_value_store_info property value. The key_value_store_info property
func (m *KeyValueStoreTypeInfoDetails) SetKeyValueStoreInfo(value KeyValueStoreInfoable) {
	m.key_value_store_info = value
}

type KeyValueStoreTypeInfoDetailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TypeInfoDetailsable
	GetKeyValueStoreInfo() KeyValueStoreInfoable
	SetKeyValueStoreInfo(value KeyValueStoreInfoable)
}
