package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LtsStreamAccountTransactionOutcomesRequest a request to retrieve a sublist of committed transactions from the ledger.
type LtsStreamAccountTransactionOutcomesRequest struct {
	// The Bech32m-encoded human readable version of the account's address
	account_address *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The from_state_version property
	from_state_version *int64
	// The maximum number of transactions that will be returned.
	limit *int32
	// The logical name of the network
	network *string
}

// NewLtsStreamAccountTransactionOutcomesRequest instantiates a new LtsStreamAccountTransactionOutcomesRequest and sets the default values.
func NewLtsStreamAccountTransactionOutcomesRequest() *LtsStreamAccountTransactionOutcomesRequest {
	m := &LtsStreamAccountTransactionOutcomesRequest{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStreamAccountTransactionOutcomesRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStreamAccountTransactionOutcomesRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStreamAccountTransactionOutcomesRequest(), nil
}

// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address
// returns a *string when successful
func (m *LtsStreamAccountTransactionOutcomesRequest) GetAccountAddress() *string {
	return m.account_address
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStreamAccountTransactionOutcomesRequest) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStreamAccountTransactionOutcomesRequest) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["account_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccountAddress(val)
		}
		return nil
	}
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
func (m *LtsStreamAccountTransactionOutcomesRequest) GetFromStateVersion() *int64 {
	return m.from_state_version
}

// GetLimit gets the limit property value. The maximum number of transactions that will be returned.
// returns a *int32 when successful
func (m *LtsStreamAccountTransactionOutcomesRequest) GetLimit() *int32 {
	return m.limit
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *LtsStreamAccountTransactionOutcomesRequest) GetNetwork() *string {
	return m.network
}

// Serialize serializes information the current object
func (m *LtsStreamAccountTransactionOutcomesRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("account_address", m.GetAccountAddress())
		if err != nil {
			return err
		}
	}
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

// SetAccountAddress sets the account_address property value. The Bech32m-encoded human readable version of the account's address
func (m *LtsStreamAccountTransactionOutcomesRequest) SetAccountAddress(value *string) {
	m.account_address = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStreamAccountTransactionOutcomesRequest) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFromStateVersion sets the from_state_version property value. The from_state_version property
func (m *LtsStreamAccountTransactionOutcomesRequest) SetFromStateVersion(value *int64) {
	m.from_state_version = value
}

// SetLimit sets the limit property value. The maximum number of transactions that will be returned.
func (m *LtsStreamAccountTransactionOutcomesRequest) SetLimit(value *int32) {
	m.limit = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *LtsStreamAccountTransactionOutcomesRequest) SetNetwork(value *string) {
	m.network = value
}

type LtsStreamAccountTransactionOutcomesRequestable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountAddress() *string
	GetFromStateVersion() *int64
	GetLimit() *int32
	GetNetwork() *string
	SetAccountAddress(value *string)
	SetFromStateVersion(value *int64)
	SetLimit(value *int32)
	SetNetwork(value *string)
}
