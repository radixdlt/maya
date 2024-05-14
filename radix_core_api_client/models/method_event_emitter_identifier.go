package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MethodEventEmitterIdentifier struct {
	EventEmitterIdentifier
	// The entity property
	entity EntityReferenceable
	// The object_module_id property
	object_module_id *ModuleId
}

// NewMethodEventEmitterIdentifier instantiates a new MethodEventEmitterIdentifier and sets the default values.
func NewMethodEventEmitterIdentifier() *MethodEventEmitterIdentifier {
	m := &MethodEventEmitterIdentifier{
		EventEmitterIdentifier: *NewEventEmitterIdentifier(),
	}
	return m
}

// CreateMethodEventEmitterIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMethodEventEmitterIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMethodEventEmitterIdentifier(), nil
}

// GetEntity gets the entity property value. The entity property
// returns a EntityReferenceable when successful
func (m *MethodEventEmitterIdentifier) GetEntity() EntityReferenceable {
	return m.entity
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MethodEventEmitterIdentifier) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.EventEmitterIdentifier.GetFieldDeserializers()
	res["entity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEntity(val.(EntityReferenceable))
		}
		return nil
	}
	res["object_module_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseModuleId)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetObjectModuleId(val.(*ModuleId))
		}
		return nil
	}
	return res
}

// GetObjectModuleId gets the object_module_id property value. The object_module_id property
// returns a *ModuleId when successful
func (m *MethodEventEmitterIdentifier) GetObjectModuleId() *ModuleId {
	return m.object_module_id
}

// Serialize serializes information the current object
func (m *MethodEventEmitterIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.EventEmitterIdentifier.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("entity", m.GetEntity())
		if err != nil {
			return err
		}
	}
	if m.GetObjectModuleId() != nil {
		cast := (*m.GetObjectModuleId()).String()
		err = writer.WriteStringValue("object_module_id", &cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetEntity sets the entity property value. The entity property
func (m *MethodEventEmitterIdentifier) SetEntity(value EntityReferenceable) {
	m.entity = value
}

// SetObjectModuleId sets the object_module_id property value. The object_module_id property
func (m *MethodEventEmitterIdentifier) SetObjectModuleId(value *ModuleId) {
	m.object_module_id = value
}

type MethodEventEmitterIdentifierable interface {
	EventEmitterIdentifierable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEntity() EntityReferenceable
	GetObjectModuleId() *ModuleId
	SetEntity(value EntityReferenceable)
	SetObjectModuleId(value *ModuleId)
}
