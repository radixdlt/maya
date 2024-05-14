package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleAuthorizedDepositorBadge struct {
	AuthorizedDepositorBadge
	// The non_fungible_global_id property
	non_fungible_global_id NonFungibleGlobalIdable
}

// NewNonFungibleAuthorizedDepositorBadge instantiates a new NonFungibleAuthorizedDepositorBadge and sets the default values.
func NewNonFungibleAuthorizedDepositorBadge() *NonFungibleAuthorizedDepositorBadge {
	m := &NonFungibleAuthorizedDepositorBadge{
		AuthorizedDepositorBadge: *NewAuthorizedDepositorBadge(),
	}
	return m
}

// CreateNonFungibleAuthorizedDepositorBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleAuthorizedDepositorBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungibleAuthorizedDepositorBadge(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleAuthorizedDepositorBadge) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.AuthorizedDepositorBadge.GetFieldDeserializers()
	res["non_fungible_global_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNonFungibleGlobalIdFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNonFungibleGlobalId(val.(NonFungibleGlobalIdable))
		}
		return nil
	}
	return res
}

// GetNonFungibleGlobalId gets the non_fungible_global_id property value. The non_fungible_global_id property
// returns a NonFungibleGlobalIdable when successful
func (m *NonFungibleAuthorizedDepositorBadge) GetNonFungibleGlobalId() NonFungibleGlobalIdable {
	return m.non_fungible_global_id
}

// Serialize serializes information the current object
func (m *NonFungibleAuthorizedDepositorBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.AuthorizedDepositorBadge.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("non_fungible_global_id", m.GetNonFungibleGlobalId())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetNonFungibleGlobalId sets the non_fungible_global_id property value. The non_fungible_global_id property
func (m *NonFungibleAuthorizedDepositorBadge) SetNonFungibleGlobalId(value NonFungibleGlobalIdable) {
	m.non_fungible_global_id = value
}

type NonFungibleAuthorizedDepositorBadgeable interface {
	AuthorizedDepositorBadgeable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNonFungibleGlobalId() NonFungibleGlobalIdable
	SetNonFungibleGlobalId(value NonFungibleGlobalIdable)
}
