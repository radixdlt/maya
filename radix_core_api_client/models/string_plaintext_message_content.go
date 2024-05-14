package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StringPlaintextMessageContent struct {
	PlaintextMessageContent
	// The value of a message that the author decided to provide as a UTF-8 string.
	value *string
}

// NewStringPlaintextMessageContent instantiates a new StringPlaintextMessageContent and sets the default values.
func NewStringPlaintextMessageContent() *StringPlaintextMessageContent {
	m := &StringPlaintextMessageContent{
		PlaintextMessageContent: *NewPlaintextMessageContent(),
	}
	return m
}

// CreateStringPlaintextMessageContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStringPlaintextMessageContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStringPlaintextMessageContent(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StringPlaintextMessageContent) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PlaintextMessageContent.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value of a message that the author decided to provide as a UTF-8 string.
// returns a *string when successful
func (m *StringPlaintextMessageContent) GetValue() *string {
	return m.value
}

// Serialize serializes information the current object
func (m *StringPlaintextMessageContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.PlaintextMessageContent.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. The value of a message that the author decided to provide as a UTF-8 string.
func (m *StringPlaintextMessageContent) SetValue(value *string) {
	m.value = value
}

type StringPlaintextMessageContentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PlaintextMessageContentable
	GetValue() *string
	SetValue(value *string)
}
