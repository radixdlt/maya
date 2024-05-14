package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericTypeParameter struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The constraints on the concrete type provided to fill the generic type parameter.Note: currently, we only support the wildcard (i.e. `Any`) generic type.
	constraints *GenericTypeParameterConstraints
}

// NewGenericTypeParameter instantiates a new GenericTypeParameter and sets the default values.
func NewGenericTypeParameter() *GenericTypeParameter {
	m := &GenericTypeParameter{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateGenericTypeParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericTypeParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenericTypeParameter(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericTypeParameter) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetConstraints gets the constraints property value. The constraints on the concrete type provided to fill the generic type parameter.Note: currently, we only support the wildcard (i.e. `Any`) generic type.
// returns a *GenericTypeParameterConstraints when successful
func (m *GenericTypeParameter) GetConstraints() *GenericTypeParameterConstraints {
	return m.constraints
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericTypeParameter) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["constraints"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseGenericTypeParameterConstraints)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetConstraints(val.(*GenericTypeParameterConstraints))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *GenericTypeParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetConstraints() != nil {
		cast := (*m.GetConstraints()).String()
		err := writer.WriteStringValue("constraints", &cast)
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
func (m *GenericTypeParameter) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetConstraints sets the constraints property value. The constraints on the concrete type provided to fill the generic type parameter.Note: currently, we only support the wildcard (i.e. `Any`) generic type.
func (m *GenericTypeParameter) SetConstraints(value *GenericTypeParameterConstraints) {
	m.constraints = value
}

type GenericTypeParameterable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetConstraints() *GenericTypeParameterConstraints
	SetConstraints(value *GenericTypeParameterConstraints)
}
