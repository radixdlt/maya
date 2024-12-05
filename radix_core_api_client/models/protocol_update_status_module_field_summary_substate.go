package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtocolUpdateStatusModuleFieldSummarySubstate struct {
    Substate
    // The protocol_version property
    protocol_version *string
    // The update_status property
    update_status ProtocolUpdateStatusable
}
// NewProtocolUpdateStatusModuleFieldSummarySubstate instantiates a new ProtocolUpdateStatusModuleFieldSummarySubstate and sets the default values.
func NewProtocolUpdateStatusModuleFieldSummarySubstate()(*ProtocolUpdateStatusModuleFieldSummarySubstate) {
    m := &ProtocolUpdateStatusModuleFieldSummarySubstate{
        Substate: *NewSubstate(),
    }
    return m
}
// CreateProtocolUpdateStatusModuleFieldSummarySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtocolUpdateStatusModuleFieldSummarySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtocolUpdateStatusModuleFieldSummarySubstate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Substate.GetFieldDeserializers()
    res["protocol_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocolVersion(val)
        }
        return nil
    }
    res["update_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtocolUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateStatus(val.(ProtocolUpdateStatusable))
        }
        return nil
    }
    return res
}
// GetProtocolVersion gets the protocol_version property value. The protocol_version property
// returns a *string when successful
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) GetProtocolVersion()(*string) {
    return m.protocol_version
}
// GetUpdateStatus gets the update_status property value. The update_status property
// returns a ProtocolUpdateStatusable when successful
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) GetUpdateStatus()(ProtocolUpdateStatusable) {
    return m.update_status
}
// Serialize serializes information the current object
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Substate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("protocol_version", m.GetProtocolVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("update_status", m.GetUpdateStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProtocolVersion sets the protocol_version property value. The protocol_version property
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) SetProtocolVersion(value *string)() {
    m.protocol_version = value
}
// SetUpdateStatus sets the update_status property value. The update_status property
func (m *ProtocolUpdateStatusModuleFieldSummarySubstate) SetUpdateStatus(value ProtocolUpdateStatusable)() {
    m.update_status = value
}
type ProtocolUpdateStatusModuleFieldSummarySubstateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Substateable
    GetProtocolVersion()(*string)
    GetUpdateStatus()(ProtocolUpdateStatusable)
    SetProtocolVersion(value *string)()
    SetUpdateStatus(value ProtocolUpdateStatusable)()
}
