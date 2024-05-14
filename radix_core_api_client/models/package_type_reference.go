package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageTypeReference struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
	full_type_id FullyScopedTypeIdable
}

// NewPackageTypeReference instantiates a new PackageTypeReference and sets the default values.
func NewPackageTypeReference() *PackageTypeReference {
	m := &PackageTypeReference{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageTypeReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageTypeReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageTypeReference(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageTypeReference) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageTypeReference) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["full_type_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFullyScopedTypeIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFullTypeId(val.(FullyScopedTypeIdable))
		}
		return nil
	}
	return res
}

// GetFullTypeId gets the full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
// returns a FullyScopedTypeIdable when successful
func (m *PackageTypeReference) GetFullTypeId() FullyScopedTypeIdable {
	return m.full_type_id
}

// Serialize serializes information the current object
func (m *PackageTypeReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("full_type_id", m.GetFullTypeId())
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
func (m *PackageTypeReference) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFullTypeId sets the full_type_id property value. An identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
func (m *PackageTypeReference) SetFullTypeId(value FullyScopedTypeIdable) {
	m.full_type_id = value
}

type PackageTypeReferenceable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFullTypeId() FullyScopedTypeIdable
	SetFullTypeId(value FullyScopedTypeIdable)
}
