package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountFieldStateValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
    default_deposit_rule *DefaultDepositRule
}
// NewAccountFieldStateValue instantiates a new AccountFieldStateValue and sets the default values.
func NewAccountFieldStateValue()(*AccountFieldStateValue) {
    m := &AccountFieldStateValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccountFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccountFieldStateValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountFieldStateValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultDepositRule gets the default_deposit_rule property value. This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
// returns a *DefaultDepositRule when successful
func (m *AccountFieldStateValue) GetDefaultDepositRule()(*DefaultDepositRule) {
    return m.default_deposit_rule
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountFieldStateValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["default_deposit_rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefaultDepositRule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDepositRule(val.(*DefaultDepositRule))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AccountFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDefaultDepositRule() != nil {
        cast := (*m.GetDefaultDepositRule()).String()
        err := writer.WriteStringValue("default_deposit_rule", &cast)
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
func (m *AccountFieldStateValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultDepositRule sets the default_deposit_rule property value. This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
func (m *AccountFieldStateValue) SetDefaultDepositRule(value *DefaultDepositRule)() {
    m.default_deposit_rule = value
}
type AccountFieldStateValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultDepositRule()(*DefaultDepositRule)
    SetDefaultDepositRule(value *DefaultDepositRule)()
}
