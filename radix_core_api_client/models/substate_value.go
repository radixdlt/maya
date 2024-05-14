package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubstateValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The substate_data property
	substate_data Substateable
	// The hex-encoded Blake2b-256 hash of the substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
	substate_data_hash *string
	// The hex-encoded, SBOR-encoded substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
	substate_hex *string
}

// NewSubstateValue instantiates a new SubstateValue and sets the default values.
func NewSubstateValue() *SubstateValue {
	m := &SubstateValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubstateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubstateValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubstateValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstateValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["substate_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSubstateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateData(val.(Substateable))
		}
		return nil
	}
	res["substate_data_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateDataHash(val)
		}
		return nil
	}
	res["substate_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubstateHex(val)
		}
		return nil
	}
	return res
}

// GetSubstateData gets the substate_data property value. The substate_data property
// returns a Substateable when successful
func (m *SubstateValue) GetSubstateData() Substateable {
	return m.substate_data
}

// GetSubstateDataHash gets the substate_data_hash property value. The hex-encoded Blake2b-256 hash of the substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
// returns a *string when successful
func (m *SubstateValue) GetSubstateDataHash() *string {
	return m.substate_data_hash
}

// GetSubstateHex gets the substate_hex property value. The hex-encoded, SBOR-encoded substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
// returns a *string when successful
func (m *SubstateValue) GetSubstateHex() *string {
	return m.substate_hex
}

// Serialize serializes information the current object
func (m *SubstateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("substate_data", m.GetSubstateData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("substate_data_hash", m.GetSubstateDataHash())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("substate_hex", m.GetSubstateHex())
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
func (m *SubstateValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSubstateData sets the substate_data property value. The substate_data property
func (m *SubstateValue) SetSubstateData(value Substateable) {
	m.substate_data = value
}

// SetSubstateDataHash sets the substate_data_hash property value. The hex-encoded Blake2b-256 hash of the substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
func (m *SubstateValue) SetSubstateDataHash(value *string) {
	m.substate_data_hash = value
}

// SetSubstateHex sets the substate_hex property value. The hex-encoded, SBOR-encoded substate data bytes. Only returned if enabled in SubstateFormatOptions on your request (default false).
func (m *SubstateValue) SetSubstateHex(value *string) {
	m.substate_hex = value
}

type SubstateValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSubstateData() Substateable
	GetSubstateDataHash() *string
	GetSubstateHex() *string
	SetSubstateData(value Substateable)
	SetSubstateDataHash(value *string)
	SetSubstateHex(value *string)
}
