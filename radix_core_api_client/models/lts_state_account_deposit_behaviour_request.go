package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStateAccountDepositBehaviourRequest struct {
    // The Bech32m-encoded human readable version of the account's address.
    account_address *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The badge property
    badge PresentedBadgeable
    // The logical name of the network
    network *string
    // The resource addresses to check the deposit behaviours of.
    resource_addresses []string
}
// NewLtsStateAccountDepositBehaviourRequest instantiates a new LtsStateAccountDepositBehaviourRequest and sets the default values.
func NewLtsStateAccountDepositBehaviourRequest()(*LtsStateAccountDepositBehaviourRequest) {
    m := &LtsStateAccountDepositBehaviourRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsStateAccountDepositBehaviourRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStateAccountDepositBehaviourRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsStateAccountDepositBehaviourRequest(), nil
}
// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address.
// returns a *string when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetAccountAddress()(*string) {
    return m.account_address
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBadge gets the badge property value. The badge property
// returns a PresentedBadgeable when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetBadge()(PresentedBadgeable) {
    return m.badge
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountAddress(val)
        }
        return nil
    }
    res["badge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePresentedBadgeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBadge(val.(PresentedBadgeable))
        }
        return nil
    }
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
    res["resource_addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResourceAddresses(res)
        }
        return nil
    }
    return res
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetNetwork()(*string) {
    return m.network
}
// GetResourceAddresses gets the resource_addresses property value. The resource addresses to check the deposit behaviours of.
// returns a []string when successful
func (m *LtsStateAccountDepositBehaviourRequest) GetResourceAddresses()([]string) {
    return m.resource_addresses
}
// Serialize serializes information the current object
func (m *LtsStateAccountDepositBehaviourRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_address", m.GetAccountAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("badge", m.GetBadge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetResourceAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("resource_addresses", m.GetResourceAddresses())
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
// SetAccountAddress sets the account_address property value. The Bech32m-encoded human readable version of the account's address.
func (m *LtsStateAccountDepositBehaviourRequest) SetAccountAddress(value *string)() {
    m.account_address = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStateAccountDepositBehaviourRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBadge sets the badge property value. The badge property
func (m *LtsStateAccountDepositBehaviourRequest) SetBadge(value PresentedBadgeable)() {
    m.badge = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *LtsStateAccountDepositBehaviourRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetResourceAddresses sets the resource_addresses property value. The resource addresses to check the deposit behaviours of.
func (m *LtsStateAccountDepositBehaviourRequest) SetResourceAddresses(value []string)() {
    m.resource_addresses = value
}
type LtsStateAccountDepositBehaviourRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountAddress()(*string)
    GetBadge()(PresentedBadgeable)
    GetNetwork()(*string)
    GetResourceAddresses()([]string)
    SetAccountAddress(value *string)()
    SetBadge(value PresentedBadgeable)()
    SetNetwork(value *string)()
    SetResourceAddresses(value []string)()
}
