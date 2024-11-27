package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParsedNotarizedTransactionV2_validation_error if the transaction is known to not be valid, this gives a reason.Different levels of validation are performed, dependent on the validation mode.Note that, even if validation mode is Static or Full, the transaction maystill be rejected or fail due to issues at runtime (e.g. if the loan cannot be repaid).
type ParsedNotarizedTransactionV2_validation_error struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the error is known to be permanent, or not.This relates to whether the transaction would be rejected permanently or temporarily if submitted.
    is_permanent *bool
    // The error message.
    reason *string
}
// NewParsedNotarizedTransactionV2_validation_error instantiates a new ParsedNotarizedTransactionV2_validation_error and sets the default values.
func NewParsedNotarizedTransactionV2_validation_error()(*ParsedNotarizedTransactionV2_validation_error) {
    m := &ParsedNotarizedTransactionV2_validation_error{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateParsedNotarizedTransactionV2_validation_errorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParsedNotarizedTransactionV2_validation_errorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParsedNotarizedTransactionV2_validation_error(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ParsedNotarizedTransactionV2_validation_error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParsedNotarizedTransactionV2_validation_error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_permanent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPermanent(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    return res
}
// GetIsPermanent gets the is_permanent property value. Whether the error is known to be permanent, or not.This relates to whether the transaction would be rejected permanently or temporarily if submitted.
// returns a *bool when successful
func (m *ParsedNotarizedTransactionV2_validation_error) GetIsPermanent()(*bool) {
    return m.is_permanent
}
// GetReason gets the reason property value. The error message.
// returns a *string when successful
func (m *ParsedNotarizedTransactionV2_validation_error) GetReason()(*string) {
    return m.reason
}
// Serialize serializes information the current object
func (m *ParsedNotarizedTransactionV2_validation_error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_permanent", m.GetIsPermanent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
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
func (m *ParsedNotarizedTransactionV2_validation_error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsPermanent sets the is_permanent property value. Whether the error is known to be permanent, or not.This relates to whether the transaction would be rejected permanently or temporarily if submitted.
func (m *ParsedNotarizedTransactionV2_validation_error) SetIsPermanent(value *bool)() {
    m.is_permanent = value
}
// SetReason sets the reason property value. The error message.
func (m *ParsedNotarizedTransactionV2_validation_error) SetReason(value *string)() {
    m.reason = value
}
type ParsedNotarizedTransactionV2_validation_errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsPermanent()(*bool)
    GetReason()(*string)
    SetIsPermanent(value *bool)()
    SetReason(value *string)()
}
