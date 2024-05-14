package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStateAccountFungibleResourceBalanceResponse struct {
	// The Bech32m-encoded human readable version of the account's address
	account_address *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The fungible_resource_balance property
	fungible_resource_balance LtsFungibleResourceBalanceable
	// The ledger_header_summary property
	ledger_header_summary LedgerHeaderSummaryable
	// The state_version property
	state_version *int64
}

// NewLtsStateAccountFungibleResourceBalanceResponse instantiates a new LtsStateAccountFungibleResourceBalanceResponse and sets the default values.
func NewLtsStateAccountFungibleResourceBalanceResponse() *LtsStateAccountFungibleResourceBalanceResponse {
	m := &LtsStateAccountFungibleResourceBalanceResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStateAccountFungibleResourceBalanceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStateAccountFungibleResourceBalanceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStateAccountFungibleResourceBalanceResponse(), nil
}

// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address
// returns a *string when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetAccountAddress() *string {
	return m.account_address
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["fungible_resource_balance"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLtsFungibleResourceBalanceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFungibleResourceBalance(val.(LtsFungibleResourceBalanceable))
		}
		return nil
	}
	res["ledger_header_summary"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerHeaderSummaryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLedgerHeaderSummary(val.(LedgerHeaderSummaryable))
		}
		return nil
	}
	res["state_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStateVersion(val)
		}
		return nil
	}
	return res
}

// GetFungibleResourceBalance gets the fungible_resource_balance property value. The fungible_resource_balance property
// returns a LtsFungibleResourceBalanceable when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetFungibleResourceBalance() LtsFungibleResourceBalanceable {
	return m.fungible_resource_balance
}

// GetLedgerHeaderSummary gets the ledger_header_summary property value. The ledger_header_summary property
// returns a LedgerHeaderSummaryable when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetLedgerHeaderSummary() LedgerHeaderSummaryable {
	return m.ledger_header_summary
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LtsStateAccountFungibleResourceBalanceResponse) GetStateVersion() *int64 {
	return m.state_version
}

// Serialize serializes information the current object
func (m *LtsStateAccountFungibleResourceBalanceResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("account_address", m.GetAccountAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("fungible_resource_balance", m.GetFungibleResourceBalance())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("ledger_header_summary", m.GetLedgerHeaderSummary())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("state_version", m.GetStateVersion())
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
func (m *LtsStateAccountFungibleResourceBalanceResponse) SetAccountAddress(value *string) {
	m.account_address = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStateAccountFungibleResourceBalanceResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFungibleResourceBalance sets the fungible_resource_balance property value. The fungible_resource_balance property
func (m *LtsStateAccountFungibleResourceBalanceResponse) SetFungibleResourceBalance(value LtsFungibleResourceBalanceable) {
	m.fungible_resource_balance = value
}

// SetLedgerHeaderSummary sets the ledger_header_summary property value. The ledger_header_summary property
func (m *LtsStateAccountFungibleResourceBalanceResponse) SetLedgerHeaderSummary(value LedgerHeaderSummaryable) {
	m.ledger_header_summary = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LtsStateAccountFungibleResourceBalanceResponse) SetStateVersion(value *int64) {
	m.state_version = value
}

type LtsStateAccountFungibleResourceBalanceResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountAddress() *string
	GetFungibleResourceBalance() LtsFungibleResourceBalanceable
	GetLedgerHeaderSummary() LedgerHeaderSummaryable
	GetStateVersion() *int64
	SetAccountAddress(value *string)
	SetFungibleResourceBalance(value LtsFungibleResourceBalanceable)
	SetLedgerHeaderSummary(value LedgerHeaderSummaryable)
	SetStateVersion(value *int64)
}
