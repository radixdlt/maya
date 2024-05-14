package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubstateId struct {
	PartitionId
	// The substate_key property
	substate_key SubstateKeyable
	// The substate_type property
	substate_type *SubstateType
}

// NewSubstateId instantiates a new SubstateId and sets the default values.
func NewSubstateId() *SubstateId {
	m := &SubstateId{
		PartitionId: *NewPartitionId(),
	}
	return m
}

// CreateSubstateIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubstateId(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstateId) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PartitionId.GetFieldDeserializers()
	res["substate_key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateKey(val.(SubstateKeyable))
		}
		return nil
	}
	res["substate_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseSubstateType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateType(val.(*SubstateType))
		}
		return nil
	}
	return res
}

// GetSubstateKey gets the substate_key property value. The substate_key property
// returns a SubstateKeyable when successful
func (m *SubstateId) GetSubstateKey() SubstateKeyable {
	return m.substate_key
}

// GetSubstateType gets the substate_type property value. The substate_type property
// returns a *SubstateType when successful
func (m *SubstateId) GetSubstateType() *SubstateType {
	return m.substate_type
}

// Serialize serializes information the current object
func (m *SubstateId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.PartitionId.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("substate_key", m.GetSubstateKey())
		if err != nil {
			return err
		}
	}
	if m.GetSubstateType() != nil {
		cast := (*m.GetSubstateType()).String()
		err = writer.WriteStringValue("substate_type", &cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetSubstateKey sets the substate_key property value. The substate_key property
func (m *SubstateId) SetSubstateKey(value SubstateKeyable) {
	m.substate_key = value
}

// SetSubstateType sets the substate_type property value. The substate_type property
func (m *SubstateId) SetSubstateType(value *SubstateType) {
	m.substate_type = value
}

type SubstateIdable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PartitionIdable
	GetSubstateKey() SubstateKeyable
	GetSubstateType() *SubstateType
	SetSubstateKey(value SubstateKeyable)
	SetSubstateType(value *SubstateType)
}
