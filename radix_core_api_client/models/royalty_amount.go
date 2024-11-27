package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoyaltyAmount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    amount *string
    // The unit property
    unit *RoyaltyAmount_unit
}
// NewRoyaltyAmount instantiates a new RoyaltyAmount and sets the default values.
func NewRoyaltyAmount()(*RoyaltyAmount) {
    m := &RoyaltyAmount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRoyaltyAmountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoyaltyAmountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoyaltyAmount(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoyaltyAmount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAmount gets the amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *RoyaltyAmount) GetAmount()(*string) {
    return m.amount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoyaltyAmount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["unit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoyaltyAmount_unit)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnit(val.(*RoyaltyAmount_unit))
        }
        return nil
    }
    return res
}
// GetUnit gets the unit property value. The unit property
// returns a *RoyaltyAmount_unit when successful
func (m *RoyaltyAmount) GetUnit()(*RoyaltyAmount_unit) {
    return m.unit
}
// Serialize serializes information the current object
func (m *RoyaltyAmount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    if m.GetUnit() != nil {
        cast := (*m.GetUnit()).String()
        err := writer.WriteStringValue("unit", &cast)
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
func (m *RoyaltyAmount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAmount sets the amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *RoyaltyAmount) SetAmount(value *string)() {
    m.amount = value
}
// SetUnit sets the unit property value. The unit property
func (m *RoyaltyAmount) SetUnit(value *RoyaltyAmount_unit)() {
    m.unit = value
}
type RoyaltyAmountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*string)
    GetUnit()(*RoyaltyAmount_unit)
    SetAmount(value *string)()
    SetUnit(value *RoyaltyAmount_unit)()
}
