package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectTypeInfoDetails struct {
	TypeInfoDetails
	// The blueprint_info property
	blueprint_info BlueprintInfoable
	// The global property
	global *bool
	// The module_versions property
	module_versions []ModuleVersionable
}

// NewObjectTypeInfoDetails instantiates a new ObjectTypeInfoDetails and sets the default values.
func NewObjectTypeInfoDetails() *ObjectTypeInfoDetails {
	m := &ObjectTypeInfoDetails{
		TypeInfoDetails: *NewTypeInfoDetails(),
	}
	return m
}

// CreateObjectTypeInfoDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectTypeInfoDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewObjectTypeInfoDetails(), nil
}

// GetBlueprintInfo gets the blueprint_info property value. The blueprint_info property
// returns a BlueprintInfoable when successful
func (m *ObjectTypeInfoDetails) GetBlueprintInfo() BlueprintInfoable {
	return m.blueprint_info
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectTypeInfoDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TypeInfoDetails.GetFieldDeserializers()
	res["blueprint_info"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintInfoFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlueprintInfo(val.(BlueprintInfoable))
		}
		return nil
	}
	res["global"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGlobal(val)
		}
		return nil
	}
	res["module_versions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateModuleVersionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ModuleVersionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ModuleVersionable)
				}
			}
			m.SetModuleVersions(res)
		}
		return nil
	}
	return res
}

// GetGlobal gets the global property value. The global property
// returns a *bool when successful
func (m *ObjectTypeInfoDetails) GetGlobal() *bool {
	return m.global
}

// GetModuleVersions gets the module_versions property value. The module_versions property
// returns a []ModuleVersionable when successful
func (m *ObjectTypeInfoDetails) GetModuleVersions() []ModuleVersionable {
	return m.module_versions
}

// Serialize serializes information the current object
func (m *ObjectTypeInfoDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TypeInfoDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("blueprint_info", m.GetBlueprintInfo())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("global", m.GetGlobal())
		if err != nil {
			return err
		}
	}
	if m.GetModuleVersions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModuleVersions()))
		for i, v := range m.GetModuleVersions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("module_versions", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBlueprintInfo sets the blueprint_info property value. The blueprint_info property
func (m *ObjectTypeInfoDetails) SetBlueprintInfo(value BlueprintInfoable) {
	m.blueprint_info = value
}

// SetGlobal sets the global property value. The global property
func (m *ObjectTypeInfoDetails) SetGlobal(value *bool) {
	m.global = value
}

// SetModuleVersions sets the module_versions property value. The module_versions property
func (m *ObjectTypeInfoDetails) SetModuleVersions(value []ModuleVersionable) {
	m.module_versions = value
}

type ObjectTypeInfoDetailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TypeInfoDetailsable
	GetBlueprintInfo() BlueprintInfoable
	GetGlobal() *bool
	GetModuleVersions() []ModuleVersionable
	SetBlueprintInfo(value BlueprintInfoable)
	SetGlobal(value *bool)
	SetModuleVersions(value []ModuleVersionable)
}
