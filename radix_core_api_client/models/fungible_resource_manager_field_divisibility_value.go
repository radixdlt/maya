package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FungibleResourceManagerFieldDivisibilityValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The divisibility property
	divisibility *int32
}

// NewFungibleResourceManagerFieldDivisibilityValue instantiates a new FungibleResourceManagerFieldDivisibilityValue and sets the default values.
func NewFungibleResourceManagerFieldDivisibilityValue() *FungibleResourceManagerFieldDivisibilityValue {
	m := &FungibleResourceManagerFieldDivisibilityValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFungibleResourceManagerFieldDivisibilityValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFungibleResourceManagerFieldDivisibilityValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFungibleResourceManagerFieldDivisibilityValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FungibleResourceManagerFieldDivisibilityValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDivisibility gets the divisibility property value. The divisibility property
// returns a *int32 when successful
func (m *FungibleResourceManagerFieldDivisibilityValue) GetDivisibility() *int32 {
	return m.divisibility
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FungibleResourceManagerFieldDivisibilityValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["divisibility"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDivisibility(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *FungibleResourceManagerFieldDivisibilityValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("divisibility", m.GetDivisibility())
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
func (m *FungibleResourceManagerFieldDivisibilityValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDivisibility sets the divisibility property value. The divisibility property
func (m *FungibleResourceManagerFieldDivisibilityValue) SetDivisibility(value *int32) {
	m.divisibility = value
}

type FungibleResourceManagerFieldDivisibilityValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDivisibility() *int32
	SetDivisibility(value *int32)
}
