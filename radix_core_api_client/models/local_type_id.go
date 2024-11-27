package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalTypeId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
    as_sbor SborDataable
    // A reference to a type, interpreted according to `kind`:- If `WellKnown`, then it is a pointer to a well known scrypto type with that ID,- If `SchemaLocal`, then it is an index into the given schema.
    id *int64
    // The location against which to resolve this type reference.
    kind *LocalTypeId_kind
}
// NewLocalTypeId instantiates a new LocalTypeId and sets the default values.
func NewLocalTypeId()(*LocalTypeId) {
    m := &LocalTypeId{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLocalTypeIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalTypeIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalTypeId(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalTypeId) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAsSbor gets the as_sbor property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *LocalTypeId) GetAsSbor()(SborDataable) {
    return m.as_sbor
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalTypeId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["as_sbor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAsSbor(val.(SborDataable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalTypeId_kind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val.(*LocalTypeId_kind))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. A reference to a type, interpreted according to `kind`:- If `WellKnown`, then it is a pointer to a well known scrypto type with that ID,- If `SchemaLocal`, then it is an index into the given schema.
// returns a *int64 when successful
func (m *LocalTypeId) GetId()(*int64) {
    return m.id
}
// GetKind gets the kind property value. The location against which to resolve this type reference.
// returns a *LocalTypeId_kind when successful
func (m *LocalTypeId) GetKind()(*LocalTypeId_kind) {
    return m.kind
}
// Serialize serializes information the current object
func (m *LocalTypeId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("as_sbor", m.GetAsSbor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetKind() != nil {
        cast := (*m.GetKind()).String()
        err := writer.WriteStringValue("kind", &cast)
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
func (m *LocalTypeId) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAsSbor sets the as_sbor property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *LocalTypeId) SetAsSbor(value SborDataable)() {
    m.as_sbor = value
}
// SetId sets the id property value. A reference to a type, interpreted according to `kind`:- If `WellKnown`, then it is a pointer to a well known scrypto type with that ID,- If `SchemaLocal`, then it is an index into the given schema.
func (m *LocalTypeId) SetId(value *int64)() {
    m.id = value
}
// SetKind sets the kind property value. The location against which to resolve this type reference.
func (m *LocalTypeId) SetKind(value *LocalTypeId_kind)() {
    m.kind = value
}
type LocalTypeIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAsSbor()(SborDataable)
    GetId()(*int64)
    GetKind()(*LocalTypeId_kind)
    SetAsSbor(value SborDataable)()
    SetId(value *int64)()
    SetKind(value *LocalTypeId_kind)()
}
