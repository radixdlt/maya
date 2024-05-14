package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModuleVersion struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The module property
	module *AttachedModuleId
	// A string of `Major.Minor.Patch` where Major, Minor and Patch are all u32s.
	version *string
}

// NewModuleVersion instantiates a new ModuleVersion and sets the default values.
func NewModuleVersion() *ModuleVersion {
	m := &ModuleVersion{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateModuleVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModuleVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewModuleVersion(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModuleVersion) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModuleVersion) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["module"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseAttachedModuleId)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetModule(val.(*AttachedModuleId))
		}
		return nil
	}
	res["version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVersion(val)
		}
		return nil
	}
	return res
}

// GetModule gets the module property value. The module property
// returns a *AttachedModuleId when successful
func (m *ModuleVersion) GetModule() *AttachedModuleId {
	return m.module
}

// GetVersion gets the version property value. A string of `Major.Minor.Patch` where Major, Minor and Patch are all u32s.
// returns a *string when successful
func (m *ModuleVersion) GetVersion() *string {
	return m.version
}

// Serialize serializes information the current object
func (m *ModuleVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetModule() != nil {
		cast := (*m.GetModule()).String()
		err := writer.WriteStringValue("module", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *ModuleVersion) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetModule sets the module property value. The module property
func (m *ModuleVersion) SetModule(value *AttachedModuleId) {
	m.module = value
}

// SetVersion sets the version property value. A string of `Major.Minor.Patch` where Major, Minor and Patch are all u32s.
func (m *ModuleVersion) SetVersion(value *string) {
	m.version = value
}

type ModuleVersionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetModule() *AttachedModuleId
	GetVersion() *string
	SetModule(value *AttachedModuleId)
	SetVersion(value *string)
}
