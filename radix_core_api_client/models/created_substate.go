package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreatedSubstate struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The substate_id property
	substate_id SubstateIdable
	// The system_structure property
	system_structure SubstateSystemStructureable
	// The value property
	value SubstateValueable
}

// NewCreatedSubstate instantiates a new CreatedSubstate and sets the default values.
func NewCreatedSubstate() *CreatedSubstate {
	m := &CreatedSubstate{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCreatedSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreatedSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCreatedSubstate(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreatedSubstate) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreatedSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["substate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateId(val.(SubstateIdable))
		}
		return nil
	}
	res["system_structure"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateSystemStructureFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSystemStructure(val.(SubstateSystemStructureable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(SubstateValueable))
		}
		return nil
	}
	return res
}

// GetSubstateId gets the substate_id property value. The substate_id property
// returns a SubstateIdable when successful
func (m *CreatedSubstate) GetSubstateId() SubstateIdable {
	return m.substate_id
}

// GetSystemStructure gets the system_structure property value. The system_structure property
// returns a SubstateSystemStructureable when successful
func (m *CreatedSubstate) GetSystemStructure() SubstateSystemStructureable {
	return m.system_structure
}

// GetValue gets the value property value. The value property
// returns a SubstateValueable when successful
func (m *CreatedSubstate) GetValue() SubstateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *CreatedSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("substate_id", m.GetSubstateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("system_structure", m.GetSystemStructure())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *CreatedSubstate) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSubstateId sets the substate_id property value. The substate_id property
func (m *CreatedSubstate) SetSubstateId(value SubstateIdable) {
	m.substate_id = value
}

// SetSystemStructure sets the system_structure property value. The system_structure property
func (m *CreatedSubstate) SetSystemStructure(value SubstateSystemStructureable) {
	m.system_structure = value
}

// SetValue sets the value property value. The value property
func (m *CreatedSubstate) SetValue(value SubstateValueable) {
	m.value = value
}

type CreatedSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSubstateId() SubstateIdable
	GetSystemStructure() SubstateSystemStructureable
	GetValue() SubstateValueable
	SetSubstateId(value SubstateIdable)
	SetSystemStructure(value SubstateSystemStructureable)
	SetValue(value SubstateValueable)
}
