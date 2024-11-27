package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldConfigValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The epoch_change_condition property
    epoch_change_condition EpochChangeConditionable
    // An integer between `0` and `10^10`, specifying the maximum number of validatorsin the active validator set.
    max_validators *int64
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    min_validator_reliability *string
    // An integer between `0` and `10^10`, specifying the minimum number of epochsbefore a fee increase takes effect.
    num_fee_increase_delay_epochs *int64
    // An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an owner can take their stake units after attempting to withdraw them.
    num_owner_stake_units_unlock_epochs *int64
    // An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an unstaker can withdraw their XRD.
    num_unstake_epochs *int64
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    total_emission_xrd_per_epoch *string
    // The defining decimal cost of a validator in USD.This is turned into an XRD cost through the current protocol-based USD/XRD multiplier.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    validator_creation_usd_equivalent_cost *string
    // The decimal amount of XRD required to be passed in a bucket to create a validator.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    validator_creation_xrd_cost *string
}
// NewConsensusManagerFieldConfigValue instantiates a new ConsensusManagerFieldConfigValue and sets the default values.
func NewConsensusManagerFieldConfigValue()(*ConsensusManagerFieldConfigValue) {
    m := &ConsensusManagerFieldConfigValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConsensusManagerFieldConfigValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldConfigValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldConfigValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldConfigValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEpochChangeCondition gets the epoch_change_condition property value. The epoch_change_condition property
// returns a EpochChangeConditionable when successful
func (m *ConsensusManagerFieldConfigValue) GetEpochChangeCondition()(EpochChangeConditionable) {
    return m.epoch_change_condition
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldConfigValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["epoch_change_condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEpochChangeConditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpochChangeCondition(val.(EpochChangeConditionable))
        }
        return nil
    }
    res["max_validators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxValidators(val)
        }
        return nil
    }
    res["min_validator_reliability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinValidatorReliability(val)
        }
        return nil
    }
    res["num_fee_increase_delay_epochs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumFeeIncreaseDelayEpochs(val)
        }
        return nil
    }
    res["num_owner_stake_units_unlock_epochs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumOwnerStakeUnitsUnlockEpochs(val)
        }
        return nil
    }
    res["num_unstake_epochs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumUnstakeEpochs(val)
        }
        return nil
    }
    res["total_emission_xrd_per_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEmissionXrdPerEpoch(val)
        }
        return nil
    }
    res["validator_creation_usd_equivalent_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorCreationUsdEquivalentCost(val)
        }
        return nil
    }
    res["validator_creation_xrd_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorCreationXrdCost(val)
        }
        return nil
    }
    return res
}
// GetMaxValidators gets the max_validators property value. An integer between `0` and `10^10`, specifying the maximum number of validatorsin the active validator set.
// returns a *int64 when successful
func (m *ConsensusManagerFieldConfigValue) GetMaxValidators()(*int64) {
    return m.max_validators
}
// GetMinValidatorReliability gets the min_validator_reliability property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ConsensusManagerFieldConfigValue) GetMinValidatorReliability()(*string) {
    return m.min_validator_reliability
}
// GetNumFeeIncreaseDelayEpochs gets the num_fee_increase_delay_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore a fee increase takes effect.
// returns a *int64 when successful
func (m *ConsensusManagerFieldConfigValue) GetNumFeeIncreaseDelayEpochs()(*int64) {
    return m.num_fee_increase_delay_epochs
}
// GetNumOwnerStakeUnitsUnlockEpochs gets the num_owner_stake_units_unlock_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an owner can take their stake units after attempting to withdraw them.
// returns a *int64 when successful
func (m *ConsensusManagerFieldConfigValue) GetNumOwnerStakeUnitsUnlockEpochs()(*int64) {
    return m.num_owner_stake_units_unlock_epochs
}
// GetNumUnstakeEpochs gets the num_unstake_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an unstaker can withdraw their XRD.
// returns a *int64 when successful
func (m *ConsensusManagerFieldConfigValue) GetNumUnstakeEpochs()(*int64) {
    return m.num_unstake_epochs
}
// GetTotalEmissionXrdPerEpoch gets the total_emission_xrd_per_epoch property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ConsensusManagerFieldConfigValue) GetTotalEmissionXrdPerEpoch()(*string) {
    return m.total_emission_xrd_per_epoch
}
// GetValidatorCreationUsdEquivalentCost gets the validator_creation_usd_equivalent_cost property value. The defining decimal cost of a validator in USD.This is turned into an XRD cost through the current protocol-based USD/XRD multiplier.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ConsensusManagerFieldConfigValue) GetValidatorCreationUsdEquivalentCost()(*string) {
    return m.validator_creation_usd_equivalent_cost
}
// GetValidatorCreationXrdCost gets the validator_creation_xrd_cost property value. The decimal amount of XRD required to be passed in a bucket to create a validator.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ConsensusManagerFieldConfigValue) GetValidatorCreationXrdCost()(*string) {
    return m.validator_creation_xrd_cost
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldConfigValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("epoch_change_condition", m.GetEpochChangeCondition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("max_validators", m.GetMaxValidators())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("min_validator_reliability", m.GetMinValidatorReliability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("num_fee_increase_delay_epochs", m.GetNumFeeIncreaseDelayEpochs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("num_owner_stake_units_unlock_epochs", m.GetNumOwnerStakeUnitsUnlockEpochs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("num_unstake_epochs", m.GetNumUnstakeEpochs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("total_emission_xrd_per_epoch", m.GetTotalEmissionXrdPerEpoch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validator_creation_usd_equivalent_cost", m.GetValidatorCreationUsdEquivalentCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validator_creation_xrd_cost", m.GetValidatorCreationXrdCost())
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
func (m *ConsensusManagerFieldConfigValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEpochChangeCondition sets the epoch_change_condition property value. The epoch_change_condition property
func (m *ConsensusManagerFieldConfigValue) SetEpochChangeCondition(value EpochChangeConditionable)() {
    m.epoch_change_condition = value
}
// SetMaxValidators sets the max_validators property value. An integer between `0` and `10^10`, specifying the maximum number of validatorsin the active validator set.
func (m *ConsensusManagerFieldConfigValue) SetMaxValidators(value *int64)() {
    m.max_validators = value
}
// SetMinValidatorReliability sets the min_validator_reliability property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ConsensusManagerFieldConfigValue) SetMinValidatorReliability(value *string)() {
    m.min_validator_reliability = value
}
// SetNumFeeIncreaseDelayEpochs sets the num_fee_increase_delay_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore a fee increase takes effect.
func (m *ConsensusManagerFieldConfigValue) SetNumFeeIncreaseDelayEpochs(value *int64)() {
    m.num_fee_increase_delay_epochs = value
}
// SetNumOwnerStakeUnitsUnlockEpochs sets the num_owner_stake_units_unlock_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an owner can take their stake units after attempting to withdraw them.
func (m *ConsensusManagerFieldConfigValue) SetNumOwnerStakeUnitsUnlockEpochs(value *int64)() {
    m.num_owner_stake_units_unlock_epochs = value
}
// SetNumUnstakeEpochs sets the num_unstake_epochs property value. An integer between `0` and `10^10`, specifying the minimum number of epochsbefore an unstaker can withdraw their XRD.
func (m *ConsensusManagerFieldConfigValue) SetNumUnstakeEpochs(value *int64)() {
    m.num_unstake_epochs = value
}
// SetTotalEmissionXrdPerEpoch sets the total_emission_xrd_per_epoch property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ConsensusManagerFieldConfigValue) SetTotalEmissionXrdPerEpoch(value *string)() {
    m.total_emission_xrd_per_epoch = value
}
// SetValidatorCreationUsdEquivalentCost sets the validator_creation_usd_equivalent_cost property value. The defining decimal cost of a validator in USD.This is turned into an XRD cost through the current protocol-based USD/XRD multiplier.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ConsensusManagerFieldConfigValue) SetValidatorCreationUsdEquivalentCost(value *string)() {
    m.validator_creation_usd_equivalent_cost = value
}
// SetValidatorCreationXrdCost sets the validator_creation_xrd_cost property value. The decimal amount of XRD required to be passed in a bucket to create a validator.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ConsensusManagerFieldConfigValue) SetValidatorCreationXrdCost(value *string)() {
    m.validator_creation_xrd_cost = value
}
type ConsensusManagerFieldConfigValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEpochChangeCondition()(EpochChangeConditionable)
    GetMaxValidators()(*int64)
    GetMinValidatorReliability()(*string)
    GetNumFeeIncreaseDelayEpochs()(*int64)
    GetNumOwnerStakeUnitsUnlockEpochs()(*int64)
    GetNumUnstakeEpochs()(*int64)
    GetTotalEmissionXrdPerEpoch()(*string)
    GetValidatorCreationUsdEquivalentCost()(*string)
    GetValidatorCreationXrdCost()(*string)
    SetEpochChangeCondition(value EpochChangeConditionable)()
    SetMaxValidators(value *int64)()
    SetMinValidatorReliability(value *string)()
    SetNumFeeIncreaseDelayEpochs(value *int64)()
    SetNumOwnerStakeUnitsUnlockEpochs(value *int64)()
    SetNumUnstakeEpochs(value *int64)()
    SetTotalEmissionXrdPerEpoch(value *string)()
    SetValidatorCreationUsdEquivalentCost(value *string)()
    SetValidatorCreationXrdCost(value *string)()
}
