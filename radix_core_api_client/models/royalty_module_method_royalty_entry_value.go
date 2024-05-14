package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoyaltyModuleMethodRoyaltyEntryValue if missing, it represents a free method.
type RoyaltyModuleMethodRoyaltyEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The royalty_amount property
	royalty_amount RoyaltyAmountable
}

// NewRoyaltyModuleMethodRoyaltyEntryValue instantiates a new RoyaltyModuleMethodRoyaltyEntryValue and sets the default values.
func NewRoyaltyModuleMethodRoyaltyEntryValue() *RoyaltyModuleMethodRoyaltyEntryValue {
	m := &RoyaltyModuleMethodRoyaltyEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateRoyaltyModuleMethodRoyaltyEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoyaltyModuleMethodRoyaltyEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoyaltyModuleMethodRoyaltyEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoyaltyModuleMethodRoyaltyEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoyaltyModuleMethodRoyaltyEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["royalty_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRoyaltyAmountFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyaltyAmount(val.(RoyaltyAmountable))
		}
		return nil
	}
	return res
}

// GetRoyaltyAmount gets the royalty_amount property value. The royalty_amount property
// returns a RoyaltyAmountable when successful
func (m *RoyaltyModuleMethodRoyaltyEntryValue) GetRoyaltyAmount() RoyaltyAmountable {
	return m.royalty_amount
}

// Serialize serializes information the current object
func (m *RoyaltyModuleMethodRoyaltyEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("royalty_amount", m.GetRoyaltyAmount())
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
func (m *RoyaltyModuleMethodRoyaltyEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetRoyaltyAmount sets the royalty_amount property value. The royalty_amount property
func (m *RoyaltyModuleMethodRoyaltyEntryValue) SetRoyaltyAmount(value RoyaltyAmountable) {
	m.royalty_amount = value
}

type RoyaltyModuleMethodRoyaltyEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRoyaltyAmount() RoyaltyAmountable
	SetRoyaltyAmount(value RoyaltyAmountable)
}
