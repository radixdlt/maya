package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedNotarizedTransaction struct {
    ParsedTransaction
    // The identifiers property
    identifiers ParsedNotarizedTransactionIdentifiersable
    // The notarized_transaction property
    notarized_transaction NotarizedTransactionable
    // If the transaction is known to not be valid, this gives a reason.Different levels of validation are performed, dependent on the validation mode.Note that, even if validation mode is Static or Full, the transaction maystill be rejected or fail due to issues at runtime (e.g. if the loan cannot be repaid).
    validation_error ParsedNotarizedTransaction_validation_errorable
}
// NewParsedNotarizedTransaction instantiates a new ParsedNotarizedTransaction and sets the default values.
func NewParsedNotarizedTransaction()(*ParsedNotarizedTransaction) {
    m := &ParsedNotarizedTransaction{
        ParsedTransaction: *NewParsedTransaction(),
    }
    return m
}
// CreateParsedNotarizedTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedNotarizedTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParsedNotarizedTransaction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedNotarizedTransaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ParsedTransaction.GetFieldDeserializers()
    res["identifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParsedNotarizedTransactionIdentifiersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifiers(val.(ParsedNotarizedTransactionIdentifiersable))
        }
        return nil
    }
    res["notarized_transaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotarizedTransactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotarizedTransaction(val.(NotarizedTransactionable))
        }
        return nil
    }
    res["validation_error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParsedNotarizedTransaction_validation_errorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationError(val.(ParsedNotarizedTransaction_validation_errorable))
        }
        return nil
    }
    return res
}
// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a ParsedNotarizedTransactionIdentifiersable when successful
func (m *ParsedNotarizedTransaction) GetIdentifiers()(ParsedNotarizedTransactionIdentifiersable) {
    return m.identifiers
}
// GetNotarizedTransaction gets the notarized_transaction property value. The notarized_transaction property
// returns a NotarizedTransactionable when successful
func (m *ParsedNotarizedTransaction) GetNotarizedTransaction()(NotarizedTransactionable) {
    return m.notarized_transaction
}
// GetValidationError gets the validation_error property value. If the transaction is known to not be valid, this gives a reason.Different levels of validation are performed, dependent on the validation mode.Note that, even if validation mode is Static or Full, the transaction maystill be rejected or fail due to issues at runtime (e.g. if the loan cannot be repaid).
// returns a ParsedNotarizedTransaction_validation_errorable when successful
func (m *ParsedNotarizedTransaction) GetValidationError()(ParsedNotarizedTransaction_validation_errorable) {
    return m.validation_error
}
// Serialize serializes information the current object
func (m *ParsedNotarizedTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("notarized_transaction", m.GetNotarizedTransaction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("validation_error", m.GetValidationError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *ParsedNotarizedTransaction) SetIdentifiers(value ParsedNotarizedTransactionIdentifiersable)() {
    m.identifiers = value
}
// SetNotarizedTransaction sets the notarized_transaction property value. The notarized_transaction property
func (m *ParsedNotarizedTransaction) SetNotarizedTransaction(value NotarizedTransactionable)() {
    m.notarized_transaction = value
}
// SetValidationError sets the validation_error property value. If the transaction is known to not be valid, this gives a reason.Different levels of validation are performed, dependent on the validation mode.Note that, even if validation mode is Static or Full, the transaction maystill be rejected or fail due to issues at runtime (e.g. if the loan cannot be repaid).
func (m *ParsedNotarizedTransaction) SetValidationError(value ParsedNotarizedTransaction_validation_errorable)() {
    m.validation_error = value
}
type ParsedNotarizedTransactionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ParsedTransactionable
    GetIdentifiers()(ParsedNotarizedTransactionIdentifiersable)
    GetNotarizedTransaction()(NotarizedTransactionable)
    GetValidationError()(ParsedNotarizedTransaction_validation_errorable)
    SetIdentifiers(value ParsedNotarizedTransactionIdentifiersable)()
    SetNotarizedTransaction(value NotarizedTransactionable)()
    SetValidationError(value ParsedNotarizedTransaction_validation_errorable)()
}
