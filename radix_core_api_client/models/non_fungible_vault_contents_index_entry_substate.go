package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleVaultContentsIndexEntrySubstate struct {
    Substate
    // The key property
    key LocalNonFungibleKeyable
    // This object is empty, and always present on this substate.Note that when a non-fungible is withdrawn from the vault, the full substate is deleted, which is markedby a DeletedSubstate action (rather than deletion of the value property in an UpdateSubstate action).This is because this is an Index entry, not a KeyValueStore entry.
    value NonFungibleVaultContentsIndexEntryValueable
}
// NewNonFungibleVaultContentsIndexEntrySubstate instantiates a new NonFungibleVaultContentsIndexEntrySubstate and sets the default values.
func NewNonFungibleVaultContentsIndexEntrySubstate()(*NonFungibleVaultContentsIndexEntrySubstate) {
    m := &NonFungibleVaultContentsIndexEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateNonFungibleVaultContentsIndexEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleVaultContentsIndexEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleVaultContentsIndexEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleVaultContentsIndexEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocalNonFungibleKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(LocalNonFungibleKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNonFungibleVaultContentsIndexEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(NonFungibleVaultContentsIndexEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a LocalNonFungibleKeyable when successful
func (m *NonFungibleVaultContentsIndexEntrySubstate) GetKey()(LocalNonFungibleKeyable) {
    return m.key
}
// GetValue gets the value property value. This object is empty, and always present on this substate.Note that when a non-fungible is withdrawn from the vault, the full substate is deleted, which is markedby a DeletedSubstate action (rather than deletion of the value property in an UpdateSubstate action).This is because this is an Index entry, not a KeyValueStore entry.
// returns a NonFungibleVaultContentsIndexEntryValueable when successful
func (m *NonFungibleVaultContentsIndexEntrySubstate) GetValue()(NonFungibleVaultContentsIndexEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *NonFungibleVaultContentsIndexEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. The key property
func (m *NonFungibleVaultContentsIndexEntrySubstate) SetKey(value LocalNonFungibleKeyable)() {
    m.key = value
}
// SetValue sets the value property value. This object is empty, and always present on this substate.Note that when a non-fungible is withdrawn from the vault, the full substate is deleted, which is markedby a DeletedSubstate action (rather than deletion of the value property in an UpdateSubstate action).This is because this is an Index entry, not a KeyValueStore entry.
func (m *NonFungibleVaultContentsIndexEntrySubstate) SetValue(value NonFungibleVaultContentsIndexEntryValueable)() {
    m.value = value
}
type NonFungibleVaultContentsIndexEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(LocalNonFungibleKeyable)
    GetValue()(NonFungibleVaultContentsIndexEntryValueable)
    SetKey(value LocalNonFungibleKeyable)()
    SetValue(value NonFungibleVaultContentsIndexEntryValueable)()
}
