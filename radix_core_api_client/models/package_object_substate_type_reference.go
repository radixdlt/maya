package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageObjectSubstateTypeReference struct {
	ObjectSubstateTypeReference
}

// NewPackageObjectSubstateTypeReference instantiates a new PackageObjectSubstateTypeReference and sets the default values.
func NewPackageObjectSubstateTypeReference() *PackageObjectSubstateTypeReference {
	m := &PackageObjectSubstateTypeReference{
		ObjectSubstateTypeReference: *NewObjectSubstateTypeReference(),
	}
	return m
}

// CreatePackageObjectSubstateTypeReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageObjectSubstateTypeReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageObjectSubstateTypeReference(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageObjectSubstateTypeReference) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ObjectSubstateTypeReference.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *PackageObjectSubstateTypeReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ObjectSubstateTypeReference.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type PackageObjectSubstateTypeReferenceable interface {
	ObjectSubstateTypeReferenceable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
