package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintInfo struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The blueprint_name property
	blueprint_name *string
	// The blueprint_version property
	blueprint_version *string
	// The features property
	features []string
	// The generic_substitutions property
	generic_substitutions []GenericSubstitutionable
	// The Bech32m-encoded human readable version of any global address
	outer_object *string
	// The Bech32m-encoded human readable version of the package address
	package_address *string
}

// NewBlueprintInfo instantiates a new BlueprintInfo and sets the default values.
func NewBlueprintInfo() *BlueprintInfo {
	m := &BlueprintInfo{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBlueprintInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBlueprintInfo(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintInfo) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBlueprintName gets the blueprint_name property value. The blueprint_name property
// returns a *string when successful
func (m *BlueprintInfo) GetBlueprintName() *string {
	return m.blueprint_name
}

// GetBlueprintVersion gets the blueprint_version property value. The blueprint_version property
// returns a *string when successful
func (m *BlueprintInfo) GetBlueprintVersion() *string {
	return m.blueprint_version
}

// GetFeatures gets the features property value. The features property
// returns a []string when successful
func (m *BlueprintInfo) GetFeatures() []string {
	return m.features
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintInfo) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["blueprint_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlueprintName(val)
		}
		return nil
	}
	res["blueprint_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBlueprintVersion(val)
		}
		return nil
	}
	res["features"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetFeatures(res)
		}
		return nil
	}
	res["generic_substitutions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateGenericSubstitutionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]GenericSubstitutionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(GenericSubstitutionable)
				}
			}
			m.SetGenericSubstitutions(res)
		}
		return nil
	}
	res["outer_object"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOuterObject(val)
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

// GetGenericSubstitutions gets the generic_substitutions property value. The generic_substitutions property
// returns a []GenericSubstitutionable when successful
func (m *BlueprintInfo) GetGenericSubstitutions() []GenericSubstitutionable {
	return m.generic_substitutions
}

// GetOuterObject gets the outer_object property value. The Bech32m-encoded human readable version of any global address
// returns a *string when successful
func (m *BlueprintInfo) GetOuterObject() *string {
	return m.outer_object
}

// GetPackageAddress gets the package_address property value. The Bech32m-encoded human readable version of the package address
// returns a *string when successful
func (m *BlueprintInfo) GetPackageAddress() *string {
	return m.package_address
}

// Serialize serializes information the current object
func (m *BlueprintInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("blueprint_name", m.GetBlueprintName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("blueprint_version", m.GetBlueprintVersion())
		if err != nil {
			return err
		}
	}
	if m.GetFeatures() != nil {
		err := writer.WriteCollectionOfStringValues("features", m.GetFeatures())
		if err != nil {
			return err
		}
	}
	if m.GetGenericSubstitutions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGenericSubstitutions()))
		for i, v := range m.GetGenericSubstitutions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("generic_substitutions", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("outer_object", m.GetOuterObject())
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
func (m *BlueprintInfo) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBlueprintName sets the blueprint_name property value. The blueprint_name property
func (m *BlueprintInfo) SetBlueprintName(value *string) {
	m.blueprint_name = value
}

// SetBlueprintVersion sets the blueprint_version property value. The blueprint_version property
func (m *BlueprintInfo) SetBlueprintVersion(value *string) {
	m.blueprint_version = value
}

// SetFeatures sets the features property value. The features property
func (m *BlueprintInfo) SetFeatures(value []string) {
	m.features = value
}

// SetGenericSubstitutions sets the generic_substitutions property value. The generic_substitutions property
func (m *BlueprintInfo) SetGenericSubstitutions(value []GenericSubstitutionable) {
	m.generic_substitutions = value
}

// SetOuterObject sets the outer_object property value. The Bech32m-encoded human readable version of any global address
func (m *BlueprintInfo) SetOuterObject(value *string) {
	m.outer_object = value
}

// SetPackageAddress sets the package_address property value. The Bech32m-encoded human readable version of the package address
func (m *BlueprintInfo) SetPackageAddress(value *string) {
	m.package_address = value
}

type BlueprintInfoable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBlueprintName() *string
	GetBlueprintVersion() *string
	GetFeatures() []string
	GetGenericSubstitutions() []GenericSubstitutionable
	GetOuterObject() *string
	GetPackageAddress() *string
	SetBlueprintName(value *string)
	SetBlueprintVersion(value *string)
	SetFeatures(value []string)
	SetGenericSubstitutions(value []GenericSubstitutionable)
	SetOuterObject(value *string)
	SetPackageAddress(value *string)
}
