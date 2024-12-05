package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BlueprintSchemaFieldPartition the fields partition of the blueprint.
type BlueprintSchemaFieldPartition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The field substates for this blueprint.
    fields []FieldSchemaable
    // The partition_description property
    partition_description PartitionDescriptionable
}
// NewBlueprintSchemaFieldPartition instantiates a new BlueprintSchemaFieldPartition and sets the default values.
func NewBlueprintSchemaFieldPartition()(*BlueprintSchemaFieldPartition) {
    m := &BlueprintSchemaFieldPartition{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintSchemaFieldPartitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintSchemaFieldPartitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintSchemaFieldPartition(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintSchemaFieldPartition) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintSchemaFieldPartition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFieldSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FieldSchemaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FieldSchemaable)
                }
            }
            m.SetFields(res)
        }
        return nil
    }
    res["partition_description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFields gets the fields property value. The field substates for this blueprint.
// returns a []FieldSchemaable when successful
func (m *BlueprintSchemaFieldPartition) GetFields()([]FieldSchemaable) {
    return m.fields
}
// GetPartitionDescription gets the partition_description property value. The partition_description property
// returns a PartitionDescriptionable when successful
func (m *BlueprintSchemaFieldPartition) GetPartitionDescription()(PartitionDescriptionable) {
    return m.partition_description
}
// Serialize serializes information the current object
func (m *BlueprintSchemaFieldPartition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
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
func (m *BlueprintSchemaFieldPartition) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFields sets the fields property value. The field substates for this blueprint.
func (m *BlueprintSchemaFieldPartition) SetFields(value []FieldSchemaable)() {
    m.fields = value
}
// SetPartitionDescription sets the partition_description property value. The partition_description property
func (m *BlueprintSchemaFieldPartition) SetPartitionDescription(value PartitionDescriptionable)() {
    m.partition_description = value
}
type BlueprintSchemaFieldPartitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFields()([]FieldSchemaable)
    GetPartitionDescription()(PartitionDescriptionable)
    SetFields(value []FieldSchemaable)()
    SetPartitionDescription(value PartitionDescriptionable)()
}
