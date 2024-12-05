package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldKernelBootValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This was added in Cuttlefish. Before that, this value was missing, but can be assumed to be V1.
    always_visible_nodes_version *AlwaysVisibleGlobalNodesVersion
}
// NewBootLoaderModuleFieldKernelBootValue instantiates a new BootLoaderModuleFieldKernelBootValue and sets the default values.
func NewBootLoaderModuleFieldKernelBootValue()(*BootLoaderModuleFieldKernelBootValue) {
    m := &BootLoaderModuleFieldKernelBootValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBootLoaderModuleFieldKernelBootValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldKernelBootValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBootLoaderModuleFieldKernelBootValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BootLoaderModuleFieldKernelBootValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlwaysVisibleNodesVersion gets the always_visible_nodes_version property value. This was added in Cuttlefish. Before that, this value was missing, but can be assumed to be V1.
// returns a *AlwaysVisibleGlobalNodesVersion when successful
func (m *BootLoaderModuleFieldKernelBootValue) GetAlwaysVisibleNodesVersion()(*AlwaysVisibleGlobalNodesVersion) {
    return m.always_visible_nodes_version
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldKernelBootValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["always_visible_nodes_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlwaysVisibleGlobalNodesVersion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysVisibleNodesVersion(val.(*AlwaysVisibleGlobalNodesVersion))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BootLoaderModuleFieldKernelBootValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAlwaysVisibleNodesVersion() != nil {
        cast := (*m.GetAlwaysVisibleNodesVersion()).String()
        err := writer.WriteStringValue("always_visible_nodes_version", &cast)
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
func (m *BootLoaderModuleFieldKernelBootValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlwaysVisibleNodesVersion sets the always_visible_nodes_version property value. This was added in Cuttlefish. Before that, this value was missing, but can be assumed to be V1.
func (m *BootLoaderModuleFieldKernelBootValue) SetAlwaysVisibleNodesVersion(value *AlwaysVisibleGlobalNodesVersion)() {
    m.always_visible_nodes_version = value
}
type BootLoaderModuleFieldKernelBootValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlwaysVisibleNodesVersion()(*AlwaysVisibleGlobalNodesVersion)
    SetAlwaysVisibleNodesVersion(value *AlwaysVisibleGlobalNodesVersion)()
}
