package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsFilterNewEpochs struct {
    StreamProofsFilter
    // The first proof to be returned should be the proof starting this epoch. If empty, it starts from the first epoch proof after genesis. The network status endpoint can be used to find the current epoch.
    from_epoch *int64
}
// NewStreamProofsFilterNewEpochs instantiates a new StreamProofsFilterNewEpochs and sets the default values.
func NewStreamProofsFilterNewEpochs()(*StreamProofsFilterNewEpochs) {
    m := &StreamProofsFilterNewEpochs{
        StreamProofsFilter: *NewStreamProofsFilter(),
    }
    return m
}
// CreateStreamProofsFilterNewEpochsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsFilterNewEpochsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamProofsFilterNewEpochs(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsFilterNewEpochs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StreamProofsFilter.GetFieldDeserializers()
    res["from_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromEpoch(val)
        }
        return nil
    }
    return res
}
// GetFromEpoch gets the from_epoch property value. The first proof to be returned should be the proof starting this epoch. If empty, it starts from the first epoch proof after genesis. The network status endpoint can be used to find the current epoch.
// returns a *int64 when successful
func (m *StreamProofsFilterNewEpochs) GetFromEpoch()(*int64) {
    return m.from_epoch
}
// Serialize serializes information the current object
func (m *StreamProofsFilterNewEpochs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StreamProofsFilter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("from_epoch", m.GetFromEpoch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromEpoch sets the from_epoch property value. The first proof to be returned should be the proof starting this epoch. If empty, it starts from the first epoch proof after genesis. The network status endpoint can be used to find the current epoch.
func (m *StreamProofsFilterNewEpochs) SetFromEpoch(value *int64)() {
    m.from_epoch = value
}
type StreamProofsFilterNewEpochsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StreamProofsFilterable
    GetFromEpoch()(*int64)
    SetFromEpoch(value *int64)()
}
