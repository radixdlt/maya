package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsResultantAccountFungibleBalances struct {
    // The Bech32m-encoded human readable version of the account's address
    account_address *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The resultant_balances property
    resultant_balances []LtsResultantFungibleBalanceable
}
// NewLtsResultantAccountFungibleBalances instantiates a new LtsResultantAccountFungibleBalances and sets the default values.
func NewLtsResultantAccountFungibleBalances()(*LtsResultantAccountFungibleBalances) {
    m := &LtsResultantAccountFungibleBalances{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsResultantAccountFungibleBalancesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsResultantAccountFungibleBalancesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsResultantAccountFungibleBalances(), nil
}
// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address
// returns a *string when successful
func (m *LtsResultantAccountFungibleBalances) GetAccountAddress()(*string) {
    return m.account_address
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsResultantAccountFungibleBalances) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsResultantAccountFungibleBalances) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountAddress(val)
        }
        return nil
    }
    res["resultant_balances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLtsResultantFungibleBalanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LtsResultantFungibleBalanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LtsResultantFungibleBalanceable)
                }
            }
            m.SetResultantBalances(res)
        }
        return nil
    }
    return res
}
// GetResultantBalances gets the resultant_balances property value. The resultant_balances property
// returns a []LtsResultantFungibleBalanceable when successful
func (m *LtsResultantAccountFungibleBalances) GetResultantBalances()([]LtsResultantFungibleBalanceable) {
    return m.resultant_balances
}
// Serialize serializes information the current object
func (m *LtsResultantAccountFungibleBalances) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_address", m.GetAccountAddress())
        if err != nil {
            return err
        }
    }
    if m.GetResultantBalances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResultantBalances()))
        for i, v := range m.GetResultantBalances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resultant_balances", cast)
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
// SetAccountAddress sets the account_address property value. The Bech32m-encoded human readable version of the account's address
func (m *LtsResultantAccountFungibleBalances) SetAccountAddress(value *string)() {
    m.account_address = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsResultantAccountFungibleBalances) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResultantBalances sets the resultant_balances property value. The resultant_balances property
func (m *LtsResultantAccountFungibleBalances) SetResultantBalances(value []LtsResultantFungibleBalanceable)() {
    m.resultant_balances = value
}
type LtsResultantAccountFungibleBalancesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountAddress()(*string)
    GetResultantBalances()([]LtsResultantFungibleBalanceable)
    SetAccountAddress(value *string)()
    SetResultantBalances(value []LtsResultantFungibleBalanceable)()
}
