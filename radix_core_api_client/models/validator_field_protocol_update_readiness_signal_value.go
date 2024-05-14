package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidatorFieldProtocolUpdateReadinessSignalValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// If present, indicates the validator is currently signalling readiness for the given protocol version.Is validated to be exactly 32 chars long (if it exists).
	protocol_version_name *string
}

// NewValidatorFieldProtocolUpdateReadinessSignalValue instantiates a new ValidatorFieldProtocolUpdateReadinessSignalValue and sets the default values.
func NewValidatorFieldProtocolUpdateReadinessSignalValue() *ValidatorFieldProtocolUpdateReadinessSignalValue {
	m := &ValidatorFieldProtocolUpdateReadinessSignalValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateValidatorFieldProtocolUpdateReadinessSignalValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidatorFieldProtocolUpdateReadinessSignalValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewValidatorFieldProtocolUpdateReadinessSignalValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["protocol_version_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProtocolVersionName(val)
		}
		return nil
	}
	return res
}

// GetProtocolVersionName gets the protocol_version_name property value. If present, indicates the validator is currently signalling readiness for the given protocol version.Is validated to be exactly 32 chars long (if it exists).
// returns a *string when successful
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) GetProtocolVersionName() *string {
	return m.protocol_version_name
}

// Serialize serializes information the current object
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("protocol_version_name", m.GetProtocolVersionName())
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
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProtocolVersionName sets the protocol_version_name property value. If present, indicates the validator is currently signalling readiness for the given protocol version.Is validated to be exactly 32 chars long (if it exists).
func (m *ValidatorFieldProtocolUpdateReadinessSignalValue) SetProtocolVersionName(value *string) {
	m.protocol_version_name = value
}

type ValidatorFieldProtocolUpdateReadinessSignalValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProtocolVersionName() *string
	SetProtocolVersionName(value *string)
}
