package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StatePackageRequest struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The logical name of the network
	network *string
	// The Bech32m-encoded human readable version of the package's global address
	package_address *string
}

// NewStatePackageRequest instantiates a new StatePackageRequest and sets the default values.
func NewStatePackageRequest() *StatePackageRequest {
	m := &StatePackageRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStatePackageRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatePackageRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStatePackageRequest(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StatePackageRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StatePackageRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["package_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPackageAddress(val)
		}
		return nil
	}
	return res
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StatePackageRequest) GetNetwork() *string {
	return m.network
}

// GetPackageAddress gets the package_address property value. The Bech32m-encoded human readable version of the package's global address
// returns a *string when successful
func (m *StatePackageRequest) GetPackageAddress() *string {
	return m.package_address
}

// Serialize serializes information the current object
func (m *StatePackageRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("package_address", m.GetPackageAddress())
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
func (m *StatePackageRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *StatePackageRequest) SetNetwork(value *string) {
	m.network = value
}

// SetPackageAddress sets the package_address property value. The Bech32m-encoded human readable version of the package's global address
func (m *StatePackageRequest) SetPackageAddress(value *string) {
	m.package_address = value
}

type StatePackageRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNetwork() *string
	GetPackageAddress() *string
	SetNetwork(value *string)
	SetPackageAddress(value *string)
}
