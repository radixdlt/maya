package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommittedIntentMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the intent was committed in a transaction with the same payload.This is a convenience field, which can also be computed using `payload_hash` by a clientknowing the payload of the submitted transaction.
    is_same_transaction *bool
    // The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
    payload_hash *string
    // The Bech32m-encoded human readable `NotarizedTransactionHash`.
    payload_hash_bech32m *string
    // The state_version property
    state_version *int64
}
// NewCommittedIntentMetadata instantiates a new CommittedIntentMetadata and sets the default values.
func NewCommittedIntentMetadata()(*CommittedIntentMetadata) {
    m := &CommittedIntentMetadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommittedIntentMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommittedIntentMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommittedIntentMetadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommittedIntentMetadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommittedIntentMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_same_transaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSameTransaction(val)
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
    res["state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateVersion(val)
        }
        return nil
    }
    return res
}
// GetIsSameTransaction gets the is_same_transaction property value. Whether the intent was committed in a transaction with the same payload.This is a convenience field, which can also be computed using `payload_hash` by a clientknowing the payload of the submitted transaction.
// returns a *bool when successful
func (m *CommittedIntentMetadata) GetIsSameTransaction()(*bool) {
    return m.is_same_transaction
}
// GetPayloadHash gets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
// returns a *string when successful
func (m *CommittedIntentMetadata) GetPayloadHash()(*string) {
    return m.payload_hash
}
// GetPayloadHashBech32m gets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
// returns a *string when successful
func (m *CommittedIntentMetadata) GetPayloadHashBech32m()(*string) {
    return m.payload_hash_bech32m
}
// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *CommittedIntentMetadata) GetStateVersion()(*int64) {
    return m.state_version
}
// Serialize serializes information the current object
func (m *CommittedIntentMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_same_transaction", m.GetIsSameTransaction())
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
        err := writer.WriteInt64Value("state_version", m.GetStateVersion())
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
func (m *CommittedIntentMetadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsSameTransaction sets the is_same_transaction property value. Whether the intent was committed in a transaction with the same payload.This is a convenience field, which can also be computed using `payload_hash` by a clientknowing the payload of the submitted transaction.
func (m *CommittedIntentMetadata) SetIsSameTransaction(value *bool)() {
    m.is_same_transaction = value
}
// SetPayloadHash sets the payload_hash property value. The hex-encoded notarized transaction hash for a user transaction.This hash identifies the full submittable notarized transaction - ie the signed intent, plus the notary signature.
func (m *CommittedIntentMetadata) SetPayloadHash(value *string)() {
    m.payload_hash = value
}
// SetPayloadHashBech32m sets the payload_hash_bech32m property value. The Bech32m-encoded human readable `NotarizedTransactionHash`.
func (m *CommittedIntentMetadata) SetPayloadHashBech32m(value *string)() {
    m.payload_hash_bech32m = value
}
// SetStateVersion sets the state_version property value. The state_version property
func (m *CommittedIntentMetadata) SetStateVersion(value *int64)() {
    m.state_version = value
}
type CommittedIntentMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsSameTransaction()(*bool)
    GetPayloadHash()(*string)
    GetPayloadHashBech32m()(*string)
    GetStateVersion()(*int64)
    SetIsSameTransaction(value *bool)()
    SetPayloadHash(value *string)()
    SetPayloadHashBech32m(value *string)()
    SetStateVersion(value *int64)()
}
