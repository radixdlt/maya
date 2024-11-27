package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecoveryProposal struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
    confirmation_role AccessRuleable
    // A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
    primary_role AccessRuleable
    // A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
    recovery_role AccessRuleable
    // An integer between `0` and `2^32 - 1`, specifying the optional proposal delay of timed recoveries.
    timed_recovery_delay_minutes *int64
}
// NewRecoveryProposal instantiates a new RecoveryProposal and sets the default values.
func NewRecoveryProposal()(*RecoveryProposal) {
    m := &RecoveryProposal{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecoveryProposalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecoveryProposalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecoveryProposal(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RecoveryProposal) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfirmationRole gets the confirmation_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
// returns a AccessRuleable when successful
func (m *RecoveryProposal) GetConfirmationRole()(AccessRuleable) {
    return m.confirmation_role
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecoveryProposal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confirmation_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmationRole(val.(AccessRuleable))
        }
        return nil
    }
    res["primary_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryRole(val.(AccessRuleable))
        }
        return nil
    }
    res["recovery_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryRole(val.(AccessRuleable))
        }
        return nil
    }
    res["timed_recovery_delay_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimedRecoveryDelayMinutes(val)
        }
        return nil
    }
    return res
}
// GetPrimaryRole gets the primary_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
// returns a AccessRuleable when successful
func (m *RecoveryProposal) GetPrimaryRole()(AccessRuleable) {
    return m.primary_role
}
// GetRecoveryRole gets the recovery_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
// returns a AccessRuleable when successful
func (m *RecoveryProposal) GetRecoveryRole()(AccessRuleable) {
    return m.recovery_role
}
// GetTimedRecoveryDelayMinutes gets the timed_recovery_delay_minutes property value. An integer between `0` and `2^32 - 1`, specifying the optional proposal delay of timed recoveries.
// returns a *int64 when successful
func (m *RecoveryProposal) GetTimedRecoveryDelayMinutes()(*int64) {
    return m.timed_recovery_delay_minutes
}
// Serialize serializes information the current object
func (m *RecoveryProposal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("confirmation_role", m.GetConfirmationRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("primary_role", m.GetPrimaryRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recovery_role", m.GetRecoveryRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("timed_recovery_delay_minutes", m.GetTimedRecoveryDelayMinutes())
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
func (m *RecoveryProposal) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfirmationRole sets the confirmation_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
func (m *RecoveryProposal) SetConfirmationRole(value AccessRuleable)() {
    m.confirmation_role = value
}
// SetPrimaryRole sets the primary_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
func (m *RecoveryProposal) SetPrimaryRole(value AccessRuleable)() {
    m.primary_role = value
}
// SetRecoveryRole sets the recovery_role property value. A representation of an access rule in the Radix Engine.Note that some of the field and discriminator names use an older naming schema,for backwards compatibility.See the [advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules) docs for more information.
func (m *RecoveryProposal) SetRecoveryRole(value AccessRuleable)() {
    m.recovery_role = value
}
// SetTimedRecoveryDelayMinutes sets the timed_recovery_delay_minutes property value. An integer between `0` and `2^32 - 1`, specifying the optional proposal delay of timed recoveries.
func (m *RecoveryProposal) SetTimedRecoveryDelayMinutes(value *int64)() {
    m.timed_recovery_delay_minutes = value
}
type RecoveryProposalable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfirmationRole()(AccessRuleable)
    GetPrimaryRole()(AccessRuleable)
    GetRecoveryRole()(AccessRuleable)
    GetTimedRecoveryDelayMinutes()(*int64)
    SetConfirmationRole(value AccessRuleable)()
    SetPrimaryRole(value AccessRuleable)()
    SetRecoveryRole(value AccessRuleable)()
    SetTimedRecoveryDelayMinutes(value *int64)()
}
