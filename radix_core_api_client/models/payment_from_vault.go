package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaymentFromVault struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The vault_entity property
	vault_entity EntityReferenceable
	// The string-encoded decimal representing the amount of fee in XRD paid by this vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	xrd_amount *string
}

// NewPaymentFromVault instantiates a new PaymentFromVault and sets the default values.
func NewPaymentFromVault() *PaymentFromVault {
	m := &PaymentFromVault{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePaymentFromVaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaymentFromVaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaymentFromVault(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaymentFromVault) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaymentFromVault) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["xrd_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetXrdAmount(val)
		}
		return nil
	}
	return res
}

// GetVaultEntity gets the vault_entity property value. The vault_entity property
// returns a EntityReferenceable when successful
func (m *PaymentFromVault) GetVaultEntity() EntityReferenceable {
	return m.vault_entity
}

// GetXrdAmount gets the xrd_amount property value. The string-encoded decimal representing the amount of fee in XRD paid by this vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *PaymentFromVault) GetXrdAmount() *string {
	return m.xrd_amount
}

// Serialize serializes information the current object
func (m *PaymentFromVault) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("vault_entity", m.GetVaultEntity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("xrd_amount", m.GetXrdAmount())
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
func (m *PaymentFromVault) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetVaultEntity sets the vault_entity property value. The vault_entity property
func (m *PaymentFromVault) SetVaultEntity(value EntityReferenceable) {
	m.vault_entity = value
}

// SetXrdAmount sets the xrd_amount property value. The string-encoded decimal representing the amount of fee in XRD paid by this vault.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *PaymentFromVault) SetXrdAmount(value *string) {
	m.xrd_amount = value
}

type PaymentFromVaultable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetVaultEntity() EntityReferenceable
	GetXrdAmount() *string
	SetVaultEntity(value EntityReferenceable)
	SetXrdAmount(value *string)
}
