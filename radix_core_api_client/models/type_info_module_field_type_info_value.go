package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TypeInfoModuleFieldTypeInfoValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details property
    details TypeInfoDetailsable
}
// NewTypeInfoModuleFieldTypeInfoValue instantiates a new TypeInfoModuleFieldTypeInfoValue and sets the default values.
func NewTypeInfoModuleFieldTypeInfoValue()(*TypeInfoModuleFieldTypeInfoValue) {
    m := &TypeInfoModuleFieldTypeInfoValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTypeInfoModuleFieldTypeInfoValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTypeInfoModuleFieldTypeInfoValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTypeInfoModuleFieldTypeInfoValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TypeInfoModuleFieldTypeInfoValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. The details property
// returns a TypeInfoDetailsable when successful
func (m *TypeInfoModuleFieldTypeInfoValue) GetDetails()(TypeInfoDetailsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TypeInfoModuleFieldTypeInfoValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTypeInfoDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(TypeInfoDetailsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TypeInfoModuleFieldTypeInfoValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
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
func (m *TypeInfoModuleFieldTypeInfoValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. The details property
func (m *TypeInfoModuleFieldTypeInfoValue) SetDetails(value TypeInfoDetailsable)() {
    m.details = value
}
type TypeInfoModuleFieldTypeInfoValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()(TypeInfoDetailsable)
    SetDetails(value TypeInfoDetailsable)()
}
