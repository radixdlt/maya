package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleResourceManagerFieldDivisibilitySubstate struct {
    Substate
    // The value property
    value FungibleResourceManagerFieldDivisibilityValueable
}
// NewFungibleResourceManagerFieldDivisibilitySubstate instantiates a new FungibleResourceManagerFieldDivisibilitySubstate and sets the default values.
func NewFungibleResourceManagerFieldDivisibilitySubstate()(*FungibleResourceManagerFieldDivisibilitySubstate) {
    m := &FungibleResourceManagerFieldDivisibilitySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateFungibleResourceManagerFieldDivisibilitySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleResourceManagerFieldDivisibilitySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFungibleResourceManagerFieldDivisibilitySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleResourceManagerFieldDivisibilitySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFungibleResourceManagerFieldDivisibilityValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(FungibleResourceManagerFieldDivisibilityValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a FungibleResourceManagerFieldDivisibilityValueable when successful
func (m *FungibleResourceManagerFieldDivisibilitySubstate) GetValue()(FungibleResourceManagerFieldDivisibilityValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *FungibleResourceManagerFieldDivisibilitySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *FungibleResourceManagerFieldDivisibilitySubstate) SetValue(value FungibleResourceManagerFieldDivisibilityValueable)() {
    m.value = value
}
type FungibleResourceManagerFieldDivisibilitySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(FungibleResourceManagerFieldDivisibilityValueable)
    SetValue(value FungibleResourceManagerFieldDivisibilityValueable)()
}
