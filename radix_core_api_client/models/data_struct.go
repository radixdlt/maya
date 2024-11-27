package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DataStruct struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The owned_entities property
    owned_entities []EntityReferenceable
    // The referenced_entities property
    referenced_entities []EntityReferenceable
    // Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
    struct_data SborDataable
}
// NewDataStruct instantiates a new DataStruct and sets the default values.
func NewDataStruct()(*DataStruct) {
    m := &DataStruct{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDataStructFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataStructFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataStruct(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DataStruct) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataStruct) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["owned_entities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOwnedEntities(res)
        }
        return nil
    }
    res["referenced_entities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetReferencedEntities(res)
        }
        return nil
    }
    res["struct_data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStructData(val.(SborDataable))
        }
        return nil
    }
    return res
}
// GetOwnedEntities gets the owned_entities property value. The owned_entities property
// returns a []EntityReferenceable when successful
func (m *DataStruct) GetOwnedEntities()([]EntityReferenceable) {
    return m.owned_entities
}
// GetReferencedEntities gets the referenced_entities property value. The referenced_entities property
// returns a []EntityReferenceable when successful
func (m *DataStruct) GetReferencedEntities()([]EntityReferenceable) {
    return m.referenced_entities
}
// GetStructData gets the struct_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *DataStruct) GetStructData()(SborDataable) {
    return m.struct_data
}
// Serialize serializes information the current object
func (m *DataStruct) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOwnedEntities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwnedEntities()))
        for i, v := range m.GetOwnedEntities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("owned_entities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReferencedEntities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferencedEntities()))
        for i, v := range m.GetReferencedEntities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("referenced_entities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("struct_data", m.GetStructData())
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
func (m *DataStruct) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOwnedEntities sets the owned_entities property value. The owned_entities property
func (m *DataStruct) SetOwnedEntities(value []EntityReferenceable)() {
    m.owned_entities = value
}
// SetReferencedEntities sets the referenced_entities property value. The referenced_entities property
func (m *DataStruct) SetReferencedEntities(value []EntityReferenceable)() {
    m.referenced_entities = value
}
// SetStructData sets the struct_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *DataStruct) SetStructData(value SborDataable)() {
    m.struct_data = value
}
type DataStructable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOwnedEntities()([]EntityReferenceable)
    GetReferencedEntities()([]EntityReferenceable)
    GetStructData()(SborDataable)
    SetOwnedEntities(value []EntityReferenceable)()
    SetReferencedEntities(value []EntityReferenceable)()
    SetStructData(value SborDataable)()
}
