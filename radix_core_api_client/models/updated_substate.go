package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdatedSubstate struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The new_value property
	new_value SubstateValueable
	// The previous_value property
	previous_value SubstateValueable
	// The substate_id property
	substate_id SubstateIdable
	// The system_structure property
	system_structure SubstateSystemStructureable
}

// NewUpdatedSubstate instantiates a new UpdatedSubstate and sets the default values.
func NewUpdatedSubstate() *UpdatedSubstate {
	m := &UpdatedSubstate{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateUpdatedSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdatedSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUpdatedSubstate(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UpdatedSubstate) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdatedSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["new_value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNewValue(val.(SubstateValueable))
		}
		return nil
	}
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

// GetNewValue gets the new_value property value. The new_value property
// returns a SubstateValueable when successful
func (m *UpdatedSubstate) GetNewValue() SubstateValueable {
	return m.new_value
}

// GetPreviousValue gets the previous_value property value. The previous_value property
// returns a SubstateValueable when successful
func (m *UpdatedSubstate) GetPreviousValue() SubstateValueable {
	return m.previous_value
}

// GetSubstateId gets the substate_id property value. The substate_id property
// returns a SubstateIdable when successful
func (m *UpdatedSubstate) GetSubstateId() SubstateIdable {
	return m.substate_id
}

// GetSystemStructure gets the system_structure property value. The system_structure property
// returns a SubstateSystemStructureable when successful
func (m *UpdatedSubstate) GetSystemStructure() SubstateSystemStructureable {
	return m.system_structure
}

// Serialize serializes information the current object
func (m *UpdatedSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("new_value", m.GetNewValue())
		if err != nil {
			return err
		}
	}
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
func (m *UpdatedSubstate) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNewValue sets the new_value property value. The new_value property
func (m *UpdatedSubstate) SetNewValue(value SubstateValueable) {
	m.new_value = value
}

// SetPreviousValue sets the previous_value property value. The previous_value property
func (m *UpdatedSubstate) SetPreviousValue(value SubstateValueable) {
	m.previous_value = value
}

// SetSubstateId sets the substate_id property value. The substate_id property
func (m *UpdatedSubstate) SetSubstateId(value SubstateIdable) {
	m.substate_id = value
}

// SetSystemStructure sets the system_structure property value. The system_structure property
func (m *UpdatedSubstate) SetSystemStructure(value SubstateSystemStructureable) {
	m.system_structure = value
}

type UpdatedSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNewValue() SubstateValueable
	GetPreviousValue() SubstateValueable
	GetSubstateId() SubstateIdable
	GetSystemStructure() SubstateSystemStructureable
	SetNewValue(value SubstateValueable)
	SetPreviousValue(value SubstateValueable)
	SetSubstateId(value SubstateIdable)
	SetSystemStructure(value SubstateSystemStructureable)
}
