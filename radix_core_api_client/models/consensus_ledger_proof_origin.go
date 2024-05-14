package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusLedgerProofOrigin struct {
	LedgerProofOrigin
	// A hex-encoded 32-byte vertex VoteData hash on the consensus side, opaque to ledger.
	opaque_hash *string
	// The timestamped_signatures property
	timestamped_signatures []TimestampedValidatorSignatureable
}

// NewConsensusLedgerProofOrigin instantiates a new ConsensusLedgerProofOrigin and sets the default values.
func NewConsensusLedgerProofOrigin() *ConsensusLedgerProofOrigin {
	m := &ConsensusLedgerProofOrigin{
		LedgerProofOrigin: *NewLedgerProofOrigin(),
	}
	return m
}

// CreateConsensusLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewConsensusLedgerProofOrigin(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusLedgerProofOrigin) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerProofOrigin.GetFieldDeserializers()
	res["opaque_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpaqueHash(val)
		}
		return nil
	}
	res["timestamped_signatures"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTimestampedValidatorSignatureFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TimestampedValidatorSignatureable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TimestampedValidatorSignatureable)
				}
			}
			m.SetTimestampedSignatures(res)
		}
		return nil
	}
	return res
}

// GetOpaqueHash gets the opaque_hash property value. A hex-encoded 32-byte vertex VoteData hash on the consensus side, opaque to ledger.
// returns a *string when successful
func (m *ConsensusLedgerProofOrigin) GetOpaqueHash() *string {
	return m.opaque_hash
}

// GetTimestampedSignatures gets the timestamped_signatures property value. The timestamped_signatures property
// returns a []TimestampedValidatorSignatureable when successful
func (m *ConsensusLedgerProofOrigin) GetTimestampedSignatures() []TimestampedValidatorSignatureable {
	return m.timestamped_signatures
}

// Serialize serializes information the current object
func (m *ConsensusLedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerProofOrigin.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("opaque_hash", m.GetOpaqueHash())
		if err != nil {
			return err
		}
	}
	if m.GetTimestampedSignatures() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimestampedSignatures()))
		for i, v := range m.GetTimestampedSignatures() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("timestamped_signatures", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetOpaqueHash sets the opaque_hash property value. A hex-encoded 32-byte vertex VoteData hash on the consensus side, opaque to ledger.
func (m *ConsensusLedgerProofOrigin) SetOpaqueHash(value *string) {
	m.opaque_hash = value
}

// SetTimestampedSignatures sets the timestamped_signatures property value. The timestamped_signatures property
func (m *ConsensusLedgerProofOrigin) SetTimestampedSignatures(value []TimestampedValidatorSignatureable) {
	m.timestamped_signatures = value
}

type ConsensusLedgerProofOriginable interface {
	LedgerProofOriginable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetOpaqueHash() *string
	GetTimestampedSignatures() []TimestampedValidatorSignatureable
	SetOpaqueHash(value *string)
	SetTimestampedSignatures(value []TimestampedValidatorSignatureable)
}
