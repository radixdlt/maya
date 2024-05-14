package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizedDepositorKey struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The badge property
	badge AuthorizedDepositorBadgeable
}

// NewAuthorizedDepositorKey instantiates a new AuthorizedDepositorKey and sets the default values.
func NewAuthorizedDepositorKey() *AuthorizedDepositorKey {
	m := &AuthorizedDepositorKey{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAuthorizedDepositorKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizedDepositorKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAuthorizedDepositorKey(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthorizedDepositorKey) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBadge gets the badge property value. The badge property
// returns a AuthorizedDepositorBadgeable when successful
func (m *AuthorizedDepositorKey) GetBadge() AuthorizedDepositorBadgeable {
	return m.badge
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthorizedDepositorKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAuthorizedDepositorBadgeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBadge(val.(AuthorizedDepositorBadgeable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *AuthorizedDepositorKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("badge", m.GetBadge())
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
func (m *AuthorizedDepositorKey) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBadge sets the badge property value. The badge property
func (m *AuthorizedDepositorKey) SetBadge(value AuthorizedDepositorBadgeable) {
	m.badge = value
}

type AuthorizedDepositorKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBadge() AuthorizedDepositorBadgeable
	SetBadge(value AuthorizedDepositorBadgeable)
}
