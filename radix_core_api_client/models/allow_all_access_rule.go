package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllowAllAccessRule struct {
	AccessRule
}

// NewAllowAllAccessRule instantiates a new AllowAllAccessRule and sets the default values.
func NewAllowAllAccessRule() *AllowAllAccessRule {
	m := &AllowAllAccessRule{
		AccessRule: *NewAccessRule(),
	}
	return m
}

// CreateAllowAllAccessRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllowAllAccessRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAllowAllAccessRule(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllowAllAccessRule) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.AccessRule.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *AllowAllAccessRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.AccessRule.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type AllowAllAccessRuleable interface {
	AccessRuleable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
