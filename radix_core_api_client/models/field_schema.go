package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FieldSchema struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The condition property
	condition FieldSchemaFeatureConditionable
	// The field_type_ref property
	field_type_ref BlueprintPayloadDefable
	// The hex-encoded default value of this field. Only present if this field is transient.
	transient_default_value_hex *string
}

// NewFieldSchema instantiates a new FieldSchema and sets the default values.
func NewFieldSchema() *FieldSchema {
	m := &FieldSchema{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFieldSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFieldSchema(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FieldSchema) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCondition gets the condition property value. The condition property
// returns a FieldSchemaFeatureConditionable when successful
func (m *FieldSchema) GetCondition() FieldSchemaFeatureConditionable {
	return m.condition
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldSchema) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["condition"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFieldSchemaFeatureConditionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCondition(val.(FieldSchemaFeatureConditionable))
		}
		return nil
	}
	res["field_type_ref"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBlueprintPayloadDefFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFieldTypeRef(val.(BlueprintPayloadDefable))
		}
		return nil
	}
	res["transient_default_value_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransientDefaultValueHex(val)
		}
		return nil
	}
	return res
}

// GetFieldTypeRef gets the field_type_ref property value. The field_type_ref property
// returns a BlueprintPayloadDefable when successful
func (m *FieldSchema) GetFieldTypeRef() BlueprintPayloadDefable {
	return m.field_type_ref
}

// GetTransientDefaultValueHex gets the transient_default_value_hex property value. The hex-encoded default value of this field. Only present if this field is transient.
// returns a *string when successful
func (m *FieldSchema) GetTransientDefaultValueHex() *string {
	return m.transient_default_value_hex
}

// Serialize serializes information the current object
func (m *FieldSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("condition", m.GetCondition())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("field_type_ref", m.GetFieldTypeRef())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("transient_default_value_hex", m.GetTransientDefaultValueHex())
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
func (m *FieldSchema) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCondition sets the condition property value. The condition property
func (m *FieldSchema) SetCondition(value FieldSchemaFeatureConditionable) {
	m.condition = value
}

// SetFieldTypeRef sets the field_type_ref property value. The field_type_ref property
func (m *FieldSchema) SetFieldTypeRef(value BlueprintPayloadDefable) {
	m.field_type_ref = value
}

// SetTransientDefaultValueHex sets the transient_default_value_hex property value. The hex-encoded default value of this field. Only present if this field is transient.
func (m *FieldSchema) SetTransientDefaultValueHex(value *string) {
	m.transient_default_value_hex = value
}

type FieldSchemaable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCondition() FieldSchemaFeatureConditionable
	GetFieldTypeRef() BlueprintPayloadDefable
	GetTransientDefaultValueHex() *string
	SetCondition(value FieldSchemaFeatureConditionable)
	SetFieldTypeRef(value BlueprintPayloadDefable)
	SetTransientDefaultValueHex(value *string)
}
