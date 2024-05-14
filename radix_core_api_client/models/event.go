package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Event event emitted by a transaction.
type Event struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
	data SborDataable
	// Identifier of a specific event schema.
	typeEscaped EventTypeIdentifierable
}

// NewEvent instantiates a new Event and sets the default values.
func NewEvent() *Event {
	m := &Event{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEvent(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Event) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetData gets the data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
// returns a SborDataable when successful
func (m *Event) GetData() SborDataable {
	return m.data
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Event) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSborDataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetData(val.(SborDataable))
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEventTypeIdentifierFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(EventTypeIdentifierable))
		}
		return nil
	}
	return res
}

// GetTypeEscaped gets the type property value. Identifier of a specific event schema.
// returns a EventTypeIdentifierable when successful
func (m *Event) GetTypeEscaped() EventTypeIdentifierable {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *Event) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("data", m.GetData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("type", m.GetTypeEscaped())
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
func (m *Event) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetData sets the data property value. Various representations of an SBOR payload.Some endpoints may allow opting in/out of each representation.
func (m *Event) SetData(value SborDataable) {
	m.data = value
}

// SetTypeEscaped sets the type property value. Identifier of a specific event schema.
func (m *Event) SetTypeEscaped(value EventTypeIdentifierable) {
	m.typeEscaped = value
}

type Eventable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetData() SborDataable
	GetTypeEscaped() EventTypeIdentifierable
	SetData(value SborDataable)
	SetTypeEscaped(value EventTypeIdentifierable)
}
