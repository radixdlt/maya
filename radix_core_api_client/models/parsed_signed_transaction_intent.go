package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedSignedTransactionIntent struct {
	ParsedTransaction
	// The identifiers property
	identifiers ParsedSignedTransactionIntentIdentifiersable
	// The signed_intent property
	signed_intent SignedTransactionIntentable
}

// NewParsedSignedTransactionIntent instantiates a new ParsedSignedTransactionIntent and sets the default values.
func NewParsedSignedTransactionIntent() *ParsedSignedTransactionIntent {
	m := &ParsedSignedTransactionIntent{
		ParsedTransaction: *NewParsedTransaction(),
	}
	return m
}

// CreateParsedSignedTransactionIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedSignedTransactionIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewParsedSignedTransactionIntent(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedSignedTransactionIntent) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ParsedTransaction.GetFieldDeserializers()
	res["identifiers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateParsedSignedTransactionIntentIdentifiersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIdentifiers(val.(ParsedSignedTransactionIntentIdentifiersable))
		}
		return nil
	}
	res["signed_intent"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSignedTransactionIntentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSignedIntent(val.(SignedTransactionIntentable))
		}
		return nil
	}
	return res
}

// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a ParsedSignedTransactionIntentIdentifiersable when successful
func (m *ParsedSignedTransactionIntent) GetIdentifiers() ParsedSignedTransactionIntentIdentifiersable {
	return m.identifiers
}

// GetSignedIntent gets the signed_intent property value. The signed_intent property
// returns a SignedTransactionIntentable when successful
func (m *ParsedSignedTransactionIntent) GetSignedIntent() SignedTransactionIntentable {
	return m.signed_intent
}

// Serialize serializes information the current object
func (m *ParsedSignedTransactionIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ParsedTransaction.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("identifiers", m.GetIdentifiers())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("signed_intent", m.GetSignedIntent())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *ParsedSignedTransactionIntent) SetIdentifiers(value ParsedSignedTransactionIntentIdentifiersable) {
	m.identifiers = value
}

// SetSignedIntent sets the signed_intent property value. The signed_intent property
func (m *ParsedSignedTransactionIntent) SetSignedIntent(value SignedTransactionIntentable) {
	m.signed_intent = value
}

type ParsedSignedTransactionIntentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ParsedTransactionable
	GetIdentifiers() ParsedSignedTransactionIntentIdentifiersable
	GetSignedIntent() SignedTransactionIntentable
	SetIdentifiers(value ParsedSignedTransactionIntentIdentifiersable)
	SetSignedIntent(value SignedTransactionIntentable)
}
