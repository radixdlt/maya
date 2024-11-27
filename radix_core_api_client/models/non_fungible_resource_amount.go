package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceAmount struct {
    ResourceAmount
    // The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    amount *string
    // The non_fungible_ids property
    non_fungible_ids []NonFungibleLocalIdable
}
// NewNonFungibleResourceAmount instantiates a new NonFungibleResourceAmount and sets the default values.
func NewNonFungibleResourceAmount()(*NonFungibleResourceAmount) {
    m := &NonFungibleResourceAmount{
        ResourceAmount: *NewResourceAmount(),
    }
    return m
}
// CreateNonFungibleResourceAmountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceAmountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleResourceAmount(), nil
}
// GetAmount gets the amount property value. The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *NonFungibleResourceAmount) GetAmount()(*string) {
    return m.amount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceAmount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResourceAmount.GetFieldDeserializers()
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
    res["non_fungible_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNonFungibleLocalIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NonFungibleLocalIdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NonFungibleLocalIdable)
                }
            }
            m.SetNonFungibleIds(res)
        }
        return nil
    }
    return res
}
// GetNonFungibleIds gets the non_fungible_ids property value. The non_fungible_ids property
// returns a []NonFungibleLocalIdable when successful
func (m *NonFungibleResourceAmount) GetNonFungibleIds()([]NonFungibleLocalIdable) {
    return m.non_fungible_ids
}
// Serialize serializes information the current object
func (m *NonFungibleResourceAmount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ResourceAmount.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    if m.GetNonFungibleIds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonFungibleIds()))
        for i, v := range m.GetNonFungibleIds() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("non_fungible_ids", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmount sets the amount property value. The string-encoded decimal representing the amount of this resource (some decimal for fungible resources, a whole integer for non-fungible resources).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *NonFungibleResourceAmount) SetAmount(value *string)() {
    m.amount = value
}
// SetNonFungibleIds sets the non_fungible_ids property value. The non_fungible_ids property
func (m *NonFungibleResourceAmount) SetNonFungibleIds(value []NonFungibleLocalIdable)() {
    m.non_fungible_ids = value
}
type NonFungibleResourceAmountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceAmountable
    GetAmount()(*string)
    GetNonFungibleIds()([]NonFungibleLocalIdable)
    SetAmount(value *string)()
    SetNonFungibleIds(value []NonFungibleLocalIdable)()
}
