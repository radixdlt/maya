package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FeeDestination struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The string-encoded decimal representing the amount of fee burnt, in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	to_burn *string
	// The string-encoded decimal representing the amount of fee in XRD paid to the proposer.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	to_proposer *string
	// A breakdown of where the royalties were paid to.
	to_royalty_recipients []PaymentToRoyaltyRecipientable
	// The string-encoded decimal representing the amount of fee in XRD paid to the validator set.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	to_validator_set *string
}

// NewFeeDestination instantiates a new FeeDestination and sets the default values.
func NewFeeDestination() *FeeDestination {
	m := &FeeDestination{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFeeDestinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFeeDestinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFeeDestination(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FeeDestination) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FeeDestination) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["to_burn"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetToBurn(val)
		}
		return nil
	}
	res["to_proposer"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetToProposer(val)
		}
		return nil
	}
	res["to_royalty_recipients"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreatePaymentToRoyaltyRecipientFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]PaymentToRoyaltyRecipientable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(PaymentToRoyaltyRecipientable)
				}
			}
			m.SetToRoyaltyRecipients(res)
		}
		return nil
	}
	res["to_validator_set"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetToValidatorSet(val)
		}
		return nil
	}
	return res
}

// GetToBurn gets the to_burn property value. The string-encoded decimal representing the amount of fee burnt, in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeDestination) GetToBurn() *string {
	return m.to_burn
}

// GetToProposer gets the to_proposer property value. The string-encoded decimal representing the amount of fee in XRD paid to the proposer.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeDestination) GetToProposer() *string {
	return m.to_proposer
}

// GetToRoyaltyRecipients gets the to_royalty_recipients property value. A breakdown of where the royalties were paid to.
// returns a []PaymentToRoyaltyRecipientable when successful
func (m *FeeDestination) GetToRoyaltyRecipients() []PaymentToRoyaltyRecipientable {
	return m.to_royalty_recipients
}

// GetToValidatorSet gets the to_validator_set property value. The string-encoded decimal representing the amount of fee in XRD paid to the validator set.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeDestination) GetToValidatorSet() *string {
	return m.to_validator_set
}

// Serialize serializes information the current object
func (m *FeeDestination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("to_burn", m.GetToBurn())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("to_proposer", m.GetToProposer())
		if err != nil {
			return err
		}
	}
	if m.GetToRoyaltyRecipients() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRoyaltyRecipients()))
		for i, v := range m.GetToRoyaltyRecipients() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("to_royalty_recipients", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("to_validator_set", m.GetToValidatorSet())
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
func (m *FeeDestination) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetToBurn sets the to_burn property value. The string-encoded decimal representing the amount of fee burnt, in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeDestination) SetToBurn(value *string) {
	m.to_burn = value
}

// SetToProposer sets the to_proposer property value. The string-encoded decimal representing the amount of fee in XRD paid to the proposer.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeDestination) SetToProposer(value *string) {
	m.to_proposer = value
}

// SetToRoyaltyRecipients sets the to_royalty_recipients property value. A breakdown of where the royalties were paid to.
func (m *FeeDestination) SetToRoyaltyRecipients(value []PaymentToRoyaltyRecipientable) {
	m.to_royalty_recipients = value
}

// SetToValidatorSet sets the to_validator_set property value. The string-encoded decimal representing the amount of fee in XRD paid to the validator set.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeDestination) SetToValidatorSet(value *string) {
	m.to_validator_set = value
}

type FeeDestinationable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetToBurn() *string
	GetToProposer() *string
	GetToRoyaltyRecipients() []PaymentToRoyaltyRecipientable
	GetToValidatorSet() *string
	SetToBurn(value *string)
	SetToProposer(value *string)
	SetToRoyaltyRecipients(value []PaymentToRoyaltyRecipientable)
	SetToValidatorSet(value *string)
}
