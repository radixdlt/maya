package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecoveryRoleRecoveryAttempt struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_timed_recovery_after property
    allow_timed_recovery_after ScryptoInstantable
    // The recovery_proposal property
    recovery_proposal RecoveryProposalable
}
// NewRecoveryRoleRecoveryAttempt instantiates a new RecoveryRoleRecoveryAttempt and sets the default values.
func NewRecoveryRoleRecoveryAttempt()(*RecoveryRoleRecoveryAttempt) {
    m := &RecoveryRoleRecoveryAttempt{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecoveryRoleRecoveryAttemptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecoveryRoleRecoveryAttemptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecoveryRoleRecoveryAttempt(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RecoveryRoleRecoveryAttempt) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowTimedRecoveryAfter gets the allow_timed_recovery_after property value. The allow_timed_recovery_after property
// returns a ScryptoInstantable when successful
func (m *RecoveryRoleRecoveryAttempt) GetAllowTimedRecoveryAfter()(ScryptoInstantable) {
    return m.allow_timed_recovery_after
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecoveryRoleRecoveryAttempt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_timed_recovery_after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScryptoInstantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTimedRecoveryAfter(val.(ScryptoInstantable))
        }
        return nil
    }
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
func (m *RecoveryRoleRecoveryAttempt) GetRecoveryProposal()(RecoveryProposalable) {
    return m.recovery_proposal
}
// Serialize serializes information the current object
func (m *RecoveryRoleRecoveryAttempt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("allow_timed_recovery_after", m.GetAllowTimedRecoveryAfter())
        if err != nil {
            return err
        }
    }
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
func (m *RecoveryRoleRecoveryAttempt) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowTimedRecoveryAfter sets the allow_timed_recovery_after property value. The allow_timed_recovery_after property
func (m *RecoveryRoleRecoveryAttempt) SetAllowTimedRecoveryAfter(value ScryptoInstantable)() {
    m.allow_timed_recovery_after = value
}
// SetRecoveryProposal sets the recovery_proposal property value. The recovery_proposal property
func (m *RecoveryRoleRecoveryAttempt) SetRecoveryProposal(value RecoveryProposalable)() {
    m.recovery_proposal = value
}
type RecoveryRoleRecoveryAttemptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowTimedRecoveryAfter()(ScryptoInstantable)
    GetRecoveryProposal()(RecoveryProposalable)
    SetAllowTimedRecoveryAfter(value ScryptoInstantable)()
    SetRecoveryProposal(value RecoveryProposalable)()
}
