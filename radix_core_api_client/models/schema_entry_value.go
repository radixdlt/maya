package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SchemaEntryValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The schema property
	schema ScryptoSchemaable
}

// NewSchemaEntryValue instantiates a new SchemaEntryValue and sets the default values.
func NewSchemaEntryValue() *SchemaEntryValue {
	m := &SchemaEntryValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSchemaEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSchemaEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSchemaEntryValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SchemaEntryValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SchemaEntryValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["schema"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateScryptoSchemaFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSchema(val.(ScryptoSchemaable))
		}
		return nil
	}
	return res
}

// GetSchema gets the schema property value. The schema property
// returns a ScryptoSchemaable when successful
func (m *SchemaEntryValue) GetSchema() ScryptoSchemaable {
	return m.schema
}

// Serialize serializes information the current object
func (m *SchemaEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("schema", m.GetSchema())
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
func (m *SchemaEntryValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSchema sets the schema property value. The schema property
func (m *SchemaEntryValue) SetSchema(value ScryptoSchemaable) {
	m.schema = value
}

type SchemaEntryValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSchema() ScryptoSchemaable
	SetSchema(value ScryptoSchemaable)
}
