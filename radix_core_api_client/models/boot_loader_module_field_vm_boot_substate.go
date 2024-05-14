package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldVmBootSubstate struct {
	Substate
	// The value property
	value BootLoaderModuleFieldVmBootValueable
}

// NewBootLoaderModuleFieldVmBootSubstate instantiates a new BootLoaderModuleFieldVmBootSubstate and sets the default values.
func NewBootLoaderModuleFieldVmBootSubstate() *BootLoaderModuleFieldVmBootSubstate {
	m := &BootLoaderModuleFieldVmBootSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateBootLoaderModuleFieldVmBootSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldVmBootSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBootLoaderModuleFieldVmBootSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldVmBootSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBootLoaderModuleFieldVmBootValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(BootLoaderModuleFieldVmBootValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a BootLoaderModuleFieldVmBootValueable when successful
func (m *BootLoaderModuleFieldVmBootSubstate) GetValue() BootLoaderModuleFieldVmBootValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *BootLoaderModuleFieldVmBootSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BootLoaderModuleFieldVmBootSubstate) SetValue(value BootLoaderModuleFieldVmBootValueable) {
	m.value = value
}

type BootLoaderModuleFieldVmBootSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() BootLoaderModuleFieldVmBootValueable
	SetValue(value BootLoaderModuleFieldVmBootValueable)
}
