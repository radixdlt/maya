package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoyaltyModuleFieldStateValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The is_enabled property
	is_enabled *bool
	// The vault_entity property
	vault_entity EntityReferenceable
}

// NewRoyaltyModuleFieldStateValue instantiates a new RoyaltyModuleFieldStateValue and sets the default values.
func NewRoyaltyModuleFieldStateValue() *RoyaltyModuleFieldStateValue {
	m := &RoyaltyModuleFieldStateValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateRoyaltyModuleFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoyaltyModuleFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoyaltyModuleFieldStateValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoyaltyModuleFieldStateValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoyaltyModuleFieldStateValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["is_enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsEnabled(val)
		}
		return nil
	}
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

// GetIsEnabled gets the is_enabled property value. The is_enabled property
// returns a *bool when successful
func (m *RoyaltyModuleFieldStateValue) GetIsEnabled() *bool {
	return m.is_enabled
}

// GetVaultEntity gets the vault_entity property value. The vault_entity property
// returns a EntityReferenceable when successful
func (m *RoyaltyModuleFieldStateValue) GetVaultEntity() EntityReferenceable {
	return m.vault_entity
}

// Serialize serializes information the current object
func (m *RoyaltyModuleFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("is_enabled", m.GetIsEnabled())
		if err != nil {
			return err
		}
	}
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
func (m *RoyaltyModuleFieldStateValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIsEnabled sets the is_enabled property value. The is_enabled property
func (m *RoyaltyModuleFieldStateValue) SetIsEnabled(value *bool) {
	m.is_enabled = value
}

// SetVaultEntity sets the vault_entity property value. The vault_entity property
func (m *RoyaltyModuleFieldStateValue) SetVaultEntity(value EntityReferenceable) {
	m.vault_entity = value
}

type RoyaltyModuleFieldStateValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIsEnabled() *bool
	GetVaultEntity() EntityReferenceable
	SetIsEnabled(value *bool)
	SetVaultEntity(value EntityReferenceable)
}
