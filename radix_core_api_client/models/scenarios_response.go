package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScenariosResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Scenarios executed as part of Genesis and Protocol Updates, in their execution order.
	executed_scenarios []ExecutedScenarioable
}

// NewScenariosResponse instantiates a new ScenariosResponse and sets the default values.
func NewScenariosResponse() *ScenariosResponse {
	m := &ScenariosResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateScenariosResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScenariosResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewScenariosResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScenariosResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetExecutedScenarios gets the executed_scenarios property value. Scenarios executed as part of Genesis and Protocol Updates, in their execution order.
// returns a []ExecutedScenarioable when successful
func (m *ScenariosResponse) GetExecutedScenarios() []ExecutedScenarioable {
	return m.executed_scenarios
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScenariosResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["executed_scenarios"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateExecutedScenarioFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ExecutedScenarioable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ExecutedScenarioable)
				}
			}
			m.SetExecutedScenarios(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *ScenariosResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetExecutedScenarios() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExecutedScenarios()))
		for i, v := range m.GetExecutedScenarios() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("executed_scenarios", cast)
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
func (m *ScenariosResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetExecutedScenarios sets the executed_scenarios property value. Scenarios executed as part of Genesis and Protocol Updates, in their execution order.
func (m *ScenariosResponse) SetExecutedScenarios(value []ExecutedScenarioable) {
	m.executed_scenarios = value
}

type ScenariosResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetExecutedScenarios() []ExecutedScenarioable
	SetExecutedScenarios(value []ExecutedScenarioable)
}
