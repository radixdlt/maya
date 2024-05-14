package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestedStateVersionOutOfBoundsErrorDetails struct {
	StreamTransactionsErrorDetails
	// The max_ledger_state_version property
	max_ledger_state_version *int64
}

// NewRequestedStateVersionOutOfBoundsErrorDetails instantiates a new RequestedStateVersionOutOfBoundsErrorDetails and sets the default values.
func NewRequestedStateVersionOutOfBoundsErrorDetails() *RequestedStateVersionOutOfBoundsErrorDetails {
	m := &RequestedStateVersionOutOfBoundsErrorDetails{
		StreamTransactionsErrorDetails: *NewStreamTransactionsErrorDetails(),
	}
	return m
}

// CreateRequestedStateVersionOutOfBoundsErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestedStateVersionOutOfBoundsErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRequestedStateVersionOutOfBoundsErrorDetails(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestedStateVersionOutOfBoundsErrorDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.StreamTransactionsErrorDetails.GetFieldDeserializers()
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
func (m *RequestedStateVersionOutOfBoundsErrorDetails) GetMaxLedgerStateVersion() *int64 {
	return m.max_ledger_state_version
}

// Serialize serializes information the current object
func (m *RequestedStateVersionOutOfBoundsErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.StreamTransactionsErrorDetails.Serialize(writer)
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
func (m *RequestedStateVersionOutOfBoundsErrorDetails) SetMaxLedgerStateVersion(value *int64) {
	m.max_ledger_state_version = value
}

type RequestedStateVersionOutOfBoundsErrorDetailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	StreamTransactionsErrorDetailsable
	GetMaxLedgerStateVersion() *int64
	SetMaxLedgerStateVersion(value *int64)
}
