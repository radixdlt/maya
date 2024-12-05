package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateComponentRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the component's global address
    component_address *string
    // The logical name of the network
    network *string
}
// NewStateComponentRequest instantiates a new StateComponentRequest and sets the default values.
func NewStateComponentRequest()(*StateComponentRequest) {
    m := &StateComponentRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStateComponentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateComponentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStateComponentRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateComponentRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComponentAddress gets the component_address property value. The Bech32m-encoded human readable version of the component's global address
// returns a *string when successful
func (m *StateComponentRequest) GetComponentAddress()(*string) {
    return m.component_address
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateComponentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *StateComponentRequest) GetNetwork()(*string) {
    return m.network
}
// Serialize serializes information the current object
func (m *StateComponentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("component_address", m.GetComponentAddress())
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
func (m *StateComponentRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComponentAddress sets the component_address property value. The Bech32m-encoded human readable version of the component's global address
func (m *StateComponentRequest) SetComponentAddress(value *string)() {
    m.component_address = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *StateComponentRequest) SetNetwork(value *string)() {
    m.network = value
}
type StateComponentRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComponentAddress()(*string)
    GetNetwork()(*string)
    SetComponentAddress(value *string)()
    SetNetwork(value *string)()
}
