package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntentSignatures signatures against the given intent.
type IntentSignatures struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The signatures property
    signatures []SignatureWithPublicKeyable
}
// NewIntentSignatures instantiates a new IntentSignatures and sets the default values.
func NewIntentSignatures()(*IntentSignatures) {
    m := &IntentSignatures{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntentSignaturesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentSignaturesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentSignatures(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntentSignatures) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntentSignatures) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignatureWithPublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignatureWithPublicKeyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SignatureWithPublicKeyable)
                }
            }
            m.SetSignatures(res)
        }
        return nil
    }
    return res
}
// GetSignatures gets the signatures property value. The signatures property
// returns a []SignatureWithPublicKeyable when successful
func (m *IntentSignatures) GetSignatures()([]SignatureWithPublicKeyable) {
    return m.signatures
}
// Serialize serializes information the current object
func (m *IntentSignatures) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSignatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignatures()))
        for i, v := range m.GetSignatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("signatures", cast)
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
func (m *IntentSignatures) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSignatures sets the signatures property value. The signatures property
func (m *IntentSignatures) SetSignatures(value []SignatureWithPublicKeyable)() {
    m.signatures = value
}
type IntentSignaturesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSignatures()([]SignatureWithPublicKeyable)
    SetSignatures(value []SignatureWithPublicKeyable)()
}
