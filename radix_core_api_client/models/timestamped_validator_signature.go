package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TimestampedValidatorSignature struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The signature property
    signature EcdsaSecp256k1Signatureable
    // An integer between `0` and `10^14`, marking the unix timestamp in ms.
    timestamp_ms *int64
    // The Bech32m-encoded human readable version of the component address
    validator_address *string
    // The validator_key property
    validator_key EcdsaSecp256k1PublicKeyable
}
// NewTimestampedValidatorSignature instantiates a new TimestampedValidatorSignature and sets the default values.
func NewTimestampedValidatorSignature()(*TimestampedValidatorSignature) {
    m := &TimestampedValidatorSignature{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimestampedValidatorSignatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimestampedValidatorSignatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimestampedValidatorSignature(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimestampedValidatorSignature) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimestampedValidatorSignature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEcdsaSecp256k1SignatureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val.(EcdsaSecp256k1Signatureable))
        }
        return nil
    }
    res["timestamp_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestampMs(val)
        }
        return nil
    }
    res["validator_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorAddress(val)
        }
        return nil
    }
    res["validator_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEcdsaSecp256k1PublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatorKey(val.(EcdsaSecp256k1PublicKeyable))
        }
        return nil
    }
    return res
}
// GetSignature gets the signature property value. The signature property
// returns a EcdsaSecp256k1Signatureable when successful
func (m *TimestampedValidatorSignature) GetSignature()(EcdsaSecp256k1Signatureable) {
    return m.signature
}
// GetTimestampMs gets the timestamp_ms property value. An integer between `0` and `10^14`, marking the unix timestamp in ms.
// returns a *int64 when successful
func (m *TimestampedValidatorSignature) GetTimestampMs()(*int64) {
    return m.timestamp_ms
}
// GetValidatorAddress gets the validator_address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *TimestampedValidatorSignature) GetValidatorAddress()(*string) {
    return m.validator_address
}
// GetValidatorKey gets the validator_key property value. The validator_key property
// returns a EcdsaSecp256k1PublicKeyable when successful
func (m *TimestampedValidatorSignature) GetValidatorKey()(EcdsaSecp256k1PublicKeyable) {
    return m.validator_key
}
// Serialize serializes information the current object
func (m *TimestampedValidatorSignature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("timestamp_ms", m.GetTimestampMs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validator_address", m.GetValidatorAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("validator_key", m.GetValidatorKey())
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
func (m *TimestampedValidatorSignature) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSignature sets the signature property value. The signature property
func (m *TimestampedValidatorSignature) SetSignature(value EcdsaSecp256k1Signatureable)() {
    m.signature = value
}
// SetTimestampMs sets the timestamp_ms property value. An integer between `0` and `10^14`, marking the unix timestamp in ms.
func (m *TimestampedValidatorSignature) SetTimestampMs(value *int64)() {
    m.timestamp_ms = value
}
// SetValidatorAddress sets the validator_address property value. The Bech32m-encoded human readable version of the component address
func (m *TimestampedValidatorSignature) SetValidatorAddress(value *string)() {
    m.validator_address = value
}
// SetValidatorKey sets the validator_key property value. The validator_key property
func (m *TimestampedValidatorSignature) SetValidatorKey(value EcdsaSecp256k1PublicKeyable)() {
    m.validator_key = value
}
type TimestampedValidatorSignatureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSignature()(EcdsaSecp256k1Signatureable)
    GetTimestampMs()(*int64)
    GetValidatorAddress()(*string)
    GetValidatorKey()(EcdsaSecp256k1PublicKeyable)
    SetSignature(value EcdsaSecp256k1Signatureable)()
    SetTimestampMs(value *int64)()
    SetValidatorAddress(value *string)()
    SetValidatorKey(value EcdsaSecp256k1PublicKeyable)()
}
