package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentProposalStatisticSubstate struct {
    Substate
    // The value property
    value ConsensusManagerFieldCurrentProposalStatisticValueable
}
// NewConsensusManagerFieldCurrentProposalStatisticSubstate instantiates a new ConsensusManagerFieldCurrentProposalStatisticSubstate and sets the default values.
func NewConsensusManagerFieldCurrentProposalStatisticSubstate()(*ConsensusManagerFieldCurrentProposalStatisticSubstate) {
    m := &ConsensusManagerFieldCurrentProposalStatisticSubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateConsensusManagerFieldCurrentProposalStatisticSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentProposalStatisticSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldCurrentProposalStatisticSubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticSubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConsensusManagerFieldCurrentProposalStatisticValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ConsensusManagerFieldCurrentProposalStatisticValueable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a ConsensusManagerFieldCurrentProposalStatisticValueable when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticSubstate) GetValue()(ConsensusManagerFieldCurrentProposalStatisticValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentProposalStatisticSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConsensusManagerFieldCurrentProposalStatisticSubstate) SetValue(value ConsensusManagerFieldCurrentProposalStatisticValueable)() {
    m.value = value
}
type ConsensusManagerFieldCurrentProposalStatisticSubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetValue()(ConsensusManagerFieldCurrentProposalStatisticValueable)
    SetValue(value ConsensusManagerFieldCurrentProposalStatisticValueable)()
}
