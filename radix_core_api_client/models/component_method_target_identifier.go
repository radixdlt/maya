package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ComponentMethodTargetIdentifier struct {
    TargetIdentifier
    // The Bech32m-encoded human readable version of the component address
    component_address *string
    // The method_name property
    method_name *string
}
// NewComponentMethodTargetIdentifier instantiates a new ComponentMethodTargetIdentifier and sets the default values.
func NewComponentMethodTargetIdentifier()(*ComponentMethodTargetIdentifier) {
    m := &ComponentMethodTargetIdentifier{
        TargetIdentifier: *NewTargetIdentifier(),
    }
    return m
}
// CreateComponentMethodTargetIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComponentMethodTargetIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComponentMethodTargetIdentifier(), nil
}
// GetComponentAddress gets the component_address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *ComponentMethodTargetIdentifier) GetComponentAddress()(*string) {
    return m.component_address
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ComponentMethodTargetIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetIdentifier.GetFieldDeserializers()
    res["component_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentAddress(val)
        }
        return nil
    }
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
    return res
}
// GetMethodName gets the method_name property value. The method_name property
// returns a *string when successful
func (m *ComponentMethodTargetIdentifier) GetMethodName()(*string) {
    return m.method_name
}
// Serialize serializes information the current object
func (m *ComponentMethodTargetIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("component_address", m.GetComponentAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("method_name", m.GetMethodName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComponentAddress sets the component_address property value. The Bech32m-encoded human readable version of the component address
func (m *ComponentMethodTargetIdentifier) SetComponentAddress(value *string)() {
    m.component_address = value
}
// SetMethodName sets the method_name property value. The method_name property
func (m *ComponentMethodTargetIdentifier) SetMethodName(value *string)() {
    m.method_name = value
}
type ComponentMethodTargetIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetIdentifierable
    GetComponentAddress()(*string)
    GetMethodName()(*string)
    SetComponentAddress(value *string)()
    SetMethodName(value *string)()
}
