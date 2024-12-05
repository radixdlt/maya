package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PendingOwnerStakeWithdrawal struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^10`, marking the epoch when the stake units are unlocked for withdrawal.
    epoch_unlocked *int64
    // A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    stake_unit_amount *string
}
// NewPendingOwnerStakeWithdrawal instantiates a new PendingOwnerStakeWithdrawal and sets the default values.
func NewPendingOwnerStakeWithdrawal()(*PendingOwnerStakeWithdrawal) {
    m := &PendingOwnerStakeWithdrawal{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePendingOwnerStakeWithdrawalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePendingOwnerStakeWithdrawalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingOwnerStakeWithdrawal(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PendingOwnerStakeWithdrawal) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEpochUnlocked gets the epoch_unlocked property value. An integer between `0` and `10^10`, marking the epoch when the stake units are unlocked for withdrawal.
// returns a *int64 when successful
func (m *PendingOwnerStakeWithdrawal) GetEpochUnlocked()(*int64) {
    return m.epoch_unlocked
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PendingOwnerStakeWithdrawal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["epoch_unlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpochUnlocked(val)
        }
        return nil
    }
    res["stake_unit_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStakeUnitAmount(val)
        }
        return nil
    }
    return res
}
// GetStakeUnitAmount gets the stake_unit_amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *PendingOwnerStakeWithdrawal) GetStakeUnitAmount()(*string) {
    return m.stake_unit_amount
}
// Serialize serializes information the current object
func (m *PendingOwnerStakeWithdrawal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("epoch_unlocked", m.GetEpochUnlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stake_unit_amount", m.GetStakeUnitAmount())
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
func (m *PendingOwnerStakeWithdrawal) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEpochUnlocked sets the epoch_unlocked property value. An integer between `0` and `10^10`, marking the epoch when the stake units are unlocked for withdrawal.
func (m *PendingOwnerStakeWithdrawal) SetEpochUnlocked(value *int64)() {
    m.epoch_unlocked = value
}
// SetStakeUnitAmount sets the stake_unit_amount property value. A string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *PendingOwnerStakeWithdrawal) SetStakeUnitAmount(value *string)() {
    m.stake_unit_amount = value
}
type PendingOwnerStakeWithdrawalable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEpochUnlocked()(*int64)
    GetStakeUnitAmount()(*string)
    SetEpochUnlocked(value *int64)()
    SetStakeUnitAmount(value *string)()
}
