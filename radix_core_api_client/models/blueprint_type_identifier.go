package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BlueprintTypeIdentifier an identifier for a defined type in the v1 blueprint version under the given package blueprint.
type BlueprintTypeIdentifier struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The blueprint_name property
    blueprint_name *string
    // The Bech32m-encoded human readable version of the package address
    package_address *string
    // The type_name property
    type_name *string
}
// NewBlueprintTypeIdentifier instantiates a new BlueprintTypeIdentifier and sets the default values.
func NewBlueprintTypeIdentifier()(*BlueprintTypeIdentifier) {
    m := &BlueprintTypeIdentifier{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintTypeIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintTypeIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintTypeIdentifier(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintTypeIdentifier) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlueprintName gets the blueprint_name property value. The blueprint_name property
// returns a *string when successful
func (m *BlueprintTypeIdentifier) GetBlueprintName()(*string) {
    return m.blueprint_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintTypeIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blueprint_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueprintName(val)
        }
        return nil
    }
    res["package_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageAddress(val)
        }
        return nil
    }
    res["type_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeName(val)
        }
        return nil
    }
    return res
}
// GetPackageAddress gets the package_address property value. The Bech32m-encoded human readable version of the package address
// returns a *string when successful
func (m *BlueprintTypeIdentifier) GetPackageAddress()(*string) {
    return m.package_address
}
// GetTypeName gets the type_name property value. The type_name property
// returns a *string when successful
func (m *BlueprintTypeIdentifier) GetTypeName()(*string) {
    return m.type_name
}
// Serialize serializes information the current object
func (m *BlueprintTypeIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("blueprint_name", m.GetBlueprintName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("package_address", m.GetPackageAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_name", m.GetTypeName())
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
func (m *BlueprintTypeIdentifier) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlueprintName sets the blueprint_name property value. The blueprint_name property
func (m *BlueprintTypeIdentifier) SetBlueprintName(value *string)() {
    m.blueprint_name = value
}
// SetPackageAddress sets the package_address property value. The Bech32m-encoded human readable version of the package address
func (m *BlueprintTypeIdentifier) SetPackageAddress(value *string)() {
    m.package_address = value
}
// SetTypeName sets the type_name property value. The type_name property
func (m *BlueprintTypeIdentifier) SetTypeName(value *string)() {
    m.type_name = value
}
type BlueprintTypeIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlueprintName()(*string)
    GetPackageAddress()(*string)
    GetTypeName()(*string)
    SetBlueprintName(value *string)()
    SetPackageAddress(value *string)()
    SetTypeName(value *string)()
}
