package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NonFungibleVaultContentsIndexEntryValue this object is empty, and always present on this substate.Note that when a non-fungible is withdrawn from the vault, the full substate is deleted, which is markedby a DeletedSubstate action (rather than deletion of the value property in an UpdateSubstate action).This is because this is an Index entry, not a KeyValueStore entry.
type NonFungibleVaultContentsIndexEntryValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This is a dummy property which is always set to true and means nothing.It exists just to ensure this object has a well-defined type in OpenAPI schemas.
    is_present *bool
}
// NewNonFungibleVaultContentsIndexEntryValue instantiates a new NonFungibleVaultContentsIndexEntryValue and sets the default values.
func NewNonFungibleVaultContentsIndexEntryValue()(*NonFungibleVaultContentsIndexEntryValue) {
    m := &NonFungibleVaultContentsIndexEntryValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNonFungibleVaultContentsIndexEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleVaultContentsIndexEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleVaultContentsIndexEntryValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleVaultContentsIndexEntryValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleVaultContentsIndexEntryValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_present"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPresent(val)
        }
        return nil
    }
    return res
}
// GetIsPresent gets the is_present property value. This is a dummy property which is always set to true and means nothing.It exists just to ensure this object has a well-defined type in OpenAPI schemas.
// returns a *bool when successful
func (m *NonFungibleVaultContentsIndexEntryValue) GetIsPresent()(*bool) {
    return m.is_present
}
// Serialize serializes information the current object
func (m *NonFungibleVaultContentsIndexEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_present", m.GetIsPresent())
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
func (m *NonFungibleVaultContentsIndexEntryValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsPresent sets the is_present property value. This is a dummy property which is always set to true and means nothing.It exists just to ensure this object has a well-defined type in OpenAPI schemas.
func (m *NonFungibleVaultContentsIndexEntryValue) SetIsPresent(value *bool)() {
    m.is_present = value
}
type NonFungibleVaultContentsIndexEntryValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsPresent()(*bool)
    SetIsPresent(value *bool)()
}
