package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsFilterProtocolUpdateExecution struct {
    StreamProofsFilter
    // The from_state_version property
    from_state_version *int64
    // The protocol version name to filter to.
    protocol_version *string
}
// NewStreamProofsFilterProtocolUpdateExecution instantiates a new StreamProofsFilterProtocolUpdateExecution and sets the default values.
func NewStreamProofsFilterProtocolUpdateExecution()(*StreamProofsFilterProtocolUpdateExecution) {
    m := &StreamProofsFilterProtocolUpdateExecution{
        StreamProofsFilter: *NewStreamProofsFilter(),
    }
    return m
}
// CreateStreamProofsFilterProtocolUpdateExecutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsFilterProtocolUpdateExecutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamProofsFilterProtocolUpdateExecution(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsFilterProtocolUpdateExecution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *StreamProofsFilterProtocolUpdateExecution) GetFromStateVersion()(*int64) {
    return m.from_state_version
}
// GetProtocolVersion gets the protocol_version property value. The protocol version name to filter to.
// returns a *string when successful
func (m *StreamProofsFilterProtocolUpdateExecution) GetProtocolVersion()(*string) {
    return m.protocol_version
}
// Serialize serializes information the current object
func (m *StreamProofsFilterProtocolUpdateExecution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("protocol_version", m.GetProtocolVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *StreamProofsFilterProtocolUpdateExecution) SetFromStateVersion(value *int64)() {
    m.from_state_version = value
}
// SetProtocolVersion sets the protocol_version property value. The protocol version name to filter to.
func (m *StreamProofsFilterProtocolUpdateExecution) SetProtocolVersion(value *string)() {
    m.protocol_version = value
}
type StreamProofsFilterProtocolUpdateExecutionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StreamProofsFilterable
    GetFromStateVersion()(*int64)
    GetProtocolVersion()(*string)
    SetFromStateVersion(value *int64)()
    SetProtocolVersion(value *string)()
}
