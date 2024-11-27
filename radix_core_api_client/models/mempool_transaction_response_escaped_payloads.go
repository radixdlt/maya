package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MempoolTransactionResponse_payloads struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Error message why `hex` field is missing: the transaction was not found in the mempool or the provided hash is invalid.
    error *string
    // The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
    hash *string
    // The Bech32m-encoded human readable `NotarizedTransactionHash`.
    hash_bech32m *string
    // The hex-encoded full notarized transaction payload - returned only if found in mempool.
    hex *string
}
// NewMempoolTransactionResponse_payloads instantiates a new MempoolTransactionResponse_payloads and sets the default values.
func NewMempoolTransactionResponse_payloads()(*MempoolTransactionResponse_payloads) {
    m := &MempoolTransactionResponse_payloads{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMempoolTransactionResponse_payloadsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMempoolTransactionResponse_payloadsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMempoolTransactionResponse_payloads(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MempoolTransactionResponse_payloads) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. Error message why `hex` field is missing: the transaction was not found in the mempool or the provided hash is invalid.
// returns a *string when successful
func (m *MempoolTransactionResponse_payloads) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MempoolTransactionResponse_payloads) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
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
    res["hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHex(val)
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
// returns a *string when successful
func (m *MempoolTransactionResponse_payloads) GetHash()(*string) {
    return m.hash
}
// GetHashBech32m gets the hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
// returns a *string when successful
func (m *MempoolTransactionResponse_payloads) GetHashBech32m()(*string) {
    return m.hash_bech32m
}
// GetHex gets the hex property value. The hex-encoded full notarized transaction payload - returned only if found in mempool.
// returns a *string when successful
func (m *MempoolTransactionResponse_payloads) GetHex()(*string) {
    return m.hex
}
// Serialize serializes information the current object
func (m *MempoolTransactionResponse_payloads) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("hex", m.GetHex())
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
func (m *MempoolTransactionResponse_payloads) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. Error message why `hex` field is missing: the transaction was not found in the mempool or the provided hash is invalid.
func (m *MempoolTransactionResponse_payloads) SetError(value *string)() {
    m.error = value
}
// SetHash sets the hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
func (m *MempoolTransactionResponse_payloads) SetHash(value *string)() {
    m.hash = value
}
// SetHashBech32m sets the hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
func (m *MempoolTransactionResponse_payloads) SetHashBech32m(value *string)() {
    m.hash_bech32m = value
}
// SetHex sets the hex property value. The hex-encoded full notarized transaction payload - returned only if found in mempool.
func (m *MempoolTransactionResponse_payloads) SetHex(value *string)() {
    m.hex = value
}
type MempoolTransactionResponse_payloadsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*string)
    GetHash()(*string)
    GetHashBech32m()(*string)
    GetHex()(*string)
    SetError(value *string)()
    SetHash(value *string)()
    SetHashBech32m(value *string)()
    SetHex(value *string)()
}
