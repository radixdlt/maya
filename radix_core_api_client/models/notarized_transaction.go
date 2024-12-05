package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotarizedTransaction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
    hash *string
    // The Bech32m-encoded human readable `NotarizedTransactionHash`.
    hash_bech32m *string
    // The notary_signature property
    notary_signature Signatureable
    // The hex-encoded full notarized transaction payload. Returning this can be disabled in TransactionFormatOptions on your request (default true).
    payload_hex *string
    // The signed_intent property
    signed_intent SignedTransactionIntentable
}
// NewNotarizedTransaction instantiates a new NotarizedTransaction and sets the default values.
func NewNotarizedTransaction()(*NotarizedTransaction) {
    m := &NotarizedTransaction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNotarizedTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotarizedTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotarizedTransaction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotarizedTransaction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotarizedTransaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["notary_signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignatureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotarySignature(val.(Signatureable))
        }
        return nil
    }
    res["payload_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadHex(val)
        }
        return nil
    }
    res["signed_intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHash gets the hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
// returns a *string when successful
func (m *NotarizedTransaction) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
// returns a *string when successful
func (m *NotarizedTransaction) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetNotarySignature gets the notary_signature property value. The notary_signature property
// returns a Signatureable when successful
func (m *NotarizedTransaction) GetNotarySignature()(Signatureable) {
    return m.notary_signature
}
// GetPayloadHex gets the payload_hex property value. The hex-encoded full notarized transaction payload. Returning this can be disabled in TransactionFormatOptions on your request (default true).
// returns a *string when successful
func (m *NotarizedTransaction) GetPayloadHex()(*string) {
    return m.payload_hex
}
// GetSignedIntent gets the signed_intent property value. The signed_intent property
// returns a SignedTransactionIntentable when successful
func (m *NotarizedTransaction) GetSignedIntent()(SignedTransactionIntentable) {
    return m.signed_intent
}
// Serialize serializes information the current object
func (m *NotarizedTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("notary_signature", m.GetNotarySignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payload_hex", m.GetPayloadHex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("signed_intent", m.GetSignedIntent())
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
func (m *NotarizedTransaction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
func (m *NotarizedTransaction) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
func (m *NotarizedTransaction) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetNotarySignature sets the notary_signature property value. The notary_signature property
func (m *NotarizedTransaction) SetNotarySignature(value Signatureable)() {
    m.notary_signature = value
}
// SetPayloadHex sets the payload_hex property value. The hex-encoded full notarized transaction payload. Returning this can be disabled in TransactionFormatOptions on your request (default true).
func (m *NotarizedTransaction) SetPayloadHex(value *string)() {
    m.payload_hex = value
}
// SetSignedIntent sets the signed_intent property value. The signed_intent property
func (m *NotarizedTransaction) SetSignedIntent(value SignedTransactionIntentable)() {
    m.signed_intent = value
}
type NotarizedTransactionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetNotarySignature()(Signatureable)
    GetPayloadHex()(*string)
    GetSignedIntent()(SignedTransactionIntentable)
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetNotarySignature(value Signatureable)()
    SetPayloadHex(value *string)()
    SetSignedIntent(value SignedTransactionIntentable)()
}
