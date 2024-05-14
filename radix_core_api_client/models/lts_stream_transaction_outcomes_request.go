package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LtsStreamTransactionOutcomesRequest a request to retrieve a sublist of committed transactions from the ledger.
type LtsStreamTransactionOutcomesRequest struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The from_state_version property
	from_state_version *int64
	// The maximum number of transactions that will be returned.
	limit *int32
	// The logical name of the network
	network *string
}

// NewLtsStreamTransactionOutcomesRequest instantiates a new LtsStreamTransactionOutcomesRequest and sets the default values.
func NewLtsStreamTransactionOutcomesRequest() *LtsStreamTransactionOutcomesRequest {
	m := &LtsStreamTransactionOutcomesRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStreamTransactionOutcomesRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStreamTransactionOutcomesRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStreamTransactionOutcomesRequest(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStreamTransactionOutcomesRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStreamTransactionOutcomesRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["from_state_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFromStateVersion(val)
		}
		return nil
	}
	res["limit"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLimit(val)
		}
		return nil
	}
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
		}
		return nil
	}
	return res
}

// GetFromStateVersion gets the from_state_version property value. The from_state_version property
// returns a *int64 when successful
func (m *LtsStreamTransactionOutcomesRequest) GetFromStateVersion() *int64 {
	return m.from_state_version
}

// GetLimit gets the limit property value. The maximum number of transactions that will be returned.
// returns a *int32 when successful
func (m *LtsStreamTransactionOutcomesRequest) GetLimit() *int32 {
	return m.limit
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *LtsStreamTransactionOutcomesRequest) GetNetwork() *string {
	return m.network
}

// Serialize serializes information the current object
func (m *LtsStreamTransactionOutcomesRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("from_state_version", m.GetFromStateVersion())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStreamTransactionOutcomesRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *LtsStreamTransactionOutcomesRequest) SetFromStateVersion(value *int64) {
	m.from_state_version = value
}

// SetLimit sets the limit property value. The maximum number of transactions that will be returned.
func (m *LtsStreamTransactionOutcomesRequest) SetLimit(value *int32) {
	m.limit = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *LtsStreamTransactionOutcomesRequest) SetNetwork(value *string) {
	m.network = value
}

type LtsStreamTransactionOutcomesRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFromStateVersion() *int64
	GetLimit() *int32
	GetNetwork() *string
	SetFromStateVersion(value *int64)
	SetLimit(value *int32)
	SetNetwork(value *string)
}
