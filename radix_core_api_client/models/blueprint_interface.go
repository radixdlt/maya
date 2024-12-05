package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlueprintInterface struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A map from the event name to the event payload type reference.
    events BlueprintInterface_eventsable
    // The features property
    features []string
    // A map from the function name to the FunctionSchema
    functions BlueprintInterface_functionsable
    // Generic (SBOR) type parameters which need to be filled by a concrete instanceof this blueprint.
    generic_type_parameters []GenericTypeParameterable
    // If true, an instantiation of this blueprint cannot be persisted. EG buckets and proofs are transient.
    is_transient *bool
    // The outer_blueprint property
    outer_blueprint *string
    // The state property
    state IndexedStateSchemaable
    // A map from the registered type name to the concrete type,resolved against a schema from the package's schema partition.
    types BlueprintInterface_typesable
}
// NewBlueprintInterface instantiates a new BlueprintInterface and sets the default values.
func NewBlueprintInterface()(*BlueprintInterface) {
    m := &BlueprintInterface{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlueprintInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlueprintInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlueprintInterface(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BlueprintInterface) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEvents gets the events property value. A map from the event name to the event payload type reference.
// returns a BlueprintInterface_eventsable when successful
func (m *BlueprintInterface) GetEvents()(BlueprintInterface_eventsable) {
    return m.events
}
// GetFeatures gets the features property value. The features property
// returns a []string when successful
func (m *BlueprintInterface) GetFeatures()([]string) {
    return m.features
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlueprintInterface) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlueprintInterface_eventsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvents(val.(BlueprintInterface_eventsable))
        }
        return nil
    }
    res["features"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["functions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlueprintInterface_functionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctions(val.(BlueprintInterface_functionsable))
        }
        return nil
    }
    res["generic_type_parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGenericTypeParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GenericTypeParameterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GenericTypeParameterable)
                }
            }
            m.SetGenericTypeParameters(res)
        }
        return nil
    }
    res["is_transient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTransient(val)
        }
        return nil
    }
    res["outer_blueprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuterBlueprint(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndexedStateSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(IndexedStateSchemaable))
        }
        return nil
    }
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlueprintInterface_typesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypes(val.(BlueprintInterface_typesable))
        }
        return nil
    }
    return res
}
// GetFunctions gets the functions property value. A map from the function name to the FunctionSchema
// returns a BlueprintInterface_functionsable when successful
func (m *BlueprintInterface) GetFunctions()(BlueprintInterface_functionsable) {
    return m.functions
}
// GetGenericTypeParameters gets the generic_type_parameters property value. Generic (SBOR) type parameters which need to be filled by a concrete instanceof this blueprint.
// returns a []GenericTypeParameterable when successful
func (m *BlueprintInterface) GetGenericTypeParameters()([]GenericTypeParameterable) {
    return m.generic_type_parameters
}
// GetIsTransient gets the is_transient property value. If true, an instantiation of this blueprint cannot be persisted. EG buckets and proofs are transient.
// returns a *bool when successful
func (m *BlueprintInterface) GetIsTransient()(*bool) {
    return m.is_transient
}
// GetOuterBlueprint gets the outer_blueprint property value. The outer_blueprint property
// returns a *string when successful
func (m *BlueprintInterface) GetOuterBlueprint()(*string) {
    return m.outer_blueprint
}
// GetState gets the state property value. The state property
// returns a IndexedStateSchemaable when successful
func (m *BlueprintInterface) GetState()(IndexedStateSchemaable) {
    return m.state
}
// GetTypes gets the types property value. A map from the registered type name to the concrete type,resolved against a schema from the package's schema partition.
// returns a BlueprintInterface_typesable when successful
func (m *BlueprintInterface) GetTypes()(BlueprintInterface_typesable) {
    return m.types
}
// Serialize serializes information the current object
func (m *BlueprintInterface) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("events", m.GetEvents())
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
    {
        err := writer.WriteObjectValue("functions", m.GetFunctions())
        if err != nil {
            return err
        }
    }
    if m.GetGenericTypeParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGenericTypeParameters()))
        for i, v := range m.GetGenericTypeParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("generic_type_parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_transient", m.GetIsTransient())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outer_blueprint", m.GetOuterBlueprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("types", m.GetTypes())
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
func (m *BlueprintInterface) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEvents sets the events property value. A map from the event name to the event payload type reference.
func (m *BlueprintInterface) SetEvents(value BlueprintInterface_eventsable)() {
    m.events = value
}
// SetFeatures sets the features property value. The features property
func (m *BlueprintInterface) SetFeatures(value []string)() {
    m.features = value
}
// SetFunctions sets the functions property value. A map from the function name to the FunctionSchema
func (m *BlueprintInterface) SetFunctions(value BlueprintInterface_functionsable)() {
    m.functions = value
}
// SetGenericTypeParameters sets the generic_type_parameters property value. Generic (SBOR) type parameters which need to be filled by a concrete instanceof this blueprint.
func (m *BlueprintInterface) SetGenericTypeParameters(value []GenericTypeParameterable)() {
    m.generic_type_parameters = value
}
// SetIsTransient sets the is_transient property value. If true, an instantiation of this blueprint cannot be persisted. EG buckets and proofs are transient.
func (m *BlueprintInterface) SetIsTransient(value *bool)() {
    m.is_transient = value
}
// SetOuterBlueprint sets the outer_blueprint property value. The outer_blueprint property
func (m *BlueprintInterface) SetOuterBlueprint(value *string)() {
    m.outer_blueprint = value
}
// SetState sets the state property value. The state property
func (m *BlueprintInterface) SetState(value IndexedStateSchemaable)() {
    m.state = value
}
// SetTypes sets the types property value. A map from the registered type name to the concrete type,resolved against a schema from the package's schema partition.
func (m *BlueprintInterface) SetTypes(value BlueprintInterface_typesable)() {
    m.types = value
}
type BlueprintInterfaceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEvents()(BlueprintInterface_eventsable)
    GetFeatures()([]string)
    GetFunctions()(BlueprintInterface_functionsable)
    GetGenericTypeParameters()([]GenericTypeParameterable)
    GetIsTransient()(*bool)
    GetOuterBlueprint()(*string)
    GetState()(IndexedStateSchemaable)
    GetTypes()(BlueprintInterface_typesable)
    SetEvents(value BlueprintInterface_eventsable)()
    SetFeatures(value []string)()
    SetFunctions(value BlueprintInterface_functionsable)()
    SetGenericTypeParameters(value []GenericTypeParameterable)()
    SetIsTransient(value *bool)()
    SetOuterBlueprint(value *string)()
    SetState(value IndexedStateSchemaable)()
    SetTypes(value BlueprintInterface_typesable)()
}
