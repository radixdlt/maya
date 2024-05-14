package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerProof struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The ledger_header property
	ledger_header LedgerHeaderable
	// The origin property
	origin LedgerProofOriginable
}

// NewLedgerProof instantiates a new LedgerProof and sets the default values.
func NewLedgerProof() *LedgerProof {
	m := &LedgerProof{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLedgerProofFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerProofFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLedgerProof(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerProof) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerProof) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["ledger_header"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerHeaderFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLedgerHeader(val.(LedgerHeaderable))
		}
		return nil
	}
	res["origin"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerProofOriginFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrigin(val.(LedgerProofOriginable))
		}
		return nil
	}
	return res
}

// GetLedgerHeader gets the ledger_header property value. The ledger_header property
// returns a LedgerHeaderable when successful
func (m *LedgerProof) GetLedgerHeader() LedgerHeaderable {
	return m.ledger_header
}

// GetOrigin gets the origin property value. The origin property
// returns a LedgerProofOriginable when successful
func (m *LedgerProof) GetOrigin() LedgerProofOriginable {
	return m.origin
}

// Serialize serializes information the current object
func (m *LedgerProof) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("ledger_header", m.GetLedgerHeader())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("origin", m.GetOrigin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LedgerProof) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLedgerHeader sets the ledger_header property value. The ledger_header property
func (m *LedgerProof) SetLedgerHeader(value LedgerHeaderable) {
	m.ledger_header = value
}

// SetOrigin sets the origin property value. The origin property
func (m *LedgerProof) SetOrigin(value LedgerProofOriginable) {
	m.origin = value
}

type LedgerProofable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLedgerHeader() LedgerHeaderable
	GetOrigin() LedgerProofOriginable
	SetLedgerHeader(value LedgerHeaderable)
	SetOrigin(value LedgerProofOriginable)
}
