package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MempoolTransactionRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The logical name of the network
    network *string
    // A list of payload hashes to attempt to read. Each hash must be either hex, or in Bech32m format.
    payload_hashes []string
}
// NewMempoolTransactionRequest instantiates a new MempoolTransactionRequest and sets the default values.
func NewMempoolTransactionRequest()(*MempoolTransactionRequest) {
    m := &MempoolTransactionRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMempoolTransactionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMempoolTransactionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMempoolTransactionRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MempoolTransactionRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MempoolTransactionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["payload_hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPayloadHashes(res)
        }
        return nil
    }
    return res
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *MempoolTransactionRequest) GetNetwork()(*string) {
    return m.network
}
// GetPayloadHashes gets the payload_hashes property value. A list of payload hashes to attempt to read. Each hash must be either hex, or in Bech32m format.
// returns a []string when successful
func (m *MempoolTransactionRequest) GetPayloadHashes()([]string) {
    return m.payload_hashes
}
// Serialize serializes information the current object
func (m *MempoolTransactionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadHashes() != nil {
        err := writer.WriteCollectionOfStringValues("payload_hashes", m.GetPayloadHashes())
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
func (m *MempoolTransactionRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *MempoolTransactionRequest) SetNetwork(value *string)() {
    m.network = value
}
// SetPayloadHashes sets the payload_hashes property value. A list of payload hashes to attempt to read. Each hash must be either hex, or in Bech32m format.
func (m *MempoolTransactionRequest) SetPayloadHashes(value []string)() {
    m.payload_hashes = value
}
type MempoolTransactionRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetwork()(*string)
    GetPayloadHashes()([]string)
    SetNetwork(value *string)()
    SetPayloadHashes(value []string)()
}
