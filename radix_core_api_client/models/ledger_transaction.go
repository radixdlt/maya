package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerTransaction struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded full ledger transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
	payload_hex *string
	// The type of the ledger transaction
	typeEscaped *LedgerTransactionType
}

// NewLedgerTransaction instantiates a new LedgerTransaction and sets the default values.
func NewLedgerTransaction() *LedgerTransaction {
	m := &LedgerTransaction{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLedgerTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("type")
		if err != nil {
			return nil, err
		}
		if mappingValueNode != nil {
			mappingValue, err := mappingValueNode.GetStringValue()
			if err != nil {
				return nil, err
			}
			if mappingValue != nil {
				switch *mappingValue {
				case "Flash":
					return NewFlashLedgerTransaction(), nil
				case "Genesis":
					return NewGenesisLedgerTransaction(), nil
				case "RoundUpdate":
					return NewRoundUpdateLedgerTransaction(), nil
				case "User":
					return NewUserLedgerTransaction(), nil
				}
			}
		}
	}
	return NewLedgerTransaction(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerTransaction) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["payload_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPayloadHex(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseLedgerTransactionType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*LedgerTransactionType))
		}
		return nil
	}
	return res
}

// GetPayloadHex gets the payload_hex property value. The hex-encoded full ledger transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
// returns a *string when successful
func (m *LedgerTransaction) GetPayloadHex() *string {
	return m.payload_hex
}

// GetTypeEscaped gets the type property value. The type of the ledger transaction
// returns a *LedgerTransactionType when successful
func (m *LedgerTransaction) GetTypeEscaped() *LedgerTransactionType {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *LedgerTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("payload_hex", m.GetPayloadHex())
		if err != nil {
			return err
		}
	}
	if m.GetTypeEscaped() != nil {
		cast := (*m.GetTypeEscaped()).String()
		err := writer.WriteStringValue("type", &cast)
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
func (m *LedgerTransaction) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetPayloadHex sets the payload_hex property value. The hex-encoded full ledger transaction payload. Only returned if enabled in TransactionFormatOptions on your request.
func (m *LedgerTransaction) SetPayloadHex(value *string) {
	m.payload_hex = value
}

// SetTypeEscaped sets the type property value. The type of the ledger transaction
func (m *LedgerTransaction) SetTypeEscaped(value *LedgerTransactionType) {
	m.typeEscaped = value
}

type LedgerTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetPayloadHex() *string
	GetTypeEscaped() *LedgerTransactionType
	SetPayloadHex(value *string)
	SetTypeEscaped(value *LedgerTransactionType)
}
