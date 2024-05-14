package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaymentToRoyaltyRecipient struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The royalty_recipient property
	royalty_recipient EntityReferenceable
	// The string-encoded decimal representing the amount of fee in XRD paid as royalty to this recipient.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	xrd_amount *string
}

// NewPaymentToRoyaltyRecipient instantiates a new PaymentToRoyaltyRecipient and sets the default values.
func NewPaymentToRoyaltyRecipient() *PaymentToRoyaltyRecipient {
	m := &PaymentToRoyaltyRecipient{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePaymentToRoyaltyRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaymentToRoyaltyRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaymentToRoyaltyRecipient(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaymentToRoyaltyRecipient) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaymentToRoyaltyRecipient) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["royalty_recipient"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyaltyRecipient(val.(EntityReferenceable))
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

// GetRoyaltyRecipient gets the royalty_recipient property value. The royalty_recipient property
// returns a EntityReferenceable when successful
func (m *PaymentToRoyaltyRecipient) GetRoyaltyRecipient() EntityReferenceable {
	return m.royalty_recipient
}

// GetXrdAmount gets the xrd_amount property value. The string-encoded decimal representing the amount of fee in XRD paid as royalty to this recipient.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *PaymentToRoyaltyRecipient) GetXrdAmount() *string {
	return m.xrd_amount
}

// Serialize serializes information the current object
func (m *PaymentToRoyaltyRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("royalty_recipient", m.GetRoyaltyRecipient())
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
func (m *PaymentToRoyaltyRecipient) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetRoyaltyRecipient sets the royalty_recipient property value. The royalty_recipient property
func (m *PaymentToRoyaltyRecipient) SetRoyaltyRecipient(value EntityReferenceable) {
	m.royalty_recipient = value
}

// SetXrdAmount sets the xrd_amount property value. The string-encoded decimal representing the amount of fee in XRD paid as royalty to this recipient.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *PaymentToRoyaltyRecipient) SetXrdAmount(value *string) {
	m.xrd_amount = value
}

type PaymentToRoyaltyRecipientable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRoyaltyRecipient() EntityReferenceable
	GetXrdAmount() *string
	SetRoyaltyRecipient(value EntityReferenceable)
	SetXrdAmount(value *string)
}
