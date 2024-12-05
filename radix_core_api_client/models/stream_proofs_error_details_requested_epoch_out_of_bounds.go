package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsErrorDetailsRequestedEpochOutOfBounds struct {
    StreamProofsErrorDetails
    // The maximum completed epoch committed to this node's ledger.*Note on the bounds:* the requested `from_epoch` cannot be greater than`max_ledger_epoch + 1`. Any greater requested value triggers this error.
    max_ledger_epoch *int64
}
// NewStreamProofsErrorDetailsRequestedEpochOutOfBounds instantiates a new StreamProofsErrorDetailsRequestedEpochOutOfBounds and sets the default values.
func NewStreamProofsErrorDetailsRequestedEpochOutOfBounds()(*StreamProofsErrorDetailsRequestedEpochOutOfBounds) {
    m := &StreamProofsErrorDetailsRequestedEpochOutOfBounds{
        StreamProofsErrorDetails: *NewStreamProofsErrorDetails(),
    }
    return m
}
// CreateStreamProofsErrorDetailsRequestedEpochOutOfBoundsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsErrorDetailsRequestedEpochOutOfBoundsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamProofsErrorDetailsRequestedEpochOutOfBounds(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsErrorDetailsRequestedEpochOutOfBounds) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StreamProofsErrorDetails.GetFieldDeserializers()
    res["max_ledger_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLedgerEpoch(val)
        }
        return nil
    }
    return res
}
// GetMaxLedgerEpoch gets the max_ledger_epoch property value. The maximum completed epoch committed to this node's ledger.*Note on the bounds:* the requested `from_epoch` cannot be greater than`max_ledger_epoch + 1`. Any greater requested value triggers this error.
// returns a *int64 when successful
func (m *StreamProofsErrorDetailsRequestedEpochOutOfBounds) GetMaxLedgerEpoch()(*int64) {
    return m.max_ledger_epoch
}
// Serialize serializes information the current object
func (m *StreamProofsErrorDetailsRequestedEpochOutOfBounds) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StreamProofsErrorDetails.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("max_ledger_epoch", m.GetMaxLedgerEpoch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaxLedgerEpoch sets the max_ledger_epoch property value. The maximum completed epoch committed to this node's ledger.*Note on the bounds:* the requested `from_epoch` cannot be greater than`max_ledger_epoch + 1`. Any greater requested value triggers this error.
func (m *StreamProofsErrorDetailsRequestedEpochOutOfBounds) SetMaxLedgerEpoch(value *int64)() {
    m.max_ledger_epoch = value
}
type StreamProofsErrorDetailsRequestedEpochOutOfBoundsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StreamProofsErrorDetailsable
    GetMaxLedgerEpoch()(*int64)
    SetMaxLedgerEpoch(value *int64)()
}
