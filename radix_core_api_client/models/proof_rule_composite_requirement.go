package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProofRuleCompositeRequirement struct {
    CompositeRequirement
    // The proof_rule property
    proof_rule BasicRequirementable
}
// NewProofRuleCompositeRequirement instantiates a new ProofRuleCompositeRequirement and sets the default values.
func NewProofRuleCompositeRequirement()(*ProofRuleCompositeRequirement) {
    m := &ProofRuleCompositeRequirement{
        CompositeRequirement: *NewCompositeRequirement(),
    }
    return m
}
// CreateProofRuleCompositeRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProofRuleCompositeRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProofRuleCompositeRequirement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProofRuleCompositeRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CompositeRequirement.GetFieldDeserializers()
    res["proof_rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBasicRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProofRule(val.(BasicRequirementable))
        }
        return nil
    }
    return res
}
// GetProofRule gets the proof_rule property value. The proof_rule property
// returns a BasicRequirementable when successful
func (m *ProofRuleCompositeRequirement) GetProofRule()(BasicRequirementable) {
    return m.proof_rule
}
// Serialize serializes information the current object
func (m *ProofRuleCompositeRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CompositeRequirement.Serialize(writer)
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
func (m *ProofRuleCompositeRequirement) SetProofRule(value BasicRequirementable)() {
    m.proof_rule = value
}
type ProofRuleCompositeRequirementable interface {
    CompositeRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProofRule()(BasicRequirementable)
    SetProofRule(value BasicRequirementable)()
}
