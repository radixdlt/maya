package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OwnerRole struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The rule property
	rule AccessRuleable
	// The updater property
	updater *OwnerRoleUpdater
}

// NewOwnerRole instantiates a new OwnerRole and sets the default values.
func NewOwnerRole() *OwnerRole {
	m := &OwnerRole{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOwnerRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOwnerRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOwnerRole(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OwnerRole) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OwnerRole) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["rule"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccessRuleFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRule(val.(AccessRuleable))
		}
		return nil
	}
	res["updater"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseOwnerRoleUpdater)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUpdater(val.(*OwnerRoleUpdater))
		}
		return nil
	}
	return res
}

// GetRule gets the rule property value. The rule property
// returns a AccessRuleable when successful
func (m *OwnerRole) GetRule() AccessRuleable {
	return m.rule
}

// GetUpdater gets the updater property value. The updater property
// returns a *OwnerRoleUpdater when successful
func (m *OwnerRole) GetUpdater() *OwnerRoleUpdater {
	return m.updater
}

// Serialize serializes information the current object
func (m *OwnerRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("rule", m.GetRule())
		if err != nil {
			return err
		}
	}
	if m.GetUpdater() != nil {
		cast := (*m.GetUpdater()).String()
		err := writer.WriteStringValue("updater", &cast)
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
func (m *OwnerRole) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetRule sets the rule property value. The rule property
func (m *OwnerRole) SetRule(value AccessRuleable) {
	m.rule = value
}

// SetUpdater sets the updater property value. The updater property
func (m *OwnerRole) SetUpdater(value *OwnerRoleUpdater) {
	m.updater = value
}

type OwnerRoleable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRule() AccessRuleable
	GetUpdater() *OwnerRoleUpdater
	SetRule(value AccessRuleable)
	SetUpdater(value *OwnerRoleUpdater)
}
