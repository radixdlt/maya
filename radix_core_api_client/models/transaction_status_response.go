package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionStatusResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The status of the transaction intent, as determined by the node.FateUncertain or FateUncertainButLikelyRejection mean that it's still possible that a payload containing the transaction
    intent_status *TransactionIntentStatus
    // An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the intent status is InMempool or Unknown and we know about a payload.
    invalid_from_epoch *int64
    // The known_payloads property
    known_payloads []TransactionPayloadDetailsable
    // An explanation as to why the intent status is resolved as it is.
    status_description *string
}
// NewTransactionStatusResponse instantiates a new TransactionStatusResponse and sets the default values.
func NewTransactionStatusResponse()(*TransactionStatusResponse) {
    m := &TransactionStatusResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionStatusResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionStatusResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionStatusResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionStatusResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionStatusResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["intent_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransactionIntentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentStatus(val.(*TransactionIntentStatus))
        }
        return nil
    }
    res["invalid_from_epoch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvalidFromEpoch(val)
        }
        return nil
    }
    res["known_payloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTransactionPayloadDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TransactionPayloadDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TransactionPayloadDetailsable)
                }
            }
            m.SetKnownPayloads(res)
        }
        return nil
    }
    res["status_description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDescription(val)
        }
        return nil
    }
    return res
}
// GetIntentStatus gets the intent_status property value. The status of the transaction intent, as determined by the node.FateUncertain or FateUncertainButLikelyRejection mean that it's still possible that a payload containing the transaction
// returns a *TransactionIntentStatus when successful
func (m *TransactionStatusResponse) GetIntentStatus()(*TransactionIntentStatus) {
    return m.intent_status
}
// GetInvalidFromEpoch gets the invalid_from_epoch property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the intent status is InMempool or Unknown and we know about a payload.
// returns a *int64 when successful
func (m *TransactionStatusResponse) GetInvalidFromEpoch()(*int64) {
    return m.invalid_from_epoch
}
// GetKnownPayloads gets the known_payloads property value. The known_payloads property
// returns a []TransactionPayloadDetailsable when successful
func (m *TransactionStatusResponse) GetKnownPayloads()([]TransactionPayloadDetailsable) {
    return m.known_payloads
}
// GetStatusDescription gets the status_description property value. An explanation as to why the intent status is resolved as it is.
// returns a *string when successful
func (m *TransactionStatusResponse) GetStatusDescription()(*string) {
    return m.status_description
}
// Serialize serializes information the current object
func (m *TransactionStatusResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIntentStatus() != nil {
        cast := (*m.GetIntentStatus()).String()
        err := writer.WriteStringValue("intent_status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("invalid_from_epoch", m.GetInvalidFromEpoch())
        if err != nil {
            return err
        }
    }
    if m.GetKnownPayloads() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKnownPayloads()))
        for i, v := range m.GetKnownPayloads() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("known_payloads", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status_description", m.GetStatusDescription())
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
func (m *TransactionStatusResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntentStatus sets the intent_status property value. The status of the transaction intent, as determined by the node.FateUncertain or FateUncertainButLikelyRejection mean that it's still possible that a payload containing the transaction
func (m *TransactionStatusResponse) SetIntentStatus(value *TransactionIntentStatus)() {
    m.intent_status = value
}
// SetInvalidFromEpoch sets the invalid_from_epoch property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the intent status is InMempool or Unknown and we know about a payload.
func (m *TransactionStatusResponse) SetInvalidFromEpoch(value *int64)() {
    m.invalid_from_epoch = value
}
// SetKnownPayloads sets the known_payloads property value. The known_payloads property
func (m *TransactionStatusResponse) SetKnownPayloads(value []TransactionPayloadDetailsable)() {
    m.known_payloads = value
}
// SetStatusDescription sets the status_description property value. An explanation as to why the intent status is resolved as it is.
func (m *TransactionStatusResponse) SetStatusDescription(value *string)() {
    m.status_description = value
}
type TransactionStatusResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntentStatus()(*TransactionIntentStatus)
    GetInvalidFromEpoch()(*int64)
    GetKnownPayloads()([]TransactionPayloadDetailsable)
    GetStatusDescription()(*string)
    SetIntentStatus(value *TransactionIntentStatus)()
    SetInvalidFromEpoch(value *int64)()
    SetKnownPayloads(value []TransactionPayloadDetailsable)()
    SetStatusDescription(value *string)()
}
