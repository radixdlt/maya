package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidatorFieldStateValue struct {
    // The accepts_delegated_stake property
    accepts_delegated_stake *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    already_unlocked_owner_stake_unit_amount *string
    // The Bech32m-encoded human readable version of the resource address
    claim_token_resource_address *string
    // The is_registered property
    is_registered *bool
    // The locked_owner_stake_unit_vault property
    locked_owner_stake_unit_vault EntityReferenceable
    // The pending_owner_stake_unit_unlock_vault property
    pending_owner_stake_unit_unlock_vault EntityReferenceable
    // The pending_owner_stake_unit_withdrawals property
    pending_owner_stake_unit_withdrawals []PendingOwnerStakeWithdrawalable
    // The pending_xrd_withdraw_vault property
    pending_xrd_withdraw_vault EntityReferenceable
    // The public_key property
    public_key EcdsaSecp256k1PublicKeyable
    // The sorted_key property
    sorted_key SubstateKeyable
    // The Bech32m-encoded human readable version of the resource address
    stake_unit_resource_address *string
    // The stake_xrd_vault property
    stake_xrd_vault EntityReferenceable
    // The validator_fee_change_request property
    validator_fee_change_request ValidatorFeeChangeRequestable
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    validator_fee_factor *string
}
// NewValidatorFieldStateValue instantiates a new ValidatorFieldStateValue and sets the default values.
func NewValidatorFieldStateValue()(*ValidatorFieldStateValue) {
    m := &ValidatorFieldStateValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidatorFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidatorFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidatorFieldStateValue(), nil
}
// GetAcceptsDelegatedStake gets the accepts_delegated_stake property value. The accepts_delegated_stake property
// returns a *bool when successful
func (m *ValidatorFieldStateValue) GetAcceptsDelegatedStake()(*bool) {
    return m.accepts_delegated_stake
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ValidatorFieldStateValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlreadyUnlockedOwnerStakeUnitAmount gets the already_unlocked_owner_stake_unit_amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ValidatorFieldStateValue) GetAlreadyUnlockedOwnerStakeUnitAmount()(*string) {
    return m.already_unlocked_owner_stake_unit_amount
}
// GetClaimTokenResourceAddress gets the claim_token_resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ValidatorFieldStateValue) GetClaimTokenResourceAddress()(*string) {
    return m.claim_token_resource_address
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidatorFieldStateValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accepts_delegated_stake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptsDelegatedStake(val)
        }
        return nil
    }
    res["already_unlocked_owner_stake_unit_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlreadyUnlockedOwnerStakeUnitAmount(val)
        }
        return nil
    }
    res["claim_token_resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimTokenResourceAddress(val)
        }
        return nil
    }
    res["is_registered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRegistered(val)
        }
        return nil
    }
    res["locked_owner_stake_unit_vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockedOwnerStakeUnitVault(val.(EntityReferenceable))
        }
        return nil
    }
    res["pending_owner_stake_unit_unlock_vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOwnerStakeUnitUnlockVault(val.(EntityReferenceable))
        }
        return nil
    }
    res["pending_owner_stake_unit_withdrawals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePendingOwnerStakeWithdrawalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PendingOwnerStakeWithdrawalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PendingOwnerStakeWithdrawalable)
                }
            }
            m.SetPendingOwnerStakeUnitWithdrawals(res)
        }
        return nil
    }
    res["pending_xrd_withdraw_vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingXrdWithdrawVault(val.(EntityReferenceable))
        }
        return nil
    }
    res["public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEcdsaSecp256k1PublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKey(val.(EcdsaSecp256k1PublicKeyable))
        }
        return nil
    }
    res["sorted_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortedKey(val.(SubstateKeyable))
        }
        return nil
    }
    res["stake_unit_resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStakeUnitResourceAddress(val)
        }
        return nil
    }
    res["stake_xrd_vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStakeXrdVault(val.(EntityReferenceable))
        }
        return nil
    }
    res["validator_fee_change_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidatorFeeChangeRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorFeeChangeRequest(val.(ValidatorFeeChangeRequestable))
        }
        return nil
    }
    res["validator_fee_factor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorFeeFactor(val)
        }
        return nil
    }
    return res
}
// GetIsRegistered gets the is_registered property value. The is_registered property
// returns a *bool when successful
func (m *ValidatorFieldStateValue) GetIsRegistered()(*bool) {
    return m.is_registered
}
// GetLockedOwnerStakeUnitVault gets the locked_owner_stake_unit_vault property value. The locked_owner_stake_unit_vault property
// returns a EntityReferenceable when successful
func (m *ValidatorFieldStateValue) GetLockedOwnerStakeUnitVault()(EntityReferenceable) {
    return m.locked_owner_stake_unit_vault
}
// GetPendingOwnerStakeUnitUnlockVault gets the pending_owner_stake_unit_unlock_vault property value. The pending_owner_stake_unit_unlock_vault property
// returns a EntityReferenceable when successful
func (m *ValidatorFieldStateValue) GetPendingOwnerStakeUnitUnlockVault()(EntityReferenceable) {
    return m.pending_owner_stake_unit_unlock_vault
}
// GetPendingOwnerStakeUnitWithdrawals gets the pending_owner_stake_unit_withdrawals property value. The pending_owner_stake_unit_withdrawals property
// returns a []PendingOwnerStakeWithdrawalable when successful
func (m *ValidatorFieldStateValue) GetPendingOwnerStakeUnitWithdrawals()([]PendingOwnerStakeWithdrawalable) {
    return m.pending_owner_stake_unit_withdrawals
}
// GetPendingXrdWithdrawVault gets the pending_xrd_withdraw_vault property value. The pending_xrd_withdraw_vault property
// returns a EntityReferenceable when successful
func (m *ValidatorFieldStateValue) GetPendingXrdWithdrawVault()(EntityReferenceable) {
    return m.pending_xrd_withdraw_vault
}
// GetPublicKey gets the public_key property value. The public_key property
// returns a EcdsaSecp256k1PublicKeyable when successful
func (m *ValidatorFieldStateValue) GetPublicKey()(EcdsaSecp256k1PublicKeyable) {
    return m.public_key
}
// GetSortedKey gets the sorted_key property value. The sorted_key property
// returns a SubstateKeyable when successful
func (m *ValidatorFieldStateValue) GetSortedKey()(SubstateKeyable) {
    return m.sorted_key
}
// GetStakeUnitResourceAddress gets the stake_unit_resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ValidatorFieldStateValue) GetStakeUnitResourceAddress()(*string) {
    return m.stake_unit_resource_address
}
// GetStakeXrdVault gets the stake_xrd_vault property value. The stake_xrd_vault property
// returns a EntityReferenceable when successful
func (m *ValidatorFieldStateValue) GetStakeXrdVault()(EntityReferenceable) {
    return m.stake_xrd_vault
}
// GetValidatorFeeChangeRequest gets the validator_fee_change_request property value. The validator_fee_change_request property
// returns a ValidatorFeeChangeRequestable when successful
func (m *ValidatorFieldStateValue) GetValidatorFeeChangeRequest()(ValidatorFeeChangeRequestable) {
    return m.validator_fee_change_request
}
// GetValidatorFeeFactor gets the validator_fee_factor property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ValidatorFieldStateValue) GetValidatorFeeFactor()(*string) {
    return m.validator_fee_factor
}
// Serialize serializes information the current object
func (m *ValidatorFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accepts_delegated_stake", m.GetAcceptsDelegatedStake())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("already_unlocked_owner_stake_unit_amount", m.GetAlreadyUnlockedOwnerStakeUnitAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("claim_token_resource_address", m.GetClaimTokenResourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_registered", m.GetIsRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locked_owner_stake_unit_vault", m.GetLockedOwnerStakeUnitVault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pending_owner_stake_unit_unlock_vault", m.GetPendingOwnerStakeUnitUnlockVault())
        if err != nil {
            return err
        }
    }
    if m.GetPendingOwnerStakeUnitWithdrawals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPendingOwnerStakeUnitWithdrawals()))
        for i, v := range m.GetPendingOwnerStakeUnitWithdrawals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("pending_owner_stake_unit_withdrawals", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pending_xrd_withdraw_vault", m.GetPendingXrdWithdrawVault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("public_key", m.GetPublicKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sorted_key", m.GetSortedKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stake_unit_resource_address", m.GetStakeUnitResourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("stake_xrd_vault", m.GetStakeXrdVault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("validator_fee_change_request", m.GetValidatorFeeChangeRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validator_fee_factor", m.GetValidatorFeeFactor())
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
// SetAcceptsDelegatedStake sets the accepts_delegated_stake property value. The accepts_delegated_stake property
func (m *ValidatorFieldStateValue) SetAcceptsDelegatedStake(value *bool)() {
    m.accepts_delegated_stake = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidatorFieldStateValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlreadyUnlockedOwnerStakeUnitAmount sets the already_unlocked_owner_stake_unit_amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ValidatorFieldStateValue) SetAlreadyUnlockedOwnerStakeUnitAmount(value *string)() {
    m.already_unlocked_owner_stake_unit_amount = value
}
// SetClaimTokenResourceAddress sets the claim_token_resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *ValidatorFieldStateValue) SetClaimTokenResourceAddress(value *string)() {
    m.claim_token_resource_address = value
}
// SetIsRegistered sets the is_registered property value. The is_registered property
func (m *ValidatorFieldStateValue) SetIsRegistered(value *bool)() {
    m.is_registered = value
}
// SetLockedOwnerStakeUnitVault sets the locked_owner_stake_unit_vault property value. The locked_owner_stake_unit_vault property
func (m *ValidatorFieldStateValue) SetLockedOwnerStakeUnitVault(value EntityReferenceable)() {
    m.locked_owner_stake_unit_vault = value
}
// SetPendingOwnerStakeUnitUnlockVault sets the pending_owner_stake_unit_unlock_vault property value. The pending_owner_stake_unit_unlock_vault property
func (m *ValidatorFieldStateValue) SetPendingOwnerStakeUnitUnlockVault(value EntityReferenceable)() {
    m.pending_owner_stake_unit_unlock_vault = value
}
// SetPendingOwnerStakeUnitWithdrawals sets the pending_owner_stake_unit_withdrawals property value. The pending_owner_stake_unit_withdrawals property
func (m *ValidatorFieldStateValue) SetPendingOwnerStakeUnitWithdrawals(value []PendingOwnerStakeWithdrawalable)() {
    m.pending_owner_stake_unit_withdrawals = value
}
// SetPendingXrdWithdrawVault sets the pending_xrd_withdraw_vault property value. The pending_xrd_withdraw_vault property
func (m *ValidatorFieldStateValue) SetPendingXrdWithdrawVault(value EntityReferenceable)() {
    m.pending_xrd_withdraw_vault = value
}
// SetPublicKey sets the public_key property value. The public_key property
func (m *ValidatorFieldStateValue) SetPublicKey(value EcdsaSecp256k1PublicKeyable)() {
    m.public_key = value
}
// SetSortedKey sets the sorted_key property value. The sorted_key property
func (m *ValidatorFieldStateValue) SetSortedKey(value SubstateKeyable)() {
    m.sorted_key = value
}
// SetStakeUnitResourceAddress sets the stake_unit_resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *ValidatorFieldStateValue) SetStakeUnitResourceAddress(value *string)() {
    m.stake_unit_resource_address = value
}
// SetStakeXrdVault sets the stake_xrd_vault property value. The stake_xrd_vault property
func (m *ValidatorFieldStateValue) SetStakeXrdVault(value EntityReferenceable)() {
    m.stake_xrd_vault = value
}
// SetValidatorFeeChangeRequest sets the validator_fee_change_request property value. The validator_fee_change_request property
func (m *ValidatorFieldStateValue) SetValidatorFeeChangeRequest(value ValidatorFeeChangeRequestable)() {
    m.validator_fee_change_request = value
}
// SetValidatorFeeFactor sets the validator_fee_factor property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ValidatorFieldStateValue) SetValidatorFeeFactor(value *string)() {
    m.validator_fee_factor = value
}
type ValidatorFieldStateValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptsDelegatedStake()(*bool)
    GetAlreadyUnlockedOwnerStakeUnitAmount()(*string)
    GetClaimTokenResourceAddress()(*string)
    GetIsRegistered()(*bool)
    GetLockedOwnerStakeUnitVault()(EntityReferenceable)
    GetPendingOwnerStakeUnitUnlockVault()(EntityReferenceable)
    GetPendingOwnerStakeUnitWithdrawals()([]PendingOwnerStakeWithdrawalable)
    GetPendingXrdWithdrawVault()(EntityReferenceable)
    GetPublicKey()(EcdsaSecp256k1PublicKeyable)
    GetSortedKey()(SubstateKeyable)
    GetStakeUnitResourceAddress()(*string)
    GetStakeXrdVault()(EntityReferenceable)
    GetValidatorFeeChangeRequest()(ValidatorFeeChangeRequestable)
    GetValidatorFeeFactor()(*string)
    SetAcceptsDelegatedStake(value *bool)()
    SetAlreadyUnlockedOwnerStakeUnitAmount(value *string)()
    SetClaimTokenResourceAddress(value *string)()
    SetIsRegistered(value *bool)()
    SetLockedOwnerStakeUnitVault(value EntityReferenceable)()
    SetPendingOwnerStakeUnitUnlockVault(value EntityReferenceable)()
    SetPendingOwnerStakeUnitWithdrawals(value []PendingOwnerStakeWithdrawalable)()
    SetPendingXrdWithdrawVault(value EntityReferenceable)()
    SetPublicKey(value EcdsaSecp256k1PublicKeyable)()
    SetSortedKey(value SubstateKeyable)()
    SetStakeUnitResourceAddress(value *string)()
    SetStakeXrdVault(value EntityReferenceable)()
    SetValidatorFeeChangeRequest(value ValidatorFeeChangeRequestable)()
    SetValidatorFeeFactor(value *string)()
}
