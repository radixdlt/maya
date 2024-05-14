package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AnyOfAccessRuleNode struct {
	AccessRuleNode
	// The access_rules property
	access_rules []AccessRuleNodeable
}

// NewAnyOfAccessRuleNode instantiates a new AnyOfAccessRuleNode and sets the default values.
func NewAnyOfAccessRuleNode() *AnyOfAccessRuleNode {
	m := &AnyOfAccessRuleNode{
		AccessRuleNode: *NewAccessRuleNode(),
	}
	return m
}

// CreateAnyOfAccessRuleNodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnyOfAccessRuleNodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAnyOfAccessRuleNode(), nil
}

// GetAccessRules gets the access_rules property value. The access_rules property
// returns a []AccessRuleNodeable when successful
func (m *AnyOfAccessRuleNode) GetAccessRules() []AccessRuleNodeable {
	return m.access_rules
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AnyOfAccessRuleNode) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.AccessRuleNode.GetFieldDeserializers()
	res["access_rules"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAccessRuleNodeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AccessRuleNodeable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AccessRuleNodeable)
				}
			}
			m.SetAccessRules(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *AnyOfAccessRuleNode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.AccessRuleNode.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetAccessRules() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessRules()))
		for i, v := range m.GetAccessRules() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("access_rules", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAccessRules sets the access_rules property value. The access_rules property
func (m *AnyOfAccessRuleNode) SetAccessRules(value []AccessRuleNodeable) {
	m.access_rules = value
}

type AnyOfAccessRuleNodeable interface {
	AccessRuleNodeable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccessRules() []AccessRuleNodeable
	SetAccessRules(value []AccessRuleNodeable)
}
