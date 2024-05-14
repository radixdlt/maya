package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProofAccessRuleNode struct {
	AccessRuleNode
	// The proof_rule property
	proof_rule ProofRuleable
}

// NewProofAccessRuleNode instantiates a new ProofAccessRuleNode and sets the default values.
func NewProofAccessRuleNode() *ProofAccessRuleNode {
	m := &ProofAccessRuleNode{
		AccessRuleNode: *NewAccessRuleNode(),
	}
	return m
}

// CreateProofAccessRuleNodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProofAccessRuleNodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProofAccessRuleNode(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProofAccessRuleNode) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.AccessRuleNode.GetFieldDeserializers()
	res["proof_rule"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateProofRuleFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProofRule(val.(ProofRuleable))
		}
		return nil
	}
	return res
}

// GetProofRule gets the proof_rule property value. The proof_rule property
// returns a ProofRuleable when successful
func (m *ProofAccessRuleNode) GetProofRule() ProofRuleable {
	return m.proof_rule
}

// Serialize serializes information the current object
func (m *ProofAccessRuleNode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.AccessRuleNode.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("proof_rule", m.GetProofRule())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetProofRule sets the proof_rule property value. The proof_rule property
func (m *ProofAccessRuleNode) SetProofRule(value ProofRuleable) {
	m.proof_rule = value
}

type ProofAccessRuleNodeable interface {
	AccessRuleNodeable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProofRule() ProofRuleable
	SetProofRule(value ProofRuleable)
}
