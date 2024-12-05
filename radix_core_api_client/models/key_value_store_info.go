package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type KeyValueStoreInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the entries of the key-value partition are allowed to own child nodes.
    allow_ownership *bool
    // The key_generic_substitution property
    key_generic_substitution GenericSubstitutionable
    // The value_generic_substitution property
    value_generic_substitution GenericSubstitutionable
}
// NewKeyValueStoreInfo instantiates a new KeyValueStoreInfo and sets the default values.
func NewKeyValueStoreInfo()(*KeyValueStoreInfo) {
    m := &KeyValueStoreInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKeyValueStoreInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateKeyValueStoreInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyValueStoreInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *KeyValueStoreInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowOwnership gets the allow_ownership property value. Whether the entries of the key-value partition are allowed to own child nodes.
// returns a *bool when successful
func (m *KeyValueStoreInfo) GetAllowOwnership()(*bool) {
    return m.allow_ownership
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *KeyValueStoreInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_ownership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowOwnership(val)
        }
        return nil
    }
    res["key_generic_substitution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGenericSubstitutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyGenericSubstitution(val.(GenericSubstitutionable))
        }
        return nil
    }
    res["value_generic_substitution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGenericSubstitutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueGenericSubstitution(val.(GenericSubstitutionable))
        }
        return nil
    }
    return res
}
// GetKeyGenericSubstitution gets the key_generic_substitution property value. The key_generic_substitution property
// returns a GenericSubstitutionable when successful
func (m *KeyValueStoreInfo) GetKeyGenericSubstitution()(GenericSubstitutionable) {
    return m.key_generic_substitution
}
// GetValueGenericSubstitution gets the value_generic_substitution property value. The value_generic_substitution property
// returns a GenericSubstitutionable when successful
func (m *KeyValueStoreInfo) GetValueGenericSubstitution()(GenericSubstitutionable) {
    return m.value_generic_substitution
}
// Serialize serializes information the current object
func (m *KeyValueStoreInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_ownership", m.GetAllowOwnership())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("key_generic_substitution", m.GetKeyGenericSubstitution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value_generic_substitution", m.GetValueGenericSubstitution())
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
func (m *KeyValueStoreInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowOwnership sets the allow_ownership property value. Whether the entries of the key-value partition are allowed to own child nodes.
func (m *KeyValueStoreInfo) SetAllowOwnership(value *bool)() {
    m.allow_ownership = value
}
// SetKeyGenericSubstitution sets the key_generic_substitution property value. The key_generic_substitution property
func (m *KeyValueStoreInfo) SetKeyGenericSubstitution(value GenericSubstitutionable)() {
    m.key_generic_substitution = value
}
// SetValueGenericSubstitution sets the value_generic_substitution property value. The value_generic_substitution property
func (m *KeyValueStoreInfo) SetValueGenericSubstitution(value GenericSubstitutionable)() {
    m.value_generic_substitution = value
}
type KeyValueStoreInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowOwnership()(*bool)
    GetKeyGenericSubstitution()(GenericSubstitutionable)
    GetValueGenericSubstitution()(GenericSubstitutionable)
    SetAllowOwnership(value *bool)()
    SetKeyGenericSubstitution(value GenericSubstitutionable)()
    SetValueGenericSubstitution(value GenericSubstitutionable)()
}
