package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewV2ErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *TransactionPreviewV2ErrorDetailsType
}
// NewTransactionPreviewV2ErrorDetails instantiates a new TransactionPreviewV2ErrorDetails and sets the default values.
func NewTransactionPreviewV2ErrorDetails()(*TransactionPreviewV2ErrorDetails) {
    m := &TransactionPreviewV2ErrorDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionPreviewV2ErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewV2ErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "Invalid":
                        return NewInvalidTransactionPreviewV2ErrorDetails(), nil
                }
            }
        }
    }
    return NewTransactionPreviewV2ErrorDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewV2ErrorDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewV2ErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionPreviewV2ErrorDetailsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*TransactionPreviewV2ErrorDetailsType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *TransactionPreviewV2ErrorDetailsType when successful
func (m *TransactionPreviewV2ErrorDetails) GetTypeEscaped()(*TransactionPreviewV2ErrorDetailsType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *TransactionPreviewV2ErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TransactionPreviewV2ErrorDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *TransactionPreviewV2ErrorDetails) SetTypeEscaped(value *TransactionPreviewV2ErrorDetailsType)() {
    m.typeEscaped = value
}
type TransactionPreviewV2ErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*TransactionPreviewV2ErrorDetailsType)
    SetTypeEscaped(value *TransactionPreviewV2ErrorDetailsType)()
}
