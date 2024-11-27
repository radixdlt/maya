package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewV2Response struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The at_ledger_state property
    at_ledger_state LedgerStateSummaryable
    // An optional field which is only provided if the `logs` flag is set to true in the`options` property of the request.If present, it gives the emitted logs from the transaction execution.
    logs []TransactionPreviewV2Response_logsable
    // An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true in the `options` property of the request.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
    radix_engine_toolkit_receipt TransactionPreviewV2Response_radix_engine_toolkit_receiptable
    // The transaction execution receipt
    receipt TransactionReceiptable
}
// NewTransactionPreviewV2Response instantiates a new TransactionPreviewV2Response and sets the default values.
func NewTransactionPreviewV2Response()(*TransactionPreviewV2Response) {
    m := &TransactionPreviewV2Response{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewV2Response(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2Response) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *TransactionPreviewV2Response) GetAtLedgerState()(LedgerStateSummaryable) {
    return m.at_ledger_state
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2Response) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["at_ledger_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtLedgerState(val.(LedgerStateSummaryable))
        }
        return nil
    }
    res["logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTransactionPreviewV2Response_logsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TransactionPreviewV2Response_logsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TransactionPreviewV2Response_logsable)
                }
            }
            m.SetLogs(res)
        }
        return nil
    }
    res["radix_engine_toolkit_receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionPreviewV2Response_radix_engine_toolkit_receiptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRadixEngineToolkitReceipt(val.(TransactionPreviewV2Response_radix_engine_toolkit_receiptable))
        }
        return nil
    }
    res["receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionReceiptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceipt(val.(TransactionReceiptable))
        }
        return nil
    }
    return res
}
// GetLogs gets the logs property value. An optional field which is only provided if the `logs` flag is set to true in the`options` property of the request.If present, it gives the emitted logs from the transaction execution.
// returns a []TransactionPreviewV2Response_logsable when successful
func (m *TransactionPreviewV2Response) GetLogs()([]TransactionPreviewV2Response_logsable) {
    return m.logs
}
// GetRadixEngineToolkitReceipt gets the radix_engine_toolkit_receipt property value. An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true in the `options` property of the request.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
// returns a TransactionPreviewV2Response_radix_engine_toolkit_receiptable when successful
func (m *TransactionPreviewV2Response) GetRadixEngineToolkitReceipt()(TransactionPreviewV2Response_radix_engine_toolkit_receiptable) {
    return m.radix_engine_toolkit_receipt
}
// GetReceipt gets the receipt property value. The transaction execution receipt
// returns a TransactionReceiptable when successful
func (m *TransactionPreviewV2Response) GetReceipt()(TransactionReceiptable) {
    return m.receipt
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2Response) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    if m.GetLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogs()))
        for i, v := range m.GetLogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("logs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("radix_engine_toolkit_receipt", m.GetRadixEngineToolkitReceipt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("receipt", m.GetReceipt())
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
func (m *TransactionPreviewV2Response) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *TransactionPreviewV2Response) SetAtLedgerState(value LedgerStateSummaryable)() {
    m.at_ledger_state = value
}
// SetLogs sets the logs property value. An optional field which is only provided if the `logs` flag is set to true in the`options` property of the request.If present, it gives the emitted logs from the transaction execution.
func (m *TransactionPreviewV2Response) SetLogs(value []TransactionPreviewV2Response_logsable)() {
    m.logs = value
}
// SetRadixEngineToolkitReceipt sets the radix_engine_toolkit_receipt property value. An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true in the `options` property of the request.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
func (m *TransactionPreviewV2Response) SetRadixEngineToolkitReceipt(value TransactionPreviewV2Response_radix_engine_toolkit_receiptable)() {
    m.radix_engine_toolkit_receipt = value
}
// SetReceipt sets the receipt property value. The transaction execution receipt
func (m *TransactionPreviewV2Response) SetReceipt(value TransactionReceiptable)() {
    m.receipt = value
}
type TransactionPreviewV2Responseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSummaryable)
    GetLogs()([]TransactionPreviewV2Response_logsable)
    GetRadixEngineToolkitReceipt()(TransactionPreviewV2Response_radix_engine_toolkit_receiptable)
    GetReceipt()(TransactionReceiptable)
    SetAtLedgerState(value LedgerStateSummaryable)()
    SetLogs(value []TransactionPreviewV2Response_logsable)()
    SetRadixEngineToolkitReceipt(value TransactionPreviewV2Response_radix_engine_toolkit_receiptable)()
    SetReceipt(value TransactionReceiptable)()
}
