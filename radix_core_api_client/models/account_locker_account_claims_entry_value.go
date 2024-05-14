package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountLockerAccountClaimsEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The resource_vaults property
	resource_vaults EntityReferenceable
}

// NewAccountLockerAccountClaimsEntryValue instantiates a new AccountLockerAccountClaimsEntryValue and sets the default values.
func NewAccountLockerAccountClaimsEntryValue() *AccountLockerAccountClaimsEntryValue {
	m := &AccountLockerAccountClaimsEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAccountLockerAccountClaimsEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountLockerAccountClaimsEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountLockerAccountClaimsEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountLockerAccountClaimsEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountLockerAccountClaimsEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["resource_vaults"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResourceVaults(val.(EntityReferenceable))
		}
		return nil
	}
	return res
}

// GetResourceVaults gets the resource_vaults property value. The resource_vaults property
// returns a EntityReferenceable when successful
func (m *AccountLockerAccountClaimsEntryValue) GetResourceVaults() EntityReferenceable {
	return m.resource_vaults
}

// Serialize serializes information the current object
func (m *AccountLockerAccountClaimsEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("resource_vaults", m.GetResourceVaults())
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
func (m *AccountLockerAccountClaimsEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetResourceVaults sets the resource_vaults property value. The resource_vaults property
func (m *AccountLockerAccountClaimsEntryValue) SetResourceVaults(value EntityReferenceable) {
	m.resource_vaults = value
}

type AccountLockerAccountClaimsEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetResourceVaults() EntityReferenceable
	SetResourceVaults(value EntityReferenceable)
}
