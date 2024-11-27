package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedTransactionIntent struct {
    ParsedTransaction
    // The identifiers property
    identifiers ParsedTransactionIntentIdentifiersable
    // The intent property
    intent TransactionIntentable
}
// NewParsedTransactionIntent instantiates a new ParsedTransactionIntent and sets the default values.
func NewParsedTransactionIntent()(*ParsedTransactionIntent) {
    m := &ParsedTransactionIntent{
        ParsedTransaction: *NewParsedTransaction(),
    }
    return m
}
// CreateParsedTransactionIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedTransactionIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParsedTransactionIntent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedTransactionIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ParsedTransaction.GetFieldDeserializers()
    res["identifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParsedTransactionIntentIdentifiersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifiers(val.(ParsedTransactionIntentIdentifiersable))
        }
        return nil
    }
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(TransactionIntentable))
        }
        return nil
    }
    return res
}
// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a ParsedTransactionIntentIdentifiersable when successful
func (m *ParsedTransactionIntent) GetIdentifiers()(ParsedTransactionIntentIdentifiersable) {
    return m.identifiers
}
// GetIntent gets the intent property value. The intent property
// returns a TransactionIntentable when successful
func (m *ParsedTransactionIntent) GetIntent()(TransactionIntentable) {
    return m.intent
}
// Serialize serializes information the current object
func (m *ParsedTransactionIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("intent", m.GetIntent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *ParsedTransactionIntent) SetIdentifiers(value ParsedTransactionIntentIdentifiersable)() {
    m.identifiers = value
}
// SetIntent sets the intent property value. The intent property
func (m *ParsedTransactionIntent) SetIntent(value TransactionIntentable)() {
    m.intent = value
}
type ParsedTransactionIntentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ParsedTransactionable
    GetIdentifiers()(ParsedTransactionIntentIdentifiersable)
    GetIntent()(TransactionIntentable)
    SetIdentifiers(value ParsedTransactionIntentIdentifiersable)()
    SetIntent(value TransactionIntentable)()
}
