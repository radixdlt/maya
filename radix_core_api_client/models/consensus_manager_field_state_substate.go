package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldStateSubstate struct {
    Substate
    // The value property
    value ConsensusManagerFieldStateValueable
}
// NewConsensusManagerFieldStateSubstate instantiates a new ConsensusManagerFieldStateSubstate and sets the default values.
func NewConsensusManagerFieldStateSubstate()(*ConsensusManagerFieldStateSubstate) {
    m := &ConsensusManagerFieldStateSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateConsensusManagerFieldStateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldStateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldStateSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldStateSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConsensusManagerFieldStateValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ConsensusManagerFieldStateValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldStateValueable when successful
func (m *ConsensusManagerFieldStateSubstate) GetValue()(ConsensusManagerFieldStateValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldStateSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConsensusManagerFieldStateSubstate) SetValue(value ConsensusManagerFieldStateValueable)() {
    m.value = value
}
type ConsensusManagerFieldStateSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(ConsensusManagerFieldStateValueable)
    SetValue(value ConsensusManagerFieldStateValueable)()
}
