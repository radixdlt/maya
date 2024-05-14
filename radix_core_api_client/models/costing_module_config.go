package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CostingModuleConfig struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Whether to apply costing for reference checks on boot.
	apply_boot_ref_check_costing *bool
	// Whether to apply execution costing for all system calls.
	apply_execution_cost_for_all_system_calls *bool
	// The string-encoded decimal representing the maximum amount of XRD configurable for a single function's royalty.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	xrd_max_per_function_royalty *string
}

// NewCostingModuleConfig instantiates a new CostingModuleConfig and sets the default values.
func NewCostingModuleConfig() *CostingModuleConfig {
	m := &CostingModuleConfig{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCostingModuleConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCostingModuleConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCostingModuleConfig(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CostingModuleConfig) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetApplyBootRefCheckCosting gets the apply_boot_ref_check_costing property value. Whether to apply costing for reference checks on boot.
// returns a *bool when successful
func (m *CostingModuleConfig) GetApplyBootRefCheckCosting() *bool {
	return m.apply_boot_ref_check_costing
}

// GetApplyExecutionCostForAllSystemCalls gets the apply_execution_cost_for_all_system_calls property value. Whether to apply execution costing for all system calls.
// returns a *bool when successful
func (m *CostingModuleConfig) GetApplyExecutionCostForAllSystemCalls() *bool {
	return m.apply_execution_cost_for_all_system_calls
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CostingModuleConfig) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["apply_boot_ref_check_costing"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplyBootRefCheckCosting(val)
		}
		return nil
	}
	res["apply_execution_cost_for_all_system_calls"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplyExecutionCostForAllSystemCalls(val)
		}
		return nil
	}
	res["xrd_max_per_function_royalty"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetXrdMaxPerFunctionRoyalty(val)
		}
		return nil
	}
	return res
}

// GetXrdMaxPerFunctionRoyalty gets the xrd_max_per_function_royalty property value. The string-encoded decimal representing the maximum amount of XRD configurable for a single function's royalty.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingModuleConfig) GetXrdMaxPerFunctionRoyalty() *string {
	return m.xrd_max_per_function_royalty
}

// Serialize serializes information the current object
func (m *CostingModuleConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("apply_boot_ref_check_costing", m.GetApplyBootRefCheckCosting())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("apply_execution_cost_for_all_system_calls", m.GetApplyExecutionCostForAllSystemCalls())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("xrd_max_per_function_royalty", m.GetXrdMaxPerFunctionRoyalty())
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
func (m *CostingModuleConfig) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetApplyBootRefCheckCosting sets the apply_boot_ref_check_costing property value. Whether to apply costing for reference checks on boot.
func (m *CostingModuleConfig) SetApplyBootRefCheckCosting(value *bool) {
	m.apply_boot_ref_check_costing = value
}

// SetApplyExecutionCostForAllSystemCalls sets the apply_execution_cost_for_all_system_calls property value. Whether to apply execution costing for all system calls.
func (m *CostingModuleConfig) SetApplyExecutionCostForAllSystemCalls(value *bool) {
	m.apply_execution_cost_for_all_system_calls = value
}

// SetXrdMaxPerFunctionRoyalty sets the xrd_max_per_function_royalty property value. The string-encoded decimal representing the maximum amount of XRD configurable for a single function's royalty.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingModuleConfig) SetXrdMaxPerFunctionRoyalty(value *string) {
	m.xrd_max_per_function_royalty = value
}

type CostingModuleConfigable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetApplyBootRefCheckCosting() *bool
	GetApplyExecutionCostForAllSystemCalls() *bool
	GetXrdMaxPerFunctionRoyalty() *string
	SetApplyBootRefCheckCosting(value *bool)
	SetApplyExecutionCostForAllSystemCalls(value *bool)
	SetXrdMaxPerFunctionRoyalty(value *string)
}
