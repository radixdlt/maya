package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllOfAccessRuleNode struct {
	AccessRuleNode
	// The access_rules property
	access_rules []AccessRuleNodeable
}

// NewAllOfAccessRuleNode instantiates a new AllOfAccessRuleNode and sets the default values.
func NewAllOfAccessRuleNode() *AllOfAccessRuleNode {
	m := &AllOfAccessRuleNode{
		AccessRuleNode: *NewAccessRuleNode(),
	}
	return m
}

// CreateAllOfAccessRuleNodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllOfAccessRuleNodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAllOfAccessRuleNode(), nil
}

// GetAccessRules gets the access_rules property value. The access_rules property
// returns a []AccessRuleNodeable when successful
func (m *AllOfAccessRuleNode) GetAccessRules() []AccessRuleNodeable {
	return m.access_rules
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllOfAccessRuleNode) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *AllOfAccessRuleNode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AllOfAccessRuleNode) SetAccessRules(value []AccessRuleNodeable) {
	m.access_rules = value
}

type AllOfAccessRuleNodeable interface {
	AccessRuleNodeable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccessRules() []AccessRuleNodeable
	SetAccessRules(value []AccessRuleNodeable)
}
