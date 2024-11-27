package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResourceRequirement struct {
    Requirement
    // The Bech32m-encoded human readable version of the resource address
    resource *string
}
// NewResourceRequirement instantiates a new ResourceRequirement and sets the default values.
func NewResourceRequirement()(*ResourceRequirement) {
    m := &ResourceRequirement{
        Requirement: *NewRequirement(),
    }
    return m
}
// CreateResourceRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceRequirement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResourceRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Requirement.GetFieldDeserializers()
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ResourceRequirement) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *ResourceRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Requirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResource sets the resource property value. The Bech32m-encoded human readable version of the resource address
func (m *ResourceRequirement) SetResource(value *string)() {
    m.resource = value
}
type ResourceRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Requirementable
    GetResource()(*string)
    SetResource(value *string)()
}
