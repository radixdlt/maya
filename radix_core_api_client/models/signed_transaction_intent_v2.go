package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SignedTransactionIntentV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
    hash *string
    // The Bech32m-encoded human readable `SignedTransactionIntentHash`.
    hash_bech32m *string
    // This gives the signatures for each subintent in `non_root_subintents` in `TransactionIntentV2`.For committed transactions, these arrays are of equal length and correspond one-to-one in order.
    non_root_subintent_signatures []IntentSignaturesable
    // The transaction_intent property
    transaction_intent TransactionIntentV2able
    // Signatures against the given intent.
    transaction_intent_signatures IntentSignaturesable
}
// NewSignedTransactionIntentV2 instantiates a new SignedTransactionIntentV2 and sets the default values.
func NewSignedTransactionIntentV2()(*SignedTransactionIntentV2) {
    m := &SignedTransactionIntentV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSignedTransactionIntentV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignedTransactionIntentV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignedTransactionIntentV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SignedTransactionIntentV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SignedTransactionIntentV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHash(val)
        }
        return nil
    }
    res["hash_bech32m"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashBech32m(val)
        }
        return nil
    }
    res["non_root_subintent_signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntentSignaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntentSignaturesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntentSignaturesable)
                }
            }
            m.SetNonRootSubintentSignatures(res)
        }
        return nil
    }
    res["transaction_intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionIntentV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionIntent(val.(TransactionIntentV2able))
        }
        return nil
    }
    res["transaction_intent_signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntentSignaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionIntentSignatures(val.(IntentSignaturesable))
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
// returns a *string when successful
func (m *SignedTransactionIntentV2) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
// returns a *string when successful
func (m *SignedTransactionIntentV2) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetNonRootSubintentSignatures gets the non_root_subintent_signatures property value. This gives the signatures for each subintent in `non_root_subintents` in `TransactionIntentV2`.For committed transactions, these arrays are of equal length and correspond one-to-one in order.
// returns a []IntentSignaturesable when successful
func (m *SignedTransactionIntentV2) GetNonRootSubintentSignatures()([]IntentSignaturesable) {
    return m.non_root_subintent_signatures
}
// GetTransactionIntent gets the transaction_intent property value. The transaction_intent property
// returns a TransactionIntentV2able when successful
func (m *SignedTransactionIntentV2) GetTransactionIntent()(TransactionIntentV2able) {
    return m.transaction_intent
}
// GetTransactionIntentSignatures gets the transaction_intent_signatures property value. Signatures against the given intent.
// returns a IntentSignaturesable when successful
func (m *SignedTransactionIntentV2) GetTransactionIntentSignatures()(IntentSignaturesable) {
    return m.transaction_intent_signatures
}
// Serialize serializes information the current object
func (m *SignedTransactionIntentV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hash", m.GetHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hash_bech32m", m.GetHashBech32m())
        if err != nil {
            return err
        }
    }
    if m.GetNonRootSubintentSignatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonRootSubintentSignatures()))
        for i, v := range m.GetNonRootSubintentSignatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("non_root_subintent_signatures", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transaction_intent", m.GetTransactionIntent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transaction_intent_signatures", m.GetTransactionIntentSignatures())
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
func (m *SignedTransactionIntentV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
func (m *SignedTransactionIntentV2) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
func (m *SignedTransactionIntentV2) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetNonRootSubintentSignatures sets the non_root_subintent_signatures property value. This gives the signatures for each subintent in `non_root_subintents` in `TransactionIntentV2`.For committed transactions, these arrays are of equal length and correspond one-to-one in order.
func (m *SignedTransactionIntentV2) SetNonRootSubintentSignatures(value []IntentSignaturesable)() {
    m.non_root_subintent_signatures = value
}
// SetTransactionIntent sets the transaction_intent property value. The transaction_intent property
func (m *SignedTransactionIntentV2) SetTransactionIntent(value TransactionIntentV2able)() {
    m.transaction_intent = value
}
// SetTransactionIntentSignatures sets the transaction_intent_signatures property value. Signatures against the given intent.
func (m *SignedTransactionIntentV2) SetTransactionIntentSignatures(value IntentSignaturesable)() {
    m.transaction_intent_signatures = value
}
type SignedTransactionIntentV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetNonRootSubintentSignatures()([]IntentSignaturesable)
    GetTransactionIntent()(TransactionIntentV2able)
    GetTransactionIntentSignatures()(IntentSignaturesable)
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetNonRootSubintentSignatures(value []IntentSignaturesable)()
    SetTransactionIntent(value TransactionIntentV2able)()
    SetTransactionIntentSignatures(value IntentSignaturesable)()
}
