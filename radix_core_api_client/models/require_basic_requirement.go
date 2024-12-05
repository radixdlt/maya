package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequireBasicRequirement struct {
    BasicRequirement
    // This is called `ResourceOrNonFungible` in the engine.
    requirement Requirementable
}
// NewRequireBasicRequirement instantiates a new RequireBasicRequirement and sets the default values.
func NewRequireBasicRequirement()(*RequireBasicRequirement) {
    m := &RequireBasicRequirement{
        BasicRequirement: *NewBasicRequirement(),
    }
    return m
}
// CreateRequireBasicRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequireBasicRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequireBasicRequirement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequireBasicRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BasicRequirement.GetFieldDeserializers()
    res["requirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequirement(val.(Requirementable))
        }
        return nil
    }
    return res
}
// GetRequirement gets the requirement property value. This is called `ResourceOrNonFungible` in the engine.
// returns a Requirementable when successful
func (m *RequireBasicRequirement) GetRequirement()(Requirementable) {
    return m.requirement
}
// Serialize serializes information the current object
func (m *RequireBasicRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BasicRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("requirement", m.GetRequirement())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRequirement sets the requirement property value. This is called `ResourceOrNonFungible` in the engine.
func (m *RequireBasicRequirement) SetRequirement(value Requirementable)() {
    m.requirement = value
}
type RequireBasicRequirementable interface {
    BasicRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequirement()(Requirementable)
    SetRequirement(value Requirementable)()
}
