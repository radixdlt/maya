package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenesisLedgerProofOrigin struct {
	LedgerProofOrigin
	// The genesis_opaque_hash property
	genesis_opaque_hash *string
}

// NewGenesisLedgerProofOrigin instantiates a new GenesisLedgerProofOrigin and sets the default values.
func NewGenesisLedgerProofOrigin() *GenesisLedgerProofOrigin {
	m := &GenesisLedgerProofOrigin{
		LedgerProofOrigin: *NewLedgerProofOrigin(),
	}
	return m
}

// CreateGenesisLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenesisLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenesisLedgerProofOrigin(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenesisLedgerProofOrigin) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LedgerProofOrigin.GetFieldDeserializers()
	res["genesis_opaque_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGenesisOpaqueHash(val)
		}
		return nil
	}
	return res
}

// GetGenesisOpaqueHash gets the genesis_opaque_hash property value. The genesis_opaque_hash property
// returns a *string when successful
func (m *GenesisLedgerProofOrigin) GetGenesisOpaqueHash() *string {
	return m.genesis_opaque_hash
}

// Serialize serializes information the current object
func (m *GenesisLedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LedgerProofOrigin.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("genesis_opaque_hash", m.GetGenesisOpaqueHash())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetGenesisOpaqueHash sets the genesis_opaque_hash property value. The genesis_opaque_hash property
func (m *GenesisLedgerProofOrigin) SetGenesisOpaqueHash(value *string) {
	m.genesis_opaque_hash = value
}

type GenesisLedgerProofOriginable interface {
	LedgerProofOriginable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetGenesisOpaqueHash() *string
	SetGenesisOpaqueHash(value *string)
}
