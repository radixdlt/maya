package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionHeader struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
    end_epoch_exclusive *int64
    // The logical id of the network
    network_id *int32
    // An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created (to enable submitting an otherwise identical/duplicate intent).As of Cuttlefish and V2 transaction models, this is now referred to in documentation as the `intent_discriminator`.
    nonce *int64
    // Specifies whether the notary public key should be included in the transaction signers list
    notary_is_signatory *bool
    // The notary_public_key property
    notary_public_key PublicKeyable
    // An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
    start_epoch_inclusive *int64
    // An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
    tip_percentage *int32
}
// NewTransactionHeader instantiates a new TransactionHeader and sets the default values.
func NewTransactionHeader()(*TransactionHeader) {
    m := &TransactionHeader{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionHeaderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionHeaderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionHeader(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionHeader) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndEpochExclusive gets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
// returns a *int64 when successful
func (m *TransactionHeader) GetEndEpochExclusive()(*int64) {
    return m.end_epoch_exclusive
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionHeader) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["nonce"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonce(val)
        }
        return nil
    }
    res["notary_is_signatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotaryIsSignatory(val)
        }
        return nil
    }
    res["notary_public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotaryPublicKey(val.(PublicKeyable))
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
    res["tip_percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTipPercentage(val)
        }
        return nil
    }
    return res
}
// GetNetworkId gets the network_id property value. The logical id of the network
// returns a *int32 when successful
func (m *TransactionHeader) GetNetworkId()(*int32) {
    return m.network_id
}
// GetNonce gets the nonce property value. An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created (to enable submitting an otherwise identical/duplicate intent).As of Cuttlefish and V2 transaction models, this is now referred to in documentation as the `intent_discriminator`.
// returns a *int64 when successful
func (m *TransactionHeader) GetNonce()(*int64) {
    return m.nonce
}
// GetNotaryIsSignatory gets the notary_is_signatory property value. Specifies whether the notary public key should be included in the transaction signers list
// returns a *bool when successful
func (m *TransactionHeader) GetNotaryIsSignatory()(*bool) {
    return m.notary_is_signatory
}
// GetNotaryPublicKey gets the notary_public_key property value. The notary_public_key property
// returns a PublicKeyable when successful
func (m *TransactionHeader) GetNotaryPublicKey()(PublicKeyable) {
    return m.notary_public_key
}
// GetStartEpochInclusive gets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
// returns a *int64 when successful
func (m *TransactionHeader) GetStartEpochInclusive()(*int64) {
    return m.start_epoch_inclusive
}
// GetTipPercentage gets the tip_percentage property value. An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
// returns a *int32 when successful
func (m *TransactionHeader) GetTipPercentage()(*int32) {
    return m.tip_percentage
}
// Serialize serializes information the current object
func (m *TransactionHeader) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end_epoch_exclusive", m.GetEndEpochExclusive())
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
        err := writer.WriteInt64Value("nonce", m.GetNonce())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notary_is_signatory", m.GetNotaryIsSignatory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notary_public_key", m.GetNotaryPublicKey())
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
        err := writer.WriteInt32Value("tip_percentage", m.GetTipPercentage())
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
func (m *TransactionHeader) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndEpochExclusive sets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be rejected.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
func (m *TransactionHeader) SetEndEpochExclusive(value *int64)() {
    m.end_epoch_exclusive = value
}
// SetNetworkId sets the network_id property value. The logical id of the network
func (m *TransactionHeader) SetNetworkId(value *int32)() {
    m.network_id = value
}
// SetNonce sets the nonce property value. An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created (to enable submitting an otherwise identical/duplicate intent).As of Cuttlefish and V2 transaction models, this is now referred to in documentation as the `intent_discriminator`.
func (m *TransactionHeader) SetNonce(value *int64)() {
    m.nonce = value
}
// SetNotaryIsSignatory sets the notary_is_signatory property value. Specifies whether the notary public key should be included in the transaction signers list
func (m *TransactionHeader) SetNotaryIsSignatory(value *bool)() {
    m.notary_is_signatory = value
}
// SetNotaryPublicKey sets the notary_public_key property value. The notary_public_key property
func (m *TransactionHeader) SetNotaryPublicKey(value PublicKeyable)() {
    m.notary_public_key = value
}
// SetStartEpochInclusive sets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch from which the transaction can be submitted.In the case of uncommitted transactions, a value of `10^10` indicates that the epoch was >= `10^10`.
func (m *TransactionHeader) SetStartEpochInclusive(value *int64)() {
    m.start_epoch_inclusive = value
}
// SetTipPercentage sets the tip_percentage property value. An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
func (m *TransactionHeader) SetTipPercentage(value *int32)() {
    m.tip_percentage = value
}
type TransactionHeaderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndEpochExclusive()(*int64)
    GetNetworkId()(*int32)
    GetNonce()(*int64)
    GetNotaryIsSignatory()(*bool)
    GetNotaryPublicKey()(PublicKeyable)
    GetStartEpochInclusive()(*int64)
    GetTipPercentage()(*int32)
    SetEndEpochExclusive(value *int64)()
    SetNetworkId(value *int32)()
    SetNonce(value *int64)()
    SetNotaryIsSignatory(value *bool)()
    SetNotaryPublicKey(value PublicKeyable)()
    SetStartEpochInclusive(value *int64)()
    SetTipPercentage(value *int32)()
}
