package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FieldSchemaFeatureConditionAlways struct {
	FieldSchemaFeatureCondition
}

// NewFieldSchemaFeatureConditionAlways instantiates a new FieldSchemaFeatureConditionAlways and sets the default values.
func NewFieldSchemaFeatureConditionAlways() *FieldSchemaFeatureConditionAlways {
	m := &FieldSchemaFeatureConditionAlways{
		FieldSchemaFeatureCondition: *NewFieldSchemaFeatureCondition(),
	}
	return m
}

// CreateFieldSchemaFeatureConditionAlwaysFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldSchemaFeatureConditionAlwaysFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFieldSchemaFeatureConditionAlways(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldSchemaFeatureConditionAlways) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.FieldSchemaFeatureCondition.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *FieldSchemaFeatureConditionAlways) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.FieldSchemaFeatureCondition.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type FieldSchemaFeatureConditionAlwaysable interface {
	FieldSchemaFeatureConditionable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
