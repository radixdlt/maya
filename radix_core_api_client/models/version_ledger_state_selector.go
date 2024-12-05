package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VersionLedgerStateSelector struct {
    LedgerStateSelector
    // The state_version property
    state_version *int64
}
// NewVersionLedgerStateSelector instantiates a new VersionLedgerStateSelector and sets the default values.
func NewVersionLedgerStateSelector()(*VersionLedgerStateSelector) {
    m := &VersionLedgerStateSelector{
        LedgerStateSelector: *NewLedgerStateSelector(),
    }
    return m
}
// CreateVersionLedgerStateSelectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVersionLedgerStateSelectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVersionLedgerStateSelector(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VersionLedgerStateSelector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LedgerStateSelector.GetFieldDeserializers()
    res["state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateVersion(val)
        }
        return nil
    }
    return res
}
// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *VersionLedgerStateSelector) GetStateVersion()(*int64) {
    return m.state_version
}
// Serialize serializes information the current object
func (m *VersionLedgerStateSelector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LedgerStateSelector.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("state_version", m.GetStateVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStateVersion sets the state_version property value. The state_version property
func (m *VersionLedgerStateSelector) SetStateVersion(value *int64)() {
    m.state_version = value
}
type VersionLedgerStateSelectorable interface {
    LedgerStateSelectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStateVersion()(*int64)
    SetStateVersion(value *int64)()
}
