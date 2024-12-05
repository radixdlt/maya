package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintMethodRoyalty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The method_name property
    method_name *string
    // The royalty_amount property
    royalty_amount RoyaltyAmountable
}
// NewBlueprintMethodRoyalty instantiates a new BlueprintMethodRoyalty and sets the default values.
func NewBlueprintMethodRoyalty()(*BlueprintMethodRoyalty) {
    m := &BlueprintMethodRoyalty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintMethodRoyaltyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintMethodRoyaltyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintMethodRoyalty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintMethodRoyalty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintMethodRoyalty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["method_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodName(val)
        }
        return nil
    }
    res["royalty_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoyaltyAmountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoyaltyAmount(val.(RoyaltyAmountable))
        }
        return nil
    }
    return res
}
// GetMethodName gets the method_name property value. The method_name property
// returns a *string when successful
func (m *BlueprintMethodRoyalty) GetMethodName()(*string) {
    return m.method_name
}
// GetRoyaltyAmount gets the royalty_amount property value. The royalty_amount property
// returns a RoyaltyAmountable when successful
func (m *BlueprintMethodRoyalty) GetRoyaltyAmount()(RoyaltyAmountable) {
    return m.royalty_amount
}
// Serialize serializes information the current object
func (m *BlueprintMethodRoyalty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("method_name", m.GetMethodName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("royalty_amount", m.GetRoyaltyAmount())
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
func (m *BlueprintMethodRoyalty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMethodName sets the method_name property value. The method_name property
func (m *BlueprintMethodRoyalty) SetMethodName(value *string)() {
    m.method_name = value
}
// SetRoyaltyAmount sets the royalty_amount property value. The royalty_amount property
func (m *BlueprintMethodRoyalty) SetRoyaltyAmount(value RoyaltyAmountable)() {
    m.royalty_amount = value
}
type BlueprintMethodRoyaltyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMethodName()(*string)
    GetRoyaltyAmount()(RoyaltyAmountable)
    SetMethodName(value *string)()
    SetRoyaltyAmount(value RoyaltyAmountable)()
}
