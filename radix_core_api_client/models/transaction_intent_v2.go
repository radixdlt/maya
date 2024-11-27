package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionIntentV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
    hash *string
    // The Bech32m-encoded human readable `TransactionIntentHash`.
    hash_bech32m *string
    // The non_root_subintents property
    non_root_subintents []SubintentV2able
    // The root_intent_core property
    root_intent_core IntentCoreV2able
    // The transaction_header property
    transaction_header TransactionHeaderV2able
}
// NewTransactionIntentV2 instantiates a new TransactionIntentV2 and sets the default values.
func NewTransactionIntentV2()(*TransactionIntentV2) {
    m := &TransactionIntentV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionIntentV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionIntentV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionIntentV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionIntentV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionIntentV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["non_root_subintents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubintentV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubintentV2able, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubintentV2able)
                }
            }
            m.SetNonRootSubintents(res)
        }
        return nil
    }
    res["root_intent_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntentCoreV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootIntentCore(val.(IntentCoreV2able))
        }
        return nil
    }
    res["transaction_header"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionHeaderV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionHeader(val.(TransactionHeaderV2able))
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
// returns a *string when successful
func (m *TransactionIntentV2) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
// returns a *string when successful
func (m *TransactionIntentV2) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetNonRootSubintents gets the non_root_subintents property value. The non_root_subintents property
// returns a []SubintentV2able when successful
func (m *TransactionIntentV2) GetNonRootSubintents()([]SubintentV2able) {
    return m.non_root_subintents
}
// GetRootIntentCore gets the root_intent_core property value. The root_intent_core property
// returns a IntentCoreV2able when successful
func (m *TransactionIntentV2) GetRootIntentCore()(IntentCoreV2able) {
    return m.root_intent_core
}
// GetTransactionHeader gets the transaction_header property value. The transaction_header property
// returns a TransactionHeaderV2able when successful
func (m *TransactionIntentV2) GetTransactionHeader()(TransactionHeaderV2able) {
    return m.transaction_header
}
// Serialize serializes information the current object
func (m *TransactionIntentV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetNonRootSubintents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonRootSubintents()))
        for i, v := range m.GetNonRootSubintents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("non_root_subintents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("root_intent_core", m.GetRootIntentCore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transaction_header", m.GetTransactionHeader())
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
func (m *TransactionIntentV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. The hex-encoded transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.
func (m *TransactionIntentV2) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `TransactionIntentHash`.
func (m *TransactionIntentV2) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetNonRootSubintents sets the non_root_subintents property value. The non_root_subintents property
func (m *TransactionIntentV2) SetNonRootSubintents(value []SubintentV2able)() {
    m.non_root_subintents = value
}
// SetRootIntentCore sets the root_intent_core property value. The root_intent_core property
func (m *TransactionIntentV2) SetRootIntentCore(value IntentCoreV2able)() {
    m.root_intent_core = value
}
// SetTransactionHeader sets the transaction_header property value. The transaction_header property
func (m *TransactionIntentV2) SetTransactionHeader(value TransactionHeaderV2able)() {
    m.transaction_header = value
}
type TransactionIntentV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetNonRootSubintents()([]SubintentV2able)
    GetRootIntentCore()(IntentCoreV2able)
    GetTransactionHeader()(TransactionHeaderV2able)
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetNonRootSubintents(value []SubintentV2able)()
    SetRootIntentCore(value IntentCoreV2able)()
    SetTransactionHeader(value TransactionHeaderV2able)()
}
