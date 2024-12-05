package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PresentedBadge struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the resource address
    resource_address *string
    // The type property
    typeEscaped *PresentedBadgeType
}
// NewPresentedBadge instantiates a new PresentedBadge and sets the default values.
func NewPresentedBadge()(*PresentedBadge) {
    m := &PresentedBadge{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePresentedBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePresentedBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                        return NewNonFungiblePresentedBadge(), nil
                    case "Resource":
                        return NewResourcePresentedBadge(), nil
                }
            }
        }
    }
    return NewPresentedBadge(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PresentedBadge) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PresentedBadge) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAddress(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePresentedBadgeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PresentedBadgeType))
        }
        return nil
    }
    return res
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *PresentedBadge) GetResourceAddress()(*string) {
    return m.resource_address
}
// GetTypeEscaped gets the type property value. The type property
// returns a *PresentedBadgeType when successful
func (m *PresentedBadge) GetTypeEscaped()(*PresentedBadgeType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *PresentedBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
        if err != nil {
            return err
        }
    }
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
func (m *PresentedBadge) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *PresentedBadge) SetResourceAddress(value *string)() {
    m.resource_address = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PresentedBadge) SetTypeEscaped(value *PresentedBadgeType)() {
    m.typeEscaped = value
}
type PresentedBadgeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceAddress()(*string)
    GetTypeEscaped()(*PresentedBadgeType)
    SetResourceAddress(value *string)()
    SetTypeEscaped(value *PresentedBadgeType)()
}
