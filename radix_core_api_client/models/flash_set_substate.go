package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FlashSetSubstate struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The substate_id property
	substate_id SubstateIdable
	// The value property
	value SubstateValueable
}

// NewFlashSetSubstate instantiates a new FlashSetSubstate and sets the default values.
func NewFlashSetSubstate() *FlashSetSubstate {
	m := &FlashSetSubstate{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFlashSetSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFlashSetSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFlashSetSubstate(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FlashSetSubstate) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FlashSetSubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *FlashSetSubstate) GetSubstateId() SubstateIdable {
	return m.substate_id
}

// GetValue gets the value property value. The value property
// returns a SubstateValueable when successful
func (m *FlashSetSubstate) GetValue() SubstateValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *FlashSetSubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("substate_id", m.GetSubstateId())
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
func (m *FlashSetSubstate) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSubstateId sets the substate_id property value. The substate_id property
func (m *FlashSetSubstate) SetSubstateId(value SubstateIdable) {
	m.substate_id = value
}

// SetValue sets the value property value. The value property
func (m *FlashSetSubstate) SetValue(value SubstateValueable) {
	m.value = value
}

type FlashSetSubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSubstateId() SubstateIdable
	GetValue() SubstateValueable
	SetSubstateId(value SubstateIdable)
	SetValue(value SubstateValueable)
}
