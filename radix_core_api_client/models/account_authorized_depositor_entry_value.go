package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountAuthorizedDepositorEntryValue empty value. The existence of the key implies the depositor is authorized.
type AccountAuthorizedDepositorEntryValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This is always true. This field is just added to ensure we return some data as the value,so a present entry is not confused by clients for a deleted/missing entry (which would implynot authorized).
    is_authorized *bool
}
// NewAccountAuthorizedDepositorEntryValue instantiates a new AccountAuthorizedDepositorEntryValue and sets the default values.
func NewAccountAuthorizedDepositorEntryValue()(*AccountAuthorizedDepositorEntryValue) {
    m := &AccountAuthorizedDepositorEntryValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccountAuthorizedDepositorEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountAuthorizedDepositorEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccountAuthorizedDepositorEntryValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountAuthorizedDepositorEntryValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountAuthorizedDepositorEntryValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_authorized"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAuthorized(val)
        }
        return nil
    }
    return res
}
// GetIsAuthorized gets the is_authorized property value. This is always true. This field is just added to ensure we return some data as the value,so a present entry is not confused by clients for a deleted/missing entry (which would implynot authorized).
// returns a *bool when successful
func (m *AccountAuthorizedDepositorEntryValue) GetIsAuthorized()(*bool) {
    return m.is_authorized
}
// Serialize serializes information the current object
func (m *AccountAuthorizedDepositorEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_authorized", m.GetIsAuthorized())
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
func (m *AccountAuthorizedDepositorEntryValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsAuthorized sets the is_authorized property value. This is always true. This field is just added to ensure we return some data as the value,so a present entry is not confused by clients for a deleted/missing entry (which would implynot authorized).
func (m *AccountAuthorizedDepositorEntryValue) SetIsAuthorized(value *bool)() {
    m.is_authorized = value
}
type AccountAuthorizedDepositorEntryValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsAuthorized()(*bool)
    SetIsAuthorized(value *bool)()
}
