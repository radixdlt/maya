package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IndexedStateSchema struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The collections property
	collections []BlueprintSchemaCollectionPartitionable
	// The fields partition of the blueprint.
	fields BlueprintSchemaFieldPartitionable
	// The num_partitions property
	num_partitions *int32
}

// NewIndexedStateSchema instantiates a new IndexedStateSchema and sets the default values.
func NewIndexedStateSchema() *IndexedStateSchema {
	m := &IndexedStateSchema{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateIndexedStateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIndexedStateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewIndexedStateSchema(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IndexedStateSchema) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCollections gets the collections property value. The collections property
// returns a []BlueprintSchemaCollectionPartitionable when successful
func (m *IndexedStateSchema) GetCollections() []BlueprintSchemaCollectionPartitionable {
	return m.collections
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IndexedStateSchema) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["collections"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateBlueprintSchemaCollectionPartitionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]BlueprintSchemaCollectionPartitionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(BlueprintSchemaCollectionPartitionable)
				}
			}
			m.SetCollections(res)
		}
		return nil
	}
	res["fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintSchemaFieldPartitionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFields(val.(BlueprintSchemaFieldPartitionable))
		}
		return nil
	}
	res["num_partitions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNumPartitions(val)
		}
		return nil
	}
	return res
}

// GetFields gets the fields property value. The fields partition of the blueprint.
// returns a BlueprintSchemaFieldPartitionable when successful
func (m *IndexedStateSchema) GetFields() BlueprintSchemaFieldPartitionable {
	return m.fields
}

// GetNumPartitions gets the num_partitions property value. The num_partitions property
// returns a *int32 when successful
func (m *IndexedStateSchema) GetNumPartitions() *int32 {
	return m.num_partitions
}

// Serialize serializes information the current object
func (m *IndexedStateSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCollections() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollections()))
		for i, v := range m.GetCollections() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("collections", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("fields", m.GetFields())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("num_partitions", m.GetNumPartitions())
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
func (m *IndexedStateSchema) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCollections sets the collections property value. The collections property
func (m *IndexedStateSchema) SetCollections(value []BlueprintSchemaCollectionPartitionable) {
	m.collections = value
}

// SetFields sets the fields property value. The fields partition of the blueprint.
func (m *IndexedStateSchema) SetFields(value BlueprintSchemaFieldPartitionable) {
	m.fields = value
}

// SetNumPartitions sets the num_partitions property value. The num_partitions property
func (m *IndexedStateSchema) SetNumPartitions(value *int32) {
	m.num_partitions = value
}

type IndexedStateSchemaable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCollections() []BlueprintSchemaCollectionPartitionable
	GetFields() BlueprintSchemaFieldPartitionable
	GetNumPartitions() *int32
	SetCollections(value []BlueprintSchemaCollectionPartitionable)
	SetFields(value BlueprintSchemaFieldPartitionable)
	SetNumPartitions(value *int32)
}
