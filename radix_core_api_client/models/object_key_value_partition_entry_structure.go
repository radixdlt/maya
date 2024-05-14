package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectKeyValuePartitionEntryStructure struct {
	SubstateSystemStructure
}

// NewObjectKeyValuePartitionEntryStructure instantiates a new ObjectKeyValuePartitionEntryStructure and sets the default values.
func NewObjectKeyValuePartitionEntryStructure() *ObjectKeyValuePartitionEntryStructure {
	m := &ObjectKeyValuePartitionEntryStructure{
		SubstateSystemStructure: *NewSubstateSystemStructure(),
	}
	return m
}

// CreateObjectKeyValuePartitionEntryStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectKeyValuePartitionEntryStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewObjectKeyValuePartitionEntryStructure(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectKeyValuePartitionEntryStructure) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SubstateSystemStructure.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *ObjectKeyValuePartitionEntryStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SubstateSystemStructure.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type ObjectKeyValuePartitionEntryStructureable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubstateSystemStructureable
}
