package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NextEpoch struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^10`, marking the new epoch
    epoch *int64
    // The significant_protocol_update_readiness property
    significant_protocol_update_readiness []SignificantProtocolUpdateReadinessEntryable
    // Active validator set for the new epoch, ordered by stake descending.
    validators []ActiveValidatorable
}
// NewNextEpoch instantiates a new NextEpoch and sets the default values.
func NewNextEpoch()(*NextEpoch) {
    m := &NextEpoch{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNextEpochFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNextEpochFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNextEpoch(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NextEpoch) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEpoch gets the epoch property value. An integer between `0` and `10^10`, marking the new epoch
// returns a *int64 when successful
func (m *NextEpoch) GetEpoch()(*int64) {
    return m.epoch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NextEpoch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEpoch(val)
        }
        return nil
    }
    res["significant_protocol_update_readiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignificantProtocolUpdateReadinessEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignificantProtocolUpdateReadinessEntryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SignificantProtocolUpdateReadinessEntryable)
                }
            }
            m.SetSignificantProtocolUpdateReadiness(res)
        }
        return nil
    }
    res["validators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActiveValidatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActiveValidatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActiveValidatorable)
                }
            }
            m.SetValidators(res)
        }
        return nil
    }
    return res
}
// GetSignificantProtocolUpdateReadiness gets the significant_protocol_update_readiness property value. The significant_protocol_update_readiness property
// returns a []SignificantProtocolUpdateReadinessEntryable when successful
func (m *NextEpoch) GetSignificantProtocolUpdateReadiness()([]SignificantProtocolUpdateReadinessEntryable) {
    return m.significant_protocol_update_readiness
}
// GetValidators gets the validators property value. Active validator set for the new epoch, ordered by stake descending.
// returns a []ActiveValidatorable when successful
func (m *NextEpoch) GetValidators()([]ActiveValidatorable) {
    return m.validators
}
// Serialize serializes information the current object
func (m *NextEpoch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("epoch", m.GetEpoch())
        if err != nil {
            return err
        }
    }
    if m.GetSignificantProtocolUpdateReadiness() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignificantProtocolUpdateReadiness()))
        for i, v := range m.GetSignificantProtocolUpdateReadiness() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("significant_protocol_update_readiness", cast)
        if err != nil {
            return err
        }
    }
    if m.GetValidators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValidators()))
        for i, v := range m.GetValidators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("validators", cast)
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
func (m *NextEpoch) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEpoch sets the epoch property value. An integer between `0` and `10^10`, marking the new epoch
func (m *NextEpoch) SetEpoch(value *int64)() {
    m.epoch = value
}
// SetSignificantProtocolUpdateReadiness sets the significant_protocol_update_readiness property value. The significant_protocol_update_readiness property
func (m *NextEpoch) SetSignificantProtocolUpdateReadiness(value []SignificantProtocolUpdateReadinessEntryable)() {
    m.significant_protocol_update_readiness = value
}
// SetValidators sets the validators property value. Active validator set for the new epoch, ordered by stake descending.
func (m *NextEpoch) SetValidators(value []ActiveValidatorable)() {
    m.validators = value
}
type NextEpochable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEpoch()(*int64)
    GetSignificantProtocolUpdateReadiness()([]SignificantProtocolUpdateReadinessEntryable)
    GetValidators()([]ActiveValidatorable)
    SetEpoch(value *int64)()
    SetSignificantProtocolUpdateReadiness(value []SignificantProtocolUpdateReadinessEntryable)()
    SetValidators(value []ActiveValidatorable)()
}
