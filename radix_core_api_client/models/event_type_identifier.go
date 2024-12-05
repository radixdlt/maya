package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventTypeIdentifier identifier of a specific event schema.
type EventTypeIdentifier struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The emitter property
    emitter EventEmitterIdentifierable
    // The name property
    name *string
    // The type_reference property
    type_reference PackageTypeReferenceable
}
// NewEventTypeIdentifier instantiates a new EventTypeIdentifier and sets the default values.
func NewEventTypeIdentifier()(*EventTypeIdentifier) {
    m := &EventTypeIdentifier{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventTypeIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventTypeIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventTypeIdentifier(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EventTypeIdentifier) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmitter gets the emitter property value. The emitter property
// returns a EventEmitterIdentifierable when successful
func (m *EventTypeIdentifier) GetEmitter()(EventEmitterIdentifierable) {
    return m.emitter
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EventTypeIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEventEmitterIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmitter(val.(EventEmitterIdentifierable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type_reference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageTypeReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeReference(val.(PackageTypeReferenceable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *EventTypeIdentifier) GetName()(*string) {
    return m.name
}
// GetTypeReference gets the type_reference property value. The type_reference property
// returns a PackageTypeReferenceable when successful
func (m *EventTypeIdentifier) GetTypeReference()(PackageTypeReferenceable) {
    return m.type_reference
}
// Serialize serializes information the current object
func (m *EventTypeIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("emitter", m.GetEmitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("type_reference", m.GetTypeReference())
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
func (m *EventTypeIdentifier) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmitter sets the emitter property value. The emitter property
func (m *EventTypeIdentifier) SetEmitter(value EventEmitterIdentifierable)() {
    m.emitter = value
}
// SetName sets the name property value. The name property
func (m *EventTypeIdentifier) SetName(value *string)() {
    m.name = value
}
// SetTypeReference sets the type_reference property value. The type_reference property
func (m *EventTypeIdentifier) SetTypeReference(value PackageTypeReferenceable)() {
    m.type_reference = value
}
type EventTypeIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmitter()(EventEmitterIdentifierable)
    GetName()(*string)
    GetTypeReference()(PackageTypeReferenceable)
    SetEmitter(value EventEmitterIdentifierable)()
    SetName(value *string)()
    SetTypeReference(value PackageTypeReferenceable)()
}
