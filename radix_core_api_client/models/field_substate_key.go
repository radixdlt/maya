package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FieldSubstateKey struct {
	SubstateKey
	// The id property
	id *int32
}

// NewFieldSubstateKey instantiates a new FieldSubstateKey and sets the default values.
func NewFieldSubstateKey() *FieldSubstateKey {
	m := &FieldSubstateKey{
		SubstateKey: *NewSubstateKey(),
	}
	return m
}

// CreateFieldSubstateKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldSubstateKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFieldSubstateKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldSubstateKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SubstateKey.GetFieldDeserializers()
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *FieldSubstateKey) GetId() *int32 {
	return m.id
}

// Serialize serializes information the current object
func (m *FieldSubstateKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SubstateKey.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteInt32Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetId sets the id property value. The id property
func (m *FieldSubstateKey) SetId(value *int32) {
	m.id = value
}

type FieldSubstateKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubstateKeyable
	GetId() *int32
	SetId(value *int32)
}
