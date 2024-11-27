package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrimaryRoleRecoveryAttempt struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The recovery_proposal property
    recovery_proposal RecoveryProposalable
}
// NewPrimaryRoleRecoveryAttempt instantiates a new PrimaryRoleRecoveryAttempt and sets the default values.
func NewPrimaryRoleRecoveryAttempt()(*PrimaryRoleRecoveryAttempt) {
    m := &PrimaryRoleRecoveryAttempt{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrimaryRoleRecoveryAttemptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrimaryRoleRecoveryAttemptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrimaryRoleRecoveryAttempt(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PrimaryRoleRecoveryAttempt) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrimaryRoleRecoveryAttempt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recovery_proposal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecoveryProposalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryProposal(val.(RecoveryProposalable))
        }
        return nil
    }
    return res
}
// GetRecoveryProposal gets the recovery_proposal property value. The recovery_proposal property
// returns a RecoveryProposalable when successful
func (m *PrimaryRoleRecoveryAttempt) GetRecoveryProposal()(RecoveryProposalable) {
    return m.recovery_proposal
}
// Serialize serializes information the current object
func (m *PrimaryRoleRecoveryAttempt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("recovery_proposal", m.GetRecoveryProposal())
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
func (m *PrimaryRoleRecoveryAttempt) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRecoveryProposal sets the recovery_proposal property value. The recovery_proposal property
func (m *PrimaryRoleRecoveryAttempt) SetRecoveryProposal(value RecoveryProposalable)() {
    m.recovery_proposal = value
}
type PrimaryRoleRecoveryAttemptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecoveryProposal()(RecoveryProposalable)
    SetRecoveryProposal(value RecoveryProposalable)()
}
