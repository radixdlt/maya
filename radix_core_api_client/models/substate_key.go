package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubstateKey struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
	db_sort_key_hex *string
	// The key_type property
	key_type *SubstateKeyType
}

// NewSubstateKey instantiates a new SubstateKey and sets the default values.
func NewSubstateKey() *SubstateKey {
	m := &SubstateKey{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubstateKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("key_type")
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
				case "Field":
					return NewFieldSubstateKey(), nil
				case "Map":
					return NewMapSubstateKey(), nil
				case "Sorted":
					return NewSortedSubstateKey(), nil
				}
			}
		}
	}
	return NewSubstateKey(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubstateKey) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDbSortKeyHex gets the db_sort_key_hex property value. The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
// returns a *string when successful
func (m *SubstateKey) GetDbSortKeyHex() *string {
	return m.db_sort_key_hex
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstateKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["db_sort_key_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDbSortKeyHex(val)
		}
		return nil
	}
	res["key_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseSubstateKeyType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyType(val.(*SubstateKeyType))
		}
		return nil
	}
	return res
}

// GetKeyType gets the key_type property value. The key_type property
// returns a *SubstateKeyType when successful
func (m *SubstateKey) GetKeyType() *SubstateKeyType {
	return m.key_type
}

// Serialize serializes information the current object
func (m *SubstateKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("db_sort_key_hex", m.GetDbSortKeyHex())
		if err != nil {
			return err
		}
	}
	if m.GetKeyType() != nil {
		cast := (*m.GetKeyType()).String()
		err := writer.WriteStringValue("key_type", &cast)
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
func (m *SubstateKey) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDbSortKeyHex sets the db_sort_key_hex property value. The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
func (m *SubstateKey) SetDbSortKeyHex(value *string) {
	m.db_sort_key_hex = value
}

// SetKeyType sets the key_type property value. The key_type property
func (m *SubstateKey) SetKeyType(value *SubstateKeyType) {
	m.key_type = value
}

type SubstateKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDbSortKeyHex() *string
	GetKeyType() *SubstateKeyType
	SetDbSortKeyHex(value *string)
	SetKeyType(value *SubstateKeyType)
}
