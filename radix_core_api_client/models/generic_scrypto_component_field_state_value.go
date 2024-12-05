package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericScryptoComponentFieldStateValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data_struct property
    data_struct DataStructable
}
// NewGenericScryptoComponentFieldStateValue instantiates a new GenericScryptoComponentFieldStateValue and sets the default values.
func NewGenericScryptoComponentFieldStateValue()(*GenericScryptoComponentFieldStateValue) {
    m := &GenericScryptoComponentFieldStateValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGenericScryptoComponentFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericScryptoComponentFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenericScryptoComponentFieldStateValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericScryptoComponentFieldStateValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDataStruct gets the data_struct property value. The data_struct property
// returns a DataStructable when successful
func (m *GenericScryptoComponentFieldStateValue) GetDataStruct()(DataStructable) {
    return m.data_struct
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericScryptoComponentFieldStateValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data_struct"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataStructFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataStruct(val.(DataStructable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GenericScryptoComponentFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data_struct", m.GetDataStruct())
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
func (m *GenericScryptoComponentFieldStateValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDataStruct sets the data_struct property value. The data_struct property
func (m *GenericScryptoComponentFieldStateValue) SetDataStruct(value DataStructable)() {
    m.data_struct = value
}
type GenericScryptoComponentFieldStateValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataStruct()(DataStructable)
    SetDataStruct(value DataStructable)()
}
