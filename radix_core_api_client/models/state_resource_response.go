package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateResourceResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The at_ledger_state property
    at_ledger_state LedgerStateSummaryable
    // The manager property
    manager StateResourceManagerable
    // The owner_role property
    owner_role Substateable
}
// NewStateResourceResponse instantiates a new StateResourceResponse and sets the default values.
func NewStateResourceResponse()(*StateResourceResponse) {
    m := &StateResourceResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateResourceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateResourceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateResourceResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateResourceResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *StateResourceResponse) GetAtLedgerState()(LedgerStateSummaryable) {
    return m.at_ledger_state
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateResourceResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["at_ledger_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtLedgerState(val.(LedgerStateSummaryable))
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStateResourceManagerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(StateResourceManagerable))
        }
        return nil
    }
    res["owner_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerRole(val.(Substateable))
        }
        return nil
    }
    return res
}
// GetManager gets the manager property value. The manager property
// returns a StateResourceManagerable when successful
func (m *StateResourceResponse) GetManager()(StateResourceManagerable) {
    return m.manager
}
// GetOwnerRole gets the owner_role property value. The owner_role property
// returns a Substateable when successful
func (m *StateResourceResponse) GetOwnerRole()(Substateable) {
    return m.owner_role
}
// Serialize serializes information the current object
func (m *StateResourceResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner_role", m.GetOwnerRole())
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
func (m *StateResourceResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *StateResourceResponse) SetAtLedgerState(value LedgerStateSummaryable)() {
    m.at_ledger_state = value
}
// SetManager sets the manager property value. The manager property
func (m *StateResourceResponse) SetManager(value StateResourceManagerable)() {
    m.manager = value
}
// SetOwnerRole sets the owner_role property value. The owner_role property
func (m *StateResourceResponse) SetOwnerRole(value Substateable)() {
    m.owner_role = value
}
type StateResourceResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSummaryable)
    GetManager()(StateResourceManagerable)
    GetOwnerRole()(Substateable)
    SetAtLedgerState(value LedgerStateSummaryable)()
    SetManager(value StateResourceManagerable)()
    SetOwnerRole(value Substateable)()
}
