package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StreamTransactionsRequest a request to retrieve a sublist of committed transactions from the ledger.
type StreamTransactionsRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The from_state_version property
    from_state_version *int64
    // Whether to include LedgerProofs (default false)
    include_proofs *bool
    // The maximum number of transactions that will be returned.
    limit *int32
    // The logical name of the network
    network *string
    // Requested SBOR formats to include in the response
    sbor_format_options SborFormatOptionsable
    // Requested substate formats to include in the response
    substate_format_options SubstateFormatOptionsable
    // Requested transaction formats to include in the response
    transaction_format_options TransactionFormatOptionsable
}
// NewStreamTransactionsRequest instantiates a new StreamTransactionsRequest and sets the default values.
func NewStreamTransactionsRequest()(*StreamTransactionsRequest) {
    m := &StreamTransactionsRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStreamTransactionsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamTransactionsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamTransactionsRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StreamTransactionsRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamTransactionsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["from_state_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromStateVersion(val)
        }
        return nil
    }
    res["include_proofs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeProofs(val)
        }
        return nil
    }
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val)
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
    res["sbor_format_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSborFormatOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSborFormatOptions(val.(SborFormatOptionsable))
        }
        return nil
    }
    res["substate_format_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubstateFormatOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubstateFormatOptions(val.(SubstateFormatOptionsable))
        }
        return nil
    }
    res["transaction_format_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionFormatOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionFormatOptions(val.(TransactionFormatOptionsable))
        }
        return nil
    }
    return res
}
// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *StreamTransactionsRequest) GetFromStateVersion()(*int64) {
    return m.from_state_version
}
// GetIncludeProofs gets the include_proofs property value. Whether to include LedgerProofs (default false)
// returns a *bool when successful
func (m *StreamTransactionsRequest) GetIncludeProofs()(*bool) {
    return m.include_proofs
}
// GetLimit gets the limit property value. The maximum number of transactions that will be returned.
// returns a *int32 when successful
func (m *StreamTransactionsRequest) GetLimit()(*int32) {
    return m.limit
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StreamTransactionsRequest) GetNetwork()(*string) {
    return m.network
}
// GetSborFormatOptions gets the sbor_format_options property value. Requested SBOR formats to include in the response
// returns a SborFormatOptionsable when successful
func (m *StreamTransactionsRequest) GetSborFormatOptions()(SborFormatOptionsable) {
    return m.sbor_format_options
}
// GetSubstateFormatOptions gets the substate_format_options property value. Requested substate formats to include in the response
// returns a SubstateFormatOptionsable when successful
func (m *StreamTransactionsRequest) GetSubstateFormatOptions()(SubstateFormatOptionsable) {
    return m.substate_format_options
}
// GetTransactionFormatOptions gets the transaction_format_options property value. Requested transaction formats to include in the response
// returns a TransactionFormatOptionsable when successful
func (m *StreamTransactionsRequest) GetTransactionFormatOptions()(TransactionFormatOptionsable) {
    return m.transaction_format_options
}
// Serialize serializes information the current object
func (m *StreamTransactionsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("from_state_version", m.GetFromStateVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("include_proofs", m.GetIncludeProofs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("limit", m.GetLimit())
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
        err := writer.WriteObjectValue("sbor_format_options", m.GetSborFormatOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("substate_format_options", m.GetSubstateFormatOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transaction_format_options", m.GetTransactionFormatOptions())
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
func (m *StreamTransactionsRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *StreamTransactionsRequest) SetFromStateVersion(value *int64)() {
    m.from_state_version = value
}
// SetIncludeProofs sets the include_proofs property value. Whether to include LedgerProofs (default false)
func (m *StreamTransactionsRequest) SetIncludeProofs(value *bool)() {
    m.include_proofs = value
}
// SetLimit sets the limit property value. The maximum number of transactions that will be returned.
func (m *StreamTransactionsRequest) SetLimit(value *int32)() {
    m.limit = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *StreamTransactionsRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetSborFormatOptions sets the sbor_format_options property value. Requested SBOR formats to include in the response
func (m *StreamTransactionsRequest) SetSborFormatOptions(value SborFormatOptionsable)() {
    m.sbor_format_options = value
}
// SetSubstateFormatOptions sets the substate_format_options property value. Requested substate formats to include in the response
func (m *StreamTransactionsRequest) SetSubstateFormatOptions(value SubstateFormatOptionsable)() {
    m.substate_format_options = value
}
// SetTransactionFormatOptions sets the transaction_format_options property value. Requested transaction formats to include in the response
func (m *StreamTransactionsRequest) SetTransactionFormatOptions(value TransactionFormatOptionsable)() {
    m.transaction_format_options = value
}
type StreamTransactionsRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFromStateVersion()(*int64)
    GetIncludeProofs()(*bool)
    GetLimit()(*int32)
    GetNetwork()(*string)
    GetSborFormatOptions()(SborFormatOptionsable)
    GetSubstateFormatOptions()(SubstateFormatOptionsable)
    GetTransactionFormatOptions()(TransactionFormatOptionsable)
    SetFromStateVersion(value *int64)()
    SetIncludeProofs(value *bool)()
    SetLimit(value *int32)()
    SetNetwork(value *string)()
    SetSborFormatOptions(value SborFormatOptionsable)()
    SetSubstateFormatOptions(value SubstateFormatOptionsable)()
    SetTransactionFormatOptions(value TransactionFormatOptionsable)()
}
