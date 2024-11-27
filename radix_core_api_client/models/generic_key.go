package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
    key_data SborDataable
}
// NewGenericKey instantiates a new GenericKey and sets the default values.
func NewGenericKey()(*GenericKey) {
    m := &GenericKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGenericKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenericKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key_data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyData(val.(SborDataable))
        }
        return nil
    }
    return res
}
// GetKeyData gets the key_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *GenericKey) GetKeyData()(SborDataable) {
    return m.key_data
}
// Serialize serializes information the current object
func (m *GenericKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("key_data", m.GetKeyData())
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
func (m *GenericKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyData sets the key_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *GenericKey) SetKeyData(value SborDataable)() {
    m.key_data = value
}
type GenericKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyData()(SborDataable)
    SetKeyData(value SborDataable)()
}
