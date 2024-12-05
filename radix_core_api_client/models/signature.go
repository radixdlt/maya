package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Signature struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The key_type property
    key_type *PublicKeyType
}
// NewSignature instantiates a new Signature and sets the default values.
func NewSignature()(*Signature) {
    m := &Signature{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSignatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("key_type")
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
                    case "EcdsaSecp256k1":
                        return NewEcdsaSecp256k1Signature(), nil
                    case "EddsaEd25519":
                        return NewEddsaEd25519Signature(), nil
                }
            }
        }
    }
    return NewSignature(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Signature) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Signature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePublicKeyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyType(val.(*PublicKeyType))
        }
        return nil
    }
    return res
}
// GetKeyType gets the key_type property value. The key_type property
// returns a *PublicKeyType when successful
func (m *Signature) GetKeyType()(*PublicKeyType) {
    return m.key_type
}
// Serialize serializes information the current object
func (m *Signature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeyType() != nil {
        cast := (*m.GetKeyType()).String()
        err := writer.WriteStringValue("key_type", &cast)
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
func (m *Signature) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyType sets the key_type property value. The key_type property
func (m *Signature) SetKeyType(value *PublicKeyType)() {
    m.key_type = value
}
type Signatureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyType()(*PublicKeyType)
    SetKeyType(value *PublicKeyType)()
}
