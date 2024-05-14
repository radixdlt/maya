package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageExport struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded code hash, capturing the vm-type and the code itself.
	code_hash *string
	// The export_name property
	export_name *string
}

// NewPackageExport instantiates a new PackageExport and sets the default values.
func NewPackageExport() *PackageExport {
	m := &PackageExport{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageExport(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageExport) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCodeHash gets the code_hash property value. The hex-encoded code hash, capturing the vm-type and the code itself.
// returns a *string when successful
func (m *PackageExport) GetCodeHash() *string {
	return m.code_hash
}

// GetExportName gets the export_name property value. The export_name property
// returns a *string when successful
func (m *PackageExport) GetExportName() *string {
	return m.export_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageExport) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["code_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCodeHash(val)
		}
		return nil
	}
	res["export_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExportName(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *PackageExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("code_hash", m.GetCodeHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("export_name", m.GetExportName())
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
func (m *PackageExport) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCodeHash sets the code_hash property value. The hex-encoded code hash, capturing the vm-type and the code itself.
func (m *PackageExport) SetCodeHash(value *string) {
	m.code_hash = value
}

// SetExportName sets the export_name property value. The export_name property
func (m *PackageExport) SetExportName(value *string) {
	m.export_name = value
}

type PackageExportable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCodeHash() *string
	GetExportName() *string
	SetCodeHash(value *string)
	SetExportName(value *string)
}
