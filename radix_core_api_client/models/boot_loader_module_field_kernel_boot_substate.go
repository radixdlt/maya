package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldKernelBootSubstate struct {
	Substate
	// The value property
	value BootLoaderModuleFieldKernelBootValueable
}

// NewBootLoaderModuleFieldKernelBootSubstate instantiates a new BootLoaderModuleFieldKernelBootSubstate and sets the default values.
func NewBootLoaderModuleFieldKernelBootSubstate() *BootLoaderModuleFieldKernelBootSubstate {
	m := &BootLoaderModuleFieldKernelBootSubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateBootLoaderModuleFieldKernelBootSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldKernelBootSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBootLoaderModuleFieldKernelBootSubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldKernelBootSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBootLoaderModuleFieldKernelBootValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(BootLoaderModuleFieldKernelBootValueable))
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a BootLoaderModuleFieldKernelBootValueable when successful
func (m *BootLoaderModuleFieldKernelBootSubstate) GetValue() BootLoaderModuleFieldKernelBootValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *BootLoaderModuleFieldKernelBootSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BootLoaderModuleFieldKernelBootSubstate) SetValue(value BootLoaderModuleFieldKernelBootValueable) {
	m.value = value
}

type BootLoaderModuleFieldKernelBootSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetValue() BootLoaderModuleFieldKernelBootValueable
	SetValue(value BootLoaderModuleFieldKernelBootValueable)
}
