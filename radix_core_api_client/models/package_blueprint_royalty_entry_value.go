package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintRoyaltyEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The royalty_config property
	royalty_config BlueprintRoyaltyConfigable
}

// NewPackageBlueprintRoyaltyEntryValue instantiates a new PackageBlueprintRoyaltyEntryValue and sets the default values.
func NewPackageBlueprintRoyaltyEntryValue() *PackageBlueprintRoyaltyEntryValue {
	m := &PackageBlueprintRoyaltyEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageBlueprintRoyaltyEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintRoyaltyEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageBlueprintRoyaltyEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageBlueprintRoyaltyEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintRoyaltyEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["royalty_config"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintRoyaltyConfigFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyaltyConfig(val.(BlueprintRoyaltyConfigable))
		}
		return nil
	}
	return res
}

// GetRoyaltyConfig gets the royalty_config property value. The royalty_config property
// returns a BlueprintRoyaltyConfigable when successful
func (m *PackageBlueprintRoyaltyEntryValue) GetRoyaltyConfig() BlueprintRoyaltyConfigable {
	return m.royalty_config
}

// Serialize serializes information the current object
func (m *PackageBlueprintRoyaltyEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("royalty_config", m.GetRoyaltyConfig())
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
func (m *PackageBlueprintRoyaltyEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetRoyaltyConfig sets the royalty_config property value. The royalty_config property
func (m *PackageBlueprintRoyaltyEntryValue) SetRoyaltyConfig(value BlueprintRoyaltyConfigable) {
	m.royalty_config = value
}

type PackageBlueprintRoyaltyEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRoyaltyConfig() BlueprintRoyaltyConfigable
	SetRoyaltyConfig(value BlueprintRoyaltyConfigable)
}
