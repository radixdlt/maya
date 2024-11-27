package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The at_ledger_state property
    at_ledger_state LedgerStateSummaryable
    // The hex-sbor-encoded receipt.This field is deprecated and will be removed from the API with the release of the next protocol update, cuttlefish. This field was provided primarily for use with the Radix Engine Toolkit and its execution summary functionality. If you still wish to use this functionality update your Radix Engine Toolkit and use the receipt provided in the `radix_engine_toolkit_receipt` field of this response.
    // Deprecated: 
    encoded_receipt *string
    // The instruction_resource_changes property
    instruction_resource_changes []InstructionResourceChangesable
    // The logs property
    logs []TransactionPreviewResponse_logsable
    // An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true when requesting a transaction preview from the API.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
    radix_engine_toolkit_receipt TransactionPreviewResponse_radix_engine_toolkit_receiptable
    // The transaction execution receipt
    receipt TransactionReceiptable
}
// NewTransactionPreviewResponse instantiates a new TransactionPreviewResponse and sets the default values.
func NewTransactionPreviewResponse()(*TransactionPreviewResponse) {
    m := &TransactionPreviewResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *TransactionPreviewResponse) GetAtLedgerState()(LedgerStateSummaryable) {
    return m.at_ledger_state
}
// GetEncodedReceipt gets the encoded_receipt property value. The hex-sbor-encoded receipt.This field is deprecated and will be removed from the API with the release of the next protocol update, cuttlefish. This field was provided primarily for use with the Radix Engine Toolkit and its execution summary functionality. If you still wish to use this functionality update your Radix Engine Toolkit and use the receipt provided in the `radix_engine_toolkit_receipt` field of this response.
// Deprecated: 
// returns a *string when successful
func (m *TransactionPreviewResponse) GetEncodedReceipt()(*string) {
    return m.encoded_receipt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["encoded_receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncodedReceipt(val)
        }
        return nil
    }
    res["instruction_resource_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInstructionResourceChangesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InstructionResourceChangesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(InstructionResourceChangesable)
                }
            }
            m.SetInstructionResourceChanges(res)
        }
        return nil
    }
    res["logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTransactionPreviewResponse_logsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TransactionPreviewResponse_logsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TransactionPreviewResponse_logsable)
                }
            }
            m.SetLogs(res)
        }
        return nil
    }
    res["radix_engine_toolkit_receipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionPreviewResponse_radix_engine_toolkit_receiptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRadixEngineToolkitReceipt(val.(TransactionPreviewResponse_radix_engine_toolkit_receiptable))
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
// GetInstructionResourceChanges gets the instruction_resource_changes property value. The instruction_resource_changes property
// returns a []InstructionResourceChangesable when successful
func (m *TransactionPreviewResponse) GetInstructionResourceChanges()([]InstructionResourceChangesable) {
    return m.instruction_resource_changes
}
// GetLogs gets the logs property value. The logs property
// returns a []TransactionPreviewResponse_logsable when successful
func (m *TransactionPreviewResponse) GetLogs()([]TransactionPreviewResponse_logsable) {
    return m.logs
}
// GetRadixEngineToolkitReceipt gets the radix_engine_toolkit_receipt property value. An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true when requesting a transaction preview from the API.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
// returns a TransactionPreviewResponse_radix_engine_toolkit_receiptable when successful
func (m *TransactionPreviewResponse) GetRadixEngineToolkitReceipt()(TransactionPreviewResponse_radix_engine_toolkit_receiptable) {
    return m.radix_engine_toolkit_receipt
}
// GetReceipt gets the receipt property value. The transaction execution receipt
// returns a TransactionReceiptable when successful
func (m *TransactionPreviewResponse) GetReceipt()(TransactionReceiptable) {
    return m.receipt
}
// Serialize serializes information the current object
func (m *TransactionPreviewResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encoded_receipt", m.GetEncodedReceipt())
        if err != nil {
            return err
        }
    }
    if m.GetInstructionResourceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstructionResourceChanges()))
        for i, v := range m.GetInstructionResourceChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("instruction_resource_changes", cast)
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
func (m *TransactionPreviewResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *TransactionPreviewResponse) SetAtLedgerState(value LedgerStateSummaryable)() {
    m.at_ledger_state = value
}
// SetEncodedReceipt sets the encoded_receipt property value. The hex-sbor-encoded receipt.This field is deprecated and will be removed from the API with the release of the next protocol update, cuttlefish. This field was provided primarily for use with the Radix Engine Toolkit and its execution summary functionality. If you still wish to use this functionality update your Radix Engine Toolkit and use the receipt provided in the `radix_engine_toolkit_receipt` field of this response.
// Deprecated: 
func (m *TransactionPreviewResponse) SetEncodedReceipt(value *string)() {
    m.encoded_receipt = value
}
// SetInstructionResourceChanges sets the instruction_resource_changes property value. The instruction_resource_changes property
func (m *TransactionPreviewResponse) SetInstructionResourceChanges(value []InstructionResourceChangesable)() {
    m.instruction_resource_changes = value
}
// SetLogs sets the logs property value. The logs property
func (m *TransactionPreviewResponse) SetLogs(value []TransactionPreviewResponse_logsable)() {
    m.logs = value
}
// SetRadixEngineToolkitReceipt sets the radix_engine_toolkit_receipt property value. An optional field which is only provided if the `radix_engine_toolkit_receipt`flag is set to true when requesting a transaction preview from the API.This receipt is primarily intended for use with the toolkit and may contain information that is already available in the receipt provided in the `receipt` field of this response.A typical client of this API is not expected to use this receipt. The primary clients this receipt is intended for is the Radix wallet or any client that needs to perform execution summaries on their transactions.
func (m *TransactionPreviewResponse) SetRadixEngineToolkitReceipt(value TransactionPreviewResponse_radix_engine_toolkit_receiptable)() {
    m.radix_engine_toolkit_receipt = value
}
// SetReceipt sets the receipt property value. The transaction execution receipt
func (m *TransactionPreviewResponse) SetReceipt(value TransactionReceiptable)() {
    m.receipt = value
}
type TransactionPreviewResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSummaryable)
    GetEncodedReceipt()(*string)
    GetInstructionResourceChanges()([]InstructionResourceChangesable)
    GetLogs()([]TransactionPreviewResponse_logsable)
    GetRadixEngineToolkitReceipt()(TransactionPreviewResponse_radix_engine_toolkit_receiptable)
    GetReceipt()(TransactionReceiptable)
    SetAtLedgerState(value LedgerStateSummaryable)()
    SetEncodedReceipt(value *string)()
    SetInstructionResourceChanges(value []InstructionResourceChangesable)()
    SetLogs(value []TransactionPreviewResponse_logsable)()
    SetRadixEngineToolkitReceipt(value TransactionPreviewResponse_radix_engine_toolkit_receiptable)()
    SetReceipt(value TransactionReceiptable)()
}
