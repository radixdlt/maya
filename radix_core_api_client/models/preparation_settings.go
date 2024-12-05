package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreparationSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The max_blobs property
    max_blobs *string
    // The max_child_subintents_per_intent property
    max_child_subintents_per_intent *string
    // The max_ledger_payload_length property
    max_ledger_payload_length *string
    // The max_subintents_per_transaction property
    max_subintents_per_transaction *string
    // The max_user_payload_length property
    max_user_payload_length *string
    // The v2_transactions_permitted property
    v2_transactions_permitted *bool
}
// NewPreparationSettings instantiates a new PreparationSettings and sets the default values.
func NewPreparationSettings()(*PreparationSettings) {
    m := &PreparationSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreparationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreparationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreparationSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreparationSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreparationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_blobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxBlobs(val)
        }
        return nil
    }
    res["max_child_subintents_per_intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxChildSubintentsPerIntent(val)
        }
        return nil
    }
    res["max_ledger_payload_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLedgerPayloadLength(val)
        }
        return nil
    }
    res["max_subintents_per_transaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSubintentsPerTransaction(val)
        }
        return nil
    }
    res["max_user_payload_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxUserPayloadLength(val)
        }
        return nil
    }
    res["v2_transactions_permitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV2TransactionsPermitted(val)
        }
        return nil
    }
    return res
}
// GetMaxBlobs gets the max_blobs property value. The max_blobs property
// returns a *string when successful
func (m *PreparationSettings) GetMaxBlobs()(*string) {
    return m.max_blobs
}
// GetMaxChildSubintentsPerIntent gets the max_child_subintents_per_intent property value. The max_child_subintents_per_intent property
// returns a *string when successful
func (m *PreparationSettings) GetMaxChildSubintentsPerIntent()(*string) {
    return m.max_child_subintents_per_intent
}
// GetMaxLedgerPayloadLength gets the max_ledger_payload_length property value. The max_ledger_payload_length property
// returns a *string when successful
func (m *PreparationSettings) GetMaxLedgerPayloadLength()(*string) {
    return m.max_ledger_payload_length
}
// GetMaxSubintentsPerTransaction gets the max_subintents_per_transaction property value. The max_subintents_per_transaction property
// returns a *string when successful
func (m *PreparationSettings) GetMaxSubintentsPerTransaction()(*string) {
    return m.max_subintents_per_transaction
}
// GetMaxUserPayloadLength gets the max_user_payload_length property value. The max_user_payload_length property
// returns a *string when successful
func (m *PreparationSettings) GetMaxUserPayloadLength()(*string) {
    return m.max_user_payload_length
}
// GetV2TransactionsPermitted gets the v2_transactions_permitted property value. The v2_transactions_permitted property
// returns a *bool when successful
func (m *PreparationSettings) GetV2TransactionsPermitted()(*bool) {
    return m.v2_transactions_permitted
}
// Serialize serializes information the current object
func (m *PreparationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("max_blobs", m.GetMaxBlobs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_child_subintents_per_intent", m.GetMaxChildSubintentsPerIntent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_ledger_payload_length", m.GetMaxLedgerPayloadLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_subintents_per_transaction", m.GetMaxSubintentsPerTransaction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_user_payload_length", m.GetMaxUserPayloadLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v2_transactions_permitted", m.GetV2TransactionsPermitted())
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
func (m *PreparationSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxBlobs sets the max_blobs property value. The max_blobs property
func (m *PreparationSettings) SetMaxBlobs(value *string)() {
    m.max_blobs = value
}
// SetMaxChildSubintentsPerIntent sets the max_child_subintents_per_intent property value. The max_child_subintents_per_intent property
func (m *PreparationSettings) SetMaxChildSubintentsPerIntent(value *string)() {
    m.max_child_subintents_per_intent = value
}
// SetMaxLedgerPayloadLength sets the max_ledger_payload_length property value. The max_ledger_payload_length property
func (m *PreparationSettings) SetMaxLedgerPayloadLength(value *string)() {
    m.max_ledger_payload_length = value
}
// SetMaxSubintentsPerTransaction sets the max_subintents_per_transaction property value. The max_subintents_per_transaction property
func (m *PreparationSettings) SetMaxSubintentsPerTransaction(value *string)() {
    m.max_subintents_per_transaction = value
}
// SetMaxUserPayloadLength sets the max_user_payload_length property value. The max_user_payload_length property
func (m *PreparationSettings) SetMaxUserPayloadLength(value *string)() {
    m.max_user_payload_length = value
}
// SetV2TransactionsPermitted sets the v2_transactions_permitted property value. The v2_transactions_permitted property
func (m *PreparationSettings) SetV2TransactionsPermitted(value *bool)() {
    m.v2_transactions_permitted = value
}
type PreparationSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxBlobs()(*string)
    GetMaxChildSubintentsPerIntent()(*string)
    GetMaxLedgerPayloadLength()(*string)
    GetMaxSubintentsPerTransaction()(*string)
    GetMaxUserPayloadLength()(*string)
    GetV2TransactionsPermitted()(*bool)
    SetMaxBlobs(value *string)()
    SetMaxChildSubintentsPerIntent(value *string)()
    SetMaxLedgerPayloadLength(value *string)()
    SetMaxSubintentsPerTransaction(value *string)()
    SetMaxUserPayloadLength(value *string)()
    SetV2TransactionsPermitted(value *bool)()
}
