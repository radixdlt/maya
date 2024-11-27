package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateComponentDescendentNode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Depth under component
    depth *int32
    // The entity property
    entity EntityReferenceable
    // The parent_entity property
    parent_entity EntityReferenceable
    // The parent_partition_number property
    parent_partition_number *int32
    // The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
    parent_substate_db_sort_key_hex *string
    // The hex-encoded bytes of the substate key, under the entity partition
    parent_substate_key_hex *string
}
// NewStateComponentDescendentNode instantiates a new StateComponentDescendentNode and sets the default values.
func NewStateComponentDescendentNode()(*StateComponentDescendentNode) {
    m := &StateComponentDescendentNode{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateComponentDescendentNodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateComponentDescendentNodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateComponentDescendentNode(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateComponentDescendentNode) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDepth gets the depth property value. Depth under component
// returns a *int32 when successful
func (m *StateComponentDescendentNode) GetDepth()(*int32) {
    return m.depth
}
// GetEntity gets the entity property value. The entity property
// returns a EntityReferenceable when successful
func (m *StateComponentDescendentNode) GetEntity()(EntityReferenceable) {
    return m.entity
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateComponentDescendentNode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["depth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepth(val)
        }
        return nil
    }
    res["entity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntity(val.(EntityReferenceable))
        }
        return nil
    }
    res["parent_entity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentEntity(val.(EntityReferenceable))
        }
        return nil
    }
    res["parent_partition_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentPartitionNumber(val)
        }
        return nil
    }
    res["parent_substate_db_sort_key_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSubstateDbSortKeyHex(val)
        }
        return nil
    }
    res["parent_substate_key_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSubstateKeyHex(val)
        }
        return nil
    }
    return res
}
// GetParentEntity gets the parent_entity property value. The parent_entity property
// returns a EntityReferenceable when successful
func (m *StateComponentDescendentNode) GetParentEntity()(EntityReferenceable) {
    return m.parent_entity
}
// GetParentPartitionNumber gets the parent_partition_number property value. The parent_partition_number property
// returns a *int32 when successful
func (m *StateComponentDescendentNode) GetParentPartitionNumber()(*int32) {
    return m.parent_partition_number
}
// GetParentSubstateDbSortKeyHex gets the parent_substate_db_sort_key_hex property value. The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
// returns a *string when successful
func (m *StateComponentDescendentNode) GetParentSubstateDbSortKeyHex()(*string) {
    return m.parent_substate_db_sort_key_hex
}
// GetParentSubstateKeyHex gets the parent_substate_key_hex property value. The hex-encoded bytes of the substate key, under the entity partition
// returns a *string when successful
func (m *StateComponentDescendentNode) GetParentSubstateKeyHex()(*string) {
    return m.parent_substate_key_hex
}
// Serialize serializes information the current object
func (m *StateComponentDescendentNode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("depth", m.GetDepth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entity", m.GetEntity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent_entity", m.GetParentEntity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("parent_partition_number", m.GetParentPartitionNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parent_substate_db_sort_key_hex", m.GetParentSubstateDbSortKeyHex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parent_substate_key_hex", m.GetParentSubstateKeyHex())
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
func (m *StateComponentDescendentNode) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDepth sets the depth property value. Depth under component
func (m *StateComponentDescendentNode) SetDepth(value *int32)() {
    m.depth = value
}
// SetEntity sets the entity property value. The entity property
func (m *StateComponentDescendentNode) SetEntity(value EntityReferenceable)() {
    m.entity = value
}
// SetParentEntity sets the parent_entity property value. The parent_entity property
func (m *StateComponentDescendentNode) SetParentEntity(value EntityReferenceable)() {
    m.parent_entity = value
}
// SetParentPartitionNumber sets the parent_partition_number property value. The parent_partition_number property
func (m *StateComponentDescendentNode) SetParentPartitionNumber(value *int32)() {
    m.parent_partition_number = value
}
// SetParentSubstateDbSortKeyHex sets the parent_substate_db_sort_key_hex property value. The hex-encoded bytes of the partially-hashed DB sort key, under the given entity partition
func (m *StateComponentDescendentNode) SetParentSubstateDbSortKeyHex(value *string)() {
    m.parent_substate_db_sort_key_hex = value
}
// SetParentSubstateKeyHex sets the parent_substate_key_hex property value. The hex-encoded bytes of the substate key, under the entity partition
func (m *StateComponentDescendentNode) SetParentSubstateKeyHex(value *string)() {
    m.parent_substate_key_hex = value
}
type StateComponentDescendentNodeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDepth()(*int32)
    GetEntity()(EntityReferenceable)
    GetParentEntity()(EntityReferenceable)
    GetParentPartitionNumber()(*int32)
    GetParentSubstateDbSortKeyHex()(*string)
    GetParentSubstateKeyHex()(*string)
    SetDepth(value *int32)()
    SetEntity(value EntityReferenceable)()
    SetParentEntity(value EntityReferenceable)()
    SetParentPartitionNumber(value *int32)()
    SetParentSubstateDbSortKeyHex(value *string)()
    SetParentSubstateKeyHex(value *string)()
}
