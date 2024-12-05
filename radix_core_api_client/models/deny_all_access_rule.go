package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DenyAllAccessRule struct {
    AccessRule
}
// NewDenyAllAccessRule instantiates a new DenyAllAccessRule and sets the default values.
func NewDenyAllAccessRule()(*DenyAllAccessRule) {
    m := &DenyAllAccessRule{
        AccessRule: *NewAccessRule(),
    }
    return m
}
// CreateDenyAllAccessRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDenyAllAccessRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDenyAllAccessRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DenyAllAccessRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessRule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DenyAllAccessRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessRule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type DenyAllAccessRuleable interface {
    AccessRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
