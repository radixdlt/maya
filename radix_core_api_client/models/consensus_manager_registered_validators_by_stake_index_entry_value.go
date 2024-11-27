package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue struct {
    // The active_validator property
    active_validator ActiveValidatorable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewConsensusManagerRegisteredValidatorsByStakeIndexEntryValue instantiates a new ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue and sets the default values.
func NewConsensusManagerRegisteredValidatorsByStakeIndexEntryValue()(*ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) {
    m := &ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConsensusManagerRegisteredValidatorsByStakeIndexEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerRegisteredValidatorsByStakeIndexEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerRegisteredValidatorsByStakeIndexEntryValue(), nil
}
// GetActiveValidator gets the active_validator property value. The active_validator property
// returns a ActiveValidatorable when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) GetActiveValidator()(ActiveValidatorable) {
    return m.active_validator
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_validator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActiveValidatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveValidator(val.(ActiveValidatorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("active_validator", m.GetActiveValidator())
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
// SetActiveValidator sets the active_validator property value. The active_validator property
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) SetActiveValidator(value ActiveValidatorable)() {
    m.active_validator = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConsensusManagerRegisteredValidatorsByStakeIndexEntryValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ConsensusManagerRegisteredValidatorsByStakeIndexEntryValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveValidator()(ActiveValidatorable)
    SetActiveValidator(value ActiveValidatorable)()
}
