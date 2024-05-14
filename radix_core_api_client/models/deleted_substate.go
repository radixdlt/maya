package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeletedSubstate struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The previous_value property
	previous_value SubstateValueable
	// The substate_id property
	substate_id SubstateIdable
	// The system_structure property
	system_structure SubstateSystemStructureable
}

// NewDeletedSubstate instantiates a new DeletedSubstate and sets the default values.
func NewDeletedSubstate() *DeletedSubstate {
	m := &DeletedSubstate{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateDeletedSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewDeletedSubstate(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeletedSubstate) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeletedSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["previous_value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreviousValue(val.(SubstateValueable))
		}
		return nil
	}
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
	return res
}

// GetPreviousValue gets the previous_value property value. The previous_value property
// returns a SubstateValueable when successful
func (m *DeletedSubstate) GetPreviousValue() SubstateValueable {
	return m.previous_value
}

// GetSubstateId gets the substate_id property value. The substate_id property
// returns a SubstateIdable when successful
func (m *DeletedSubstate) GetSubstateId() SubstateIdable {
	return m.substate_id
}

// GetSystemStructure gets the system_structure property value. The system_structure property
// returns a SubstateSystemStructureable when successful
func (m *DeletedSubstate) GetSystemStructure() SubstateSystemStructureable {
	return m.system_structure
}

// Serialize serializes information the current object
func (m *DeletedSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("previous_value", m.GetPreviousValue())
		if err != nil {
			return err
		}
	}
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeletedSubstate) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetPreviousValue sets the previous_value property value. The previous_value property
func (m *DeletedSubstate) SetPreviousValue(value SubstateValueable) {
	m.previous_value = value
}

// SetSubstateId sets the substate_id property value. The substate_id property
func (m *DeletedSubstate) SetSubstateId(value SubstateIdable) {
	m.substate_id = value
}

// SetSystemStructure sets the system_structure property value. The system_structure property
func (m *DeletedSubstate) SetSystemStructure(value SubstateSystemStructureable) {
	m.system_structure = value
}

type DeletedSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetPreviousValue() SubstateValueable
	GetSubstateId() SubstateIdable
	GetSystemStructure() SubstateSystemStructureable
	SetPreviousValue(value SubstateValueable)
	SetSubstateId(value SubstateIdable)
	SetSystemStructure(value SubstateSystemStructureable)
}
