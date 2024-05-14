package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintDependenciesEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The dependencies property
	dependencies BlueprintDependenciesable
}

// NewPackageBlueprintDependenciesEntryValue instantiates a new PackageBlueprintDependenciesEntryValue and sets the default values.
func NewPackageBlueprintDependenciesEntryValue() *PackageBlueprintDependenciesEntryValue {
	m := &PackageBlueprintDependenciesEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageBlueprintDependenciesEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintDependenciesEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageBlueprintDependenciesEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageBlueprintDependenciesEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDependencies gets the dependencies property value. The dependencies property
// returns a BlueprintDependenciesable when successful
func (m *PackageBlueprintDependenciesEntryValue) GetDependencies() BlueprintDependenciesable {
	return m.dependencies
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintDependenciesEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["dependencies"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintDependenciesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDependencies(val.(BlueprintDependenciesable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *PackageBlueprintDependenciesEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("dependencies", m.GetDependencies())
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
func (m *PackageBlueprintDependenciesEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDependencies sets the dependencies property value. The dependencies property
func (m *PackageBlueprintDependenciesEntryValue) SetDependencies(value BlueprintDependenciesable) {
	m.dependencies = value
}

type PackageBlueprintDependenciesEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDependencies() BlueprintDependenciesable
	SetDependencies(value BlueprintDependenciesable)
}
