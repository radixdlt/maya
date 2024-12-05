package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageCodeVmTypeEntryValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The vm_type property
    vm_type *VmType
}
// NewPackageCodeVmTypeEntryValue instantiates a new PackageCodeVmTypeEntryValue and sets the default values.
func NewPackageCodeVmTypeEntryValue()(*PackageCodeVmTypeEntryValue) {
    m := &PackageCodeVmTypeEntryValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageCodeVmTypeEntryValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageCodeVmTypeEntryValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageCodeVmTypeEntryValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageCodeVmTypeEntryValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageCodeVmTypeEntryValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["vm_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmType(val.(*VmType))
        }
        return nil
    }
    return res
}
// GetVmType gets the vm_type property value. The vm_type property
// returns a *VmType when successful
func (m *PackageCodeVmTypeEntryValue) GetVmType()(*VmType) {
    return m.vm_type
}
// Serialize serializes information the current object
func (m *PackageCodeVmTypeEntryValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetVmType() != nil {
        cast := (*m.GetVmType()).String()
        err := writer.WriteStringValue("vm_type", &cast)
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
func (m *PackageCodeVmTypeEntryValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVmType sets the vm_type property value. The vm_type property
func (m *PackageCodeVmTypeEntryValue) SetVmType(value *VmType)() {
    m.vm_type = value
}
type PackageCodeVmTypeEntryValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVmType()(*VmType)
    SetVmType(value *VmType)()
}
