package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LedgerStateSelector an optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
type LedgerStateSelector struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *LedgerStateSelectorType
}
// NewLedgerStateSelector instantiates a new LedgerStateSelector and sets the default values.
func NewLedgerStateSelector()(*LedgerStateSelector) {
    m := &LedgerStateSelector{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLedgerStateSelectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerStateSelectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "ByStateVersion":
                        return NewVersionLedgerStateSelector(), nil
                }
            }
        }
    }
    return NewLedgerStateSelector(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerStateSelector) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerStateSelector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLedgerStateSelectorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*LedgerStateSelectorType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *LedgerStateSelectorType when successful
func (m *LedgerStateSelector) GetTypeEscaped()(*LedgerStateSelectorType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *LedgerStateSelector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *LedgerStateSelector) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *LedgerStateSelector) SetTypeEscaped(value *LedgerStateSelectorType)() {
    m.typeEscaped = value
}
type LedgerStateSelectorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*LedgerStateSelectorType)
    SetTypeEscaped(value *LedgerStateSelectorType)()
}
