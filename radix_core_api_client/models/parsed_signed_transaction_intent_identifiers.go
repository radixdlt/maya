package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParsedSignedTransactionIntentIdentifiers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
    intent_hash *string
    // The Bech32m-encoded human readable `TransactionIntentHash`.
    intent_hash_bech32m *string
    // The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
    signed_intent_hash *string
    // The Bech32m-encoded human readable `SignedTransactionIntentHash`.
    signed_intent_hash_bech32m *string
}
// NewParsedSignedTransactionIntentIdentifiers instantiates a new ParsedSignedTransactionIntentIdentifiers and sets the default values.
func NewParsedSignedTransactionIntentIdentifiers()(*ParsedSignedTransactionIntentIdentifiers) {
    m := &ParsedSignedTransactionIntentIdentifiers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateParsedSignedTransactionIntentIdentifiersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedSignedTransactionIntentIdentifiersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParsedSignedTransactionIntentIdentifiers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["intent_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentHash(val)
        }
        return nil
    }
    res["intent_hash_bech32m"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentHashBech32m(val)
        }
        return nil
    }
    res["signed_intent_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedIntentHash(val)
        }
        return nil
    }
    res["signed_intent_hash_bech32m"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedIntentHashBech32m(val)
        }
        return nil
    }
    return res
}
// GetIntentHash gets the intent_hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
// returns a *string when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetIntentHash()(*string) {
    return m.intent_hash
}
// GetIntentHashBech32m gets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
// returns a *string when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetIntentHashBech32m()(*string) {
    return m.intent_hash_bech32m
}
// GetSignedIntentHash gets the signed_intent_hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
// returns a *string when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetSignedIntentHash()(*string) {
    return m.signed_intent_hash
}
// GetSignedIntentHashBech32m gets the signed_intent_hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
// returns a *string when successful
func (m *ParsedSignedTransactionIntentIdentifiers) GetSignedIntentHashBech32m()(*string) {
    return m.signed_intent_hash_bech32m
}
// Serialize serializes information the current object
func (m *ParsedSignedTransactionIntentIdentifiers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("intent_hash", m.GetIntentHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("intent_hash_bech32m", m.GetIntentHashBech32m())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signed_intent_hash", m.GetSignedIntentHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signed_intent_hash_bech32m", m.GetSignedIntentHashBech32m())
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
func (m *ParsedSignedTransactionIntentIdentifiers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntentHash sets the intent_hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *ParsedSignedTransactionIntentIdentifiers) SetIntentHash(value *string)() {
    m.intent_hash = value
}
// SetIntentHashBech32m sets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
func (m *ParsedSignedTransactionIntentIdentifiers) SetIntentHashBech32m(value *string)() {
    m.intent_hash_bech32m = value
}
// SetSignedIntentHash sets the signed_intent_hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
func (m *ParsedSignedTransactionIntentIdentifiers) SetSignedIntentHash(value *string)() {
    m.signed_intent_hash = value
}
// SetSignedIntentHashBech32m sets the signed_intent_hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
func (m *ParsedSignedTransactionIntentIdentifiers) SetSignedIntentHashBech32m(value *string)() {
    m.signed_intent_hash_bech32m = value
}
type ParsedSignedTransactionIntentIdentifiersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntentHash()(*string)
    GetIntentHashBech32m()(*string)
    GetSignedIntentHash()(*string)
    GetSignedIntentHashBech32m()(*string)
    SetIntentHash(value *string)()
    SetIntentHashBech32m(value *string)()
    SetSignedIntentHash(value *string)()
    SetSignedIntentHashBech32m(value *string)()
}
