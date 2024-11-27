package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A map from a function name to AccessRule.Only exists if `function_auth_type` is set to `FunctionAccessRules`.
    function_access_rules AuthConfig_function_access_rulesable
    // The function_auth_type property
    function_auth_type *FunctionAuthType
    // The method_auth_type property
    method_auth_type *MethodAuthType
    // The method_roles property
    method_roles StaticRoleDefinitionAuthTemplateable
}
// NewAuthConfig instantiates a new AuthConfig and sets the default values.
func NewAuthConfig()(*AuthConfig) {
    m := &AuthConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["function_access_rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthConfig_function_access_rulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctionAccessRules(val.(AuthConfig_function_access_rulesable))
        }
        return nil
    }
    res["function_auth_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFunctionAuthType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctionAuthType(val.(*FunctionAuthType))
        }
        return nil
    }
    res["method_auth_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMethodAuthType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodAuthType(val.(*MethodAuthType))
        }
        return nil
    }
    res["method_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStaticRoleDefinitionAuthTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodRoles(val.(StaticRoleDefinitionAuthTemplateable))
        }
        return nil
    }
    return res
}
// GetFunctionAccessRules gets the function_access_rules property value. A map from a function name to AccessRule.Only exists if `function_auth_type` is set to `FunctionAccessRules`.
// returns a AuthConfig_function_access_rulesable when successful
func (m *AuthConfig) GetFunctionAccessRules()(AuthConfig_function_access_rulesable) {
    return m.function_access_rules
}
// GetFunctionAuthType gets the function_auth_type property value. The function_auth_type property
// returns a *FunctionAuthType when successful
func (m *AuthConfig) GetFunctionAuthType()(*FunctionAuthType) {
    return m.function_auth_type
}
// GetMethodAuthType gets the method_auth_type property value. The method_auth_type property
// returns a *MethodAuthType when successful
func (m *AuthConfig) GetMethodAuthType()(*MethodAuthType) {
    return m.method_auth_type
}
// GetMethodRoles gets the method_roles property value. The method_roles property
// returns a StaticRoleDefinitionAuthTemplateable when successful
func (m *AuthConfig) GetMethodRoles()(StaticRoleDefinitionAuthTemplateable) {
    return m.method_roles
}
// Serialize serializes information the current object
func (m *AuthConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("function_access_rules", m.GetFunctionAccessRules())
        if err != nil {
            return err
        }
    }
    if m.GetFunctionAuthType() != nil {
        cast := (*m.GetFunctionAuthType()).String()
        err := writer.WriteStringValue("function_auth_type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMethodAuthType() != nil {
        cast := (*m.GetMethodAuthType()).String()
        err := writer.WriteStringValue("method_auth_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("method_roles", m.GetMethodRoles())
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
func (m *AuthConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFunctionAccessRules sets the function_access_rules property value. A map from a function name to AccessRule.Only exists if `function_auth_type` is set to `FunctionAccessRules`.
func (m *AuthConfig) SetFunctionAccessRules(value AuthConfig_function_access_rulesable)() {
    m.function_access_rules = value
}
// SetFunctionAuthType sets the function_auth_type property value. The function_auth_type property
func (m *AuthConfig) SetFunctionAuthType(value *FunctionAuthType)() {
    m.function_auth_type = value
}
// SetMethodAuthType sets the method_auth_type property value. The method_auth_type property
func (m *AuthConfig) SetMethodAuthType(value *MethodAuthType)() {
    m.method_auth_type = value
}
// SetMethodRoles sets the method_roles property value. The method_roles property
func (m *AuthConfig) SetMethodRoles(value StaticRoleDefinitionAuthTemplateable)() {
    m.method_roles = value
}
type AuthConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFunctionAccessRules()(AuthConfig_function_access_rulesable)
    GetFunctionAuthType()(*FunctionAuthType)
    GetMethodAuthType()(*MethodAuthType)
    GetMethodRoles()(StaticRoleDefinitionAuthTemplateable)
    SetFunctionAccessRules(value AuthConfig_function_access_rulesable)()
    SetFunctionAuthType(value *FunctionAuthType)()
    SetMethodAuthType(value *MethodAuthType)()
    SetMethodRoles(value StaticRoleDefinitionAuthTemplateable)()
}
