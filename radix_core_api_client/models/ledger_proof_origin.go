package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerProofOrigin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *LedgerProofOriginType
}
// NewLedgerProofOrigin instantiates a new LedgerProofOrigin and sets the default values.
func NewLedgerProofOrigin()(*LedgerProofOrigin) {
    m := &LedgerProofOrigin{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "Consensus":
                        return NewConsensusLedgerProofOrigin(), nil
                    case "Genesis":
                        return NewGenesisLedgerProofOrigin(), nil
                    case "ProtocolUpdate":
                        return NewProtocolUpdateLedgerProofOrigin(), nil
                }
            }
        }
    }
    return NewLedgerProofOrigin(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerProofOrigin) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerProofOrigin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLedgerProofOriginType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*LedgerProofOriginType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *LedgerProofOriginType when successful
func (m *LedgerProofOrigin) GetTypeEscaped()(*LedgerProofOriginType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *LedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *LedgerProofOrigin) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *LedgerProofOrigin) SetTypeEscaped(value *LedgerProofOriginType)() {
    m.typeEscaped = value
}
type LedgerProofOriginable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*LedgerProofOriginType)
    SetTypeEscaped(value *LedgerProofOriginType)()
}
