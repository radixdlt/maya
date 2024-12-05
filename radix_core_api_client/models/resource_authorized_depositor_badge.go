package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResourceAuthorizedDepositorBadge struct {
    AuthorizedDepositorBadge
    // The Bech32m-encoded human readable version of the resource address
    resource_address *string
}
// NewResourceAuthorizedDepositorBadge instantiates a new ResourceAuthorizedDepositorBadge and sets the default values.
func NewResourceAuthorizedDepositorBadge()(*ResourceAuthorizedDepositorBadge) {
    m := &ResourceAuthorizedDepositorBadge{
        AuthorizedDepositorBadge: *NewAuthorizedDepositorBadge(),
    }
    return m
}
// CreateResourceAuthorizedDepositorBadgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceAuthorizedDepositorBadgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAuthorizedDepositorBadge(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResourceAuthorizedDepositorBadge) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizedDepositorBadge.GetFieldDeserializers()
    res["resource_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAddress(val)
        }
        return nil
    }
    return res
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource address
// returns a *string when successful
func (m *ResourceAuthorizedDepositorBadge) GetResourceAddress()(*string) {
    return m.resource_address
}
// Serialize serializes information the current object
func (m *ResourceAuthorizedDepositorBadge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizedDepositorBadge.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resource_address", m.GetResourceAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource address
func (m *ResourceAuthorizedDepositorBadge) SetResourceAddress(value *string)() {
    m.resource_address = value
}
type ResourceAuthorizedDepositorBadgeable interface {
    AuthorizedDepositorBadgeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceAddress()(*string)
    SetResourceAddress(value *string)()
}
