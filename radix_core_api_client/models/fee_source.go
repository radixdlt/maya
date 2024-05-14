package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FeeSource struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A breakdown of which vaults were used to pay the fee.
	from_vaults []PaymentFromVaultable
}

// NewFeeSource instantiates a new FeeSource and sets the default values.
func NewFeeSource() *FeeSource {
	m := &FeeSource{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFeeSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFeeSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFeeSource(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FeeSource) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FeeSource) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["from_vaults"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreatePaymentFromVaultFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]PaymentFromVaultable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(PaymentFromVaultable)
				}
			}
			m.SetFromVaults(res)
		}
		return nil
	}
	return res
}

// GetFromVaults gets the from_vaults property value. A breakdown of which vaults were used to pay the fee.
// returns a []PaymentFromVaultable when successful
func (m *FeeSource) GetFromVaults() []PaymentFromVaultable {
	return m.from_vaults
}

// Serialize serializes information the current object
func (m *FeeSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetFromVaults() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFromVaults()))
		for i, v := range m.GetFromVaults() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("from_vaults", cast)
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
func (m *FeeSource) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFromVaults sets the from_vaults property value. A breakdown of which vaults were used to pay the fee.
func (m *FeeSource) SetFromVaults(value []PaymentFromVaultable) {
	m.from_vaults = value
}

type FeeSourceable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFromVaults() []PaymentFromVaultable
	SetFromVaults(value []PaymentFromVaultable)
}
