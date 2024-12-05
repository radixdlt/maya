package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InvalidTransactionPreviewV2ErrorDetails struct {
    TransactionPreviewV2ErrorDetails
    // The validation_error property
    validation_error *string
}
// NewInvalidTransactionPreviewV2ErrorDetails instantiates a new InvalidTransactionPreviewV2ErrorDetails and sets the default values.
func NewInvalidTransactionPreviewV2ErrorDetails()(*InvalidTransactionPreviewV2ErrorDetails) {
    m := &InvalidTransactionPreviewV2ErrorDetails{
        TransactionPreviewV2ErrorDetails: *NewTransactionPreviewV2ErrorDetails(),
    }
    return m
}
// CreateInvalidTransactionPreviewV2ErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInvalidTransactionPreviewV2ErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvalidTransactionPreviewV2ErrorDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InvalidTransactionPreviewV2ErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TransactionPreviewV2ErrorDetails.GetFieldDeserializers()
    res["validation_error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationError(val)
        }
        return nil
    }
    return res
}
// GetValidationError gets the validation_error property value. The validation_error property
// returns a *string when successful
func (m *InvalidTransactionPreviewV2ErrorDetails) GetValidationError()(*string) {
    return m.validation_error
}
// Serialize serializes information the current object
func (m *InvalidTransactionPreviewV2ErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TransactionPreviewV2ErrorDetails.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("validation_error", m.GetValidationError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValidationError sets the validation_error property value. The validation_error property
func (m *InvalidTransactionPreviewV2ErrorDetails) SetValidationError(value *string)() {
    m.validation_error = value
}
type InvalidTransactionPreviewV2ErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TransactionPreviewV2ErrorDetailsable
    GetValidationError()(*string)
    SetValidationError(value *string)()
}
