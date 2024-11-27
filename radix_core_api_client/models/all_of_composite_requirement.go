package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllOfCompositeRequirement struct {
    CompositeRequirement
    // The access_rules property
    access_rules []CompositeRequirementable
}
// NewAllOfCompositeRequirement instantiates a new AllOfCompositeRequirement and sets the default values.
func NewAllOfCompositeRequirement()(*AllOfCompositeRequirement) {
    m := &AllOfCompositeRequirement{
        CompositeRequirement: *NewCompositeRequirement(),
    }
    return m
}
// CreateAllOfCompositeRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllOfCompositeRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllOfCompositeRequirement(), nil
}
// GetAccessRules gets the access_rules property value. The access_rules property
// returns a []CompositeRequirementable when successful
func (m *AllOfCompositeRequirement) GetAccessRules()([]CompositeRequirementable) {
    return m.access_rules
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllOfCompositeRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CompositeRequirement.GetFieldDeserializers()
    res["access_rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompositeRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompositeRequirementable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompositeRequirementable)
                }
            }
            m.SetAccessRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AllOfCompositeRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CompositeRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessRules()))
        for i, v := range m.GetAccessRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("access_rules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessRules sets the access_rules property value. The access_rules property
func (m *AllOfCompositeRequirement) SetAccessRules(value []CompositeRequirementable)() {
    m.access_rules = value
}
type AllOfCompositeRequirementable interface {
    CompositeRequirementable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRules()([]CompositeRequirementable)
    SetAccessRules(value []CompositeRequirementable)()
}
