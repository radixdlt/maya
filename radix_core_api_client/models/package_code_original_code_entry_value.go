package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageCodeOriginalCodeEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Either the hex-encoded WASM package code (if Scrypto), or the native package identifier.
	code_hex *string
}

// NewPackageCodeOriginalCodeEntryValue instantiates a new PackageCodeOriginalCodeEntryValue and sets the default values.
func NewPackageCodeOriginalCodeEntryValue() *PackageCodeOriginalCodeEntryValue {
	m := &PackageCodeOriginalCodeEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageCodeOriginalCodeEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageCodeOriginalCodeEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageCodeOriginalCodeEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageCodeOriginalCodeEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCodeHex gets the code_hex property value. Either the hex-encoded WASM package code (if Scrypto), or the native package identifier.
// returns a *string when successful
func (m *PackageCodeOriginalCodeEntryValue) GetCodeHex() *string {
	return m.code_hex
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageCodeOriginalCodeEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["code_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCodeHex(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *PackageCodeOriginalCodeEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("code_hex", m.GetCodeHex())
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
func (m *PackageCodeOriginalCodeEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCodeHex sets the code_hex property value. Either the hex-encoded WASM package code (if Scrypto), or the native package identifier.
func (m *PackageCodeOriginalCodeEntryValue) SetCodeHex(value *string) {
	m.code_hex = value
}

type PackageCodeOriginalCodeEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCodeHex() *string
	SetCodeHex(value *string)
}
