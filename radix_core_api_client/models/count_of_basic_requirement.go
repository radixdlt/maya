package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CountOfBasicRequirement struct {
    BasicRequirement
    // The count property
    count *int32
    // The list property
    list []Requirementable
}
// NewCountOfBasicRequirement instantiates a new CountOfBasicRequirement and sets the default values.
func NewCountOfBasicRequirement()(*CountOfBasicRequirement) {
    m := &CountOfBasicRequirement{
        BasicRequirement: *NewBasicRequirement(),
    }
    return m
}
// CreateCountOfBasicRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCountOfBasicRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCountOfBasicRequirement(), nil
}
// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *CountOfBasicRequirement) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CountOfBasicRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BasicRequirement.GetFieldDeserializers()
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
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
func (m *CountOfBasicRequirement) GetList()([]Requirementable) {
    return m.list
}
// Serialize serializes information the current object
func (m *CountOfBasicRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BasicRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
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
// SetCount sets the count property value. The count property
func (m *CountOfBasicRequirement) SetCount(value *int32)() {
    m.count = value
}
// SetList sets the list property value. The list property
func (m *CountOfBasicRequirement) SetList(value []Requirementable)() {
    m.list = value
}
type CountOfBasicRequirementable interface {
    BasicRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetList()([]Requirementable)
    SetCount(value *int32)()
    SetList(value []Requirementable)()
}
