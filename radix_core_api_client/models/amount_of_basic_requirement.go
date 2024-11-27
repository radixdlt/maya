package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AmountOfBasicRequirement struct {
    BasicRequirement
    // The amount property
    amount *string
    // The Bech32m-encoded human readable version of the resource address
    resource *string
}
// NewAmountOfBasicRequirement instantiates a new AmountOfBasicRequirement and sets the default values.
func NewAmountOfBasicRequirement()(*AmountOfBasicRequirement) {
    m := &AmountOfBasicRequirement{
        BasicRequirement: *NewBasicRequirement(),
    }
    return m
}
// CreateAmountOfBasicRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAmountOfBasicRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAmountOfBasicRequirement(), nil
}
// GetAmount gets the amount property value. The amount property
// returns a *string when successful
func (m *AmountOfBasicRequirement) GetAmount()(*string) {
    return m.amount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AmountOfBasicRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BasicRequirement.GetFieldDeserializers()
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *AmountOfBasicRequirement) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *AmountOfBasicRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BasicRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("amount", m.GetAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmount sets the amount property value. The amount property
func (m *AmountOfBasicRequirement) SetAmount(value *string)() {
    m.amount = value
}
// SetResource sets the resource property value. The Bech32m-encoded human readable version of the resource address
func (m *AmountOfBasicRequirement) SetResource(value *string)() {
    m.resource = value
}
type AmountOfBasicRequirementable interface {
    BasicRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*string)
    GetResource()(*string)
    SetAmount(value *string)()
    SetResource(value *string)()
}
