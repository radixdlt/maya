package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkConfigurationResponse_version different versions regarding the node, network and api.
type NetworkConfigurationResponse_version struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The api_version property
	api_version *string
	// The core_version property
	core_version *string
}

// NewNetworkConfigurationResponse_version instantiates a new NetworkConfigurationResponse_version and sets the default values.
func NewNetworkConfigurationResponse_version() *NetworkConfigurationResponse_version {
	m := &NetworkConfigurationResponse_version{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNetworkConfigurationResponse_versionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkConfigurationResponse_versionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNetworkConfigurationResponse_version(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkConfigurationResponse_version) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetApiVersion gets the api_version property value. The api_version property
// returns a *string when successful
func (m *NetworkConfigurationResponse_version) GetApiVersion() *string {
	return m.api_version
}

// GetCoreVersion gets the core_version property value. The core_version property
// returns a *string when successful
func (m *NetworkConfigurationResponse_version) GetCoreVersion() *string {
	return m.core_version
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkConfigurationResponse_version) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["api_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApiVersion(val)
		}
		return nil
	}
	res["core_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCoreVersion(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *NetworkConfigurationResponse_version) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("api_version", m.GetApiVersion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("core_version", m.GetCoreVersion())
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
func (m *NetworkConfigurationResponse_version) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetApiVersion sets the api_version property value. The api_version property
func (m *NetworkConfigurationResponse_version) SetApiVersion(value *string) {
	m.api_version = value
}

// SetCoreVersion sets the core_version property value. The core_version property
func (m *NetworkConfigurationResponse_version) SetCoreVersion(value *string) {
	m.core_version = value
}

type NetworkConfigurationResponse_versionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetApiVersion() *string
	GetCoreVersion() *string
	SetApiVersion(value *string)
	SetCoreVersion(value *string)
}
