package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllOfBasicRequirement struct {
    BasicRequirement
    // The list property
    list []Requirementable
}
// NewAllOfBasicRequirement instantiates a new AllOfBasicRequirement and sets the default values.
func NewAllOfBasicRequirement()(*AllOfBasicRequirement) {
    m := &AllOfBasicRequirement{
        BasicRequirement: *NewBasicRequirement(),
    }
    return m
}
// CreateAllOfBasicRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllOfBasicRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllOfBasicRequirement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllOfBasicRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BasicRequirement.GetFieldDeserializers()
    res["list"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Requirementable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Requirementable)
                }
            }
            m.SetList(res)
        }
        return nil
    }
    return res
}
// GetList gets the list property value. The list property
// returns a []Requirementable when successful
func (m *AllOfBasicRequirement) GetList()([]Requirementable) {
    return m.list
}
// Serialize serializes information the current object
func (m *AllOfBasicRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BasicRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetList()))
        for i, v := range m.GetList() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("list", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetList sets the list property value. The list property
func (m *AllOfBasicRequirement) SetList(value []Requirementable)() {
    m.list = value
}
type AllOfBasicRequirementable interface {
    BasicRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetList()([]Requirementable)
    SetList(value []Requirementable)()
}
