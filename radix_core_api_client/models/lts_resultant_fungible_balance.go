package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsResultantFungibleBalance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the fungible resource's address
    resource_address *string
    // The string-encoded decimal representing the resultant balance of the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    resultant_balance *string
}
// NewLtsResultantFungibleBalance instantiates a new LtsResultantFungibleBalance and sets the default values.
func NewLtsResultantFungibleBalance()(*LtsResultantFungibleBalance) {
    m := &LtsResultantFungibleBalance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsResultantFungibleBalanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsResultantFungibleBalanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsResultantFungibleBalance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsResultantFungibleBalance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsResultantFungibleBalance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resultant_balance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultantBalance(val)
        }
        return nil
    }
    return res
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the fungible resource's address
// returns a *string when successful
func (m *LtsResultantFungibleBalance) GetResourceAddress()(*string) {
    return m.resource_address
}
// GetResultantBalance gets the resultant_balance property value. The string-encoded decimal representing the resultant balance of the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *LtsResultantFungibleBalance) GetResultantBalance()(*string) {
    return m.resultant_balance
}
// Serialize serializes information the current object
func (m *LtsResultantFungibleBalance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resultant_balance", m.GetResultantBalance())
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
func (m *LtsResultantFungibleBalance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the fungible resource's address
func (m *LtsResultantFungibleBalance) SetResourceAddress(value *string)() {
    m.resource_address = value
}
// SetResultantBalance sets the resultant_balance property value. The string-encoded decimal representing the resultant balance of the fungible resource.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *LtsResultantFungibleBalance) SetResultantBalance(value *string)() {
    m.resultant_balance = value
}
type LtsResultantFungibleBalanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceAddress()(*string)
    GetResultantBalance()(*string)
    SetResourceAddress(value *string)()
    SetResultantBalance(value *string)()
}
