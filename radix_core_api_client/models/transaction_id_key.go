package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionIdKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
    intent_hash *string
    // The Bech32m-encoded human readable `TransactionIntentHash`.
    intent_hash_bech32m *string
}
// NewTransactionIdKey instantiates a new TransactionIdKey and sets the default values.
func NewTransactionIdKey()(*TransactionIdKey) {
    m := &TransactionIdKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionIdKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionIdKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionIdKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionIdKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionIdKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetIntentHash gets the intent_hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
// returns a *string when successful
func (m *TransactionIdKey) GetIntentHash()(*string) {
    return m.intent_hash
}
// GetIntentHashBech32m gets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
// returns a *string when successful
func (m *TransactionIdKey) GetIntentHashBech32m()(*string) {
    return m.intent_hash_bech32m
}
// Serialize serializes information the current object
func (m *TransactionIdKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransactionIdKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntentHash sets the intent_hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *TransactionIdKey) SetIntentHash(value *string)() {
    m.intent_hash = value
}
// SetIntentHashBech32m sets the intent_hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
func (m *TransactionIdKey) SetIntentHashBech32m(value *string)() {
    m.intent_hash_bech32m = value
}
type TransactionIdKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntentHash()(*string)
    GetIntentHashBech32m()(*string)
    SetIntentHash(value *string)()
    SetIntentHashBech32m(value *string)()
}
