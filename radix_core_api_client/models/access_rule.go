package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessRule struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The type property
	typeEscaped *AccessRuleType
}

// NewAccessRule instantiates a new AccessRule and sets the default values.
func NewAccessRule() *AccessRule {
	m := &AccessRule{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAccessRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
				case "AllowAll":
					return NewAllowAllAccessRule(), nil
				case "DenyAll":
					return NewDenyAllAccessRule(), nil
				case "Protected":
					return NewProtectedAccessRule(), nil
				}
			}
		}
	}
	return NewAccessRule(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessRule) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessRule) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseAccessRuleType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*AccessRuleType))
		}
		return nil
	}
	return res
}

// GetTypeEscaped gets the type property value. The type property
// returns a *AccessRuleType when successful
func (m *AccessRule) GetTypeEscaped() *AccessRuleType {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *AccessRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AccessRule) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *AccessRule) SetTypeEscaped(value *AccessRuleType) {
	m.typeEscaped = value
}

type AccessRuleable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetTypeEscaped() *AccessRuleType
	SetTypeEscaped(value *AccessRuleType)
}
