package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StaticRoleDefinitionAuthTemplate_roles a map from role name to role details
type StaticRoleDefinitionAuthTemplate_roles struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewStaticRoleDefinitionAuthTemplate_roles instantiates a new StaticRoleDefinitionAuthTemplate_roles and sets the default values.
func NewStaticRoleDefinitionAuthTemplate_roles() *StaticRoleDefinitionAuthTemplate_roles {
	m := &StaticRoleDefinitionAuthTemplate_roles{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStaticRoleDefinitionAuthTemplate_rolesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStaticRoleDefinitionAuthTemplate_rolesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStaticRoleDefinitionAuthTemplate_roles(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StaticRoleDefinitionAuthTemplate_roles) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StaticRoleDefinitionAuthTemplate_roles) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	return res
}

// Serialize serializes information the current object
func (m *StaticRoleDefinitionAuthTemplate_roles) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StaticRoleDefinitionAuthTemplate_roles) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

type StaticRoleDefinitionAuthTemplate_rolesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
