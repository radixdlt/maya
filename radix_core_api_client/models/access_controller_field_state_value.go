package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessControllerFieldStateValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The controlled_vault property
	controlled_vault EntityReferenceable
	// Whether the primary role badge withdraw is currently being attempted.
	has_primary_role_badge_withdraw_attempt *bool
	// Whether the recovery role badge withdraw is currently being attempted.
	has_recovery_role_badge_withdraw_attempt *bool
	// Whether the primary role is currently locked.
	is_primary_role_locked *bool
	// The primary_role_recovery_attempt property
	primary_role_recovery_attempt PrimaryRoleRecoveryAttemptable
	// The Bech32m-encoded human readable version of the resource address
	recovery_badge_resource_address *string
	// The recovery_role_recovery_attempt property
	recovery_role_recovery_attempt RecoveryRoleRecoveryAttemptable
	// An integer between `0` and `2^32 - 1`, specifying the amount of time (in minutes) thatit takes for timed recovery to be done. When not present, then timed recovery can not beperformed through this access controller.
	timed_recovery_delay_minutes *int64
	// The xrd_fee_vault property
	xrd_fee_vault EntityReferenceable
}

// NewAccessControllerFieldStateValue instantiates a new AccessControllerFieldStateValue and sets the default values.
func NewAccessControllerFieldStateValue() *AccessControllerFieldStateValue {
	m := &AccessControllerFieldStateValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAccessControllerFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessControllerFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccessControllerFieldStateValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessControllerFieldStateValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetControlledVault gets the controlled_vault property value. The controlled_vault property
// returns a EntityReferenceable when successful
func (m *AccessControllerFieldStateValue) GetControlledVault() EntityReferenceable {
	return m.controlled_vault
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessControllerFieldStateValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["controlled_vault"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetControlledVault(val.(EntityReferenceable))
		}
		return nil
	}
	res["has_primary_role_badge_withdraw_attempt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHasPrimaryRoleBadgeWithdrawAttempt(val)
		}
		return nil
	}
	res["has_recovery_role_badge_withdraw_attempt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHasRecoveryRoleBadgeWithdrawAttempt(val)
		}
		return nil
	}
	res["is_primary_role_locked"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsPrimaryRoleLocked(val)
		}
		return nil
	}
	res["primary_role_recovery_attempt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePrimaryRoleRecoveryAttemptFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrimaryRoleRecoveryAttempt(val.(PrimaryRoleRecoveryAttemptable))
		}
		return nil
	}
	res["recovery_badge_resource_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecoveryBadgeResourceAddress(val)
		}
		return nil
	}
	res["recovery_role_recovery_attempt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRecoveryRoleRecoveryAttemptFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecoveryRoleRecoveryAttempt(val.(RecoveryRoleRecoveryAttemptable))
		}
		return nil
	}
	res["timed_recovery_delay_minutes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTimedRecoveryDelayMinutes(val)
		}
		return nil
	}
	res["xrd_fee_vault"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetXrdFeeVault(val.(EntityReferenceable))
		}
		return nil
	}
	return res
}

// GetHasPrimaryRoleBadgeWithdrawAttempt gets the has_primary_role_badge_withdraw_attempt property value. Whether the primary role badge withdraw is currently being attempted.
// returns a *bool when successful
func (m *AccessControllerFieldStateValue) GetHasPrimaryRoleBadgeWithdrawAttempt() *bool {
	return m.has_primary_role_badge_withdraw_attempt
}

// GetHasRecoveryRoleBadgeWithdrawAttempt gets the has_recovery_role_badge_withdraw_attempt property value. Whether the recovery role badge withdraw is currently being attempted.
// returns a *bool when successful
func (m *AccessControllerFieldStateValue) GetHasRecoveryRoleBadgeWithdrawAttempt() *bool {
	return m.has_recovery_role_badge_withdraw_attempt
}

// GetIsPrimaryRoleLocked gets the is_primary_role_locked property value. Whether the primary role is currently locked.
// returns a *bool when successful
func (m *AccessControllerFieldStateValue) GetIsPrimaryRoleLocked() *bool {
	return m.is_primary_role_locked
}

// GetPrimaryRoleRecoveryAttempt gets the primary_role_recovery_attempt property value. The primary_role_recovery_attempt property
// returns a PrimaryRoleRecoveryAttemptable when successful
func (m *AccessControllerFieldStateValue) GetPrimaryRoleRecoveryAttempt() PrimaryRoleRecoveryAttemptable {
	return m.primary_role_recovery_attempt
}

// GetRecoveryBadgeResourceAddress gets the recovery_badge_resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *AccessControllerFieldStateValue) GetRecoveryBadgeResourceAddress() *string {
	return m.recovery_badge_resource_address
}

// GetRecoveryRoleRecoveryAttempt gets the recovery_role_recovery_attempt property value. The recovery_role_recovery_attempt property
// returns a RecoveryRoleRecoveryAttemptable when successful
func (m *AccessControllerFieldStateValue) GetRecoveryRoleRecoveryAttempt() RecoveryRoleRecoveryAttemptable {
	return m.recovery_role_recovery_attempt
}

// GetTimedRecoveryDelayMinutes gets the timed_recovery_delay_minutes property value. An integer between `0` and `2^32 - 1`, specifying the amount of time (in minutes) thatit takes for timed recovery to be done. When not present, then timed recovery can not beperformed through this access controller.
// returns a *int64 when successful
func (m *AccessControllerFieldStateValue) GetTimedRecoveryDelayMinutes() *int64 {
	return m.timed_recovery_delay_minutes
}

// GetXrdFeeVault gets the xrd_fee_vault property value. The xrd_fee_vault property
// returns a EntityReferenceable when successful
func (m *AccessControllerFieldStateValue) GetXrdFeeVault() EntityReferenceable {
	return m.xrd_fee_vault
}

// Serialize serializes information the current object
func (m *AccessControllerFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("controlled_vault", m.GetControlledVault())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("has_primary_role_badge_withdraw_attempt", m.GetHasPrimaryRoleBadgeWithdrawAttempt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("has_recovery_role_badge_withdraw_attempt", m.GetHasRecoveryRoleBadgeWithdrawAttempt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_primary_role_locked", m.GetIsPrimaryRoleLocked())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("primary_role_recovery_attempt", m.GetPrimaryRoleRecoveryAttempt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("recovery_badge_resource_address", m.GetRecoveryBadgeResourceAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("recovery_role_recovery_attempt", m.GetRecoveryRoleRecoveryAttempt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("timed_recovery_delay_minutes", m.GetTimedRecoveryDelayMinutes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("xrd_fee_vault", m.GetXrdFeeVault())
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
func (m *AccessControllerFieldStateValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetControlledVault sets the controlled_vault property value. The controlled_vault property
func (m *AccessControllerFieldStateValue) SetControlledVault(value EntityReferenceable) {
	m.controlled_vault = value
}

// SetHasPrimaryRoleBadgeWithdrawAttempt sets the has_primary_role_badge_withdraw_attempt property value. Whether the primary role badge withdraw is currently being attempted.
func (m *AccessControllerFieldStateValue) SetHasPrimaryRoleBadgeWithdrawAttempt(value *bool) {
	m.has_primary_role_badge_withdraw_attempt = value
}

// SetHasRecoveryRoleBadgeWithdrawAttempt sets the has_recovery_role_badge_withdraw_attempt property value. Whether the recovery role badge withdraw is currently being attempted.
func (m *AccessControllerFieldStateValue) SetHasRecoveryRoleBadgeWithdrawAttempt(value *bool) {
	m.has_recovery_role_badge_withdraw_attempt = value
}

// SetIsPrimaryRoleLocked sets the is_primary_role_locked property value. Whether the primary role is currently locked.
func (m *AccessControllerFieldStateValue) SetIsPrimaryRoleLocked(value *bool) {
	m.is_primary_role_locked = value
}

// SetPrimaryRoleRecoveryAttempt sets the primary_role_recovery_attempt property value. The primary_role_recovery_attempt property
func (m *AccessControllerFieldStateValue) SetPrimaryRoleRecoveryAttempt(value PrimaryRoleRecoveryAttemptable) {
	m.primary_role_recovery_attempt = value
}

// SetRecoveryBadgeResourceAddress sets the recovery_badge_resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *AccessControllerFieldStateValue) SetRecoveryBadgeResourceAddress(value *string) {
	m.recovery_badge_resource_address = value
}

// SetRecoveryRoleRecoveryAttempt sets the recovery_role_recovery_attempt property value. The recovery_role_recovery_attempt property
func (m *AccessControllerFieldStateValue) SetRecoveryRoleRecoveryAttempt(value RecoveryRoleRecoveryAttemptable) {
	m.recovery_role_recovery_attempt = value
}

// SetTimedRecoveryDelayMinutes sets the timed_recovery_delay_minutes property value. An integer between `0` and `2^32 - 1`, specifying the amount of time (in minutes) thatit takes for timed recovery to be done. When not present, then timed recovery can not beperformed through this access controller.
func (m *AccessControllerFieldStateValue) SetTimedRecoveryDelayMinutes(value *int64) {
	m.timed_recovery_delay_minutes = value
}

// SetXrdFeeVault sets the xrd_fee_vault property value. The xrd_fee_vault property
func (m *AccessControllerFieldStateValue) SetXrdFeeVault(value EntityReferenceable) {
	m.xrd_fee_vault = value
}

type AccessControllerFieldStateValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetControlledVault() EntityReferenceable
	GetHasPrimaryRoleBadgeWithdrawAttempt() *bool
	GetHasRecoveryRoleBadgeWithdrawAttempt() *bool
	GetIsPrimaryRoleLocked() *bool
	GetPrimaryRoleRecoveryAttempt() PrimaryRoleRecoveryAttemptable
	GetRecoveryBadgeResourceAddress() *string
	GetRecoveryRoleRecoveryAttempt() RecoveryRoleRecoveryAttemptable
	GetTimedRecoveryDelayMinutes() *int64
	GetXrdFeeVault() EntityReferenceable
	SetControlledVault(value EntityReferenceable)
	SetHasPrimaryRoleBadgeWithdrawAttempt(value *bool)
	SetHasRecoveryRoleBadgeWithdrawAttempt(value *bool)
	SetIsPrimaryRoleLocked(value *bool)
	SetPrimaryRoleRecoveryAttempt(value PrimaryRoleRecoveryAttemptable)
	SetRecoveryBadgeResourceAddress(value *string)
	SetRecoveryRoleRecoveryAttempt(value RecoveryRoleRecoveryAttemptable)
	SetTimedRecoveryDelayMinutes(value *int64)
	SetXrdFeeVault(value EntityReferenceable)
}
