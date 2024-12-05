package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InProgressProtocolUpdateStatus struct {
    ProtocolUpdateStatus
    // The latest_commit property
    latest_commit ProtocolUpdateStatusLatestCommitable
}
// NewInProgressProtocolUpdateStatus instantiates a new InProgressProtocolUpdateStatus and sets the default values.
func NewInProgressProtocolUpdateStatus()(*InProgressProtocolUpdateStatus) {
    m := &InProgressProtocolUpdateStatus{
        ProtocolUpdateStatus: *NewProtocolUpdateStatus(),
    }
    return m
}
// CreateInProgressProtocolUpdateStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInProgressProtocolUpdateStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInProgressProtocolUpdateStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InProgressProtocolUpdateStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtocolUpdateStatus.GetFieldDeserializers()
    res["latest_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtocolUpdateStatusLatestCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestCommit(val.(ProtocolUpdateStatusLatestCommitable))
        }
        return nil
    }
    return res
}
// GetLatestCommit gets the latest_commit property value. The latest_commit property
// returns a ProtocolUpdateStatusLatestCommitable when successful
func (m *InProgressProtocolUpdateStatus) GetLatestCommit()(ProtocolUpdateStatusLatestCommitable) {
    return m.latest_commit
}
// Serialize serializes information the current object
func (m *InProgressProtocolUpdateStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtocolUpdateStatus.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("latest_commit", m.GetLatestCommit())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLatestCommit sets the latest_commit property value. The latest_commit property
func (m *InProgressProtocolUpdateStatus) SetLatestCommit(value ProtocolUpdateStatusLatestCommitable)() {
    m.latest_commit = value
}
type InProgressProtocolUpdateStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtocolUpdateStatusable
    GetLatestCommit()(ProtocolUpdateStatusLatestCommitable)
    SetLatestCommit(value ProtocolUpdateStatusLatestCommitable)()
}
