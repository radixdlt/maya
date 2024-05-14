package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SborData various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
type SborData struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The hex-encoded, raw SBOR-encoded data
	hex *string
	// The (untyped) unannotated programmatic SBOR JSON
	programmatic_json SborData_programmatic_jsonable
}

// NewSborData instantiates a new SborData and sets the default values.
func NewSborData() *SborData {
	m := &SborData{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSborDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSborDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSborData(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SborData) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SborData) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHex(val)
		}
		return nil
	}
	res["programmatic_json"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSborData_programmatic_jsonFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProgrammaticJson(val.(SborData_programmatic_jsonable))
		}
		return nil
	}
	return res
}

// GetHex gets the hex property value. The hex-encoded, raw SBOR-encoded data
// returns a *string when successful
func (m *SborData) GetHex() *string {
	return m.hex
}

// GetProgrammaticJson gets the programmatic_json property value. The (untyped) unannotated programmatic SBOR JSON
// returns a SborData_programmatic_jsonable when successful
func (m *SborData) GetProgrammaticJson() SborData_programmatic_jsonable {
	return m.programmatic_json
}

// Serialize serializes information the current object
func (m *SborData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("hex", m.GetHex())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("programmatic_json", m.GetProgrammaticJson())
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
func (m *SborData) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetHex sets the hex property value. The hex-encoded, raw SBOR-encoded data
func (m *SborData) SetHex(value *string) {
	m.hex = value
}

// SetProgrammaticJson sets the programmatic_json property value. The (untyped) unannotated programmatic SBOR JSON
func (m *SborData) SetProgrammaticJson(value SborData_programmatic_jsonable) {
	m.programmatic_json = value
}

type SborDataable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetHex() *string
	GetProgrammaticJson() SborData_programmatic_jsonable
	SetHex(value *string)
	SetProgrammaticJson(value SborData_programmatic_jsonable)
}
