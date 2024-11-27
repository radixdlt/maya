package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
    at_ledger_state LedgerStateSelectorable
    // An array of hex-encoded blob data, if referenced by the manifest.
    blobs_hex []string
    // An integer between `0` and `10^10`, marking the epoch at which the transaction is nolonger valid.If not provided, a maximum epoch (relative to the `start_epoch_inclusive`) will be used.
    end_epoch_exclusive *int64
    // The flags property
    flags PreviewFlagsable
    // A text representation of a transaction manifest.
    manifest *string
    // The message property
    message TransactionMessageable
    // The logical name of the network
    network *string
    // An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created(to enable submitting an otherwise identical/duplicate intent).If not provided, this defaults to 0.
    nonce *int64
    // Whether the notary should be used as a signer (optional).If not provided, this defaults to false.
    notary_is_signatory *bool
    // The notary_public_key property
    notary_public_key PublicKeyable
    // A set of flags to configure the response of the transaction preview.
    options TransactionPreviewResponseOptionsable
    // A list of public keys to be used as transaction signers.If not provided, this defaults to an empty array.
    signer_public_keys []PublicKeyable
    // An integer between `0` and `10^10`, marking the epoch at which the transaction startsbeing valid.If not provided, the current epoch will be used (taking into account the`at_ledger_state`, if specified).
    start_epoch_inclusive *int64
    // An integer between `0` and `65535`, giving the validator tip as a percentage amount.A value of `1` corresponds to a 1% fee.If not provided, this defaults to 0.
    tip_percentage *int32
}
// NewTransactionPreviewRequest instantiates a new TransactionPreviewRequest and sets the default values.
func NewTransactionPreviewRequest()(*TransactionPreviewRequest) {
    m := &TransactionPreviewRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
// returns a LedgerStateSelectorable when successful
func (m *TransactionPreviewRequest) GetAtLedgerState()(LedgerStateSelectorable) {
    return m.at_ledger_state
}
// GetBlobsHex gets the blobs_hex property value. An array of hex-encoded blob data, if referenced by the manifest.
// returns a []string when successful
func (m *TransactionPreviewRequest) GetBlobsHex()([]string) {
    return m.blobs_hex
}
// GetEndEpochExclusive gets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch at which the transaction is nolonger valid.If not provided, a maximum epoch (relative to the `start_epoch_inclusive`) will be used.
// returns a *int64 when successful
func (m *TransactionPreviewRequest) GetEndEpochExclusive()(*int64) {
    return m.end_epoch_exclusive
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["at_ledger_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerStateSelectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtLedgerState(val.(LedgerStateSelectorable))
        }
        return nil
    }
    res["blobs_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetBlobsHex(res)
        }
        return nil
    }
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
    res["flags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreviewFlagsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlags(val.(PreviewFlagsable))
        }
        return nil
    }
    res["manifest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifest(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(TransactionMessageable))
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
    res["options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionPreviewResponseOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptions(val.(TransactionPreviewResponseOptionsable))
        }
        return nil
    }
    res["signer_public_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicKeyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PublicKeyable)
                }
            }
            m.SetSignerPublicKeys(res)
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
// GetFlags gets the flags property value. The flags property
// returns a PreviewFlagsable when successful
func (m *TransactionPreviewRequest) GetFlags()(PreviewFlagsable) {
    return m.flags
}
// GetManifest gets the manifest property value. A text representation of a transaction manifest.
// returns a *string when successful
func (m *TransactionPreviewRequest) GetManifest()(*string) {
    return m.manifest
}
// GetMessage gets the message property value. The message property
// returns a TransactionMessageable when successful
func (m *TransactionPreviewRequest) GetMessage()(TransactionMessageable) {
    return m.message
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionPreviewRequest) GetNetwork()(*string) {
    return m.network
}
// GetNonce gets the nonce property value. An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created(to enable submitting an otherwise identical/duplicate intent).If not provided, this defaults to 0.
// returns a *int64 when successful
func (m *TransactionPreviewRequest) GetNonce()(*int64) {
    return m.nonce
}
// GetNotaryIsSignatory gets the notary_is_signatory property value. Whether the notary should be used as a signer (optional).If not provided, this defaults to false.
// returns a *bool when successful
func (m *TransactionPreviewRequest) GetNotaryIsSignatory()(*bool) {
    return m.notary_is_signatory
}
// GetNotaryPublicKey gets the notary_public_key property value. The notary_public_key property
// returns a PublicKeyable when successful
func (m *TransactionPreviewRequest) GetNotaryPublicKey()(PublicKeyable) {
    return m.notary_public_key
}
// GetOptions gets the options property value. A set of flags to configure the response of the transaction preview.
// returns a TransactionPreviewResponseOptionsable when successful
func (m *TransactionPreviewRequest) GetOptions()(TransactionPreviewResponseOptionsable) {
    return m.options
}
// GetSignerPublicKeys gets the signer_public_keys property value. A list of public keys to be used as transaction signers.If not provided, this defaults to an empty array.
// returns a []PublicKeyable when successful
func (m *TransactionPreviewRequest) GetSignerPublicKeys()([]PublicKeyable) {
    return m.signer_public_keys
}
// GetStartEpochInclusive gets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch at which the transaction startsbeing valid.If not provided, the current epoch will be used (taking into account the`at_ledger_state`, if specified).
// returns a *int64 when successful
func (m *TransactionPreviewRequest) GetStartEpochInclusive()(*int64) {
    return m.start_epoch_inclusive
}
// GetTipPercentage gets the tip_percentage property value. An integer between `0` and `65535`, giving the validator tip as a percentage amount.A value of `1` corresponds to a 1% fee.If not provided, this defaults to 0.
// returns a *int32 when successful
func (m *TransactionPreviewRequest) GetTipPercentage()(*int32) {
    return m.tip_percentage
}
// Serialize serializes information the current object
func (m *TransactionPreviewRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    if m.GetBlobsHex() != nil {
        err := writer.WriteCollectionOfStringValues("blobs_hex", m.GetBlobsHex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("end_epoch_exclusive", m.GetEndEpochExclusive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("flags", m.GetFlags())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manifest", m.GetManifest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
        err := writer.WriteObjectValue("options", m.GetOptions())
        if err != nil {
            return err
        }
    }
    if m.GetSignerPublicKeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignerPublicKeys()))
        for i, v := range m.GetSignerPublicKeys() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("signer_public_keys", cast)
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
func (m *TransactionPreviewRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
func (m *TransactionPreviewRequest) SetAtLedgerState(value LedgerStateSelectorable)() {
    m.at_ledger_state = value
}
// SetBlobsHex sets the blobs_hex property value. An array of hex-encoded blob data, if referenced by the manifest.
func (m *TransactionPreviewRequest) SetBlobsHex(value []string)() {
    m.blobs_hex = value
}
// SetEndEpochExclusive sets the end_epoch_exclusive property value. An integer between `0` and `10^10`, marking the epoch at which the transaction is nolonger valid.If not provided, a maximum epoch (relative to the `start_epoch_inclusive`) will be used.
func (m *TransactionPreviewRequest) SetEndEpochExclusive(value *int64)() {
    m.end_epoch_exclusive = value
}
// SetFlags sets the flags property value. The flags property
func (m *TransactionPreviewRequest) SetFlags(value PreviewFlagsable)() {
    m.flags = value
}
// SetManifest sets the manifest property value. A text representation of a transaction manifest.
func (m *TransactionPreviewRequest) SetManifest(value *string)() {
    m.manifest = value
}
// SetMessage sets the message property value. The message property
func (m *TransactionPreviewRequest) SetMessage(value TransactionMessageable)() {
    m.message = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionPreviewRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetNonce sets the nonce property value. An integer between `0` and `2^32 - 1`, chosen to allow a unique intent to be created(to enable submitting an otherwise identical/duplicate intent).If not provided, this defaults to 0.
func (m *TransactionPreviewRequest) SetNonce(value *int64)() {
    m.nonce = value
}
// SetNotaryIsSignatory sets the notary_is_signatory property value. Whether the notary should be used as a signer (optional).If not provided, this defaults to false.
func (m *TransactionPreviewRequest) SetNotaryIsSignatory(value *bool)() {
    m.notary_is_signatory = value
}
// SetNotaryPublicKey sets the notary_public_key property value. The notary_public_key property
func (m *TransactionPreviewRequest) SetNotaryPublicKey(value PublicKeyable)() {
    m.notary_public_key = value
}
// SetOptions sets the options property value. A set of flags to configure the response of the transaction preview.
func (m *TransactionPreviewRequest) SetOptions(value TransactionPreviewResponseOptionsable)() {
    m.options = value
}
// SetSignerPublicKeys sets the signer_public_keys property value. A list of public keys to be used as transaction signers.If not provided, this defaults to an empty array.
func (m *TransactionPreviewRequest) SetSignerPublicKeys(value []PublicKeyable)() {
    m.signer_public_keys = value
}
// SetStartEpochInclusive sets the start_epoch_inclusive property value. An integer between `0` and `10^10`, marking the epoch at which the transaction startsbeing valid.If not provided, the current epoch will be used (taking into account the`at_ledger_state`, if specified).
func (m *TransactionPreviewRequest) SetStartEpochInclusive(value *int64)() {
    m.start_epoch_inclusive = value
}
// SetTipPercentage sets the tip_percentage property value. An integer between `0` and `65535`, giving the validator tip as a percentage amount.A value of `1` corresponds to a 1% fee.If not provided, this defaults to 0.
func (m *TransactionPreviewRequest) SetTipPercentage(value *int32)() {
    m.tip_percentage = value
}
type TransactionPreviewRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSelectorable)
    GetBlobsHex()([]string)
    GetEndEpochExclusive()(*int64)
    GetFlags()(PreviewFlagsable)
    GetManifest()(*string)
    GetMessage()(TransactionMessageable)
    GetNetwork()(*string)
    GetNonce()(*int64)
    GetNotaryIsSignatory()(*bool)
    GetNotaryPublicKey()(PublicKeyable)
    GetOptions()(TransactionPreviewResponseOptionsable)
    GetSignerPublicKeys()([]PublicKeyable)
    GetStartEpochInclusive()(*int64)
    GetTipPercentage()(*int32)
    SetAtLedgerState(value LedgerStateSelectorable)()
    SetBlobsHex(value []string)()
    SetEndEpochExclusive(value *int64)()
    SetFlags(value PreviewFlagsable)()
    SetManifest(value *string)()
    SetMessage(value TransactionMessageable)()
    SetNetwork(value *string)()
    SetNonce(value *int64)()
    SetNotaryIsSignatory(value *bool)()
    SetNotaryPublicKey(value PublicKeyable)()
    SetOptions(value TransactionPreviewResponseOptionsable)()
    SetSignerPublicKeys(value []PublicKeyable)()
    SetStartEpochInclusive(value *int64)()
    SetTipPercentage(value *int32)()
}
