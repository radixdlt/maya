package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintRoyaltyConfig struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The is_enabled property
	is_enabled *bool
	// The royalty rules by method. The array is only present if royalties are enabled.
	method_rules []BlueprintMethodRoyaltyable
}

// NewBlueprintRoyaltyConfig instantiates a new BlueprintRoyaltyConfig and sets the default values.
func NewBlueprintRoyaltyConfig() *BlueprintRoyaltyConfig {
	m := &BlueprintRoyaltyConfig{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBlueprintRoyaltyConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintRoyaltyConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBlueprintRoyaltyConfig(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintRoyaltyConfig) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintRoyaltyConfig) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["is_enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsEnabled(val)
		}
		return nil
	}
	res["method_rules"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateBlueprintMethodRoyaltyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]BlueprintMethodRoyaltyable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(BlueprintMethodRoyaltyable)
				}
			}
			m.SetMethodRules(res)
		}
		return nil
	}
	return res
}

// GetIsEnabled gets the is_enabled property value. The is_enabled property
// returns a *bool when successful
func (m *BlueprintRoyaltyConfig) GetIsEnabled() *bool {
	return m.is_enabled
}

// GetMethodRules gets the method_rules property value. The royalty rules by method. The array is only present if royalties are enabled.
// returns a []BlueprintMethodRoyaltyable when successful
func (m *BlueprintRoyaltyConfig) GetMethodRules() []BlueprintMethodRoyaltyable {
	return m.method_rules
}

// Serialize serializes information the current object
func (m *BlueprintRoyaltyConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("is_enabled", m.GetIsEnabled())
		if err != nil {
			return err
		}
	}
	if m.GetMethodRules() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMethodRules()))
		for i, v := range m.GetMethodRules() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("method_rules", cast)
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
func (m *BlueprintRoyaltyConfig) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIsEnabled sets the is_enabled property value. The is_enabled property
func (m *BlueprintRoyaltyConfig) SetIsEnabled(value *bool) {
	m.is_enabled = value
}

// SetMethodRules sets the method_rules property value. The royalty rules by method. The array is only present if royalties are enabled.
func (m *BlueprintRoyaltyConfig) SetMethodRules(value []BlueprintMethodRoyaltyable) {
	m.method_rules = value
}

type BlueprintRoyaltyConfigable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIsEnabled() *bool
	GetMethodRules() []BlueprintMethodRoyaltyable
	SetIsEnabled(value *bool)
	SetMethodRules(value []BlueprintMethodRoyaltyable)
}
