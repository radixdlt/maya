package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageFieldRoyaltyAccumulatorValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The vault_entity property
	vault_entity EntityReferenceable
}

// NewPackageFieldRoyaltyAccumulatorValue instantiates a new PackageFieldRoyaltyAccumulatorValue and sets the default values.
func NewPackageFieldRoyaltyAccumulatorValue() *PackageFieldRoyaltyAccumulatorValue {
	m := &PackageFieldRoyaltyAccumulatorValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageFieldRoyaltyAccumulatorValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageFieldRoyaltyAccumulatorValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageFieldRoyaltyAccumulatorValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageFieldRoyaltyAccumulatorValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageFieldRoyaltyAccumulatorValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["vault_entity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVaultEntity(val.(EntityReferenceable))
		}
		return nil
	}
	return res
}

// GetVaultEntity gets the vault_entity property value. The vault_entity property
// returns a EntityReferenceable when successful
func (m *PackageFieldRoyaltyAccumulatorValue) GetVaultEntity() EntityReferenceable {
	return m.vault_entity
}

// Serialize serializes information the current object
func (m *PackageFieldRoyaltyAccumulatorValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("vault_entity", m.GetVaultEntity())
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
func (m *PackageFieldRoyaltyAccumulatorValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetVaultEntity sets the vault_entity property value. The vault_entity property
func (m *PackageFieldRoyaltyAccumulatorValue) SetVaultEntity(value EntityReferenceable) {
	m.vault_entity = value
}

type PackageFieldRoyaltyAccumulatorValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetVaultEntity() EntityReferenceable
	SetVaultEntity(value EntityReferenceable)
}
