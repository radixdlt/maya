package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScryptoSchema struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
	sbor_data SborDataable
}

// NewScryptoSchema instantiates a new ScryptoSchema and sets the default values.
func NewScryptoSchema() *ScryptoSchema {
	m := &ScryptoSchema{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateScryptoSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScryptoSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewScryptoSchema(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScryptoSchema) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScryptoSchema) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["sbor_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSborData(val.(SborDataable))
		}
		return nil
	}
	return res
}

// GetSborData gets the sbor_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *ScryptoSchema) GetSborData() SborDataable {
	return m.sbor_data
}

// Serialize serializes information the current object
func (m *ScryptoSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("sbor_data", m.GetSborData())
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
func (m *ScryptoSchema) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSborData sets the sbor_data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *ScryptoSchema) SetSborData(value SborDataable) {
	m.sbor_data = value
}

type ScryptoSchemaable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSborData() SborDataable
	SetSborData(value SborDataable)
}
