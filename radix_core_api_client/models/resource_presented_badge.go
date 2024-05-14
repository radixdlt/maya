package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResourcePresentedBadge struct {
	PresentedBadge
}

// NewResourcePresentedBadge instantiates a new ResourcePresentedBadge and sets the default values.
func NewResourcePresentedBadge() *ResourcePresentedBadge {
	m := &ResourcePresentedBadge{
		PresentedBadge: *NewPresentedBadge(),
	}
	return m
}

// CreateResourcePresentedBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourcePresentedBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewResourcePresentedBadge(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResourcePresentedBadge) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PresentedBadge.GetFieldDeserializers()
	return res
}

// Serialize serializes information the current object
func (m *ResourcePresentedBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.PresentedBadge.Serialize(writer)
	if err != nil {
		return err
	}
	return nil
}

type ResourcePresentedBadgeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PresentedBadgeable
}
