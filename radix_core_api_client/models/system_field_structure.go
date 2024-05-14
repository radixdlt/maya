package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SystemFieldStructure struct {
	SubstateSystemStructure
	// The boot_loader_type property
	boot_loader_type *BootLoaderType
	// The field_kind property
	field_kind *SystemFieldKind
}

// NewSystemFieldStructure instantiates a new SystemFieldStructure and sets the default values.
func NewSystemFieldStructure() *SystemFieldStructure {
	m := &SystemFieldStructure{
		SubstateSystemStructure: *NewSubstateSystemStructure(),
	}
	return m
}

// CreateSystemFieldStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSystemFieldStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSystemFieldStructure(), nil
}

// GetBootLoaderType gets the boot_loader_type property value. The boot_loader_type property
// returns a *BootLoaderType when successful
func (m *SystemFieldStructure) GetBootLoaderType() *BootLoaderType {
	return m.boot_loader_type
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SystemFieldStructure) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SubstateSystemStructure.GetFieldDeserializers()
	res["boot_loader_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseBootLoaderType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBootLoaderType(val.(*BootLoaderType))
		}
		return nil
	}
	res["field_kind"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseSystemFieldKind)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFieldKind(val.(*SystemFieldKind))
		}
		return nil
	}
	return res
}

// GetFieldKind gets the field_kind property value. The field_kind property
// returns a *SystemFieldKind when successful
func (m *SystemFieldStructure) GetFieldKind() *SystemFieldKind {
	return m.field_kind
}

// Serialize serializes information the current object
func (m *SystemFieldStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SubstateSystemStructure.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetBootLoaderType() != nil {
		cast := (*m.GetBootLoaderType()).String()
		err = writer.WriteStringValue("boot_loader_type", &cast)
		if err != nil {
			return err
		}
	}
	if m.GetFieldKind() != nil {
		cast := (*m.GetFieldKind()).String()
		err = writer.WriteStringValue("field_kind", &cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBootLoaderType sets the boot_loader_type property value. The boot_loader_type property
func (m *SystemFieldStructure) SetBootLoaderType(value *BootLoaderType) {
	m.boot_loader_type = value
}

// SetFieldKind sets the field_kind property value. The field_kind property
func (m *SystemFieldStructure) SetFieldKind(value *SystemFieldKind) {
	m.field_kind = value
}

type SystemFieldStructureable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubstateSystemStructureable
	GetBootLoaderType() *BootLoaderType
	GetFieldKind() *SystemFieldKind
	SetBootLoaderType(value *BootLoaderType)
	SetFieldKind(value *SystemFieldKind)
}
