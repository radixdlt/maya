package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiResourcePoolFieldStateValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the resource address
    pool_unit_resource_address *string
    // The vaults property
    vaults []PoolVaultable
}
// NewMultiResourcePoolFieldStateValue instantiates a new MultiResourcePoolFieldStateValue and sets the default values.
func NewMultiResourcePoolFieldStateValue()(*MultiResourcePoolFieldStateValue) {
    m := &MultiResourcePoolFieldStateValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMultiResourcePoolFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiResourcePoolFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiResourcePoolFieldStateValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MultiResourcePoolFieldStateValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiResourcePoolFieldStateValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pool_unit_resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoolUnitResourceAddress(val)
        }
        return nil
    }
    res["vaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePoolVaultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PoolVaultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PoolVaultable)
                }
            }
            m.SetVaults(res)
        }
        return nil
    }
    return res
}
// GetPoolUnitResourceAddress gets the pool_unit_resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *MultiResourcePoolFieldStateValue) GetPoolUnitResourceAddress()(*string) {
    return m.pool_unit_resource_address
}
// GetVaults gets the vaults property value. The vaults property
// returns a []PoolVaultable when successful
func (m *MultiResourcePoolFieldStateValue) GetVaults()([]PoolVaultable) {
    return m.vaults
}
// Serialize serializes information the current object
func (m *MultiResourcePoolFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pool_unit_resource_address", m.GetPoolUnitResourceAddress())
        if err != nil {
            return err
        }
    }
    if m.GetVaults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVaults()))
        for i, v := range m.GetVaults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("vaults", cast)
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
func (m *MultiResourcePoolFieldStateValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPoolUnitResourceAddress sets the pool_unit_resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *MultiResourcePoolFieldStateValue) SetPoolUnitResourceAddress(value *string)() {
    m.pool_unit_resource_address = value
}
// SetVaults sets the vaults property value. The vaults property
func (m *MultiResourcePoolFieldStateValue) SetVaults(value []PoolVaultable)() {
    m.vaults = value
}
type MultiResourcePoolFieldStateValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPoolUnitResourceAddress()(*string)
    GetVaults()([]PoolVaultable)
    SetPoolUnitResourceAddress(value *string)()
    SetVaults(value []PoolVaultable)()
}
