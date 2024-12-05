package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompleteProtocolUpdateStatus struct {
    ProtocolUpdateStatus
}
// NewCompleteProtocolUpdateStatus instantiates a new CompleteProtocolUpdateStatus and sets the default values.
func NewCompleteProtocolUpdateStatus()(*CompleteProtocolUpdateStatus) {
    m := &CompleteProtocolUpdateStatus{
        ProtocolUpdateStatus: *NewProtocolUpdateStatus(),
    }
    return m
}
// CreateCompleteProtocolUpdateStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompleteProtocolUpdateStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompleteProtocolUpdateStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompleteProtocolUpdateStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtocolUpdateStatus.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CompleteProtocolUpdateStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtocolUpdateStatus.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type CompleteProtocolUpdateStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtocolUpdateStatusable
}
