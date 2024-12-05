package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HookExport struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The export property
    export PackageExportable
    // The object_hook property
    object_hook *ObjectHook
}
// NewHookExport instantiates a new HookExport and sets the default values.
func NewHookExport()(*HookExport) {
    m := &HookExport{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookExportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookExportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookExport(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HookExport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExport gets the export property value. The export property
// returns a PackageExportable when successful
func (m *HookExport) GetExport()(PackageExportable) {
    return m.export
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HookExport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["export"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageExportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExport(val.(PackageExportable))
        }
        return nil
    }
    res["object_hook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectHook)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectHook(val.(*ObjectHook))
        }
        return nil
    }
    return res
}
// GetObjectHook gets the object_hook property value. The object_hook property
// returns a *ObjectHook when successful
func (m *HookExport) GetObjectHook()(*ObjectHook) {
    return m.object_hook
}
// Serialize serializes information the current object
func (m *HookExport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("export", m.GetExport())
        if err != nil {
            return err
        }
    }
    if m.GetObjectHook() != nil {
        cast := (*m.GetObjectHook()).String()
        err := writer.WriteStringValue("object_hook", &cast)
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
func (m *HookExport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExport sets the export property value. The export property
func (m *HookExport) SetExport(value PackageExportable)() {
    m.export = value
}
// SetObjectHook sets the object_hook property value. The object_hook property
func (m *HookExport) SetObjectHook(value *ObjectHook)() {
    m.object_hook = value
}
type HookExportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExport()(PackageExportable)
    GetObjectHook()(*ObjectHook)
    SetExport(value PackageExportable)()
    SetObjectHook(value *ObjectHook)()
}
