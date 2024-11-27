package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LimitParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcall depth allowed during transaction execution.
    max_call_depth *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single emitted event.
    max_event_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept on the heap during a single transaction's execution.
    max_heap_substate_total_bytes *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single call's input parameters.
    max_invoke_input_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single logged line.
    max_log_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of events emitted during a single transaction's execution.
    max_number_of_events *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of log lines emitted during a single transaction's execution.
    max_number_of_logs *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single panic message.
    max_panic_message_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's key in the low-level Substate database.
    max_substate_key_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's value in the low-level Substate database.
    max_substate_value_size *string
    // A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept in the track during a single transaction's execution.
    max_track_substate_total_bytes *string
}
// NewLimitParameters instantiates a new LimitParameters and sets the default values.
func NewLimitParameters()(*LimitParameters) {
    m := &LimitParameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLimitParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLimitParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLimitParameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LimitParameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LimitParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_call_depth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCallDepth(val)
        }
        return nil
    }
    res["max_event_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxEventSize(val)
        }
        return nil
    }
    res["max_heap_substate_total_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxHeapSubstateTotalBytes(val)
        }
        return nil
    }
    res["max_invoke_input_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxInvokeInputSize(val)
        }
        return nil
    }
    res["max_log_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLogSize(val)
        }
        return nil
    }
    res["max_number_of_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxNumberOfEvents(val)
        }
        return nil
    }
    res["max_number_of_logs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxNumberOfLogs(val)
        }
        return nil
    }
    res["max_panic_message_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPanicMessageSize(val)
        }
        return nil
    }
    res["max_substate_key_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSubstateKeySize(val)
        }
        return nil
    }
    res["max_substate_value_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSubstateValueSize(val)
        }
        return nil
    }
    res["max_track_substate_total_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTrackSubstateTotalBytes(val)
        }
        return nil
    }
    return res
}
// GetMaxCallDepth gets the max_call_depth property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcall depth allowed during transaction execution.
// returns a *string when successful
func (m *LimitParameters) GetMaxCallDepth()(*string) {
    return m.max_call_depth
}
// GetMaxEventSize gets the max_event_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single emitted event.
// returns a *string when successful
func (m *LimitParameters) GetMaxEventSize()(*string) {
    return m.max_event_size
}
// GetMaxHeapSubstateTotalBytes gets the max_heap_substate_total_bytes property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept on the heap during a single transaction's execution.
// returns a *string when successful
func (m *LimitParameters) GetMaxHeapSubstateTotalBytes()(*string) {
    return m.max_heap_substate_total_bytes
}
// GetMaxInvokeInputSize gets the max_invoke_input_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single call's input parameters.
// returns a *string when successful
func (m *LimitParameters) GetMaxInvokeInputSize()(*string) {
    return m.max_invoke_input_size
}
// GetMaxLogSize gets the max_log_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single logged line.
// returns a *string when successful
func (m *LimitParameters) GetMaxLogSize()(*string) {
    return m.max_log_size
}
// GetMaxNumberOfEvents gets the max_number_of_events property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of events emitted during a single transaction's execution.
// returns a *string when successful
func (m *LimitParameters) GetMaxNumberOfEvents()(*string) {
    return m.max_number_of_events
}
// GetMaxNumberOfLogs gets the max_number_of_logs property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of log lines emitted during a single transaction's execution.
// returns a *string when successful
func (m *LimitParameters) GetMaxNumberOfLogs()(*string) {
    return m.max_number_of_logs
}
// GetMaxPanicMessageSize gets the max_panic_message_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single panic message.
// returns a *string when successful
func (m *LimitParameters) GetMaxPanicMessageSize()(*string) {
    return m.max_panic_message_size
}
// GetMaxSubstateKeySize gets the max_substate_key_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's key in the low-level Substate database.
// returns a *string when successful
func (m *LimitParameters) GetMaxSubstateKeySize()(*string) {
    return m.max_substate_key_size
}
// GetMaxSubstateValueSize gets the max_substate_value_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's value in the low-level Substate database.
// returns a *string when successful
func (m *LimitParameters) GetMaxSubstateValueSize()(*string) {
    return m.max_substate_value_size
}
// GetMaxTrackSubstateTotalBytes gets the max_track_substate_total_bytes property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept in the track during a single transaction's execution.
// returns a *string when successful
func (m *LimitParameters) GetMaxTrackSubstateTotalBytes()(*string) {
    return m.max_track_substate_total_bytes
}
// Serialize serializes information the current object
func (m *LimitParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("max_call_depth", m.GetMaxCallDepth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_event_size", m.GetMaxEventSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_heap_substate_total_bytes", m.GetMaxHeapSubstateTotalBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_invoke_input_size", m.GetMaxInvokeInputSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_log_size", m.GetMaxLogSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_number_of_events", m.GetMaxNumberOfEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_number_of_logs", m.GetMaxNumberOfLogs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_panic_message_size", m.GetMaxPanicMessageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_substate_key_size", m.GetMaxSubstateKeySize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_substate_value_size", m.GetMaxSubstateValueSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_track_substate_total_bytes", m.GetMaxTrackSubstateTotalBytes())
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
func (m *LimitParameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxCallDepth sets the max_call_depth property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcall depth allowed during transaction execution.
func (m *LimitParameters) SetMaxCallDepth(value *string)() {
    m.max_call_depth = value
}
// SetMaxEventSize sets the max_event_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single emitted event.
func (m *LimitParameters) SetMaxEventSize(value *string)() {
    m.max_event_size = value
}
// SetMaxHeapSubstateTotalBytes sets the max_heap_substate_total_bytes property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept on the heap during a single transaction's execution.
func (m *LimitParameters) SetMaxHeapSubstateTotalBytes(value *string)() {
    m.max_heap_substate_total_bytes = value
}
// SetMaxInvokeInputSize sets the max_invoke_input_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single call's input parameters.
func (m *LimitParameters) SetMaxInvokeInputSize(value *string)() {
    m.max_invoke_input_size = value
}
// SetMaxLogSize sets the max_log_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single logged line.
func (m *LimitParameters) SetMaxLogSize(value *string)() {
    m.max_log_size = value
}
// SetMaxNumberOfEvents sets the max_number_of_events property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of events emitted during a single transaction's execution.
func (m *LimitParameters) SetMaxNumberOfEvents(value *string)() {
    m.max_number_of_events = value
}
// SetMaxNumberOfLogs sets the max_number_of_logs property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumcount of log lines emitted during a single transaction's execution.
func (m *LimitParameters) SetMaxNumberOfLogs(value *string)() {
    m.max_number_of_logs = value
}
// SetMaxPanicMessageSize sets the max_panic_message_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a single panic message.
func (m *LimitParameters) SetMaxPanicMessageSize(value *string)() {
    m.max_panic_message_size = value
}
// SetMaxSubstateKeySize sets the max_substate_key_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's key in the low-level Substate database.
func (m *LimitParameters) SetMaxSubstateKeySize(value *string)() {
    m.max_substate_key_size = value
}
// SetMaxSubstateValueSize sets the max_substate_value_size property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of a Substate's value in the low-level Substate database.
func (m *LimitParameters) SetMaxSubstateValueSize(value *string)() {
    m.max_substate_value_size = value
}
// SetMaxTrackSubstateTotalBytes sets the max_track_substate_total_bytes property value. A decimal string-encoded 64-bit unsigned integer, representing the configured maximumbyte size of all substates kept in the track during a single transaction's execution.
func (m *LimitParameters) SetMaxTrackSubstateTotalBytes(value *string)() {
    m.max_track_substate_total_bytes = value
}
type LimitParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxCallDepth()(*string)
    GetMaxEventSize()(*string)
    GetMaxHeapSubstateTotalBytes()(*string)
    GetMaxInvokeInputSize()(*string)
    GetMaxLogSize()(*string)
    GetMaxNumberOfEvents()(*string)
    GetMaxNumberOfLogs()(*string)
    GetMaxPanicMessageSize()(*string)
    GetMaxSubstateKeySize()(*string)
    GetMaxSubstateValueSize()(*string)
    GetMaxTrackSubstateTotalBytes()(*string)
    SetMaxCallDepth(value *string)()
    SetMaxEventSize(value *string)()
    SetMaxHeapSubstateTotalBytes(value *string)()
    SetMaxInvokeInputSize(value *string)()
    SetMaxLogSize(value *string)()
    SetMaxNumberOfEvents(value *string)()
    SetMaxNumberOfLogs(value *string)()
    SetMaxPanicMessageSize(value *string)()
    SetMaxSubstateKeySize(value *string)()
    SetMaxSubstateValueSize(value *string)()
    SetMaxTrackSubstateTotalBytes(value *string)()
}
