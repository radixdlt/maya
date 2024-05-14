package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StateResourceRequest struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The logical name of the network
	network *string
	// The Bech32m-encoded human readable version of the resource's global address
	resource_address *string
}

// NewStateResourceRequest instantiates a new StateResourceRequest and sets the default values.
func NewStateResourceRequest() *StateResourceRequest {
	m := &StateResourceRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStateResourceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStateResourceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStateResourceRequest(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StateResourceRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StateResourceRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
		}
		return nil
	}
	res["resource_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *StateResourceRequest) GetNetwork() *string {
	return m.network
}

// GetResourceAddress gets the resource_address property value. The Bech32m-encoded human readable version of the resource's global address
// returns a *string when successful
func (m *StateResourceRequest) GetResourceAddress() *string {
	return m.resource_address
}

// Serialize serializes information the current object
func (m *StateResourceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
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
func (m *StateResourceRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *StateResourceRequest) SetNetwork(value *string) {
	m.network = value
}

// SetResourceAddress sets the resource_address property value. The Bech32m-encoded human readable version of the resource's global address
func (m *StateResourceRequest) SetResourceAddress(value *string) {
	m.resource_address = value
}

type StateResourceRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNetwork() *string
	GetResourceAddress() *string
	SetNetwork(value *string)
	SetResourceAddress(value *string)
}
