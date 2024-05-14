package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SystemParameters struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The costing_module_config property
	costing_module_config CostingModuleConfigable
	// The costing_parameters property
	costing_parameters SystemCostingParametersable
	// The limit_parameters property
	limit_parameters LimitParametersable
	// The network_definition property
	network_definition NetworkDefinitionable
}

// NewSystemParameters instantiates a new SystemParameters and sets the default values.
func NewSystemParameters() *SystemParameters {
	m := &SystemParameters{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSystemParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSystemParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSystemParameters(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SystemParameters) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCostingModuleConfig gets the costing_module_config property value. The costing_module_config property
// returns a CostingModuleConfigable when successful
func (m *SystemParameters) GetCostingModuleConfig() CostingModuleConfigable {
	return m.costing_module_config
}

// GetCostingParameters gets the costing_parameters property value. The costing_parameters property
// returns a SystemCostingParametersable when successful
func (m *SystemParameters) GetCostingParameters() SystemCostingParametersable {
	return m.costing_parameters
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SystemParameters) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["costing_module_config"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCostingModuleConfigFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCostingModuleConfig(val.(CostingModuleConfigable))
		}
		return nil
	}
	res["costing_parameters"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSystemCostingParametersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCostingParameters(val.(SystemCostingParametersable))
		}
		return nil
	}
	res["limit_parameters"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLimitParametersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLimitParameters(val.(LimitParametersable))
		}
		return nil
	}
	res["network_definition"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNetworkDefinitionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetworkDefinition(val.(NetworkDefinitionable))
		}
		return nil
	}
	return res
}

// GetLimitParameters gets the limit_parameters property value. The limit_parameters property
// returns a LimitParametersable when successful
func (m *SystemParameters) GetLimitParameters() LimitParametersable {
	return m.limit_parameters
}

// GetNetworkDefinition gets the network_definition property value. The network_definition property
// returns a NetworkDefinitionable when successful
func (m *SystemParameters) GetNetworkDefinition() NetworkDefinitionable {
	return m.network_definition
}

// Serialize serializes information the current object
func (m *SystemParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("costing_module_config", m.GetCostingModuleConfig())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("costing_parameters", m.GetCostingParameters())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("limit_parameters", m.GetLimitParameters())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("network_definition", m.GetNetworkDefinition())
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
func (m *SystemParameters) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCostingModuleConfig sets the costing_module_config property value. The costing_module_config property
func (m *SystemParameters) SetCostingModuleConfig(value CostingModuleConfigable) {
	m.costing_module_config = value
}

// SetCostingParameters sets the costing_parameters property value. The costing_parameters property
func (m *SystemParameters) SetCostingParameters(value SystemCostingParametersable) {
	m.costing_parameters = value
}

// SetLimitParameters sets the limit_parameters property value. The limit_parameters property
func (m *SystemParameters) SetLimitParameters(value LimitParametersable) {
	m.limit_parameters = value
}

// SetNetworkDefinition sets the network_definition property value. The network_definition property
func (m *SystemParameters) SetNetworkDefinition(value NetworkDefinitionable) {
	m.network_definition = value
}

type SystemParametersable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCostingModuleConfig() CostingModuleConfigable
	GetCostingParameters() SystemCostingParametersable
	GetLimitParameters() LimitParametersable
	GetNetworkDefinition() NetworkDefinitionable
	SetCostingModuleConfig(value CostingModuleConfigable)
	SetCostingParameters(value SystemCostingParametersable)
	SetLimitParameters(value LimitParametersable)
	SetNetworkDefinition(value NetworkDefinitionable)
}
