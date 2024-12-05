package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ErrorResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A numeric code corresponding to the given HTTP error code.
    code *int32
    // The error_type property
    error_type *ErrorResponseType
    // A human-readable error message.
    message *string
    // A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
    trace_id *string
}
// NewErrorResponse instantiates a new ErrorResponse and sets the default values.
func NewErrorResponse()(*ErrorResponse) {
    m := &ErrorResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateErrorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateErrorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("error_type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "Basic":
                        return NewBasicErrorResponse(), nil
                    case "LtsTransactionSubmit":
                        return NewLtsTransactionSubmitErrorResponse(), nil
                    case "StreamProofs":
                        return NewStreamProofsErrorResponse(), nil
                    case "StreamTransactions":
                        return NewStreamTransactionsErrorResponse(), nil
                    case "TransactionPreviewV2":
                        return NewTransactionPreviewV2ErrorResponse(), nil
                    case "TransactionSubmit":
                        return NewTransactionSubmitErrorResponse(), nil
                }
            }
        }
    }
    return NewErrorResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ErrorResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. A numeric code corresponding to the given HTTP error code.
// returns a *int32 when successful
func (m *ErrorResponse) GetCode()(*int32) {
    return m.code
}
// GetErrorType gets the error_type property value. The error_type property
// returns a *ErrorResponseType when successful
func (m *ErrorResponse) GetErrorType()(*ErrorResponseType) {
    return m.error_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ErrorResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ErrorResponse) GetMessage()(*string) {
    return m.message
}
// GetTraceId gets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
// returns a *string when successful
func (m *ErrorResponse) GetTraceId()(*string) {
    return m.trace_id
}
// Serialize serializes information the current object
func (m *ErrorResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
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
func (m *ErrorResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. A numeric code corresponding to the given HTTP error code.
func (m *ErrorResponse) SetCode(value *int32)() {
    m.code = value
}
// SetErrorType sets the error_type property value. The error_type property
func (m *ErrorResponse) SetErrorType(value *ErrorResponseType)() {
    m.error_type = value
}
// SetMessage sets the message property value. A human-readable error message.
func (m *ErrorResponse) SetMessage(value *string)() {
    m.message = value
}
// SetTraceId sets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
func (m *ErrorResponse) SetTraceId(value *string)() {
    m.trace_id = value
}
type ErrorResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*int32)
    GetErrorType()(*ErrorResponseType)
    GetMessage()(*string)
    GetTraceId()(*string)
    SetCode(value *int32)()
    SetErrorType(value *ErrorResponseType)()
    SetMessage(value *string)()
    SetTraceId(value *string)()
}
