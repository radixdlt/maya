package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubstateSystemStructure struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *SubstateSystemStructureType
}
// NewSubstateSystemStructure instantiates a new SubstateSystemStructure and sets the default values.
func NewSubstateSystemStructure()(*SubstateSystemStructure) {
    m := &SubstateSystemStructure{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubstateSystemStructureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateSystemStructureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "KeyValueStoreEntry":
                        return NewKeyValueStoreEntryStructure(), nil
                    case "ObjectField":
                        return NewObjectFieldStructure(), nil
                    case "SystemField":
                        return NewSystemFieldStructure(), nil
                    case "SystemSchema":
                        return NewSystemSchemaStructure(), nil
                }
            }
        }
    }
    return NewSubstateSystemStructure(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubstateSystemStructure) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstateSystemStructure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubstateSystemStructureType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SubstateSystemStructureType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *SubstateSystemStructureType when successful
func (m *SubstateSystemStructure) GetTypeEscaped()(*SubstateSystemStructureType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SubstateSystemStructure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SubstateSystemStructure) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *SubstateSystemStructure) SetTypeEscaped(value *SubstateSystemStructureType)() {
    m.typeEscaped = value
}
type SubstateSystemStructureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*SubstateSystemStructureType)
    SetTypeEscaped(value *SubstateSystemStructureType)()
}
