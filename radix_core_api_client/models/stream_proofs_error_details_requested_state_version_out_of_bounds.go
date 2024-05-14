package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsErrorDetailsRequestedStateVersionOutOfBounds struct {
	StreamProofsErrorDetails
	// The max_ledger_state_version property
	max_ledger_state_version *int64
}

// NewStreamProofsErrorDetailsRequestedStateVersionOutOfBounds instantiates a new StreamProofsErrorDetailsRequestedStateVersionOutOfBounds and sets the default values.
func NewStreamProofsErrorDetailsRequestedStateVersionOutOfBounds() *StreamProofsErrorDetailsRequestedStateVersionOutOfBounds {
	m := &StreamProofsErrorDetailsRequestedStateVersionOutOfBounds{
		StreamProofsErrorDetails: *NewStreamProofsErrorDetails(),
	}
	return m
}

// CreateStreamProofsErrorDetailsRequestedStateVersionOutOfBoundsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsErrorDetailsRequestedStateVersionOutOfBoundsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStreamProofsErrorDetailsRequestedStateVersionOutOfBounds(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsErrorDetailsRequestedStateVersionOutOfBounds) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.StreamProofsErrorDetails.GetFieldDeserializers()
	res["max_ledger_state_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxLedgerStateVersion(val)
		}
		return nil
	}
	return res
}

// GetMaxLedgerStateVersion gets the max_ledger_state_version property value. The max_ledger_state_version property
// returns a *int64 when successful
func (m *StreamProofsErrorDetailsRequestedStateVersionOutOfBounds) GetMaxLedgerStateVersion() *int64 {
	return m.max_ledger_state_version
}

// Serialize serializes information the current object
func (m *StreamProofsErrorDetailsRequestedStateVersionOutOfBounds) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.StreamProofsErrorDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteInt64Value("max_ledger_state_version", m.GetMaxLedgerStateVersion())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetMaxLedgerStateVersion sets the max_ledger_state_version property value. The max_ledger_state_version property
func (m *StreamProofsErrorDetailsRequestedStateVersionOutOfBounds) SetMaxLedgerStateVersion(value *int64) {
	m.max_ledger_state_version = value
}

type StreamProofsErrorDetailsRequestedStateVersionOutOfBoundsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	StreamProofsErrorDetailsable
	GetMaxLedgerStateVersion() *int64
	SetMaxLedgerStateVersion(value *int64)
}
