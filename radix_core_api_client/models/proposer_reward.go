package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProposerReward struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The validator_index property
    validator_index ActiveValidatorIndexable
    // The string-encoded decimal representing the amount of reward in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_amount *string
}
// NewProposerReward instantiates a new ProposerReward and sets the default values.
func NewProposerReward()(*ProposerReward) {
    m := &ProposerReward{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProposerRewardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProposerRewardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProposerReward(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProposerReward) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProposerReward) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["validator_index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActiveValidatorIndexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorIndex(val.(ActiveValidatorIndexable))
        }
        return nil
    }
    res["xrd_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdAmount(val)
        }
        return nil
    }
    return res
}
// GetValidatorIndex gets the validator_index property value. The validator_index property
// returns a ActiveValidatorIndexable when successful
func (m *ProposerReward) GetValidatorIndex()(ActiveValidatorIndexable) {
    return m.validator_index
}
// GetXrdAmount gets the xrd_amount property value. The string-encoded decimal representing the amount of reward in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *ProposerReward) GetXrdAmount()(*string) {
    return m.xrd_amount
}
// Serialize serializes information the current object
func (m *ProposerReward) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("validator_index", m.GetValidatorIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_amount", m.GetXrdAmount())
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
func (m *ProposerReward) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValidatorIndex sets the validator_index property value. The validator_index property
func (m *ProposerReward) SetValidatorIndex(value ActiveValidatorIndexable)() {
    m.validator_index = value
}
// SetXrdAmount sets the xrd_amount property value. The string-encoded decimal representing the amount of reward in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *ProposerReward) SetXrdAmount(value *string)() {
    m.xrd_amount = value
}
type ProposerRewardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValidatorIndex()(ActiveValidatorIndexable)
    GetXrdAmount()(*string)
    SetValidatorIndex(value ActiveValidatorIndexable)()
    SetXrdAmount(value *string)()
}
