package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScopedTypeId an identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
type ScopedTypeId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The local_type_id property
    local_type_id LocalTypeIdable
    // The hex-encoded schema hash, capturing the identity of an SBOR schema.
    schema_hash *string
}
// NewScopedTypeId instantiates a new ScopedTypeId and sets the default values.
func NewScopedTypeId()(*ScopedTypeId) {
    m := &ScopedTypeId{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScopedTypeIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScopedTypeIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScopedTypeId(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScopedTypeId) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScopedTypeId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["local_type_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocalTypeIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalTypeId(val.(LocalTypeIdable))
        }
        return nil
    }
    res["schema_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaHash(val)
        }
        return nil
    }
    return res
}
// GetLocalTypeId gets the local_type_id property value. The local_type_id property
// returns a LocalTypeIdable when successful
func (m *ScopedTypeId) GetLocalTypeId()(LocalTypeIdable) {
    return m.local_type_id
}
// GetSchemaHash gets the schema_hash property value. The hex-encoded schema hash, capturing the identity of an SBOR schema.
// returns a *string when successful
func (m *ScopedTypeId) GetSchemaHash()(*string) {
    return m.schema_hash
}
// Serialize serializes information the current object
func (m *ScopedTypeId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("local_type_id", m.GetLocalTypeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schema_hash", m.GetSchemaHash())
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
func (m *ScopedTypeId) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocalTypeId sets the local_type_id property value. The local_type_id property
func (m *ScopedTypeId) SetLocalTypeId(value LocalTypeIdable)() {
    m.local_type_id = value
}
// SetSchemaHash sets the schema_hash property value. The hex-encoded schema hash, capturing the identity of an SBOR schema.
func (m *ScopedTypeId) SetSchemaHash(value *string)() {
    m.schema_hash = value
}
type ScopedTypeIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocalTypeId()(LocalTypeIdable)
    GetSchemaHash()(*string)
    SetLocalTypeId(value LocalTypeIdable)()
    SetSchemaHash(value *string)()
}
