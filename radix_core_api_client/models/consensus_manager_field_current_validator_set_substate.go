package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentValidatorSetSubstate struct {
    Substate
    // The value property
    value ConsensusManagerFieldCurrentValidatorSetValueable
}
// NewConsensusManagerFieldCurrentValidatorSetSubstate instantiates a new ConsensusManagerFieldCurrentValidatorSetSubstate and sets the default values.
func NewConsensusManagerFieldCurrentValidatorSetSubstate()(*ConsensusManagerFieldCurrentValidatorSetSubstate) {
    m := &ConsensusManagerFieldCurrentValidatorSetSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateConsensusManagerFieldCurrentValidatorSetSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentValidatorSetSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldCurrentValidatorSetSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentValidatorSetSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConsensusManagerFieldCurrentValidatorSetValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ConsensusManagerFieldCurrentValidatorSetValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldCurrentValidatorSetValueable when successful
func (m *ConsensusManagerFieldCurrentValidatorSetSubstate) GetValue()(ConsensusManagerFieldCurrentValidatorSetValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentValidatorSetSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ConsensusManagerFieldCurrentValidatorSetSubstate) SetValue(value ConsensusManagerFieldCurrentValidatorSetValueable)() {
    m.value = value
}
type ConsensusManagerFieldCurrentValidatorSetSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(ConsensusManagerFieldCurrentValidatorSetValueable)
    SetValue(value ConsensusManagerFieldCurrentValidatorSetValueable)()
}
