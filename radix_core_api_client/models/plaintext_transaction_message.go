package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlaintextTransactionMessage an unencrypted message.
type PlaintextTransactionMessage struct {
	TransactionMessage
	// The content property
	content PlaintextMessageContentable
	// Intended to represent the RFC 2046 MIME type of the `content`.A client cannot trust that this field is a valid mime type - in particular, thechoice between `String` or `Binary` representation of the content is not enforced bythis `mime_type`.
	mime_type *string
}

// NewPlaintextTransactionMessage instantiates a new PlaintextTransactionMessage and sets the default values.
func NewPlaintextTransactionMessage() *PlaintextTransactionMessage {
	m := &PlaintextTransactionMessage{
		TransactionMessage: *NewTransactionMessage(),
	}
	return m
}

// CreatePlaintextTransactionMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlaintextTransactionMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPlaintextTransactionMessage(), nil
}

// GetContent gets the content property value. The content property
// returns a PlaintextMessageContentable when successful
func (m *PlaintextTransactionMessage) GetContent() PlaintextMessageContentable {
	return m.content
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlaintextTransactionMessage) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TransactionMessage.GetFieldDeserializers()
	res["content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePlaintextMessageContentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContent(val.(PlaintextMessageContentable))
		}
		return nil
	}
	res["mime_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMimeType(val)
		}
		return nil
	}
	return res
}

// GetMimeType gets the mime_type property value. Intended to represent the RFC 2046 MIME type of the `content`.A client cannot trust that this field is a valid mime type - in particular, thechoice between `String` or `Binary` representation of the content is not enforced bythis `mime_type`.
// returns a *string when successful
func (m *PlaintextTransactionMessage) GetMimeType() *string {
	return m.mime_type
}

// Serialize serializes information the current object
func (m *PlaintextTransactionMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TransactionMessage.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("content", m.GetContent())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("mime_type", m.GetMimeType())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetContent sets the content property value. The content property
func (m *PlaintextTransactionMessage) SetContent(value PlaintextMessageContentable) {
	m.content = value
}

// SetMimeType sets the mime_type property value. Intended to represent the RFC 2046 MIME type of the `content`.A client cannot trust that this field is a valid mime type - in particular, thechoice between `String` or `Binary` representation of the content is not enforced bythis `mime_type`.
func (m *PlaintextTransactionMessage) SetMimeType(value *string) {
	m.mime_type = value
}

type PlaintextTransactionMessageable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TransactionMessageable
	GetContent() PlaintextMessageContentable
	GetMimeType() *string
	SetContent(value PlaintextMessageContentable)
	SetMimeType(value *string)
}
