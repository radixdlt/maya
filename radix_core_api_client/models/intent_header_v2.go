package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntentHeaderV2 the metadata common to both transaction intents and subintents.The `min_proposer_timestamp_inclusive` and `max_proposer_timestamp_exclusive` are both optional.
type IntentHeaderV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
    end_epoch_exclusive *int64
    // The string representation of an integer between `0` and `2^64 - 1`, which can be chosen to ensure thata unique intent can be created. It only needs to be unique for a particular intent content and epoch/timestamp,and can be safely selected randomly to very high probability.This field was called `nonce` in the V1 models, but was a misleading name, as it got confused with acryptographic nonce or an Ethereum-style nonce, and it is neither.
    intent_discriminator *string
    // The max_proposer_timestamp_exclusive property
    max_proposer_timestamp_exclusive ScryptoInstantable
    // The min_proposer_timestamp_inclusive property
    min_proposer_timestamp_inclusive ScryptoInstantable
    // The logical id of the network
    network_id *int32
    // An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
    start_epoch_inclusive *int64
}
// NewIntentHeaderV2 instantiates a new IntentHeaderV2 and sets the default values.
func NewIntentHeaderV2()(*IntentHeaderV2) {
    m := &IntentHeaderV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntentHeaderV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentHeaderV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentHeaderV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntentHeaderV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndEpochExclusive gets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
// returns a *int64 when successful
func (m *IntentHeaderV2) GetEndEpochExclusive()(*int64) {
    return m.end_epoch_exclusive
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntentHeaderV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end_epoch_exclusive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndEpochExclusive(val)
        }
        return nil
    }
    res["intent_discriminator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentDiscriminator(val)
        }
        return nil
    }
    res["max_proposer_timestamp_exclusive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScryptoInstantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxProposerTimestampExclusive(val.(ScryptoInstantable))
        }
        return nil
    }
    res["min_proposer_timestamp_inclusive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScryptoInstantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinProposerTimestampInclusive(val.(ScryptoInstantable))
        }
        return nil
    }
    res["network_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkId(val)
        }
        return nil
    }
    res["start_epoch_inclusive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartEpochInclusive(val)
        }
        return nil
    }
    return res
}
// GetIntentDiscriminator gets the intent_discriminator property value. The string representation of an integer between `0` and `2^64 - 1`, which can be chosen to ensure thata unique intent can be created. It only needs to be unique for a particular intent content and epoch/timestamp,and can be safely selected randomly to very high probability.This field was called `nonce` in the V1 models, but was a misleading name, as it got confused with acryptographic nonce or an Ethereum-style nonce, and it is neither.
// returns a *string when successful
func (m *IntentHeaderV2) GetIntentDiscriminator()(*string) {
    return m.intent_discriminator
}
// GetMaxProposerTimestampExclusive gets the max_proposer_timestamp_exclusive property value. The max_proposer_timestamp_exclusive property
// returns a ScryptoInstantable when successful
func (m *IntentHeaderV2) GetMaxProposerTimestampExclusive()(ScryptoInstantable) {
    return m.max_proposer_timestamp_exclusive
}
// GetMinProposerTimestampInclusive gets the min_proposer_timestamp_inclusive property value. The min_proposer_timestamp_inclusive property
// returns a ScryptoInstantable when successful
func (m *IntentHeaderV2) GetMinProposerTimestampInclusive()(ScryptoInstantable) {
    return m.min_proposer_timestamp_inclusive
}
// GetNetworkId gets the network_id property value. The logical id of the network
// returns a *int32 when successful
func (m *IntentHeaderV2) GetNetworkId()(*int32) {
    return m.network_id
}
// GetStartEpochInclusive gets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
// returns a *int64 when successful
func (m *IntentHeaderV2) GetStartEpochInclusive()(*int64) {
    return m.start_epoch_inclusive
}
// Serialize serializes information the current object
func (m *IntentHeaderV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end_epoch_exclusive", m.GetEndEpochExclusive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("intent_discriminator", m.GetIntentDiscriminator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("max_proposer_timestamp_exclusive", m.GetMaxProposerTimestampExclusive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("min_proposer_timestamp_inclusive", m.GetMinProposerTimestampInclusive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("network_id", m.GetNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start_epoch_inclusive", m.GetStartEpochInclusive())
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
func (m *IntentHeaderV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndEpochExclusive sets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
func (m *IntentHeaderV2) SetEndEpochExclusive(value *int64)() {
    m.end_epoch_exclusive = value
}
// SetIntentDiscriminator sets the intent_discriminator property value. The string representation of an integer between `0` and `2^64 - 1`, which can be chosen to ensure thata unique intent can be created. It only needs to be unique for a particular intent content and epoch/timestamp,and can be safely selected randomly to very high probability.This field was called `nonce` in the V1 models, but was a misleading name, as it got confused with acryptographic nonce or an Ethereum-style nonce, and it is neither.
func (m *IntentHeaderV2) SetIntentDiscriminator(value *string)() {
    m.intent_discriminator = value
}
// SetMaxProposerTimestampExclusive sets the max_proposer_timestamp_exclusive property value. The max_proposer_timestamp_exclusive property
func (m *IntentHeaderV2) SetMaxProposerTimestampExclusive(value ScryptoInstantable)() {
    m.max_proposer_timestamp_exclusive = value
}
// SetMinProposerTimestampInclusive sets the min_proposer_timestamp_inclusive property value. The min_proposer_timestamp_inclusive property
func (m *IntentHeaderV2) SetMinProposerTimestampInclusive(value ScryptoInstantable)() {
    m.min_proposer_timestamp_inclusive = value
}
// SetNetworkId sets the network_id property value. The logical id of the network
func (m *IntentHeaderV2) SetNetworkId(value *int32)() {
    m.network_id = value
}
// SetStartEpochInclusive sets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
func (m *IntentHeaderV2) SetStartEpochInclusive(value *int64)() {
    m.start_epoch_inclusive = value
}
type IntentHeaderV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndEpochExclusive()(*int64)
    GetIntentDiscriminator()(*string)
    GetMaxProposerTimestampExclusive()(ScryptoInstantable)
    GetMinProposerTimestampInclusive()(ScryptoInstantable)
    GetNetworkId()(*int32)
    GetStartEpochInclusive()(*int64)
    SetEndEpochExclusive(value *int64)()
    SetIntentDiscriminator(value *string)()
    SetMaxProposerTimestampExclusive(value ScryptoInstantable)()
    SetMinProposerTimestampInclusive(value ScryptoInstantable)()
    SetNetworkId(value *int32)()
    SetStartEpochInclusive(value *int64)()
}
