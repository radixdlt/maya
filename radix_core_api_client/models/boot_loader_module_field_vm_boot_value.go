package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BootLoaderModuleFieldVmBootValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The scrypto_v1_minor_version property
	scrypto_v1_minor_version *int64
}

// NewBootLoaderModuleFieldVmBootValue instantiates a new BootLoaderModuleFieldVmBootValue and sets the default values.
func NewBootLoaderModuleFieldVmBootValue() *BootLoaderModuleFieldVmBootValue {
	m := &BootLoaderModuleFieldVmBootValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBootLoaderModuleFieldVmBootValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBootLoaderModuleFieldVmBootValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBootLoaderModuleFieldVmBootValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BootLoaderModuleFieldVmBootValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BootLoaderModuleFieldVmBootValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["scrypto_v1_minor_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetScryptoV1MinorVersion(val)
		}
		return nil
	}
	return res
}

// GetScryptoV1MinorVersion gets the scrypto_v1_minor_version property value. The scrypto_v1_minor_version property
// returns a *int64 when successful
func (m *BootLoaderModuleFieldVmBootValue) GetScryptoV1MinorVersion() *int64 {
	return m.scrypto_v1_minor_version
}

// Serialize serializes information the current object
func (m *BootLoaderModuleFieldVmBootValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("scrypto_v1_minor_version", m.GetScryptoV1MinorVersion())
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
func (m *BootLoaderModuleFieldVmBootValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetScryptoV1MinorVersion sets the scrypto_v1_minor_version property value. The scrypto_v1_minor_version property
func (m *BootLoaderModuleFieldVmBootValue) SetScryptoV1MinorVersion(value *int64) {
	m.scrypto_v1_minor_version = value
}

type BootLoaderModuleFieldVmBootValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetScryptoV1MinorVersion() *int64
	SetScryptoV1MinorVersion(value *int64)
}
