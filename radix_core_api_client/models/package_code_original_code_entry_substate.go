package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageCodeOriginalCodeEntrySubstate struct {
	Substate
	// The key property
	key PackageCodeKeyable
	// The value property
	value PackageCodeOriginalCodeEntryValueable
}

// NewPackageCodeOriginalCodeEntrySubstate instantiates a new PackageCodeOriginalCodeEntrySubstate and sets the default values.
func NewPackageCodeOriginalCodeEntrySubstate() *PackageCodeOriginalCodeEntrySubstate {
	m := &PackageCodeOriginalCodeEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreatePackageCodeOriginalCodeEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageCodeOriginalCodeEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageCodeOriginalCodeEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageCodeOriginalCodeEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePackageCodeKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(PackageCodeKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePackageCodeOriginalCodeEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(PackageCodeOriginalCodeEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a PackageCodeKeyable when successful
func (m *PackageCodeOriginalCodeEntrySubstate) GetKey() PackageCodeKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a PackageCodeOriginalCodeEntryValueable when successful
func (m *PackageCodeOriginalCodeEntrySubstate) GetValue() PackageCodeOriginalCodeEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *PackageCodeOriginalCodeEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *PackageCodeOriginalCodeEntrySubstate) SetKey(value PackageCodeKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *PackageCodeOriginalCodeEntrySubstate) SetValue(value PackageCodeOriginalCodeEntryValueable) {
	m.value = value
}

type PackageCodeOriginalCodeEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() PackageCodeKeyable
	GetValue() PackageCodeOriginalCodeEntryValueable
	SetKey(value PackageCodeKeyable)
	SetValue(value PackageCodeOriginalCodeEntryValueable)
}
