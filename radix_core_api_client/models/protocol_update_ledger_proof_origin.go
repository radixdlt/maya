package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtocolUpdateLedgerProofOrigin struct {
	LedgerProofOrigin
	// The batch_idx property
	batch_idx *int64
	// The protocol_version_name property
	protocol_version_name *string
}

// NewProtocolUpdateLedgerProofOrigin instantiates a new ProtocolUpdateLedgerProofOrigin and sets the default values.
func NewProtocolUpdateLedgerProofOrigin() *ProtocolUpdateLedgerProofOrigin {
	m := &ProtocolUpdateLedgerProofOrigin{
		LedgerProofOrigin: *NewLedgerProofOrigin(),
	}
	return m
}

// CreateProtocolUpdateLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtocolUpdateLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProtocolUpdateLedgerProofOrigin(), nil
}

// GetBatchIdx gets the batch_idx property value. The batch_idx property
// returns a *int64 when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetBatchIdx() *int64 {
	return m.batch_idx
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerProofOrigin.GetFieldDeserializers()
	res["batch_idx"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBatchIdx(val)
		}
		return nil
	}
	res["protocol_version_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProtocolVersionName(val)
		}
		return nil
	}
	return res
}

// GetProtocolVersionName gets the protocol_version_name property value. The protocol_version_name property
// returns a *string when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetProtocolVersionName() *string {
	return m.protocol_version_name
}

// Serialize serializes information the current object
func (m *ProtocolUpdateLedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerProofOrigin.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteInt64Value("batch_idx", m.GetBatchIdx())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("protocol_version_name", m.GetProtocolVersionName())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBatchIdx sets the batch_idx property value. The batch_idx property
func (m *ProtocolUpdateLedgerProofOrigin) SetBatchIdx(value *int64) {
	m.batch_idx = value
}

// SetProtocolVersionName sets the protocol_version_name property value. The protocol_version_name property
func (m *ProtocolUpdateLedgerProofOrigin) SetProtocolVersionName(value *string) {
	m.protocol_version_name = value
}

type ProtocolUpdateLedgerProofOriginable interface {
	LedgerProofOriginable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBatchIdx() *int64
	GetProtocolVersionName() *string
	SetBatchIdx(value *int64)
	SetProtocolVersionName(value *string)
}
