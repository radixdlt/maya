package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FieldSchemaFeatureConditionIfOwnFeature struct {
    FieldSchemaFeatureCondition
    // The feature_name property
    feature_name *string
}
// NewFieldSchemaFeatureConditionIfOwnFeature instantiates a new FieldSchemaFeatureConditionIfOwnFeature and sets the default values.
func NewFieldSchemaFeatureConditionIfOwnFeature()(*FieldSchemaFeatureConditionIfOwnFeature) {
    m := &FieldSchemaFeatureConditionIfOwnFeature{
        FieldSchemaFeatureCondition: *NewFieldSchemaFeatureCondition(),
    }
    return m
}
// CreateFieldSchemaFeatureConditionIfOwnFeatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldSchemaFeatureConditionIfOwnFeatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFieldSchemaFeatureConditionIfOwnFeature(), nil
}
// GetFeatureName gets the feature_name property value. The feature_name property
// returns a *string when successful
func (m *FieldSchemaFeatureConditionIfOwnFeature) GetFeatureName()(*string) {
    return m.feature_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldSchemaFeatureConditionIfOwnFeature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FieldSchemaFeatureCondition.GetFieldDeserializers()
    res["feature_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FieldSchemaFeatureConditionIfOwnFeature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FieldSchemaFeatureCondition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("feature_name", m.GetFeatureName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFeatureName sets the feature_name property value. The feature_name property
func (m *FieldSchemaFeatureConditionIfOwnFeature) SetFeatureName(value *string)() {
    m.feature_name = value
}
type FieldSchemaFeatureConditionIfOwnFeatureable interface {
    FieldSchemaFeatureConditionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFeatureName()(*string)
    SetFeatureName(value *string)()
}
