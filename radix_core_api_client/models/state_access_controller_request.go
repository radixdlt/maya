package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateAccessControllerRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the component address
    controller_address *string
    // The logical name of the network
    network *string
}
// NewStateAccessControllerRequest instantiates a new StateAccessControllerRequest and sets the default values.
func NewStateAccessControllerRequest()(*StateAccessControllerRequest) {
    m := &StateAccessControllerRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateAccessControllerRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateAccessControllerRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateAccessControllerRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateAccessControllerRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetControllerAddress gets the controller_address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *StateAccessControllerRequest) GetControllerAddress()(*string) {
    return m.controller_address
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateAccessControllerRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["controller_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControllerAddress(val)
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
    return res
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StateAccessControllerRequest) GetNetwork()(*string) {
    return m.network
}
// Serialize serializes information the current object
func (m *StateAccessControllerRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("controller_address", m.GetControllerAddress())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StateAccessControllerRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetControllerAddress sets the controller_address property value. The Bech32m-encoded human readable version of the component address
func (m *StateAccessControllerRequest) SetControllerAddress(value *string)() {
    m.controller_address = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *StateAccessControllerRequest) SetNetwork(value *string)() {
    m.network = value
}
type StateAccessControllerRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetControllerAddress()(*string)
    GetNetwork()(*string)
    SetControllerAddress(value *string)()
    SetNetwork(value *string)()
}
