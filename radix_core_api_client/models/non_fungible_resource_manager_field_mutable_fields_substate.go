package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleResourceManagerFieldMutableFieldsSubstate struct {
    Substate
    // The value property
    value NonFungibleResourceManagerFieldMutableFieldsValueable
}
// NewNonFungibleResourceManagerFieldMutableFieldsSubstate instantiates a new NonFungibleResourceManagerFieldMutableFieldsSubstate and sets the default values.
func NewNonFungibleResourceManagerFieldMutableFieldsSubstate()(*NonFungibleResourceManagerFieldMutableFieldsSubstate) {
    m := &NonFungibleResourceManagerFieldMutableFieldsSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateNonFungibleResourceManagerFieldMutableFieldsSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleResourceManagerFieldMutableFieldsSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleResourceManagerFieldMutableFieldsSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleResourceManagerFieldMutableFieldsSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNonFungibleResourceManagerFieldMutableFieldsValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(NonFungibleResourceManagerFieldMutableFieldsValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a NonFungibleResourceManagerFieldMutableFieldsValueable when successful
func (m *NonFungibleResourceManagerFieldMutableFieldsSubstate) GetValue()(NonFungibleResourceManagerFieldMutableFieldsValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *NonFungibleResourceManagerFieldMutableFieldsSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *NonFungibleResourceManagerFieldMutableFieldsSubstate) SetValue(value NonFungibleResourceManagerFieldMutableFieldsValueable)() {
    m.value = value
}
type NonFungibleResourceManagerFieldMutableFieldsSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(NonFungibleResourceManagerFieldMutableFieldsValueable)
    SetValue(value NonFungibleResourceManagerFieldMutableFieldsValueable)()
}
