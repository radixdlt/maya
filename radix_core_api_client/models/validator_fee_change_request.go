package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidatorFeeChangeRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^10`, marking the epoch at which the fee change becomes effective.
    epoch_effective *int64
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    new_fee_factor *string
}
// NewValidatorFeeChangeRequest instantiates a new ValidatorFeeChangeRequest and sets the default values.
func NewValidatorFeeChangeRequest()(*ValidatorFeeChangeRequest) {
    m := &ValidatorFeeChangeRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidatorFeeChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidatorFeeChangeRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidatorFeeChangeRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ValidatorFeeChangeRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEpochEffective gets the epoch_effective property value. An integer between `0` and `10^10`, marking the epoch at which the fee change becomes effective.
// returns a *int64 when successful
func (m *ValidatorFeeChangeRequest) GetEpochEffective()(*int64) {
    return m.epoch_effective
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidatorFeeChangeRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["epoch_effective"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpochEffective(val)
        }
        return nil
    }
    res["new_fee_factor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewFeeFactor(val)
        }
        return nil
    }
    return res
}
// GetNewFeeFactor gets the new_fee_factor property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ValidatorFeeChangeRequest) GetNewFeeFactor()(*string) {
    return m.new_fee_factor
}
// Serialize serializes information the current object
func (m *ValidatorFeeChangeRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("epoch_effective", m.GetEpochEffective())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("new_fee_factor", m.GetNewFeeFactor())
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
func (m *ValidatorFeeChangeRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEpochEffective sets the epoch_effective property value. An integer between `0` and `10^10`, marking the epoch at which the fee change becomes effective.
func (m *ValidatorFeeChangeRequest) SetEpochEffective(value *int64)() {
    m.epoch_effective = value
}
// SetNewFeeFactor sets the new_fee_factor property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ValidatorFeeChangeRequest) SetNewFeeFactor(value *string)() {
    m.new_fee_factor = value
}
type ValidatorFeeChangeRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEpochEffective()(*int64)
    GetNewFeeFactor()(*string)
    SetEpochEffective(value *int64)()
    SetNewFeeFactor(value *string)()
}
