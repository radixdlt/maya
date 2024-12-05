package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleVaultFieldBalanceValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The string-encoded decimal representing the token amount in the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    amount *string
}
// NewFungibleVaultFieldBalanceValue instantiates a new FungibleVaultFieldBalanceValue and sets the default values.
func NewFungibleVaultFieldBalanceValue()(*FungibleVaultFieldBalanceValue) {
    m := &FungibleVaultFieldBalanceValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFungibleVaultFieldBalanceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleVaultFieldBalanceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFungibleVaultFieldBalanceValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FungibleVaultFieldBalanceValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAmount gets the amount property value. The string-encoded decimal representing the token amount in the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FungibleVaultFieldBalanceValue) GetAmount()(*string) {
    return m.amount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleVaultFieldBalanceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *FungibleVaultFieldBalanceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("amount", m.GetAmount())
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
func (m *FungibleVaultFieldBalanceValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAmount sets the amount property value. The string-encoded decimal representing the token amount in the vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FungibleVaultFieldBalanceValue) SetAmount(value *string)() {
    m.amount = value
}
type FungibleVaultFieldBalanceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*string)
    SetAmount(value *string)()
}
