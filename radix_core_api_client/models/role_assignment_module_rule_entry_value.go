package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RoleAssignmentModuleRuleEntryValue struct {
	// The access_rule property
	access_rule AccessRuleable
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewRoleAssignmentModuleRuleEntryValue instantiates a new RoleAssignmentModuleRuleEntryValue and sets the default values.
func NewRoleAssignmentModuleRuleEntryValue() *RoleAssignmentModuleRuleEntryValue {
	m := &RoleAssignmentModuleRuleEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateRoleAssignmentModuleRuleEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleAssignmentModuleRuleEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRoleAssignmentModuleRuleEntryValue(), nil
}

// GetAccessRule gets the access_rule property value. The access_rule property
// returns a AccessRuleable when successful
func (m *RoleAssignmentModuleRuleEntryValue) GetAccessRule() AccessRuleable {
	return m.access_rule
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoleAssignmentModuleRuleEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoleAssignmentModuleRuleEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["access_rule"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccessRuleFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccessRule(val.(AccessRuleable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *RoleAssignmentModuleRuleEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("access_rule", m.GetAccessRule())
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

// SetAccessRule sets the access_rule property value. The access_rule property
func (m *RoleAssignmentModuleRuleEntryValue) SetAccessRule(value AccessRuleable) {
	m.access_rule = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleAssignmentModuleRuleEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

type RoleAssignmentModuleRuleEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccessRule() AccessRuleable
	SetAccessRule(value AccessRuleable)
}
