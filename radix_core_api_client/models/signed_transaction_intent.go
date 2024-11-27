package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SignedTransactionIntent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
    hash *string
    // The Bech32m-encoded human readable `SignedTransactionIntentHash`.
    hash_bech32m *string
    // The intent property
    intent TransactionIntentable
    // The intent_signatures property
    intent_signatures []SignatureWithPublicKeyable
}
// NewSignedTransactionIntent instantiates a new SignedTransactionIntent and sets the default values.
func NewSignedTransactionIntent()(*SignedTransactionIntent) {
    m := &SignedTransactionIntent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSignedTransactionIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignedTransactionIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignedTransactionIntent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SignedTransactionIntent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SignedTransactionIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["intent_signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignatureWithPublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignatureWithPublicKeyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SignatureWithPublicKeyable)
                }
            }
            m.SetIntentSignatures(res)
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
// returns a *string when successful
func (m *SignedTransactionIntent) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
// returns a *string when successful
func (m *SignedTransactionIntent) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetIntent gets the intent property value. The intent property
// returns a TransactionIntentable when successful
func (m *SignedTransactionIntent) GetIntent()(TransactionIntentable) {
    return m.intent
}
// GetIntentSignatures gets the intent_signatures property value. The intent_signatures property
// returns a []SignatureWithPublicKeyable when successful
func (m *SignedTransactionIntent) GetIntentSignatures()([]SignatureWithPublicKeyable) {
    return m.intent_signatures
}
// Serialize serializes information the current object
func (m *SignedTransactionIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteObjectValue("intent", m.GetIntent())
        if err != nil {
            return err
        }
    }
    if m.GetIntentSignatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntentSignatures()))
        for i, v := range m.GetIntentSignatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("intent_signatures", cast)
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
func (m *SignedTransactionIntent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. The hex-encoded signed intent hash for a user transaction.This hash identifies the transaction intent, plus additional signatures.This hash is signed by the notary, to create the submittable `NotarizedTransaction`.
func (m *SignedTransactionIntent) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `SignedTransactionIntentHash`.
func (m *SignedTransactionIntent) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetIntent sets the intent property value. The intent property
func (m *SignedTransactionIntent) SetIntent(value TransactionIntentable)() {
    m.intent = value
}
// SetIntentSignatures sets the intent_signatures property value. The intent_signatures property
func (m *SignedTransactionIntent) SetIntentSignatures(value []SignatureWithPublicKeyable)() {
    m.intent_signatures = value
}
type SignedTransactionIntentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetIntent()(TransactionIntentable)
    GetIntentSignatures()([]SignatureWithPublicKeyable)
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetIntent(value TransactionIntentable)()
    SetIntentSignatures(value []SignatureWithPublicKeyable)()
}
