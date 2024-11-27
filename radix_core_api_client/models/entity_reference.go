package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EntityReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
    entity_address *string
    // To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
    entity_type *EntityType
    // The is_global property
    is_global *bool
}
// NewEntityReference instantiates a new EntityReference and sets the default values.
func NewEntityReference()(*EntityReference) {
    m := &EntityReference{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEntityReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntityReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntityReference(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EntityReference) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEntityAddress gets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
// returns a *string when successful
func (m *EntityReference) GetEntityAddress()(*string) {
    return m.entity_address
}
// GetEntityType gets the entity_type property value. To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
// returns a *EntityType when successful
func (m *EntityReference) GetEntityType()(*EntityType) {
    return m.entity_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EntityReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entity_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityAddress(val)
        }
        return nil
    }
    res["entity_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val.(*EntityType))
        }
        return nil
    }
    res["is_global"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobal(val)
        }
        return nil
    }
    return res
}
// GetIsGlobal gets the is_global property value. The is_global property
// returns a *bool when successful
func (m *EntityReference) GetIsGlobal()(*bool) {
    return m.is_global
}
// Serialize serializes information the current object
func (m *EntityReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("entity_address", m.GetEntityAddress())
        if err != nil {
            return err
        }
    }
    if m.GetEntityType() != nil {
        cast := (*m.GetEntityType()).String()
        err := writer.WriteStringValue("entity_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_global", m.GetIsGlobal())
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
func (m *EntityReference) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEntityAddress sets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
func (m *EntityReference) SetEntityAddress(value *string)() {
    m.entity_address = value
}
// SetEntityType sets the entity_type property value. To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
func (m *EntityReference) SetEntityType(value *EntityType)() {
    m.entity_type = value
}
// SetIsGlobal sets the is_global property value. The is_global property
func (m *EntityReference) SetIsGlobal(value *bool)() {
    m.is_global = value
}
type EntityReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntityAddress()(*string)
    GetEntityType()(*EntityType)
    GetIsGlobal()(*bool)
    SetEntityAddress(value *string)()
    SetEntityType(value *EntityType)()
    SetIsGlobal(value *bool)()
}
