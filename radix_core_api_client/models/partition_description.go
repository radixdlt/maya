package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PartitionDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *PartitionDescriptionType
    // An absolute or relative partition description, depending on the `type`:- if `Physical`, then this is a partition number,- if `Logical`, then this is a partition offset.
    value *int32
}
// NewPartitionDescription instantiates a new PartitionDescription and sets the default values.
func NewPartitionDescription()(*PartitionDescription) {
    m := &PartitionDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePartitionDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartitionDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPartitionDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PartitionDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartitionDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePartitionDescriptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PartitionDescriptionType))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *PartitionDescriptionType when successful
func (m *PartitionDescription) GetTypeEscaped()(*PartitionDescriptionType) {
    return m.typeEscaped
}
// GetValue gets the value property value. An absolute or relative partition description, depending on the `type`:- if `Physical`, then this is a partition number,- if `Logical`, then this is a partition offset.
// returns a *int32 when successful
func (m *PartitionDescription) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *PartitionDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *PartitionDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PartitionDescription) SetTypeEscaped(value *PartitionDescriptionType)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. An absolute or relative partition description, depending on the `type`:- if `Physical`, then this is a partition number,- if `Logical`, then this is a partition offset.
func (m *PartitionDescription) SetValue(value *int32)() {
    m.value = value
}
type PartitionDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*PartitionDescriptionType)
    GetValue()(*int32)
    SetTypeEscaped(value *PartitionDescriptionType)()
    SetValue(value *int32)()
}
