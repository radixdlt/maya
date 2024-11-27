package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtocolVersionReadiness struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A name identifying a protocol version. May be absent to denote no readiness signalled bythe `signalling_validators`.
    signalled_protocol_version *string
    // References to some of the current validators (i.e. a subset of `current_validator_set`)which have signalled readiness for the `signalled_protocol_version`.
    signalling_validators []SignallingValidatorable
    // A sum of `active_stake_proportion` across `signalling_validators` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    total_active_stake_proportion *string
}
// NewProtocolVersionReadiness instantiates a new ProtocolVersionReadiness and sets the default values.
func NewProtocolVersionReadiness()(*ProtocolVersionReadiness) {
    m := &ProtocolVersionReadiness{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtocolVersionReadinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtocolVersionReadinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtocolVersionReadiness(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtocolVersionReadiness) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtocolVersionReadiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["signalled_protocol_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignalledProtocolVersion(val)
        }
        return nil
    }
    res["signalling_validators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignallingValidatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignallingValidatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SignallingValidatorable)
                }
            }
            m.SetSignallingValidators(res)
        }
        return nil
    }
    res["total_active_stake_proportion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveStakeProportion(val)
        }
        return nil
    }
    return res
}
// GetSignalledProtocolVersion gets the signalled_protocol_version property value. A name identifying a protocol version. May be absent to denote no readiness signalled bythe `signalling_validators`.
// returns a *string when successful
func (m *ProtocolVersionReadiness) GetSignalledProtocolVersion()(*string) {
    return m.signalled_protocol_version
}
// GetSignallingValidators gets the signalling_validators property value. References to some of the current validators (i.e. a subset of `current_validator_set`)which have signalled readiness for the `signalled_protocol_version`.
// returns a []SignallingValidatorable when successful
func (m *ProtocolVersionReadiness) GetSignallingValidators()([]SignallingValidatorable) {
    return m.signalling_validators
}
// GetTotalActiveStakeProportion gets the total_active_stake_proportion property value. A sum of `active_stake_proportion` across `signalling_validators` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ProtocolVersionReadiness) GetTotalActiveStakeProportion()(*string) {
    return m.total_active_stake_proportion
}
// Serialize serializes information the current object
func (m *ProtocolVersionReadiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("signalled_protocol_version", m.GetSignalledProtocolVersion())
        if err != nil {
            return err
        }
    }
    if m.GetSignallingValidators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignallingValidators()))
        for i, v := range m.GetSignallingValidators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("signalling_validators", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("total_active_stake_proportion", m.GetTotalActiveStakeProportion())
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
func (m *ProtocolVersionReadiness) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSignalledProtocolVersion sets the signalled_protocol_version property value. A name identifying a protocol version. May be absent to denote no readiness signalled bythe `signalling_validators`.
func (m *ProtocolVersionReadiness) SetSignalledProtocolVersion(value *string)() {
    m.signalled_protocol_version = value
}
// SetSignallingValidators sets the signalling_validators property value. References to some of the current validators (i.e. a subset of `current_validator_set`)which have signalled readiness for the `signalled_protocol_version`.
func (m *ProtocolVersionReadiness) SetSignallingValidators(value []SignallingValidatorable)() {
    m.signalling_validators = value
}
// SetTotalActiveStakeProportion sets the total_active_stake_proportion property value. A sum of `active_stake_proportion` across `signalling_validators` (i.e. an easily-computable convenience field).This is a string-encoded fixed-precision decimal to 18 decimal places.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ProtocolVersionReadiness) SetTotalActiveStakeProportion(value *string)() {
    m.total_active_stake_proportion = value
}
type ProtocolVersionReadinessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSignalledProtocolVersion()(*string)
    GetSignallingValidators()([]SignallingValidatorable)
    GetTotalActiveStakeProportion()(*string)
    SetSignalledProtocolVersion(value *string)()
    SetSignallingValidators(value []SignallingValidatorable)()
    SetTotalActiveStakeProportion(value *string)()
}
