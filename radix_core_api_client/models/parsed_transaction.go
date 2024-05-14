package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedTransaction struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The type of the parsed (partial) transaction payload
	typeEscaped *ParsedTransactionType
}

// NewParsedTransaction instantiates a new ParsedTransaction and sets the default values.
func NewParsedTransaction() *ParsedTransaction {
	m := &ParsedTransaction{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateParsedTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
				case "LedgerTransaction":
					return NewParsedLedgerTransaction(), nil
				case "NotarizedTransaction":
					return NewParsedNotarizedTransaction(), nil
				case "SignedTransactionIntent":
					return NewParsedSignedTransactionIntent(), nil
				case "TransactionIntent":
					return NewParsedTransactionIntent(), nil
				}
			}
		}
	}
	return NewParsedTransaction(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ParsedTransaction) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedTransaction) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseParsedTransactionType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*ParsedTransactionType))
		}
		return nil
	}
	return res
}

// GetTypeEscaped gets the type property value. The type of the parsed (partial) transaction payload
// returns a *ParsedTransactionType when successful
func (m *ParsedTransaction) GetTypeEscaped() *ParsedTransactionType {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *ParsedTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ParsedTransaction) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetTypeEscaped sets the type property value. The type of the parsed (partial) transaction payload
func (m *ParsedTransaction) SetTypeEscaped(value *ParsedTransactionType) {
	m.typeEscaped = value
}

type ParsedTransactionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetTypeEscaped() *ParsedTransactionType
	SetTypeEscaped(value *ParsedTransactionType)
}
