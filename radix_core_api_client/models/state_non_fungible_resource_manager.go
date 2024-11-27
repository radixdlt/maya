package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateNonFungibleResourceManager struct {
    StateResourceManager
    // The id_type property
    id_type Substateable
    // The mutable_fields property
    mutable_fields Substateable
    // The total_supply property
    total_supply Substateable
}
// NewStateNonFungibleResourceManager instantiates a new StateNonFungibleResourceManager and sets the default values.
func NewStateNonFungibleResourceManager()(*StateNonFungibleResourceManager) {
    m := &StateNonFungibleResourceManager{
        StateResourceManager: *NewStateResourceManager(),
    }
    return m
}
// CreateStateNonFungibleResourceManagerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateNonFungibleResourceManagerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateNonFungibleResourceManager(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateNonFungibleResourceManager) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StateResourceManager.GetFieldDeserializers()
    res["id_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdType(val.(Substateable))
        }
        return nil
    }
    res["mutable_fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMutableFields(val.(Substateable))
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
// GetIdType gets the id_type property value. The id_type property
// returns a Substateable when successful
func (m *StateNonFungibleResourceManager) GetIdType()(Substateable) {
    return m.id_type
}
// GetMutableFields gets the mutable_fields property value. The mutable_fields property
// returns a Substateable when successful
func (m *StateNonFungibleResourceManager) GetMutableFields()(Substateable) {
    return m.mutable_fields
}
// GetTotalSupply gets the total_supply property value. The total_supply property
// returns a Substateable when successful
func (m *StateNonFungibleResourceManager) GetTotalSupply()(Substateable) {
    return m.total_supply
}
// Serialize serializes information the current object
func (m *StateNonFungibleResourceManager) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StateResourceManager.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("id_type", m.GetIdType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mutable_fields", m.GetMutableFields())
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
// SetIdType sets the id_type property value. The id_type property
func (m *StateNonFungibleResourceManager) SetIdType(value Substateable)() {
    m.id_type = value
}
// SetMutableFields sets the mutable_fields property value. The mutable_fields property
func (m *StateNonFungibleResourceManager) SetMutableFields(value Substateable)() {
    m.mutable_fields = value
}
// SetTotalSupply sets the total_supply property value. The total_supply property
func (m *StateNonFungibleResourceManager) SetTotalSupply(value Substateable)() {
    m.total_supply = value
}
type StateNonFungibleResourceManagerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StateResourceManagerable
    GetIdType()(Substateable)
    GetMutableFields()(Substateable)
    GetTotalSupply()(Substateable)
    SetIdType(value Substateable)()
    SetMutableFields(value Substateable)()
    SetTotalSupply(value Substateable)()
}
