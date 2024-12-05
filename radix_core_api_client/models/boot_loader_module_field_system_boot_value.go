package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldSystemBootValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The system_parameters property
    system_parameters SystemParametersable
    // The SystemVersion was added at Cuttlefish. Before that it can be assumed to be V1.
    system_version *SystemVersion
}
// NewBootLoaderModuleFieldSystemBootValue instantiates a new BootLoaderModuleFieldSystemBootValue and sets the default values.
func NewBootLoaderModuleFieldSystemBootValue()(*BootLoaderModuleFieldSystemBootValue) {
    m := &BootLoaderModuleFieldSystemBootValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBootLoaderModuleFieldSystemBootValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldSystemBootValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBootLoaderModuleFieldSystemBootValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BootLoaderModuleFieldSystemBootValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldSystemBootValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["system_parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSystemParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemParameters(val.(SystemParametersable))
        }
        return nil
    }
    res["system_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSystemVersion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemVersion(val.(*SystemVersion))
        }
        return nil
    }
    return res
}
// GetSystemParameters gets the system_parameters property value. The system_parameters property
// returns a SystemParametersable when successful
func (m *BootLoaderModuleFieldSystemBootValue) GetSystemParameters()(SystemParametersable) {
    return m.system_parameters
}
// GetSystemVersion gets the system_version property value. The SystemVersion was added at Cuttlefish. Before that it can be assumed to be V1.
// returns a *SystemVersion when successful
func (m *BootLoaderModuleFieldSystemBootValue) GetSystemVersion()(*SystemVersion) {
    return m.system_version
}
// Serialize serializes information the current object
func (m *BootLoaderModuleFieldSystemBootValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("system_parameters", m.GetSystemParameters())
        if err != nil {
            return err
        }
    }
    if m.GetSystemVersion() != nil {
        cast := (*m.GetSystemVersion()).String()
        err := writer.WriteStringValue("system_version", &cast)
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
func (m *BootLoaderModuleFieldSystemBootValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSystemParameters sets the system_parameters property value. The system_parameters property
func (m *BootLoaderModuleFieldSystemBootValue) SetSystemParameters(value SystemParametersable)() {
    m.system_parameters = value
}
// SetSystemVersion sets the system_version property value. The SystemVersion was added at Cuttlefish. Before that it can be assumed to be V1.
func (m *BootLoaderModuleFieldSystemBootValue) SetSystemVersion(value *SystemVersion)() {
    m.system_version = value
}
type BootLoaderModuleFieldSystemBootValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSystemParameters()(SystemParametersable)
    GetSystemVersion()(*SystemVersion)
    SetSystemParameters(value SystemParametersable)()
    SetSystemVersion(value *SystemVersion)()
}
