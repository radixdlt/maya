package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintAuthTemplateEntrySubstate struct {
	Substate
	// The key property
	key BlueprintVersionKeyable
	// The value property
	value PackageBlueprintAuthTemplateEntryValueable
}

// NewPackageBlueprintAuthTemplateEntrySubstate instantiates a new PackageBlueprintAuthTemplateEntrySubstate and sets the default values.
func NewPackageBlueprintAuthTemplateEntrySubstate() *PackageBlueprintAuthTemplateEntrySubstate {
	m := &PackageBlueprintAuthTemplateEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreatePackageBlueprintAuthTemplateEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintAuthTemplateEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageBlueprintAuthTemplateEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintAuthTemplateEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetObjectValue(CreatePackageBlueprintAuthTemplateEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(PackageBlueprintAuthTemplateEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a BlueprintVersionKeyable when successful
func (m *PackageBlueprintAuthTemplateEntrySubstate) GetKey() BlueprintVersionKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a PackageBlueprintAuthTemplateEntryValueable when successful
func (m *PackageBlueprintAuthTemplateEntrySubstate) GetValue() PackageBlueprintAuthTemplateEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *PackageBlueprintAuthTemplateEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *PackageBlueprintAuthTemplateEntrySubstate) SetKey(value BlueprintVersionKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *PackageBlueprintAuthTemplateEntrySubstate) SetValue(value PackageBlueprintAuthTemplateEntryValueable) {
	m.value = value
}

type PackageBlueprintAuthTemplateEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() BlueprintVersionKeyable
	GetValue() PackageBlueprintAuthTemplateEntryValueable
	SetKey(value BlueprintVersionKeyable)
	SetValue(value PackageBlueprintAuthTemplateEntryValueable)
}
