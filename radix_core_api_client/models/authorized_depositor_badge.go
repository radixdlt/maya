package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizedDepositorBadge struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The type property
	typeEscaped *AuthorizedDepositorBadgeType
}

// NewAuthorizedDepositorBadge instantiates a new AuthorizedDepositorBadge and sets the default values.
func NewAuthorizedDepositorBadge() *AuthorizedDepositorBadge {
	m := &AuthorizedDepositorBadge{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAuthorizedDepositorBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizedDepositorBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("type")
		if err != nil {
			return nil, err
		}
		if mappingValueNode != nil {
			mappingValue, err := mappingValueNode.GetStringValue()
			if err != nil {
				return nil, err
			}
			if mappingValue != nil {
				switch *mappingValue {
				case "NonFungible":
					return NewNonFungibleAuthorizedDepositorBadge(), nil
				case "Resource":
					return NewResourceAuthorizedDepositorBadge(), nil
				}
			}
		}
	}
	return NewAuthorizedDepositorBadge(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthorizedDepositorBadge) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthorizedDepositorBadge) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseAuthorizedDepositorBadgeType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*AuthorizedDepositorBadgeType))
		}
		return nil
	}
	return res
}

// GetTypeEscaped gets the type property value. The type property
// returns a *AuthorizedDepositorBadgeType when successful
func (m *AuthorizedDepositorBadge) GetTypeEscaped() *AuthorizedDepositorBadgeType {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *AuthorizedDepositorBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetTypeEscaped() != nil {
		cast := (*m.GetTypeEscaped()).String()
		err := writer.WriteStringValue("type", &cast)
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
func (m *AuthorizedDepositorBadge) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *AuthorizedDepositorBadge) SetTypeEscaped(value *AuthorizedDepositorBadgeType) {
	m.typeEscaped = value
}

type AuthorizedDepositorBadgeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetTypeEscaped() *AuthorizedDepositorBadgeType
	SetTypeEscaped(value *AuthorizedDepositorBadgeType)
}
