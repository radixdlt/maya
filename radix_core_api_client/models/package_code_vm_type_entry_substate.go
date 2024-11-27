package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageCodeVmTypeEntrySubstate struct {
    Substate
    // The key property
    key PackageCodeKeyable
    // The value property
    value PackageCodeVmTypeEntryValueable
}
// NewPackageCodeVmTypeEntrySubstate instantiates a new PackageCodeVmTypeEntrySubstate and sets the default values.
func NewPackageCodeVmTypeEntrySubstate()(*PackageCodeVmTypeEntrySubstate) {
    m := &PackageCodeVmTypeEntrySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreatePackageCodeVmTypeEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageCodeVmTypeEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageCodeVmTypeEntrySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageCodeVmTypeEntrySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageCodeKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(PackageCodeKeyable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageCodeVmTypeEntryValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(PackageCodeVmTypeEntryValueable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a PackageCodeKeyable when successful
func (m *PackageCodeVmTypeEntrySubstate) GetKey()(PackageCodeKeyable) {
    return m.key
}
// GetValue gets the value property value. The value property
// returns a PackageCodeVmTypeEntryValueable when successful
func (m *PackageCodeVmTypeEntrySubstate) GetValue()(PackageCodeVmTypeEntryValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PackageCodeVmTypeEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PackageCodeVmTypeEntrySubstate) SetKey(value PackageCodeKeyable)() {
    m.key = value
}
// SetValue sets the value property value. The value property
func (m *PackageCodeVmTypeEntrySubstate) SetValue(value PackageCodeVmTypeEntryValueable)() {
    m.value = value
}
type PackageCodeVmTypeEntrySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetKey()(PackageCodeKeyable)
    GetValue()(PackageCodeVmTypeEntryValueable)
    SetKey(value PackageCodeKeyable)()
    SetValue(value PackageCodeVmTypeEntryValueable)()
}
