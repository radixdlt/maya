package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FullyScopedTypeId an identifier for a type in the context of a schema in an entity's schema partition.Note - this type provides a schema context even for well-known types where this contextis effectively irrelevant.
type FullyScopedTypeId struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
	entity_address *string
	// The local_type_id property
	local_type_id LocalTypeIdable
	// The hex-encoded schema hash, capturing the identity of an SBOR schema.
	schema_hash *string
}

// NewFullyScopedTypeId instantiates a new FullyScopedTypeId and sets the default values.
func NewFullyScopedTypeId() *FullyScopedTypeId {
	m := &FullyScopedTypeId{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFullyScopedTypeIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFullyScopedTypeIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFullyScopedTypeId(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FullyScopedTypeId) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEntityAddress gets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
// returns a *string when successful
func (m *FullyScopedTypeId) GetEntityAddress() *string {
	return m.entity_address
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FullyScopedTypeId) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["entity_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEntityAddress(val)
		}
		return nil
	}
	res["local_type_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLocalTypeIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocalTypeId(val.(LocalTypeIdable))
		}
		return nil
	}
	res["schema_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSchemaHash(val)
		}
		return nil
	}
	return res
}

// GetLocalTypeId gets the local_type_id property value. The local_type_id property
// returns a LocalTypeIdable when successful
func (m *FullyScopedTypeId) GetLocalTypeId() LocalTypeIdable {
	return m.local_type_id
}

// GetSchemaHash gets the schema_hash property value. The hex-encoded schema hash, capturing the identity of an SBOR schema.
// returns a *string when successful
func (m *FullyScopedTypeId) GetSchemaHash() *string {
	return m.schema_hash
}

// Serialize serializes information the current object
func (m *FullyScopedTypeId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("entity_address", m.GetEntityAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("local_type_id", m.GetLocalTypeId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("schema_hash", m.GetSchemaHash())
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
func (m *FullyScopedTypeId) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEntityAddress sets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
func (m *FullyScopedTypeId) SetEntityAddress(value *string) {
	m.entity_address = value
}

// SetLocalTypeId sets the local_type_id property value. The local_type_id property
func (m *FullyScopedTypeId) SetLocalTypeId(value LocalTypeIdable) {
	m.local_type_id = value
}

// SetSchemaHash sets the schema_hash property value. The hex-encoded schema hash, capturing the identity of an SBOR schema.
func (m *FullyScopedTypeId) SetSchemaHash(value *string) {
	m.schema_hash = value
}

type FullyScopedTypeIdable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEntityAddress() *string
	GetLocalTypeId() LocalTypeIdable
	GetSchemaHash() *string
	SetEntityAddress(value *string)
	SetLocalTypeId(value LocalTypeIdable)
	SetSchemaHash(value *string)
}
