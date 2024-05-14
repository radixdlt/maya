package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldSystemBootSubstate struct {
	Substate
	// The value property
	value BootLoaderModuleFieldSystemBootValueable
}

// NewBootLoaderModuleFieldSystemBootSubstate instantiates a new BootLoaderModuleFieldSystemBootSubstate and sets the default values.
func NewBootLoaderModuleFieldSystemBootSubstate() *BootLoaderModuleFieldSystemBootSubstate {
	m := &BootLoaderModuleFieldSystemBootSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateBootLoaderModuleFieldSystemBootSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldSystemBootSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBootLoaderModuleFieldSystemBootSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldSystemBootSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBootLoaderModuleFieldSystemBootValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(BootLoaderModuleFieldSystemBootValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a BootLoaderModuleFieldSystemBootValueable when successful
func (m *BootLoaderModuleFieldSystemBootSubstate) GetValue() BootLoaderModuleFieldSystemBootValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *BootLoaderModuleFieldSystemBootSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BootLoaderModuleFieldSystemBootSubstate) SetValue(value BootLoaderModuleFieldSystemBootValueable) {
	m.value = value
}

type BootLoaderModuleFieldSystemBootSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() BootLoaderModuleFieldSystemBootValueable
	SetValue(value BootLoaderModuleFieldSystemBootValueable)
}
