package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintFunctionTargetIdentifier struct {
    TargetIdentifier
    // The blueprint_name property
    blueprint_name *string
    // The function_name property
    function_name *string
    // The Bech32m-encoded human readable version of the package address
    package_address *string
}
// NewBlueprintFunctionTargetIdentifier instantiates a new BlueprintFunctionTargetIdentifier and sets the default values.
func NewBlueprintFunctionTargetIdentifier()(*BlueprintFunctionTargetIdentifier) {
    m := &BlueprintFunctionTargetIdentifier{
        TargetIdentifier: *NewTargetIdentifier(),
    }
    return m
}
// CreateBlueprintFunctionTargetIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintFunctionTargetIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintFunctionTargetIdentifier(), nil
}
// GetBlueprintName gets the blueprint_name property value. The blueprint_name property
// returns a *string when successful
func (m *BlueprintFunctionTargetIdentifier) GetBlueprintName()(*string) {
    return m.blueprint_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintFunctionTargetIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetIdentifier.GetFieldDeserializers()
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
    res["function_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctionName(val)
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
    return res
}
// GetFunctionName gets the function_name property value. The function_name property
// returns a *string when successful
func (m *BlueprintFunctionTargetIdentifier) GetFunctionName()(*string) {
    return m.function_name
}
// GetPackageAddress gets the package_address property value. The Bech32m-encoded human readable version of the package address
// returns a *string when successful
func (m *BlueprintFunctionTargetIdentifier) GetPackageAddress()(*string) {
    return m.package_address
}
// Serialize serializes information the current object
func (m *BlueprintFunctionTargetIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("blueprint_name", m.GetBlueprintName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("function_name", m.GetFunctionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("package_address", m.GetPackageAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlueprintName sets the blueprint_name property value. The blueprint_name property
func (m *BlueprintFunctionTargetIdentifier) SetBlueprintName(value *string)() {
    m.blueprint_name = value
}
// SetFunctionName sets the function_name property value. The function_name property
func (m *BlueprintFunctionTargetIdentifier) SetFunctionName(value *string)() {
    m.function_name = value
}
// SetPackageAddress sets the package_address property value. The Bech32m-encoded human readable version of the package address
func (m *BlueprintFunctionTargetIdentifier) SetPackageAddress(value *string)() {
    m.package_address = value
}
type BlueprintFunctionTargetIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetIdentifierable
    GetBlueprintName()(*string)
    GetFunctionName()(*string)
    GetPackageAddress()(*string)
    SetBlueprintName(value *string)()
    SetFunctionName(value *string)()
    SetPackageAddress(value *string)()
}
