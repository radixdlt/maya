package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PartitionId struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
	entity_address *string
	// The entity_module property
	entity_module *EntityModule
	// The entity_type property
	entity_type *EntityType
	// The type of the partition:- Field  - Has key: TupleKey(u8) also known as an offset  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- KeyValue ("ConcurrentMap")  - Has key: MapKey(Vec<u8>)  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- Index  - Has key: MapKey(Vec<u8>)  - Iteration exposed to engine via the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety- SortedIndex  - Has key: SortedU16Key(U16, Vec<u8>)  - Iteration exposed to engine via the user-controlled U16 prefix and then the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety
	partition_kind *PartitionKind
	// The partition_number property
	partition_number *int32
}

// NewPartitionId instantiates a new PartitionId and sets the default values.
func NewPartitionId() *PartitionId {
	m := &PartitionId{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePartitionIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartitionIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPartitionId(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PartitionId) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEntityAddress gets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
// returns a *string when successful
func (m *PartitionId) GetEntityAddress() *string {
	return m.entity_address
}

// GetEntityModule gets the entity_module property value. The entity_module property
// returns a *EntityModule when successful
func (m *PartitionId) GetEntityModule() *EntityModule {
	return m.entity_module
}

// GetEntityType gets the entity_type property value. The entity_type property
// returns a *EntityType when successful
func (m *PartitionId) GetEntityType() *EntityType {
	return m.entity_type
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartitionId) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["entity_module"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseEntityModule)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEntityModule(val.(*EntityModule))
		}
		return nil
	}
	res["entity_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseEntityType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEntityType(val.(*EntityType))
		}
		return nil
	}
	res["partition_kind"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParsePartitionKind)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartitionKind(val.(*PartitionKind))
		}
		return nil
	}
	res["partition_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartitionNumber(val)
		}
		return nil
	}
	return res
}

// GetPartitionKind gets the partition_kind property value. The type of the partition:- Field  - Has key: TupleKey(u8) also known as an offset  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- KeyValue ("ConcurrentMap")  - Has key: MapKey(Vec<u8>)  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- Index  - Has key: MapKey(Vec<u8>)  - Iteration exposed to engine via the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety- SortedIndex  - Has key: SortedU16Key(U16, Vec<u8>)  - Iteration exposed to engine via the user-controlled U16 prefix and then the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety
// returns a *PartitionKind when successful
func (m *PartitionId) GetPartitionKind() *PartitionKind {
	return m.partition_kind
}

// GetPartitionNumber gets the partition_number property value. The partition_number property
// returns a *int32 when successful
func (m *PartitionId) GetPartitionNumber() *int32 {
	return m.partition_number
}

// Serialize serializes information the current object
func (m *PartitionId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("entity_address", m.GetEntityAddress())
		if err != nil {
			return err
		}
	}
	if m.GetEntityModule() != nil {
		cast := (*m.GetEntityModule()).String()
		err := writer.WriteStringValue("entity_module", &cast)
		if err != nil {
			return err
		}
	}
	if m.GetEntityType() != nil {
		cast := (*m.GetEntityType()).String()
		err := writer.WriteStringValue("entity_type", &cast)
		if err != nil {
			return err
		}
	}
	if m.GetPartitionKind() != nil {
		cast := (*m.GetPartitionKind()).String()
		err := writer.WriteStringValue("partition_kind", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("partition_number", m.GetPartitionNumber())
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
func (m *PartitionId) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEntityAddress sets the entity_address property value. Bech32m-encoded human readable version of the entity's address (ie the entity's node id)
func (m *PartitionId) SetEntityAddress(value *string) {
	m.entity_address = value
}

// SetEntityModule sets the entity_module property value. The entity_module property
func (m *PartitionId) SetEntityModule(value *EntityModule) {
	m.entity_module = value
}

// SetEntityType sets the entity_type property value. The entity_type property
func (m *PartitionId) SetEntityType(value *EntityType) {
	m.entity_type = value
}

// SetPartitionKind sets the partition_kind property value. The type of the partition:- Field  - Has key: TupleKey(u8) also known as an offset  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- KeyValue ("ConcurrentMap")  - Has key: MapKey(Vec<u8>)  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- Index  - Has key: MapKey(Vec<u8>)  - Iteration exposed to engine via the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety- SortedIndex  - Has key: SortedU16Key(U16, Vec<u8>)  - Iteration exposed to engine via the user-controlled U16 prefix and then the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety
func (m *PartitionId) SetPartitionKind(value *PartitionKind) {
	m.partition_kind = value
}

// SetPartitionNumber sets the partition_number property value. The partition_number property
func (m *PartitionId) SetPartitionNumber(value *int32) {
	m.partition_number = value
}

type PartitionIdable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEntityAddress() *string
	GetEntityModule() *EntityModule
	GetEntityType() *EntityType
	GetPartitionKind() *PartitionKind
	GetPartitionNumber() *int32
	SetEntityAddress(value *string)
	SetEntityModule(value *EntityModule)
	SetEntityType(value *EntityType)
	SetPartitionKind(value *PartitionKind)
	SetPartitionNumber(value *int32)
}
