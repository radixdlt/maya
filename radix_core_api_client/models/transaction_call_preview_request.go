package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionCallPreviewRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Argument list
    arguments []string
    // An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
    at_ledger_state LedgerStateSelectorable
    // The logical name of the network
    network *string
    // The target property
    target TargetIdentifierable
}
// NewTransactionCallPreviewRequest instantiates a new TransactionCallPreviewRequest and sets the default values.
func NewTransactionCallPreviewRequest()(*TransactionCallPreviewRequest) {
    m := &TransactionCallPreviewRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionCallPreviewRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionCallPreviewRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionCallPreviewRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionCallPreviewRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArguments gets the arguments property value. Argument list
// returns a []string when successful
func (m *TransactionCallPreviewRequest) GetArguments()([]string) {
    return m.arguments
}
// GetAtLedgerState gets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
// returns a LedgerStateSelectorable when successful
func (m *TransactionCallPreviewRequest) GetAtLedgerState()(LedgerStateSelectorable) {
    return m.at_ledger_state
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionCallPreviewRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["arguments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetArguments(res)
        }
        return nil
    }
    res["at_ledger_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLedgerStateSelectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtLedgerState(val.(LedgerStateSelectorable))
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTargetIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(TargetIdentifierable))
        }
        return nil
    }
    return res
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *TransactionCallPreviewRequest) GetNetwork()(*string) {
    return m.network
}
// GetTarget gets the target property value. The target property
// returns a TargetIdentifierable when successful
func (m *TransactionCallPreviewRequest) GetTarget()(TargetIdentifierable) {
    return m.target
}
// Serialize serializes information the current object
func (m *TransactionCallPreviewRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetArguments() != nil {
        err := writer.WriteCollectionOfStringValues("arguments", m.GetArguments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("at_ledger_state", m.GetAtLedgerState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("target", m.GetTarget())
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
func (m *TransactionCallPreviewRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArguments sets the arguments property value. Argument list
func (m *TransactionCallPreviewRequest) SetArguments(value []string)() {
    m.arguments = value
}
// SetAtLedgerState sets the at_ledger_state property value. An optional specification of a historical ledger state at which to execute the request.The "historical state" feature (see the `db.historical_substate_values.enable` flag) must beenabled on the Node, and the requested point in history must be recent enough (in accordancewith the Node's configured `state_hash_tree.state_version_history_length`).
func (m *TransactionCallPreviewRequest) SetAtLedgerState(value LedgerStateSelectorable)() {
    m.at_ledger_state = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *TransactionCallPreviewRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetTarget sets the target property value. The target property
func (m *TransactionCallPreviewRequest) SetTarget(value TargetIdentifierable)() {
    m.target = value
}
type TransactionCallPreviewRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]string)
    GetAtLedgerState()(LedgerStateSelectorable)
    GetNetwork()(*string)
    GetTarget()(TargetIdentifierable)
    SetArguments(value []string)()
    SetAtLedgerState(value LedgerStateSelectorable)()
    SetNetwork(value *string)()
    SetTarget(value TargetIdentifierable)()
}
