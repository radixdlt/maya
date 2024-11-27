package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResourceChange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The string-encoded decimal representing the XRD amount put or taken from the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    amount *string
    // The component_entity property
    component_entity EntityReferenceable
    // The Bech32m-encoded human readable version of the resource address
    resource_address *string
    // The vault_entity property
    vault_entity EntityReferenceable
}
// NewResourceChange instantiates a new ResourceChange and sets the default values.
func NewResourceChange()(*ResourceChange) {
    m := &ResourceChange{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResourceChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceChangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceChange(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ResourceChange) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAmount gets the amount property value. The string-encoded decimal representing the XRD amount put or taken from the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ResourceChange) GetAmount()(*string) {
    return m.amount
}
// GetComponentEntity gets the component_entity property value. The component_entity property
// returns a EntityReferenceable when successful
func (m *ResourceChange) GetComponentEntity()(EntityReferenceable) {
    return m.component_entity
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResourceChange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmount(val)
        }
        return nil
    }
    res["component_entity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentEntity(val.(EntityReferenceable))
        }
        return nil
    }
    res["resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAddress(val)
        }
        return nil
    }
    res["vault_entity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultEntity(val.(EntityReferenceable))
        }
        return nil
    }
    return res
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ResourceChange) GetResourceAddress()(*string) {
    return m.resource_address
}
// GetVaultEntity gets the vault_entity property value. The vault_entity property
// returns a EntityReferenceable when successful
func (m *ResourceChange) GetVaultEntity()(EntityReferenceable) {
    return m.vault_entity
}
// Serialize serializes information the current object
func (m *ResourceChange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("component_entity", m.GetComponentEntity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vault_entity", m.GetVaultEntity())
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
func (m *ResourceChange) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAmount sets the amount property value. The string-encoded decimal representing the XRD amount put or taken from the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ResourceChange) SetAmount(value *string)() {
    m.amount = value
}
// SetComponentEntity sets the component_entity property value. The component_entity property
func (m *ResourceChange) SetComponentEntity(value EntityReferenceable)() {
    m.component_entity = value
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *ResourceChange) SetResourceAddress(value *string)() {
    m.resource_address = value
}
// SetVaultEntity sets the vault_entity property value. The vault_entity property
func (m *ResourceChange) SetVaultEntity(value EntityReferenceable)() {
    m.vault_entity = value
}
type ResourceChangeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*string)
    GetComponentEntity()(EntityReferenceable)
    GetResourceAddress()(*string)
    GetVaultEntity()(EntityReferenceable)
    SetAmount(value *string)()
    SetComponentEntity(value EntityReferenceable)()
    SetResourceAddress(value *string)()
    SetVaultEntity(value EntityReferenceable)()
}
