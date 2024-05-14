package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllOfProofRule struct {
	ProofRule
	// The list property
	list []Requirementable
}

// NewAllOfProofRule instantiates a new AllOfProofRule and sets the default values.
func NewAllOfProofRule() *AllOfProofRule {
	m := &AllOfProofRule{
		ProofRule: *NewProofRule(),
	}
	return m
}

// CreateAllOfProofRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllOfProofRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAllOfProofRule(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllOfProofRule) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.ProofRule.GetFieldDeserializers()
	res["list"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateRequirementFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]Requirementable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(Requirementable)
				}
			}
			m.SetList(res)
		}
		return nil
	}
	return res
}

// GetList gets the list property value. The list property
// returns a []Requirementable when successful
func (m *AllOfProofRule) GetList() []Requirementable {
	return m.list
}

// Serialize serializes information the current object
func (m *AllOfProofRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.ProofRule.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetList() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetList()))
		for i, v := range m.GetList() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("list", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetList sets the list property value. The list property
func (m *AllOfProofRule) SetList(value []Requirementable) {
	m.list = value
}

type AllOfProofRuleable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ProofRuleable
	GetList() []Requirementable
	SetList(value []Requirementable)
}
