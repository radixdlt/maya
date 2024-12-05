package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionParseRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The logical name of the network
    network *string
    // The type of transaction payload that should be assumed. If omitted, "Any" is used - where the payload isattempted to be parsed as each of the following in turn: Notarized, Signed, Unsigned, Ledger.
    parse_mode *TransactionParseRequest_parse_mode
    // A hex-encoded payload of a full transaction or a partial transaction - either a notarized transaction,a signed transaction intent an unsigned transaction intent, or a ledger payload.
    payload_hex *string
    // The amount of information to return in the response."Basic" includes the type, validity information, and any relevant identifiers."Full" also includes the fully parsed information.If omitted, "Full" is used.
    response_mode *TransactionParseRequest_response_mode
    // Requested transaction formats to include in the response
    transaction_format_options TransactionFormatOptionsable
    // The type of validation that should be performed, if the payload correctly decompiles as a Notarized Transaction.This is only relevant for Notarized payloads. If omitted, "Static" is used.
    validation_mode *TransactionParseRequest_validation_mode
}
// NewTransactionParseRequest instantiates a new TransactionParseRequest and sets the default values.
func NewTransactionParseRequest()(*TransactionParseRequest) {
    m := &TransactionParseRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionParseRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionParseRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionParseRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionParseRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionParseRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["parse_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionParseRequest_parse_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParseMode(val.(*TransactionParseRequest_parse_mode))
        }
        return nil
    }
    res["payload_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadHex(val)
        }
        return nil
    }
    res["response_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionParseRequest_response_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseMode(val.(*TransactionParseRequest_response_mode))
        }
        return nil
    }
    res["transaction_format_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionFormatOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionFormatOptions(val.(TransactionFormatOptionsable))
        }
        return nil
    }
    res["validation_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionParseRequest_validation_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationMode(val.(*TransactionParseRequest_validation_mode))
        }
        return nil
    }
    return res
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionParseRequest) GetNetwork()(*string) {
    return m.network
}
// GetParseMode gets the parse_mode property value. The type of transaction payload that should be assumed. If omitted, "Any" is used - where the payload isattempted to be parsed as each of the following in turn: Notarized, Signed, Unsigned, Ledger.
// returns a *TransactionParseRequest_parse_mode when successful
func (m *TransactionParseRequest) GetParseMode()(*TransactionParseRequest_parse_mode) {
    return m.parse_mode
}
// GetPayloadHex gets the payload_hex property value. A hex-encoded payload of a full transaction or a partial transaction - either a notarized transaction,a signed transaction intent an unsigned transaction intent, or a ledger payload.
// returns a *string when successful
func (m *TransactionParseRequest) GetPayloadHex()(*string) {
    return m.payload_hex
}
// GetResponseMode gets the response_mode property value. The amount of information to return in the response."Basic" includes the type, validity information, and any relevant identifiers."Full" also includes the fully parsed information.If omitted, "Full" is used.
// returns a *TransactionParseRequest_response_mode when successful
func (m *TransactionParseRequest) GetResponseMode()(*TransactionParseRequest_response_mode) {
    return m.response_mode
}
// GetTransactionFormatOptions gets the transaction_format_options property value. Requested transaction formats to include in the response
// returns a TransactionFormatOptionsable when successful
func (m *TransactionParseRequest) GetTransactionFormatOptions()(TransactionFormatOptionsable) {
    return m.transaction_format_options
}
// GetValidationMode gets the validation_mode property value. The type of validation that should be performed, if the payload correctly decompiles as a Notarized Transaction.This is only relevant for Notarized payloads. If omitted, "Static" is used.
// returns a *TransactionParseRequest_validation_mode when successful
func (m *TransactionParseRequest) GetValidationMode()(*TransactionParseRequest_validation_mode) {
    return m.validation_mode
}
// Serialize serializes information the current object
func (m *TransactionParseRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetParseMode() != nil {
        cast := (*m.GetParseMode()).String()
        err := writer.WriteStringValue("parse_mode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payload_hex", m.GetPayloadHex())
        if err != nil {
            return err
        }
    }
    if m.GetResponseMode() != nil {
        cast := (*m.GetResponseMode()).String()
        err := writer.WriteStringValue("response_mode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transaction_format_options", m.GetTransactionFormatOptions())
        if err != nil {
            return err
        }
    }
    if m.GetValidationMode() != nil {
        cast := (*m.GetValidationMode()).String()
        err := writer.WriteStringValue("validation_mode", &cast)
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
func (m *TransactionParseRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionParseRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetParseMode sets the parse_mode property value. The type of transaction payload that should be assumed. If omitted, "Any" is used - where the payload isattempted to be parsed as each of the following in turn: Notarized, Signed, Unsigned, Ledger.
func (m *TransactionParseRequest) SetParseMode(value *TransactionParseRequest_parse_mode)() {
    m.parse_mode = value
}
// SetPayloadHex sets the payload_hex property value. A hex-encoded payload of a full transaction or a partial transaction - either a notarized transaction,a signed transaction intent an unsigned transaction intent, or a ledger payload.
func (m *TransactionParseRequest) SetPayloadHex(value *string)() {
    m.payload_hex = value
}
// SetResponseMode sets the response_mode property value. The amount of information to return in the response."Basic" includes the type, validity information, and any relevant identifiers."Full" also includes the fully parsed information.If omitted, "Full" is used.
func (m *TransactionParseRequest) SetResponseMode(value *TransactionParseRequest_response_mode)() {
    m.response_mode = value
}
// SetTransactionFormatOptions sets the transaction_format_options property value. Requested transaction formats to include in the response
func (m *TransactionParseRequest) SetTransactionFormatOptions(value TransactionFormatOptionsable)() {
    m.transaction_format_options = value
}
// SetValidationMode sets the validation_mode property value. The type of validation that should be performed, if the payload correctly decompiles as a Notarized Transaction.This is only relevant for Notarized payloads. If omitted, "Static" is used.
func (m *TransactionParseRequest) SetValidationMode(value *TransactionParseRequest_validation_mode)() {
    m.validation_mode = value
}
type TransactionParseRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetwork()(*string)
    GetParseMode()(*TransactionParseRequest_parse_mode)
    GetPayloadHex()(*string)
    GetResponseMode()(*TransactionParseRequest_response_mode)
    GetTransactionFormatOptions()(TransactionFormatOptionsable)
    GetValidationMode()(*TransactionParseRequest_validation_mode)
    SetNetwork(value *string)()
    SetParseMode(value *TransactionParseRequest_parse_mode)()
    SetPayloadHex(value *string)()
    SetResponseMode(value *TransactionParseRequest_response_mode)()
    SetTransactionFormatOptions(value TransactionFormatOptionsable)()
    SetValidationMode(value *TransactionParseRequest_validation_mode)()
}
