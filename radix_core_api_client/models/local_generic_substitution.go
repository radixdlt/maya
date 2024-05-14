package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalGenericSubstitution the generic substitution is provided by the instance itself.The `scoped_type_id` can be expanded to a `FullyScopedTypeId` by including the current entity's address.
type LocalGenericSubstitution struct {
	GenericSubstitution
	// An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
	scoped_type_id ScopedTypeIdable
}

// NewLocalGenericSubstitution instantiates a new LocalGenericSubstitution and sets the default values.
func NewLocalGenericSubstitution() *LocalGenericSubstitution {
	m := &LocalGenericSubstitution{
		GenericSubstitution: *NewGenericSubstitution(),
	}
	return m
}

// CreateLocalGenericSubstitutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalGenericSubstitutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalGenericSubstitution(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalGenericSubstitution) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.GenericSubstitution.GetFieldDeserializers()
	res["scoped_type_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateScopedTypeIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetScopedTypeId(val.(ScopedTypeIdable))
		}
		return nil
	}
	return res
}

// GetScopedTypeId gets the scoped_type_id property value. An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
// returns a ScopedTypeIdable when successful
func (m *LocalGenericSubstitution) GetScopedTypeId() ScopedTypeIdable {
	return m.scoped_type_id
}

// Serialize serializes information the current object
func (m *LocalGenericSubstitution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.GenericSubstitution.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("scoped_type_id", m.GetScopedTypeId())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetScopedTypeId sets the scoped_type_id property value. An identifier for a type in the context of a schema.The location of the schema store to locate the schema is not included, andis known from context. Currently the schema store will be in theschema partition under a node (typically a package).Note - this type provides scoping to a schema even for well-known types wherethe schema is irrelevant.
func (m *LocalGenericSubstitution) SetScopedTypeId(value ScopedTypeIdable) {
	m.scoped_type_id = value
}

type LocalGenericSubstitutionable interface {
	GenericSubstitutionable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetScopedTypeId() ScopedTypeIdable
	SetScopedTypeId(value ScopedTypeIdable)
}
