package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PoolVault struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the resource address
    resource_address *string
    // The vault property
    vault EntityReferenceable
}
// NewPoolVault instantiates a new PoolVault and sets the default values.
func NewPoolVault()(*PoolVault) {
    m := &PoolVault{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePoolVaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePoolVaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPoolVault(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PoolVault) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PoolVault) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVault(val.(EntityReferenceable))
        }
        return nil
    }
    return res
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *PoolVault) GetResourceAddress()(*string) {
    return m.resource_address
}
// GetVault gets the vault property value. The vault property
// returns a EntityReferenceable when successful
func (m *PoolVault) GetVault()(EntityReferenceable) {
    return m.vault
}
// Serialize serializes information the current object
func (m *PoolVault) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vault", m.GetVault())
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
func (m *PoolVault) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *PoolVault) SetResourceAddress(value *string)() {
    m.resource_address = value
}
// SetVault sets the vault property value. The vault property
func (m *PoolVault) SetVault(value EntityReferenceable)() {
    m.vault = value
}
type PoolVaultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceAddress()(*string)
    GetVault()(EntityReferenceable)
    SetResourceAddress(value *string)()
    SetVault(value EntityReferenceable)()
}
