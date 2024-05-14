package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FrozenStatus struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The is_burn_frozen property
	is_burn_frozen *bool
	// The is_deposit_frozen property
	is_deposit_frozen *bool
	// The is_withdraw_frozen property
	is_withdraw_frozen *bool
}

// NewFrozenStatus instantiates a new FrozenStatus and sets the default values.
func NewFrozenStatus() *FrozenStatus {
	m := &FrozenStatus{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFrozenStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFrozenStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFrozenStatus(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FrozenStatus) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FrozenStatus) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["is_burn_frozen"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsBurnFrozen(val)
		}
		return nil
	}
	res["is_deposit_frozen"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsDepositFrozen(val)
		}
		return nil
	}
	res["is_withdraw_frozen"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsWithdrawFrozen(val)
		}
		return nil
	}
	return res
}

// GetIsBurnFrozen gets the is_burn_frozen property value. The is_burn_frozen property
// returns a *bool when successful
func (m *FrozenStatus) GetIsBurnFrozen() *bool {
	return m.is_burn_frozen
}

// GetIsDepositFrozen gets the is_deposit_frozen property value. The is_deposit_frozen property
// returns a *bool when successful
func (m *FrozenStatus) GetIsDepositFrozen() *bool {
	return m.is_deposit_frozen
}

// GetIsWithdrawFrozen gets the is_withdraw_frozen property value. The is_withdraw_frozen property
// returns a *bool when successful
func (m *FrozenStatus) GetIsWithdrawFrozen() *bool {
	return m.is_withdraw_frozen
}

// Serialize serializes information the current object
func (m *FrozenStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("is_burn_frozen", m.GetIsBurnFrozen())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_deposit_frozen", m.GetIsDepositFrozen())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_withdraw_frozen", m.GetIsWithdrawFrozen())
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
func (m *FrozenStatus) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIsBurnFrozen sets the is_burn_frozen property value. The is_burn_frozen property
func (m *FrozenStatus) SetIsBurnFrozen(value *bool) {
	m.is_burn_frozen = value
}

// SetIsDepositFrozen sets the is_deposit_frozen property value. The is_deposit_frozen property
func (m *FrozenStatus) SetIsDepositFrozen(value *bool) {
	m.is_deposit_frozen = value
}

// SetIsWithdrawFrozen sets the is_withdraw_frozen property value. The is_withdraw_frozen property
func (m *FrozenStatus) SetIsWithdrawFrozen(value *bool) {
	m.is_withdraw_frozen = value
}

type FrozenStatusable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIsBurnFrozen() *bool
	GetIsDepositFrozen() *bool
	GetIsWithdrawFrozen() *bool
	SetIsBurnFrozen(value *bool)
	SetIsDepositFrozen(value *bool)
	SetIsWithdrawFrozen(value *bool)
}
