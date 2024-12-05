package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateFungibleResourceManager struct {
    StateResourceManager
    // The divisibility property
    divisibility Substateable
    // The total_supply property
    total_supply Substateable
}
// NewStateFungibleResourceManager instantiates a new StateFungibleResourceManager and sets the default values.
func NewStateFungibleResourceManager()(*StateFungibleResourceManager) {
    m := &StateFungibleResourceManager{
        StateResourceManager: *NewStateResourceManager(),
    }
    return m
}
// CreateStateFungibleResourceManagerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateFungibleResourceManagerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateFungibleResourceManager(), nil
}
// GetDivisibility gets the divisibility property value. The divisibility property
// returns a Substateable when successful
func (m *StateFungibleResourceManager) GetDivisibility()(Substateable) {
    return m.divisibility
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateFungibleResourceManager) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StateResourceManager.GetFieldDeserializers()
    res["divisibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDivisibility(val.(Substateable))
        }
        return nil
    }
    res["total_supply"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSupply(val.(Substateable))
        }
        return nil
    }
    return res
}
// GetTotalSupply gets the total_supply property value. The total_supply property
// returns a Substateable when successful
func (m *StateFungibleResourceManager) GetTotalSupply()(Substateable) {
    return m.total_supply
}
// Serialize serializes information the current object
func (m *StateFungibleResourceManager) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StateResourceManager.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("divisibility", m.GetDivisibility())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("total_supply", m.GetTotalSupply())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDivisibility sets the divisibility property value. The divisibility property
func (m *StateFungibleResourceManager) SetDivisibility(value Substateable)() {
    m.divisibility = value
}
// SetTotalSupply sets the total_supply property value. The total_supply property
func (m *StateFungibleResourceManager) SetTotalSupply(value Substateable)() {
    m.total_supply = value
}
type StateFungibleResourceManagerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StateResourceManagerable
    GetDivisibility()(Substateable)
    GetTotalSupply()(Substateable)
    SetDivisibility(value Substateable)()
    SetTotalSupply(value Substateable)()
}
