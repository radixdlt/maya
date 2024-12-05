package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SborFormatOptions requested SBOR formats to include in the response
type SborFormatOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to return the programmatic json format (normally default true, defaults false for streamed transactions)
    programmatic_json *bool
    // Whether to return the raw hex-encoded bytes (default true)
    raw *bool
}
// NewSborFormatOptions instantiates a new SborFormatOptions and sets the default values.
func NewSborFormatOptions()(*SborFormatOptions) {
    m := &SborFormatOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSborFormatOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSborFormatOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSborFormatOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SborFormatOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SborFormatOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["programmatic_json"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgrammaticJson(val)
        }
        return nil
    }
    res["raw"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRaw(val)
        }
        return nil
    }
    return res
}
// GetProgrammaticJson gets the programmatic_json property value. Whether to return the programmatic json format (normally default true, defaults false for streamed transactions)
// returns a *bool when successful
func (m *SborFormatOptions) GetProgrammaticJson()(*bool) {
    return m.programmatic_json
}
// GetRaw gets the raw property value. Whether to return the raw hex-encoded bytes (default true)
// returns a *bool when successful
func (m *SborFormatOptions) GetRaw()(*bool) {
    return m.raw
}
// Serialize serializes information the current object
func (m *SborFormatOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("programmatic_json", m.GetProgrammaticJson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("raw", m.GetRaw())
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
func (m *SborFormatOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProgrammaticJson sets the programmatic_json property value. Whether to return the programmatic json format (normally default true, defaults false for streamed transactions)
func (m *SborFormatOptions) SetProgrammaticJson(value *bool)() {
    m.programmatic_json = value
}
// SetRaw sets the raw property value. Whether to return the raw hex-encoded bytes (default true)
func (m *SborFormatOptions) SetRaw(value *bool)() {
    m.raw = value
}
type SborFormatOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProgrammaticJson()(*bool)
    GetRaw()(*bool)
    SetProgrammaticJson(value *bool)()
    SetRaw(value *bool)()
}
