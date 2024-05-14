package models

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionSubmitErrorResponse struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A numeric code corresponding to the given HTTP error code.
	code *int32
	// The details property
	details LtsTransactionSubmitErrorDetailsable
	// The error_type property
	error_type *ErrorResponseType
	// A human-readable error message.
	message *string
	// A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
	trace_id *string
}

// NewLtsTransactionSubmitErrorResponse instantiates a new LtsTransactionSubmitErrorResponse and sets the default values.
func NewLtsTransactionSubmitErrorResponse() *LtsTransactionSubmitErrorResponse {
	m := &LtsTransactionSubmitErrorResponse{
		ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
	}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsTransactionSubmitErrorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionSubmitErrorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionSubmitErrorResponse(), nil
}

// Error the primary error message.
// returns a string when successful
func (m *LtsTransactionSubmitErrorResponse) Error() string {
	return m.ApiError.Error()
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsTransactionSubmitErrorResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCode gets the code property value. A numeric code corresponding to the given HTTP error code.
// returns a *int32 when successful
func (m *LtsTransactionSubmitErrorResponse) GetCode() *int32 {
	return m.code
}

// GetDetails gets the details property value. The details property
// returns a LtsTransactionSubmitErrorDetailsable when successful
func (m *LtsTransactionSubmitErrorResponse) GetDetails() LtsTransactionSubmitErrorDetailsable {
	return m.details
}

// GetErrorType gets the error_type property value. The error_type property
// returns a *ErrorResponseType when successful
func (m *LtsTransactionSubmitErrorResponse) GetErrorType() *ErrorResponseType {
	return m.error_type
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionSubmitErrorResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCode(val)
		}
		return nil
	}
	res["details"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLtsTransactionSubmitErrorDetailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDetails(val.(LtsTransactionSubmitErrorDetailsable))
		}
		return nil
	}
	res["error_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseErrorResponseType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrorType(val.(*ErrorResponseType))
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val)
		}
		return nil
	}
	res["trace_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *LtsTransactionSubmitErrorResponse) GetMessage() *string {
	return m.message
}

// GetTraceId gets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
// returns a *string when successful
func (m *LtsTransactionSubmitErrorResponse) GetTraceId() *string {
	return m.trace_id
}

// Serialize serializes information the current object
func (m *LtsTransactionSubmitErrorResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *LtsTransactionSubmitErrorResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCode sets the code property value. A numeric code corresponding to the given HTTP error code.
func (m *LtsTransactionSubmitErrorResponse) SetCode(value *int32) {
	m.code = value
}

// SetDetails sets the details property value. The details property
func (m *LtsTransactionSubmitErrorResponse) SetDetails(value LtsTransactionSubmitErrorDetailsable) {
	m.details = value
}

// SetErrorType sets the error_type property value. The error_type property
func (m *LtsTransactionSubmitErrorResponse) SetErrorType(value *ErrorResponseType) {
	m.error_type = value
}

// SetMessage sets the message property value. A human-readable error message.
func (m *LtsTransactionSubmitErrorResponse) SetMessage(value *string) {
	m.message = value
}

// SetTraceId sets the trace_id property value. A GUID to be used when reporting errors, to allow correlation with the Core API's error logs, in the case where the Core API details are hidden.
func (m *LtsTransactionSubmitErrorResponse) SetTraceId(value *string) {
	m.trace_id = value
}

type LtsTransactionSubmitErrorResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCode() *int32
	GetDetails() LtsTransactionSubmitErrorDetailsable
	GetErrorType() *ErrorResponseType
	GetMessage() *string
	GetTraceId() *string
	SetCode(value *int32)
	SetDetails(value LtsTransactionSubmitErrorDetailsable)
	SetErrorType(value *ErrorResponseType)
	SetMessage(value *string)
	SetTraceId(value *string)
}
