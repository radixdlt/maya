package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultBalance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The resource_amount property
    resource_amount ResourceAmountable
    // The vault_entity property
    vault_entity EntityReferenceable
}
// NewVaultBalance instantiates a new VaultBalance and sets the default values.
func NewVaultBalance()(*VaultBalance) {
    m := &VaultBalance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVaultBalanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultBalanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultBalance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VaultBalance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultBalance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resource_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceAmountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAmount(val.(ResourceAmountable))
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
// GetResourceAmount gets the resource_amount property value. The resource_amount property
// returns a ResourceAmountable when successful
func (m *VaultBalance) GetResourceAmount()(ResourceAmountable) {
    return m.resource_amount
}
// GetVaultEntity gets the vault_entity property value. The vault_entity property
// returns a EntityReferenceable when successful
func (m *VaultBalance) GetVaultEntity()(EntityReferenceable) {
    return m.vault_entity
}
// Serialize serializes information the current object
func (m *VaultBalance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("resource_amount", m.GetResourceAmount())
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
func (m *VaultBalance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResourceAmount sets the resource_amount property value. The resource_amount property
func (m *VaultBalance) SetResourceAmount(value ResourceAmountable)() {
    m.resource_amount = value
}
// SetVaultEntity sets the vault_entity property value. The vault_entity property
func (m *VaultBalance) SetVaultEntity(value EntityReferenceable)() {
    m.vault_entity = value
}
type VaultBalanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceAmount()(ResourceAmountable)
    GetVaultEntity()(EntityReferenceable)
    SetResourceAmount(value ResourceAmountable)()
    SetVaultEntity(value EntityReferenceable)()
}
