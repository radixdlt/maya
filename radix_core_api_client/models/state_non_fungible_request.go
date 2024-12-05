package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateNonFungibleRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The logical name of the network
    network *string
    // The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
    non_fungible_id *string
    // The Bech32m-encoded human readable version of the resource's global address
    resource_address *string
}
// NewStateNonFungibleRequest instantiates a new StateNonFungibleRequest and sets the default values.
func NewStateNonFungibleRequest()(*StateNonFungibleRequest) {
    m := &StateNonFungibleRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateNonFungibleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateNonFungibleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateNonFungibleRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateNonFungibleRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateNonFungibleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["non_fungible_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonFungibleId(val)
        }
        return nil
    }
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
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StateNonFungibleRequest) GetNetwork()(*string) {
    return m.network
}
// GetNonFungibleId gets the non_fungible_id property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
// returns a *string when successful
func (m *StateNonFungibleRequest) GetNonFungibleId()(*string) {
    return m.non_fungible_id
}
// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource's global address
// returns a *string when successful
func (m *StateNonFungibleRequest) GetResourceAddress()(*string) {
    return m.resource_address
}
// Serialize serializes information the current object
func (m *StateNonFungibleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("non_fungible_id", m.GetNonFungibleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource_address", m.GetResourceAddress())
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
func (m *StateNonFungibleRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *StateNonFungibleRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetNonFungibleId sets the non_fungible_id property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
func (m *StateNonFungibleRequest) SetNonFungibleId(value *string)() {
    m.non_fungible_id = value
}
// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource's global address
func (m *StateNonFungibleRequest) SetResourceAddress(value *string)() {
    m.resource_address = value
}
type StateNonFungibleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetwork()(*string)
    GetNonFungibleId()(*string)
    GetResourceAddress()(*string)
    SetNetwork(value *string)()
    SetNonFungibleId(value *string)()
    SetResourceAddress(value *string)()
}
