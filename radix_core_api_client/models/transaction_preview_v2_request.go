package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewV2Request struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
    at_ledger_state LedgerStateSelectorable
    // The flags property
    flags PreviewFlagsable
    // The logical name of the network
    network *string
    // A set of flags to configure the response of the transaction preview.
    options TransactionPreviewV2ResponseOptionsable
    // The preview_transaction property
    preview_transaction PreviewTransactionable
}
// NewTransactionPreviewV2Request instantiates a new TransactionPreviewV2Request and sets the default values.
func NewTransactionPreviewV2Request()(*TransactionPreviewV2Request) {
    m := &TransactionPreviewV2Request{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2RequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2RequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewV2Request(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2Request) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
// returns a LedgerStateSelectorable when successful
func (m *TransactionPreviewV2Request) GetAtLedgerState()(LedgerStateSelectorable) {
    return m.at_ledger_state
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2Request) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionPreviewV2ResponseOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptions(val.(TransactionPreviewV2ResponseOptionsable))
        }
        return nil
    }
    res["preview_transaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreviewTransactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewTransaction(val.(PreviewTransactionable))
        }
        return nil
    }
    return res
}
// GetFlags gets the flags property value. The flags property
// returns a PreviewFlagsable when successful
func (m *TransactionPreviewV2Request) GetFlags()(PreviewFlagsable) {
    return m.flags
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionPreviewV2Request) GetNetwork()(*string) {
    return m.network
}
// GetOptions gets the options property value. A set of flags to configure the response of the transaction preview.
// returns a TransactionPreviewV2ResponseOptionsable when successful
func (m *TransactionPreviewV2Request) GetOptions()(TransactionPreviewV2ResponseOptionsable) {
    return m.options
}
// GetPreviewTransaction gets the preview_transaction property value. The preview_transaction property
// returns a PreviewTransactionable when successful
func (m *TransactionPreviewV2Request) GetPreviewTransaction()(PreviewTransactionable) {
    return m.preview_transaction
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2Request) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
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
        err := writer.WriteStringValue("network", m.GetNetwork())
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
    {
        err := writer.WriteObjectValue("preview_transaction", m.GetPreviewTransaction())
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
func (m *TransactionPreviewV2Request) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
func (m *TransactionPreviewV2Request) SetAtLedgerState(value LedgerStateSelectorable)() {
    m.at_ledger_state = value
}
// SetFlags sets the flags property value. The flags property
func (m *TransactionPreviewV2Request) SetFlags(value PreviewFlagsable)() {
    m.flags = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionPreviewV2Request) SetNetwork(value *string)() {
    m.network = value
}
// SetOptions sets the options property value. A set of flags to configure the response of the transaction preview.
func (m *TransactionPreviewV2Request) SetOptions(value TransactionPreviewV2ResponseOptionsable)() {
    m.options = value
}
// SetPreviewTransaction sets the preview_transaction property value. The preview_transaction property
func (m *TransactionPreviewV2Request) SetPreviewTransaction(value PreviewTransactionable)() {
    m.preview_transaction = value
}
type TransactionPreviewV2Requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSelectorable)
    GetFlags()(PreviewFlagsable)
    GetNetwork()(*string)
    GetOptions()(TransactionPreviewV2ResponseOptionsable)
    GetPreviewTransaction()(PreviewTransactionable)
    SetAtLedgerState(value LedgerStateSelectorable)()
    SetFlags(value PreviewFlagsable)()
    SetNetwork(value *string)()
    SetOptions(value TransactionPreviewV2ResponseOptionsable)()
    SetPreviewTransaction(value PreviewTransactionable)()
}
