package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FunctionEventEmitterIdentifier struct {
	EventEmitterIdentifier
	// The blueprint under the package which emitted the event.
	blueprint_name *string
	// The Bech32m-encoded human readable version of the package address
	package_address *string
}

// NewFunctionEventEmitterIdentifier instantiates a new FunctionEventEmitterIdentifier and sets the default values.
func NewFunctionEventEmitterIdentifier() *FunctionEventEmitterIdentifier {
	m := &FunctionEventEmitterIdentifier{
		EventEmitterIdentifier: *NewEventEmitterIdentifier(),
	}
	return m
}

// CreateFunctionEventEmitterIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFunctionEventEmitterIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFunctionEventEmitterIdentifier(), nil
}

// GetBlueprintName gets the blueprint_name property value. The blueprint under the package which emitted the event.
// returns a *string when successful
func (m *FunctionEventEmitterIdentifier) GetBlueprintName() *string {
	return m.blueprint_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FunctionEventEmitterIdentifier) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.EventEmitterIdentifier.GetFieldDeserializers()
	res["blueprint_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlueprintName(val)
		}
		return nil
	}
	res["package_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPackageAddress(val)
		}
		return nil
	}
	return res
}

// GetPackageAddress gets the package_address property value. The Bech32m-encoded human readable version of the package address
// returns a *string when successful
func (m *FunctionEventEmitterIdentifier) GetPackageAddress() *string {
	return m.package_address
}

// Serialize serializes information the current object
func (m *FunctionEventEmitterIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.EventEmitterIdentifier.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("blueprint_name", m.GetBlueprintName())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("package_address", m.GetPackageAddress())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBlueprintName sets the blueprint_name property value. The blueprint under the package which emitted the event.
func (m *FunctionEventEmitterIdentifier) SetBlueprintName(value *string) {
	m.blueprint_name = value
}

// SetPackageAddress sets the package_address property value. The Bech32m-encoded human readable version of the package address
func (m *FunctionEventEmitterIdentifier) SetPackageAddress(value *string) {
	m.package_address = value
}

type FunctionEventEmitterIdentifierable interface {
	EventEmitterIdentifierable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBlueprintName() *string
	GetPackageAddress() *string
	SetBlueprintName(value *string)
	SetPackageAddress(value *string)
}
