package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompositeRequirement this type was historically called `AccessRuleNode`.
type CompositeRequirement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // `ProofRule` is now called `BasicRequirement`.
    typeEscaped *CompositeRequirementType
}
// NewCompositeRequirement instantiates a new CompositeRequirement and sets the default values.
func NewCompositeRequirement()(*CompositeRequirement) {
    m := &CompositeRequirement{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCompositeRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompositeRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "AllOf":
                        return NewAllOfCompositeRequirement(), nil
                    case "AnyOf":
                        return NewAnyOfCompositeRequirement(), nil
                    case "ProofRule":
                        return NewProofRuleCompositeRequirement(), nil
                }
            }
        }
    }
    return NewCompositeRequirement(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompositeRequirement) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompositeRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCompositeRequirementType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*CompositeRequirementType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. `ProofRule` is now called `BasicRequirement`.
// returns a *CompositeRequirementType when successful
func (m *CompositeRequirement) GetTypeEscaped()(*CompositeRequirementType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *CompositeRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CompositeRequirement) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. `ProofRule` is now called `BasicRequirement`.
func (m *CompositeRequirement) SetTypeEscaped(value *CompositeRequirementType)() {
    m.typeEscaped = value
}
type CompositeRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*CompositeRequirementType)
    SetTypeEscaped(value *CompositeRequirementType)()
}
