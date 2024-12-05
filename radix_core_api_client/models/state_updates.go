package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StateUpdates transaction state updates (only present if status is Succeeded or Failed)
type StateUpdates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_substates property
    created_substates []CreatedSubstateable
    // The deleted_partitions property
    deleted_partitions []PartitionIdable
    // The deleted_substates property
    deleted_substates []DeletedSubstateable
    // The new_global_entities property
    new_global_entities []EntityReferenceable
    // The updated_substates property
    updated_substates []UpdatedSubstateable
}
// NewStateUpdates instantiates a new StateUpdates and sets the default values.
func NewStateUpdates()(*StateUpdates) {
    m := &StateUpdates{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateUpdatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateUpdates(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateUpdates) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedSubstates gets the created_substates property value. The created_substates property
// returns a []CreatedSubstateable when successful
func (m *StateUpdates) GetCreatedSubstates()([]CreatedSubstateable) {
    return m.created_substates
}
// GetDeletedPartitions gets the deleted_partitions property value. The deleted_partitions property
// returns a []PartitionIdable when successful
func (m *StateUpdates) GetDeletedPartitions()([]PartitionIdable) {
    return m.deleted_partitions
}
// GetDeletedSubstates gets the deleted_substates property value. The deleted_substates property
// returns a []DeletedSubstateable when successful
func (m *StateUpdates) GetDeletedSubstates()([]DeletedSubstateable) {
    return m.deleted_substates
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateUpdates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_substates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCreatedSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CreatedSubstateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CreatedSubstateable)
                }
            }
            m.SetCreatedSubstates(res)
        }
        return nil
    }
    res["deleted_partitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePartitionIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PartitionIdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PartitionIdable)
                }
            }
            m.SetDeletedPartitions(res)
        }
        return nil
    }
    res["deleted_substates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedSubstateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedSubstateable)
                }
            }
            m.SetDeletedSubstates(res)
        }
        return nil
    }
    res["new_global_entities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EntityReferenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EntityReferenceable)
                }
            }
            m.SetNewGlobalEntities(res)
        }
        return nil
    }
    res["updated_substates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatedSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatedSubstateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatedSubstateable)
                }
            }
            m.SetUpdatedSubstates(res)
        }
        return nil
    }
    return res
}
// GetNewGlobalEntities gets the new_global_entities property value. The new_global_entities property
// returns a []EntityReferenceable when successful
func (m *StateUpdates) GetNewGlobalEntities()([]EntityReferenceable) {
    return m.new_global_entities
}
// GetUpdatedSubstates gets the updated_substates property value. The updated_substates property
// returns a []UpdatedSubstateable when successful
func (m *StateUpdates) GetUpdatedSubstates()([]UpdatedSubstateable) {
    return m.updated_substates
}
// Serialize serializes information the current object
func (m *StateUpdates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCreatedSubstates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCreatedSubstates()))
        for i, v := range m.GetCreatedSubstates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("created_substates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedPartitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedPartitions()))
        for i, v := range m.GetDeletedPartitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("deleted_partitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedSubstates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedSubstates()))
        for i, v := range m.GetDeletedSubstates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("deleted_substates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNewGlobalEntities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNewGlobalEntities()))
        for i, v := range m.GetNewGlobalEntities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("new_global_entities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatedSubstates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdatedSubstates()))
        for i, v := range m.GetUpdatedSubstates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("updated_substates", cast)
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
func (m *StateUpdates) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedSubstates sets the created_substates property value. The created_substates property
func (m *StateUpdates) SetCreatedSubstates(value []CreatedSubstateable)() {
    m.created_substates = value
}
// SetDeletedPartitions sets the deleted_partitions property value. The deleted_partitions property
func (m *StateUpdates) SetDeletedPartitions(value []PartitionIdable)() {
    m.deleted_partitions = value
}
// SetDeletedSubstates sets the deleted_substates property value. The deleted_substates property
func (m *StateUpdates) SetDeletedSubstates(value []DeletedSubstateable)() {
    m.deleted_substates = value
}
// SetNewGlobalEntities sets the new_global_entities property value. The new_global_entities property
func (m *StateUpdates) SetNewGlobalEntities(value []EntityReferenceable)() {
    m.new_global_entities = value
}
// SetUpdatedSubstates sets the updated_substates property value. The updated_substates property
func (m *StateUpdates) SetUpdatedSubstates(value []UpdatedSubstateable)() {
    m.updated_substates = value
}
type StateUpdatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedSubstates()([]CreatedSubstateable)
    GetDeletedPartitions()([]PartitionIdable)
    GetDeletedSubstates()([]DeletedSubstateable)
    GetNewGlobalEntities()([]EntityReferenceable)
    GetUpdatedSubstates()([]UpdatedSubstateable)
    SetCreatedSubstates(value []CreatedSubstateable)()
    SetDeletedPartitions(value []PartitionIdable)()
    SetDeletedSubstates(value []DeletedSubstateable)()
    SetNewGlobalEntities(value []EntityReferenceable)()
    SetUpdatedSubstates(value []UpdatedSubstateable)()
}
