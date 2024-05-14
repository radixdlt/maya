package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BlueprintSchemaCollectionPartition the fields partition of the blueprint.
type BlueprintSchemaCollectionPartition struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The collection_schema property
	collection_schema BlueprintCollectionSchemaable
	// The partition_description property
	partition_description PartitionDescriptionable
}

// NewBlueprintSchemaCollectionPartition instantiates a new BlueprintSchemaCollectionPartition and sets the default values.
func NewBlueprintSchemaCollectionPartition() *BlueprintSchemaCollectionPartition {
	m := &BlueprintSchemaCollectionPartition{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBlueprintSchemaCollectionPartitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintSchemaCollectionPartitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBlueprintSchemaCollectionPartition(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintSchemaCollectionPartition) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCollectionSchema gets the collection_schema property value. The collection_schema property
// returns a BlueprintCollectionSchemaable when successful
func (m *BlueprintSchemaCollectionPartition) GetCollectionSchema() BlueprintCollectionSchemaable {
	return m.collection_schema
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintSchemaCollectionPartition) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["collection_schema"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintCollectionSchemaFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCollectionSchema(val.(BlueprintCollectionSchemaable))
		}
		return nil
	}
	res["partition_description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePartitionDescriptionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartitionDescription(val.(PartitionDescriptionable))
		}
		return nil
	}
	return res
}

// GetPartitionDescription gets the partition_description property value. The partition_description property
// returns a PartitionDescriptionable when successful
func (m *BlueprintSchemaCollectionPartition) GetPartitionDescription() PartitionDescriptionable {
	return m.partition_description
}

// Serialize serializes information the current object
func (m *BlueprintSchemaCollectionPartition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("collection_schema", m.GetCollectionSchema())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("partition_description", m.GetPartitionDescription())
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
func (m *BlueprintSchemaCollectionPartition) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCollectionSchema sets the collection_schema property value. The collection_schema property
func (m *BlueprintSchemaCollectionPartition) SetCollectionSchema(value BlueprintCollectionSchemaable) {
	m.collection_schema = value
}

// SetPartitionDescription sets the partition_description property value. The partition_description property
func (m *BlueprintSchemaCollectionPartition) SetPartitionDescription(value PartitionDescriptionable) {
	m.partition_description = value
}

type BlueprintSchemaCollectionPartitionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCollectionSchema() BlueprintCollectionSchemaable
	GetPartitionDescription() PartitionDescriptionable
	SetCollectionSchema(value BlueprintCollectionSchemaable)
	SetPartitionDescription(value PartitionDescriptionable)
}
