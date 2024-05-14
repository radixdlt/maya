package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageBlueprintAuthTemplateEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The auth_config property
	auth_config AuthConfigable
}

// NewPackageBlueprintAuthTemplateEntryValue instantiates a new PackageBlueprintAuthTemplateEntryValue and sets the default values.
func NewPackageBlueprintAuthTemplateEntryValue() *PackageBlueprintAuthTemplateEntryValue {
	m := &PackageBlueprintAuthTemplateEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePackageBlueprintAuthTemplateEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageBlueprintAuthTemplateEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPackageBlueprintAuthTemplateEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageBlueprintAuthTemplateEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAuthConfig gets the auth_config property value. The auth_config property
// returns a AuthConfigable when successful
func (m *PackageBlueprintAuthTemplateEntryValue) GetAuthConfig() AuthConfigable {
	return m.auth_config
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageBlueprintAuthTemplateEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["auth_config"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAuthConfigFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAuthConfig(val.(AuthConfigable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *PackageBlueprintAuthTemplateEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("auth_config", m.GetAuthConfig())
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
func (m *PackageBlueprintAuthTemplateEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAuthConfig sets the auth_config property value. The auth_config property
func (m *PackageBlueprintAuthTemplateEntryValue) SetAuthConfig(value AuthConfigable) {
	m.auth_config = value
}

type PackageBlueprintAuthTemplateEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAuthConfig() AuthConfigable
	SetAuthConfig(value AuthConfigable)
}
