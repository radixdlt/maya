package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionStatusRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
    intent_hash *string
    // The logical name of the network
    network *string
}
// NewLtsTransactionStatusRequest instantiates a new LtsTransactionStatusRequest and sets the default values.
func NewLtsTransactionStatusRequest()(*LtsTransactionStatusRequest) {
    m := &LtsTransactionStatusRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsTransactionStatusRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionStatusRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsTransactionStatusRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsTransactionStatusRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionStatusRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    return res
}
// GetIntentHash gets the intent_hash property value. The transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
// returns a *string when successful
func (m *LtsTransactionStatusRequest) GetIntentHash()(*string) {
    return m.intent_hash
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *LtsTransactionStatusRequest) GetNetwork()(*string) {
    return m.network
}
// Serialize serializes information the current object
func (m *LtsTransactionStatusRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("intent_hash", m.GetIntentHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
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
func (m *LtsTransactionStatusRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntentHash sets the intent_hash property value. The transaction intent hash for a user transaction, also known as the transaction id.This hash identifies the core "intent" of the transaction. Each transaction intent can only be committed once.This hash gets signed by any signatories on the transaction, to create the signed intent.Either hex or Bech32m-encoded strings are supported.
func (m *LtsTransactionStatusRequest) SetIntentHash(value *string)() {
    m.intent_hash = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *LtsTransactionStatusRequest) SetNetwork(value *string)() {
    m.network = value
}
type LtsTransactionStatusRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntentHash()(*string)
    GetNetwork()(*string)
    SetIntentHash(value *string)()
    SetNetwork(value *string)()
}
