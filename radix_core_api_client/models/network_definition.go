package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NetworkDefinition struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The network suffix of Bech32m HRPs used for addressing.
	hrp_suffix *string
	// The logical id of the network
	id *int32
	// The logical name of the network
	logical_name *string
}

// NewNetworkDefinition instantiates a new NetworkDefinition and sets the default values.
func NewNetworkDefinition() *NetworkDefinition {
	m := &NetworkDefinition{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNetworkDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNetworkDefinition(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkDefinition) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkDefinition) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["hrp_suffix"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHrpSuffix(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["logical_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLogicalName(val)
		}
		return nil
	}
	return res
}

// GetHrpSuffix gets the hrp_suffix property value. The network suffix of Bech32m HRPs used for addressing.
// returns a *string when successful
func (m *NetworkDefinition) GetHrpSuffix() *string {
	return m.hrp_suffix
}

// GetId gets the id property value. The logical id of the network
// returns a *int32 when successful
func (m *NetworkDefinition) GetId() *int32 {
	return m.id
}

// GetLogicalName gets the logical_name property value. The logical name of the network
// returns a *string when successful
func (m *NetworkDefinition) GetLogicalName() *string {
	return m.logical_name
}

// Serialize serializes information the current object
func (m *NetworkDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("hrp_suffix", m.GetHrpSuffix())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("logical_name", m.GetLogicalName())
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
func (m *NetworkDefinition) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetHrpSuffix sets the hrp_suffix property value. The network suffix of Bech32m HRPs used for addressing.
func (m *NetworkDefinition) SetHrpSuffix(value *string) {
	m.hrp_suffix = value
}

// SetId sets the id property value. The logical id of the network
func (m *NetworkDefinition) SetId(value *int32) {
	m.id = value
}

// SetLogicalName sets the logical_name property value. The logical name of the network
func (m *NetworkDefinition) SetLogicalName(value *string) {
	m.logical_name = value
}

type NetworkDefinitionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetHrpSuffix() *string
	GetId() *int32
	GetLogicalName() *string
	SetHrpSuffix(value *string)
	SetId(value *int32)
	SetLogicalName(value *string)
}
