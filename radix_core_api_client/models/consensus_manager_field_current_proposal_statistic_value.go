package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldCurrentProposalStatisticValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of successfully completed proposals this epoch for each validator, indexed by the validator order in the active set.
    completed []int64
    // The number of missed proposals this epoch for each validator, indexed by the validator order in the active set.
    missed []int64
}
// NewConsensusManagerFieldCurrentProposalStatisticValue instantiates a new ConsensusManagerFieldCurrentProposalStatisticValue and sets the default values.
func NewConsensusManagerFieldCurrentProposalStatisticValue()(*ConsensusManagerFieldCurrentProposalStatisticValue) {
    m := &ConsensusManagerFieldCurrentProposalStatisticValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConsensusManagerFieldCurrentProposalStatisticValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldCurrentProposalStatisticValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldCurrentProposalStatisticValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompleted gets the completed property value. The number of successfully completed proposals this epoch for each validator, indexed by the validator order in the active set.
// returns a []int64 when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) GetCompleted()([]int64) {
    return m.completed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int64))
                }
            }
            m.SetCompleted(res)
        }
        return nil
    }
    res["missed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int64))
                }
            }
            m.SetMissed(res)
        }
        return nil
    }
    return res
}
// GetMissed gets the missed property value. The number of missed proposals this epoch for each validator, indexed by the validator order in the active set.
// returns a []int64 when successful
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) GetMissed()([]int64) {
    return m.missed
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompleted() != nil {
        err := writer.WriteCollectionOfInt64Values("completed", m.GetCompleted())
        if err != nil {
            return err
        }
    }
    if m.GetMissed() != nil {
        err := writer.WriteCollectionOfInt64Values("missed", m.GetMissed())
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
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompleted sets the completed property value. The number of successfully completed proposals this epoch for each validator, indexed by the validator order in the active set.
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) SetCompleted(value []int64)() {
    m.completed = value
}
// SetMissed sets the missed property value. The number of missed proposals this epoch for each validator, indexed by the validator order in the active set.
func (m *ConsensusManagerFieldCurrentProposalStatisticValue) SetMissed(value []int64)() {
    m.missed = value
}
type ConsensusManagerFieldCurrentProposalStatisticValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompleted()([]int64)
    GetMissed()([]int64)
    SetCompleted(value []int64)()
    SetMissed(value []int64)()
}
