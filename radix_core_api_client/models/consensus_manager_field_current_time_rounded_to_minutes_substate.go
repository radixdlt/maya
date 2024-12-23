package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate struct {
    Substate
    // The value property
    value ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable
}
// NewConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate instantiates a new ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate and sets the default values.
func NewConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate()(*ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate) {
    m := &ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateConsensusManagerFieldCurrentTimeRoundedToMinutesSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentTimeRoundedToMinutesSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConsensusManagerFieldCurrentTimeRoundedToMinutesValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable when successful
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate) GetValue()(ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate) SetValue(value ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable)() {
    m.value = value
}
type ConsensusManagerFieldCurrentTimeRoundedToMinutesSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable)
    SetValue(value ConsensusManagerFieldCurrentTimeRoundedToMinutesValueable)()
}
