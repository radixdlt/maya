package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubintentV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hex-encoded subintent hash for a subintent, also known as the partial transaction id.This hash identifies the subintent. Each subintent can only be *successfully* committed once,but unlike a transaction intent, could be committed as a failure zero or more times first.This hash gets signed by any signatories on subintent.
    hash *string
    // The Bech32m-encoded human readable `SubintentHash`.
    hash_bech32m *string
    // The intent_core property
    intent_core IntentCoreV2able
}
// NewSubintentV2 instantiates a new SubintentV2 and sets the default values.
func NewSubintentV2()(*SubintentV2) {
    m := &SubintentV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubintentV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubintentV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubintentV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubintentV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubintentV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["intent_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntentCoreV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentCore(val.(IntentCoreV2able))
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hex-encoded subintent hash for a subintent, also known as the partial transaction id.This hash identifies the subintent. Each subintent can only be *successfully* committed once,but unlike a transaction intent, could be committed as a failure zero or more times first.This hash gets signed by any signatories on subintent.
// returns a *string when successful
func (m *SubintentV2) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `SubintentHash`.
// returns a *string when successful
func (m *SubintentV2) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetIntentCore gets the intent_core property value. The intent_core property
// returns a IntentCoreV2able when successful
func (m *SubintentV2) GetIntentCore()(IntentCoreV2able) {
    return m.intent_core
}
// Serialize serializes information the current object
func (m *SubintentV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("intent_core", m.GetIntentCore())
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
func (m *SubintentV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. The hex-encoded subintent hash for a subintent, also known as the partial transaction id.This hash identifies the subintent. Each subintent can only be *successfully* committed once,but unlike a transaction intent, could be committed as a failure zero or more times first.This hash gets signed by any signatories on subintent.
func (m *SubintentV2) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `SubintentHash`.
func (m *SubintentV2) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetIntentCore sets the intent_core property value. The intent_core property
func (m *SubintentV2) SetIntentCore(value IntentCoreV2able)() {
    m.intent_core = value
}
type SubintentV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetIntentCore()(IntentCoreV2able)
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetIntentCore(value IntentCoreV2able)()
}
