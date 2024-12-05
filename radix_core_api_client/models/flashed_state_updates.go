package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FlashedStateUpdates direct state updates performed by a Flash Transaction.
type FlashedStateUpdates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deleted_partitions property
    deleted_partitions []PartitionIdable
    // The deleted_substates property
    deleted_substates []SubstateIdable
    // The set_substates property
    set_substates []FlashSetSubstateable
}
// NewFlashedStateUpdates instantiates a new FlashedStateUpdates and sets the default values.
func NewFlashedStateUpdates()(*FlashedStateUpdates) {
    m := &FlashedStateUpdates{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFlashedStateUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFlashedStateUpdatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFlashedStateUpdates(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FlashedStateUpdates) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeletedPartitions gets the deleted_partitions property value. The deleted_partitions property
// returns a []PartitionIdable when successful
func (m *FlashedStateUpdates) GetDeletedPartitions()([]PartitionIdable) {
    return m.deleted_partitions
}
// GetDeletedSubstates gets the deleted_substates property value. The deleted_substates property
// returns a []SubstateIdable when successful
func (m *FlashedStateUpdates) GetDeletedSubstates()([]SubstateIdable) {
    return m.deleted_substates
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FlashedStateUpdates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetCollectionOfObjectValues(CreateSubstateIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubstateIdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubstateIdable)
                }
            }
            m.SetDeletedSubstates(res)
        }
        return nil
    }
    res["set_substates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFlashSetSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FlashSetSubstateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FlashSetSubstateable)
                }
            }
            m.SetSetSubstates(res)
        }
        return nil
    }
    return res
}
// GetSetSubstates gets the set_substates property value. The set_substates property
// returns a []FlashSetSubstateable when successful
func (m *FlashedStateUpdates) GetSetSubstates()([]FlashSetSubstateable) {
    return m.set_substates
}
// Serialize serializes information the current object
func (m *FlashedStateUpdates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSetSubstates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSetSubstates()))
        for i, v := range m.GetSetSubstates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("set_substates", cast)
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
func (m *FlashedStateUpdates) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeletedPartitions sets the deleted_partitions property value. The deleted_partitions property
func (m *FlashedStateUpdates) SetDeletedPartitions(value []PartitionIdable)() {
    m.deleted_partitions = value
}
// SetDeletedSubstates sets the deleted_substates property value. The deleted_substates property
func (m *FlashedStateUpdates) SetDeletedSubstates(value []SubstateIdable)() {
    m.deleted_substates = value
}
// SetSetSubstates sets the set_substates property value. The set_substates property
func (m *FlashedStateUpdates) SetSetSubstates(value []FlashSetSubstateable)() {
    m.set_substates = value
}
type FlashedStateUpdatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedPartitions()([]PartitionIdable)
    GetDeletedSubstates()([]SubstateIdable)
    GetSetSubstates()([]FlashSetSubstateable)
    SetDeletedPartitions(value []PartitionIdable)()
    SetDeletedSubstates(value []SubstateIdable)()
    SetSetSubstates(value []FlashSetSubstateable)()
}
