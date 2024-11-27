package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubstateFormatOptions requested substate formats to include in the response
type SubstateFormatOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to return the raw substate value bytes hash (default false)
    hash *bool
    // Whether to return the previous substate value for updates and deletes (default false)
    previous *bool
    // Whether to return the raw substate value bytes (default false)
    raw *bool
    // Whether to return the typed substate information (default true)
    typed *bool
}
// NewSubstateFormatOptions instantiates a new SubstateFormatOptions and sets the default values.
func NewSubstateFormatOptions()(*SubstateFormatOptions) {
    m := &SubstateFormatOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubstateFormatOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateFormatOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubstateFormatOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubstateFormatOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubstateFormatOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHash(val)
        }
        return nil
    }
    res["previous"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrevious(val)
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
    res["typed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTyped(val)
        }
        return nil
    }
    return res
}
// GetHash gets the hash property value. Whether to return the raw substate value bytes hash (default false)
// returns a *bool when successful
func (m *SubstateFormatOptions) GetHash()(*bool) {
    return m.hash
}
// GetPrevious gets the previous property value. Whether to return the previous substate value for updates and deletes (default false)
// returns a *bool when successful
func (m *SubstateFormatOptions) GetPrevious()(*bool) {
    return m.previous
}
// GetRaw gets the raw property value. Whether to return the raw substate value bytes (default false)
// returns a *bool when successful
func (m *SubstateFormatOptions) GetRaw()(*bool) {
    return m.raw
}
// GetTyped gets the typed property value. Whether to return the typed substate information (default true)
// returns a *bool when successful
func (m *SubstateFormatOptions) GetTyped()(*bool) {
    return m.typed
}
// Serialize serializes information the current object
func (m *SubstateFormatOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hash", m.GetHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("previous", m.GetPrevious())
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
        err := writer.WriteBoolValue("typed", m.GetTyped())
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
func (m *SubstateFormatOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHash sets the hash property value. Whether to return the raw substate value bytes hash (default false)
func (m *SubstateFormatOptions) SetHash(value *bool)() {
    m.hash = value
}
// SetPrevious sets the previous property value. Whether to return the previous substate value for updates and deletes (default false)
func (m *SubstateFormatOptions) SetPrevious(value *bool)() {
    m.previous = value
}
// SetRaw sets the raw property value. Whether to return the raw substate value bytes (default false)
func (m *SubstateFormatOptions) SetRaw(value *bool)() {
    m.raw = value
}
// SetTyped sets the typed property value. Whether to return the typed substate information (default true)
func (m *SubstateFormatOptions) SetTyped(value *bool)() {
    m.typed = value
}
type SubstateFormatOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHash()(*bool)
    GetPrevious()(*bool)
    GetRaw()(*bool)
    GetTyped()(*bool)
    SetHash(value *bool)()
    SetPrevious(value *bool)()
    SetRaw(value *bool)()
    SetTyped(value *bool)()
}
