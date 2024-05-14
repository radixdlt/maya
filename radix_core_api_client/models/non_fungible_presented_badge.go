package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungiblePresentedBadge struct {
	PresentedBadge
	// The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
	local_id *string
}

// NewNonFungiblePresentedBadge instantiates a new NonFungiblePresentedBadge and sets the default values.
func NewNonFungiblePresentedBadge() *NonFungiblePresentedBadge {
	m := &NonFungiblePresentedBadge{
		PresentedBadge: *NewPresentedBadge(),
	}
	return m
}

// CreateNonFungiblePresentedBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungiblePresentedBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNonFungiblePresentedBadge(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungiblePresentedBadge) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.PresentedBadge.GetFieldDeserializers()
	res["local_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocalId(val)
		}
		return nil
	}
	return res
}

// GetLocalId gets the local_id property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
// returns a *string when successful
func (m *NonFungiblePresentedBadge) GetLocalId() *string {
	return m.local_id
}

// Serialize serializes information the current object
func (m *NonFungiblePresentedBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.PresentedBadge.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("local_id", m.GetLocalId())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetLocalId sets the local_id property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
func (m *NonFungiblePresentedBadge) SetLocalId(value *string) {
	m.local_id = value
}

type NonFungiblePresentedBadgeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PresentedBadgeable
	GetLocalId() *string
	SetLocalId(value *string)
}
