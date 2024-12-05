package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionReceipt the transaction execution receipt
type TransactionReceipt struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The costing_parameters property
    costing_parameters CostingParametersable
    // The error message. This property is only present if the status is `Failed` or `Rejected`.
    error_message *string
    // The events property
    events []Eventable
    // The fee_destination property
    fee_destination FeeDestinationable
    // The fee_source property
    fee_source FeeSourceable
    // The fee_summary property
    fee_summary FeeSummaryable
    // The next_epoch property
    next_epoch NextEpochable
    // The return data for each line of the transaction intent's manifest.This property is only present if the `status` is `Succeeded`.
    output []SborDataable
    // Transaction state updates (only present if status is Succeeded or Failed)
    state_updates StateUpdatesable
    // The status of the transaction
    status *TransactionStatus
}
// NewTransactionReceipt instantiates a new TransactionReceipt and sets the default values.
func NewTransactionReceipt()(*TransactionReceipt) {
    m := &TransactionReceipt{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionReceiptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionReceiptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionReceipt(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionReceipt) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCostingParameters gets the costing_parameters property value. The costing_parameters property
// returns a CostingParametersable when successful
func (m *TransactionReceipt) GetCostingParameters()(CostingParametersable) {
    return m.costing_parameters
}
// GetErrorMessage gets the error_message property value. The error message. This property is only present if the status is `Failed` or `Rejected`.
// returns a *string when successful
func (m *TransactionReceipt) GetErrorMessage()(*string) {
    return m.error_message
}
// GetEvents gets the events property value. The events property
// returns a []Eventable when successful
func (m *TransactionReceipt) GetEvents()([]Eventable) {
    return m.events
}
// GetFeeDestination gets the fee_destination property value. The fee_destination property
// returns a FeeDestinationable when successful
func (m *TransactionReceipt) GetFeeDestination()(FeeDestinationable) {
    return m.fee_destination
}
// GetFeeSource gets the fee_source property value. The fee_source property
// returns a FeeSourceable when successful
func (m *TransactionReceipt) GetFeeSource()(FeeSourceable) {
    return m.fee_source
}
// GetFeeSummary gets the fee_summary property value. The fee_summary property
// returns a FeeSummaryable when successful
func (m *TransactionReceipt) GetFeeSummary()(FeeSummaryable) {
    return m.fee_summary
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionReceipt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["costing_parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCostingParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCostingParameters(val.(CostingParametersable))
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
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Eventable)
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["fee_destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeeDestinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeeDestination(val.(FeeDestinationable))
        }
        return nil
    }
    res["fee_source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeeSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeeSource(val.(FeeSourceable))
        }
        return nil
    }
    res["fee_summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeeSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeeSummary(val.(FeeSummaryable))
        }
        return nil
    }
    res["next_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNextEpochFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextEpoch(val.(NextEpochable))
        }
        return nil
    }
    res["output"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSborDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SborDataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SborDataable)
                }
            }
            m.SetOutput(res)
        }
        return nil
    }
    res["state_updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStateUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateUpdates(val.(StateUpdatesable))
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
// GetNextEpoch gets the next_epoch property value. The next_epoch property
// returns a NextEpochable when successful
func (m *TransactionReceipt) GetNextEpoch()(NextEpochable) {
    return m.next_epoch
}
// GetOutput gets the output property value. The return data for each line of the transaction intent's manifest.This property is only present if the `status` is `Succeeded`.
// returns a []SborDataable when successful
func (m *TransactionReceipt) GetOutput()([]SborDataable) {
    return m.output
}
// GetStateUpdates gets the state_updates property value. Transaction state updates (only present if status is Succeeded or Failed)
// returns a StateUpdatesable when successful
func (m *TransactionReceipt) GetStateUpdates()(StateUpdatesable) {
    return m.state_updates
}
// GetStatus gets the status property value. The status of the transaction
// returns a *TransactionStatus when successful
func (m *TransactionReceipt) GetStatus()(*TransactionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *TransactionReceipt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("costing_parameters", m.GetCostingParameters())
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
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fee_destination", m.GetFeeDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fee_source", m.GetFeeSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fee_summary", m.GetFeeSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("next_epoch", m.GetNextEpoch())
        if err != nil {
            return err
        }
    }
    if m.GetOutput() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOutput()))
        for i, v := range m.GetOutput() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("output", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("state_updates", m.GetStateUpdates())
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
func (m *TransactionReceipt) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCostingParameters sets the costing_parameters property value. The costing_parameters property
func (m *TransactionReceipt) SetCostingParameters(value CostingParametersable)() {
    m.costing_parameters = value
}
// SetErrorMessage sets the error_message property value. The error message. This property is only present if the status is `Failed` or `Rejected`.
func (m *TransactionReceipt) SetErrorMessage(value *string)() {
    m.error_message = value
}
// SetEvents sets the events property value. The events property
func (m *TransactionReceipt) SetEvents(value []Eventable)() {
    m.events = value
}
// SetFeeDestination sets the fee_destination property value. The fee_destination property
func (m *TransactionReceipt) SetFeeDestination(value FeeDestinationable)() {
    m.fee_destination = value
}
// SetFeeSource sets the fee_source property value. The fee_source property
func (m *TransactionReceipt) SetFeeSource(value FeeSourceable)() {
    m.fee_source = value
}
// SetFeeSummary sets the fee_summary property value. The fee_summary property
func (m *TransactionReceipt) SetFeeSummary(value FeeSummaryable)() {
    m.fee_summary = value
}
// SetNextEpoch sets the next_epoch property value. The next_epoch property
func (m *TransactionReceipt) SetNextEpoch(value NextEpochable)() {
    m.next_epoch = value
}
// SetOutput sets the output property value. The return data for each line of the transaction intent's manifest.This property is only present if the `status` is `Succeeded`.
func (m *TransactionReceipt) SetOutput(value []SborDataable)() {
    m.output = value
}
// SetStateUpdates sets the state_updates property value. Transaction state updates (only present if status is Succeeded or Failed)
func (m *TransactionReceipt) SetStateUpdates(value StateUpdatesable)() {
    m.state_updates = value
}
// SetStatus sets the status property value. The status of the transaction
func (m *TransactionReceipt) SetStatus(value *TransactionStatus)() {
    m.status = value
}
type TransactionReceiptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCostingParameters()(CostingParametersable)
    GetErrorMessage()(*string)
    GetEvents()([]Eventable)
    GetFeeDestination()(FeeDestinationable)
    GetFeeSource()(FeeSourceable)
    GetFeeSummary()(FeeSummaryable)
    GetNextEpoch()(NextEpochable)
    GetOutput()([]SborDataable)
    GetStateUpdates()(StateUpdatesable)
    GetStatus()(*TransactionStatus)
    SetCostingParameters(value CostingParametersable)()
    SetErrorMessage(value *string)()
    SetEvents(value []Eventable)()
    SetFeeDestination(value FeeDestinationable)()
    SetFeeSource(value FeeSourceable)()
    SetFeeSummary(value FeeSummaryable)()
    SetNextEpoch(value NextEpochable)()
    SetOutput(value []SborDataable)()
    SetStateUpdates(value StateUpdatesable)()
    SetStatus(value *TransactionStatus)()
}
