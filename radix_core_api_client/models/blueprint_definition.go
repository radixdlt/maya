package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintDefinition struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A map from the function name to its export
	function_exports BlueprintDefinition_function_exportsable
	// A map from certain object lifecycle hooks to a callback "package export".There is at most one callback registered for each `ObjectHook`.
	hook_exports []HookExportable
	// The interface property
	interfaceEscaped BlueprintInterfaceable
}

// NewBlueprintDefinition instantiates a new BlueprintDefinition and sets the default values.
func NewBlueprintDefinition() *BlueprintDefinition {
	m := &BlueprintDefinition{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBlueprintDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBlueprintDefinition(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintDefinition) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintDefinition) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["function_exports"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintDefinition_function_exportsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFunctionExports(val.(BlueprintDefinition_function_exportsable))
		}
		return nil
	}
	res["hook_exports"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateHookExportFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]HookExportable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(HookExportable)
				}
			}
			m.SetHookExports(res)
		}
		return nil
	}
	res["interface"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintInterfaceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInterfaceEscaped(val.(BlueprintInterfaceable))
		}
		return nil
	}
	return res
}

// GetFunctionExports gets the function_exports property value. A map from the function name to its export
// returns a BlueprintDefinition_function_exportsable when successful
func (m *BlueprintDefinition) GetFunctionExports() BlueprintDefinition_function_exportsable {
	return m.function_exports
}

// GetHookExports gets the hook_exports property value. A map from certain object lifecycle hooks to a callback "package export".There is at most one callback registered for each `ObjectHook`.
// returns a []HookExportable when successful
func (m *BlueprintDefinition) GetHookExports() []HookExportable {
	return m.hook_exports
}

// GetInterfaceEscaped gets the interface property value. The interface property
// returns a BlueprintInterfaceable when successful
func (m *BlueprintDefinition) GetInterfaceEscaped() BlueprintInterfaceable {
	return m.interfaceEscaped
}

// Serialize serializes information the current object
func (m *BlueprintDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("function_exports", m.GetFunctionExports())
		if err != nil {
			return err
		}
	}
	if m.GetHookExports() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHookExports()))
		for i, v := range m.GetHookExports() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("hook_exports", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("interface", m.GetInterfaceEscaped())
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
func (m *BlueprintDefinition) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFunctionExports sets the function_exports property value. A map from the function name to its export
func (m *BlueprintDefinition) SetFunctionExports(value BlueprintDefinition_function_exportsable) {
	m.function_exports = value
}

// SetHookExports sets the hook_exports property value. A map from certain object lifecycle hooks to a callback "package export".There is at most one callback registered for each `ObjectHook`.
func (m *BlueprintDefinition) SetHookExports(value []HookExportable) {
	m.hook_exports = value
}

// SetInterfaceEscaped sets the interface property value. The interface property
func (m *BlueprintDefinition) SetInterfaceEscaped(value BlueprintInterfaceable) {
	m.interfaceEscaped = value
}

type BlueprintDefinitionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFunctionExports() BlueprintDefinition_function_exportsable
	GetHookExports() []HookExportable
	GetInterfaceEscaped() BlueprintInterfaceable
	SetFunctionExports(value BlueprintDefinition_function_exportsable)
	SetHookExports(value []HookExportable)
	SetInterfaceEscaped(value BlueprintInterfaceable)
}
