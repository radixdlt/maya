package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsFilterAny struct {
    StreamProofsFilter
    // The from_state_version property
    from_state_version *int64
}
// NewStreamProofsFilterAny instantiates a new StreamProofsFilterAny and sets the default values.
func NewStreamProofsFilterAny()(*StreamProofsFilterAny) {
    m := &StreamProofsFilterAny{
        StreamProofsFilter: *NewStreamProofsFilter(),
    }
    return m
}
// CreateStreamProofsFilterAnyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsFilterAnyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamProofsFilterAny(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsFilterAny) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StreamProofsFilter.GetFieldDeserializers()
    res["from_state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromStateVersion(val)
        }
        return nil
    }
    return res
}
// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *StreamProofsFilterAny) GetFromStateVersion()(*int64) {
    return m.from_state_version
}
// Serialize serializes information the current object
func (m *StreamProofsFilterAny) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StreamProofsFilter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("from_state_version", m.GetFromStateVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *StreamProofsFilterAny) SetFromStateVersion(value *int64)() {
    m.from_state_version = value
}
type StreamProofsFilterAnyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StreamProofsFilterable
    GetFromStateVersion()(*int64)
    SetFromStateVersion(value *int64)()
}
