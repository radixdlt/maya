package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectedAccessRule struct {
    AccessRule
    // This type was historically called `AccessRuleNode`.
    access_rule CompositeRequirementable
}
// NewProtectedAccessRule instantiates a new ProtectedAccessRule and sets the default values.
func NewProtectedAccessRule()(*ProtectedAccessRule) {
    m := &ProtectedAccessRule{
        AccessRule: *NewAccessRule(),
    }
    return m
}
// CreateProtectedAccessRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedAccessRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedAccessRule(), nil
}
// GetAccessRule gets the access_rule property value. This type was historically called `AccessRuleNode`.
// returns a CompositeRequirementable when successful
func (m *ProtectedAccessRule) GetAccessRule()(CompositeRequirementable) {
    return m.access_rule
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedAccessRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessRule.GetFieldDeserializers()
    res["access_rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCompositeRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRule(val.(CompositeRequirementable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtectedAccessRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("access_rule", m.GetAccessRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessRule sets the access_rule property value. This type was historically called `AccessRuleNode`.
func (m *ProtectedAccessRule) SetAccessRule(value CompositeRequirementable)() {
    m.access_rule = value
}
type ProtectedAccessRuleable interface {
    AccessRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRule()(CompositeRequirementable)
    SetAccessRule(value CompositeRequirementable)()
}
