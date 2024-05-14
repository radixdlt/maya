package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequireProofRule struct {
	ProofRule
	// The requirement property
	requirement Requirementable
}

// NewRequireProofRule instantiates a new RequireProofRule and sets the default values.
func NewRequireProofRule() *RequireProofRule {
	m := &RequireProofRule{
		ProofRule: *NewProofRule(),
	}
	return m
}

// CreateRequireProofRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequireProofRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewRequireProofRule(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequireProofRule) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ProofRule.GetFieldDeserializers()
	res["requirement"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateRequirementFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRequirement(val.(Requirementable))
		}
		return nil
	}
	return res
}

// GetRequirement gets the requirement property value. The requirement property
// returns a Requirementable when successful
func (m *RequireProofRule) GetRequirement() Requirementable {
	return m.requirement
}

// Serialize serializes information the current object
func (m *RequireProofRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ProofRule.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("requirement", m.GetRequirement())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRequirement sets the requirement property value. The requirement property
func (m *RequireProofRule) SetRequirement(value Requirementable) {
	m.requirement = value
}

type RequireProofRuleable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ProofRuleable
	GetRequirement() Requirementable
	SetRequirement(value Requirementable)
}
