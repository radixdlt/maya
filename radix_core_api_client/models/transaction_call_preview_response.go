package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionCallPreviewResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The at_ledger_state property
    at_ledger_state LedgerStateSummaryable
    // Error message (only present if status is Failed or Rejected)
    error_message *string
    // Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
    output SborDataable
    // The status of the transaction
    status *TransactionStatus
}
// NewTransactionCallPreviewResponse instantiates a new TransactionCallPreviewResponse and sets the default values.
func NewTransactionCallPreviewResponse()(*TransactionCallPreviewResponse) {
    m := &TransactionCallPreviewResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionCallPreviewResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionCallPreviewResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionCallPreviewResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionCallPreviewResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtLedgerState gets the at_ledger_state property value. The at_ledger_state property
// returns a LedgerStateSummaryable when successful
func (m *TransactionCallPreviewResponse) GetAtLedgerState()(LedgerStateSummaryable) {
    return m.at_ledger_state
}
// GetErrorMessage gets the error_message property value. Error message (only present if status is Failed or Rejected)
// returns a *string when successful
func (m *TransactionCallPreviewResponse) GetErrorMessage()(*string) {
    return m.error_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionCallPreviewResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["error_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["output"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutput(val.(SborDataable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TransactionStatus))
        }
        return nil
    }
    return res
}
// GetOutput gets the output property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *TransactionCallPreviewResponse) GetOutput()(SborDataable) {
    return m.output
}
// GetStatus gets the status property value. The status of the transaction
// returns a *TransactionStatus when successful
func (m *TransactionCallPreviewResponse) GetStatus()(*TransactionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *TransactionCallPreviewResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error_message", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("output", m.GetOutput())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *TransactionCallPreviewResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtLedgerState sets the at_ledger_state property value. The at_ledger_state property
func (m *TransactionCallPreviewResponse) SetAtLedgerState(value LedgerStateSummaryable)() {
    m.at_ledger_state = value
}
// SetErrorMessage sets the error_message property value. Error message (only present if status is Failed or Rejected)
func (m *TransactionCallPreviewResponse) SetErrorMessage(value *string)() {
    m.error_message = value
}
// SetOutput sets the output property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *TransactionCallPreviewResponse) SetOutput(value SborDataable)() {
    m.output = value
}
// SetStatus sets the status property value. The status of the transaction
func (m *TransactionCallPreviewResponse) SetStatus(value *TransactionStatus)() {
    m.status = value
}
type TransactionCallPreviewResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAtLedgerState()(LedgerStateSummaryable)
    GetErrorMessage()(*string)
    GetOutput()(SborDataable)
    GetStatus()(*TransactionStatus)
    SetAtLedgerState(value LedgerStateSummaryable)()
    SetErrorMessage(value *string)()
    SetOutput(value SborDataable)()
    SetStatus(value *TransactionStatus)()
}
