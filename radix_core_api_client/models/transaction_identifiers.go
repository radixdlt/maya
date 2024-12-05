package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionIdentifiers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
    intent_hash *string
    // The Bech32m-encoded human readable `TransactionIntentHash`.
    intent_hash_bech32m *string
    // The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
    payload_hash *string
    // The Bech32m-encoded human readable `NotarizedTransactionHash`.
    payload_hash_bech32m *string
    // The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
    signed_intent_hash *string
    // The Bech32m-encoded human readable `SignedTransactionIntentHash`.
    signed_intent_hash_bech32m *string
}
// NewTransactionIdentifiers instantiates a new TransactionIdentifiers and sets the default values.
func NewTransactionIdentifiers()(*TransactionIdentifiers) {
    m := &TransactionIdentifiers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionIdentifiersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionIdentifiersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionIdentifiers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionIdentifiers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionIdentifiers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["payload_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadHash(val)
        }
        return nil
    }
    res["payload_hash_bech32m"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadHashBech32m(val)
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
func (m *TransactionIdentifiers) GetIntentHash()(*string) {
    return m.intent_hash
}
// GetIntentHashBech32m gets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
// returns a *string when successful
func (m *TransactionIdentifiers) GetIntentHashBech32m()(*string) {
    return m.intent_hash_bech32m
}
// GetPayloadHash gets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
// returns a *string when successful
func (m *TransactionIdentifiers) GetPayloadHash()(*string) {
    return m.payload_hash
}
// GetPayloadHashBech32m gets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
// returns a *string when successful
func (m *TransactionIdentifiers) GetPayloadHashBech32m()(*string) {
    return m.payload_hash_bech32m
}
// GetSignedIntentHash gets the signed_intent_hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
// returns a *string when successful
func (m *TransactionIdentifiers) GetSignedIntentHash()(*string) {
    return m.signed_intent_hash
}
// GetSignedIntentHashBech32m gets the signed_intent_hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
// returns a *string when successful
func (m *TransactionIdentifiers) GetSignedIntentHashBech32m()(*string) {
    return m.signed_intent_hash_bech32m
}
// Serialize serializes information the current object
func (m *TransactionIdentifiers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("payload_hash", m.GetPayloadHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payload_hash_bech32m", m.GetPayloadHashBech32m())
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
func (m *TransactionIdentifiers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntentHash sets the intent_hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *TransactionIdentifiers) SetIntentHash(value *string)() {
    m.intent_hash = value
}
// SetIntentHashBech32m sets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
func (m *TransactionIdentifiers) SetIntentHashBech32m(value *string)() {
    m.intent_hash_bech32m = value
}
// SetPayloadHash sets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
func (m *TransactionIdentifiers) SetPayloadHash(value *string)() {
    m.payload_hash = value
}
// SetPayloadHashBech32m sets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
func (m *TransactionIdentifiers) SetPayloadHashBech32m(value *string)() {
    m.payload_hash_bech32m = value
}
// SetSignedIntentHash sets the signed_intent_hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
func (m *TransactionIdentifiers) SetSignedIntentHash(value *string)() {
    m.signed_intent_hash = value
}
// SetSignedIntentHashBech32m sets the signed_intent_hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
func (m *TransactionIdentifiers) SetSignedIntentHashBech32m(value *string)() {
    m.signed_intent_hash_bech32m = value
}
type TransactionIdentifiersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntentHash()(*string)
    GetIntentHashBech32m()(*string)
    GetPayloadHash()(*string)
    GetPayloadHashBech32m()(*string)
    GetSignedIntentHash()(*string)
    GetSignedIntentHashBech32m()(*string)
    SetIntentHash(value *string)()
    SetIntentHashBech32m(value *string)()
    SetPayloadHash(value *string)()
    SetPayloadHashBech32m(value *string)()
    SetSignedIntentHash(value *string)()
    SetSignedIntentHashBech32m(value *string)()
}
