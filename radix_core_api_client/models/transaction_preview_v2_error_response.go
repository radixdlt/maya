package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewV2ErrorResponse struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A numeric code corresponding to the given HTTP error code.
    code *int32
    // The details property
    details TransactionPreviewV2ErrorDetailsable
    // The error_type property
    error_type *ErrorResponseType
    // A human-readable error message.
    message *string
    // A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
    trace_id *string
}
// NewTransactionPreviewV2ErrorResponse instantiates a new TransactionPreviewV2ErrorResponse and sets the default values.
func NewTransactionPreviewV2ErrorResponse()(*TransactionPreviewV2ErrorResponse) {
    m := &TransactionPreviewV2ErrorResponse{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2ErrorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2ErrorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionPreviewV2ErrorResponse(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *TransactionPreviewV2ErrorResponse) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2ErrorResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. A numeric code corresponding to the given HTTP error code.
// returns a *int32 when successful
func (m *TransactionPreviewV2ErrorResponse) GetCode()(*int32) {
    return m.code
}
// GetDetails gets the details property value. The details property
// returns a TransactionPreviewV2ErrorDetailsable when successful
func (m *TransactionPreviewV2ErrorResponse) GetDetails()(TransactionPreviewV2ErrorDetailsable) {
    return m.details
}
// GetErrorType gets the error_type property value. The error_type property
// returns a *ErrorResponseType when successful
func (m *TransactionPreviewV2ErrorResponse) GetErrorType()(*ErrorResponseType) {
    return m.error_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2ErrorResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionPreviewV2ErrorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(TransactionPreviewV2ErrorDetailsable))
        }
        return nil
    }
    res["error_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseErrorResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorType(val.(*ErrorResponseType))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["trace_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTraceId(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. A human-readable error message.
// returns a *string when successful
func (m *TransactionPreviewV2ErrorResponse) GetMessage()(*string) {
    return m.message
}
// GetTraceId gets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
// returns a *string when successful
func (m *TransactionPreviewV2ErrorResponse) GetTraceId()(*string) {
    return m.trace_id
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2ErrorResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    if m.GetErrorType() != nil {
        cast := (*m.GetErrorType()).String()
        err := writer.WriteStringValue("error_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trace_id", m.GetTraceId())
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
func (m *TransactionPreviewV2ErrorResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. A numeric code corresponding to the given HTTP error code.
func (m *TransactionPreviewV2ErrorResponse) SetCode(value *int32)() {
    m.code = value
}
// SetDetails sets the details property value. The details property
func (m *TransactionPreviewV2ErrorResponse) SetDetails(value TransactionPreviewV2ErrorDetailsable)() {
    m.details = value
}
// SetErrorType sets the error_type property value. The error_type property
func (m *TransactionPreviewV2ErrorResponse) SetErrorType(value *ErrorResponseType)() {
    m.error_type = value
}
// SetMessage sets the message property value. A human-readable error message.
func (m *TransactionPreviewV2ErrorResponse) SetMessage(value *string)() {
    m.message = value
}
// SetTraceId sets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
func (m *TransactionPreviewV2ErrorResponse) SetTraceId(value *string)() {
    m.trace_id = value
}
type TransactionPreviewV2ErrorResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*int32)
    GetDetails()(TransactionPreviewV2ErrorDetailsable)
    GetErrorType()(*ErrorResponseType)
    GetMessage()(*string)
    GetTraceId()(*string)
    SetCode(value *int32)()
    SetDetails(value TransactionPreviewV2ErrorDetailsable)()
    SetErrorType(value *ErrorResponseType)()
    SetMessage(value *string)()
    SetTraceId(value *string)()
}
