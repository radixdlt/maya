package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectSortedIndexPartitionEntryStructure struct {
	SubstateSystemStructure
}

// NewObjectSortedIndexPartitionEntryStructure instantiates a new ObjectSortedIndexPartitionEntryStructure and sets the default values.
func NewObjectSortedIndexPartitionEntryStructure() *ObjectSortedIndexPartitionEntryStructure {
	m := &ObjectSortedIndexPartitionEntryStructure{
		SubstateSystemStructure: *NewSubstateSystemStructure(),
	}
	return m
}

// CreateObjectSortedIndexPartitionEntryStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectSortedIndexPartitionEntryStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewObjectSortedIndexPartitionEntryStructure(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectSortedIndexPartitionEntryStructure) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SubstateSystemStructure.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *ObjectSortedIndexPartitionEntryStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SubstateSystemStructure.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type ObjectSortedIndexPartitionEntryStructureable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubstateSystemStructureable
}
