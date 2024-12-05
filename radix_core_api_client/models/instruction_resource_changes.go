package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InstructionResourceChanges struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The index property
    index *int64
    // The resource_changes property
    resource_changes []ResourceChangeable
}
// NewInstructionResourceChanges instantiates a new InstructionResourceChanges and sets the default values.
func NewInstructionResourceChanges()(*InstructionResourceChanges) {
    m := &InstructionResourceChanges{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstructionResourceChangesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstructionResourceChangesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstructionResourceChanges(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InstructionResourceChanges) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InstructionResourceChanges) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["resource_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceChangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceChangeable)
                }
            }
            m.SetResourceChanges(res)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index property
// returns a *int64 when successful
func (m *InstructionResourceChanges) GetIndex()(*int64) {
    return m.index
}
// GetResourceChanges gets the resource_changes property value. The resource_changes property
// returns a []ResourceChangeable when successful
func (m *InstructionResourceChanges) GetResourceChanges()([]ResourceChangeable) {
    return m.resource_changes
}
// Serialize serializes information the current object
func (m *InstructionResourceChanges) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    if m.GetResourceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceChanges()))
        for i, v := range m.GetResourceChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resource_changes", cast)
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
func (m *InstructionResourceChanges) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndex sets the index property value. The index property
func (m *InstructionResourceChanges) SetIndex(value *int64)() {
    m.index = value
}
// SetResourceChanges sets the resource_changes property value. The resource_changes property
func (m *InstructionResourceChanges) SetResourceChanges(value []ResourceChangeable)() {
    m.resource_changes = value
}
type InstructionResourceChangesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndex()(*int64)
    GetResourceChanges()([]ResourceChangeable)
    SetIndex(value *int64)()
    SetResourceChanges(value []ResourceChangeable)()
}
