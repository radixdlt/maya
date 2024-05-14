package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MempoolTransactionResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An integer giving the total count of payload hashes checked in the returned response.
	count *int32
	// An array containing pairs of payload hash (query) and payload hex or error (response).Note that this response is bounded - this means it is not guaranteed all queries will be processed.Please query missing payload hashes again.
	payloads []MempoolTransactionResponse_payloadsable
}

// NewMempoolTransactionResponse instantiates a new MempoolTransactionResponse and sets the default values.
func NewMempoolTransactionResponse() *MempoolTransactionResponse {
	m := &MempoolTransactionResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateMempoolTransactionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMempoolTransactionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMempoolTransactionResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MempoolTransactionResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. An integer giving the total count of payload hashes checked in the returned response.
// returns a *int32 when successful
func (m *MempoolTransactionResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MempoolTransactionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["payloads"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateMempoolTransactionResponse_payloadsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]MempoolTransactionResponse_payloadsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(MempoolTransactionResponse_payloadsable)
				}
			}
			m.SetPayloads(res)
		}
		return nil
	}
	return res
}

// GetPayloads gets the payloads property value. An array containing pairs of payload hash (query) and payload hex or error (response).Note that this response is bounded - this means it is not guaranteed all queries will be processed.Please query missing payload hashes again.
// returns a []MempoolTransactionResponse_payloadsable when successful
func (m *MempoolTransactionResponse) GetPayloads() []MempoolTransactionResponse_payloadsable {
	return m.payloads
}

// Serialize serializes information the current object
func (m *MempoolTransactionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	if m.GetPayloads() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPayloads()))
		for i, v := range m.GetPayloads() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("payloads", cast)
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
func (m *MempoolTransactionResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. An integer giving the total count of payload hashes checked in the returned response.
func (m *MempoolTransactionResponse) SetCount(value *int32) {
	m.count = value
}

// SetPayloads sets the payloads property value. An array containing pairs of payload hash (query) and payload hex or error (response).Note that this response is bounded - this means it is not guaranteed all queries will be processed.Please query missing payload hashes again.
func (m *MempoolTransactionResponse) SetPayloads(value []MempoolTransactionResponse_payloadsable) {
	m.payloads = value
}

type MempoolTransactionResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetPayloads() []MempoolTransactionResponse_payloadsable
	SetCount(value *int32)
	SetPayloads(value []MempoolTransactionResponse_payloadsable)
}
