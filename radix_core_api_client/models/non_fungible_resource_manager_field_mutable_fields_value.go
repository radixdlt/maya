package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldMutableFieldsValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fields of the NF Metadata which are mutable.
    mutable_fields []MutableFieldable
}
// NewNonFungibleResourceManagerFieldMutableFieldsValue instantiates a new NonFungibleResourceManagerFieldMutableFieldsValue and sets the default values.
func NewNonFungibleResourceManagerFieldMutableFieldsValue()(*NonFungibleResourceManagerFieldMutableFieldsValue) {
    m := &NonFungibleResourceManagerFieldMutableFieldsValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNonFungibleResourceManagerFieldMutableFieldsValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldMutableFieldsValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleResourceManagerFieldMutableFieldsValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mutable_fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMutableFieldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MutableFieldable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MutableFieldable)
                }
            }
            m.SetMutableFields(res)
        }
        return nil
    }
    return res
}
// GetMutableFields gets the mutable_fields property value. The fields of the NF Metadata which are mutable.
// returns a []MutableFieldable when successful
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) GetMutableFields()([]MutableFieldable) {
    return m.mutable_fields
}
// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMutableFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMutableFields()))
        for i, v := range m.GetMutableFields() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("mutable_fields", cast)
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
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMutableFields sets the mutable_fields property value. The fields of the NF Metadata which are mutable.
func (m *NonFungibleResourceManagerFieldMutableFieldsValue) SetMutableFields(value []MutableFieldable)() {
    m.mutable_fields = value
}
type NonFungibleResourceManagerFieldMutableFieldsValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMutableFields()([]MutableFieldable)
    SetMutableFields(value []MutableFieldable)()
}
