package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintDependenciesEntrySubstate struct {
	Substate
	// The key property
	key BlueprintVersionKeyable
	// The value property
	value PackageBlueprintDependenciesEntryValueable
}

// NewPackageBlueprintDependenciesEntrySubstate instantiates a new PackageBlueprintDependenciesEntrySubstate and sets the default values.
func NewPackageBlueprintDependenciesEntrySubstate() *PackageBlueprintDependenciesEntrySubstate {
	m := &PackageBlueprintDependenciesEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreatePackageBlueprintDependenciesEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintDependenciesEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageBlueprintDependenciesEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintDependenciesEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintVersionKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(BlueprintVersionKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePackageBlueprintDependenciesEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(PackageBlueprintDependenciesEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a BlueprintVersionKeyable when successful
func (m *PackageBlueprintDependenciesEntrySubstate) GetKey() BlueprintVersionKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a PackageBlueprintDependenciesEntryValueable when successful
func (m *PackageBlueprintDependenciesEntrySubstate) GetValue() PackageBlueprintDependenciesEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *PackageBlueprintDependenciesEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *PackageBlueprintDependenciesEntrySubstate) SetKey(value BlueprintVersionKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *PackageBlueprintDependenciesEntrySubstate) SetValue(value PackageBlueprintDependenciesEntryValueable) {
	m.value = value
}

type PackageBlueprintDependenciesEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() BlueprintVersionKeyable
	GetValue() PackageBlueprintDependenciesEntryValueable
	SetKey(value BlueprintVersionKeyable)
	SetValue(value PackageBlueprintDependenciesEntryValueable)
}
