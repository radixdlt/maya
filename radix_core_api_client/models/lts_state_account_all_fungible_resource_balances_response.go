package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStateAccountAllFungibleResourceBalancesResponse struct {
	// The Bech32m-encoded human readable version of the account's address
	account_address *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A list containing all resource balances for the requested account.
	fungible_resource_balances []LtsFungibleResourceBalanceable
	// The ledger_header_summary property
	ledger_header_summary LedgerHeaderSummaryable
	// The state_version property
	state_version *int64
}

// NewLtsStateAccountAllFungibleResourceBalancesResponse instantiates a new LtsStateAccountAllFungibleResourceBalancesResponse and sets the default values.
func NewLtsStateAccountAllFungibleResourceBalancesResponse() *LtsStateAccountAllFungibleResourceBalancesResponse {
	m := &LtsStateAccountAllFungibleResourceBalancesResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStateAccountAllFungibleResourceBalancesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStateAccountAllFungibleResourceBalancesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStateAccountAllFungibleResourceBalancesResponse(), nil
}

// GetAccountAddress gets the account_address property value. The Bech32m-encoded human readable version of the account's address
// returns a *string when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetAccountAddress() *string {
	return m.account_address
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["fungible_resource_balances"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsFungibleResourceBalanceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsFungibleResourceBalanceable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsFungibleResourceBalanceable)
				}
			}
			m.SetFungibleResourceBalances(res)
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

// GetFungibleResourceBalances gets the fungible_resource_balances property value. A list containing all resource balances for the requested account.
// returns a []LtsFungibleResourceBalanceable when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetFungibleResourceBalances() []LtsFungibleResourceBalanceable {
	return m.fungible_resource_balances
}

// GetLedgerHeaderSummary gets the ledger_header_summary property value. The ledger_header_summary property
// returns a LedgerHeaderSummaryable when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetLedgerHeaderSummary() LedgerHeaderSummaryable {
	return m.ledger_header_summary
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) GetStateVersion() *int64 {
	return m.state_version
}

// Serialize serializes information the current object
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("account_address", m.GetAccountAddress())
		if err != nil {
			return err
		}
	}
	if m.GetFungibleResourceBalances() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFungibleResourceBalances()))
		for i, v := range m.GetFungibleResourceBalances() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("fungible_resource_balances", cast)
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
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) SetAccountAddress(value *string) {
	m.account_address = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFungibleResourceBalances sets the fungible_resource_balances property value. A list containing all resource balances for the requested account.
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) SetFungibleResourceBalances(value []LtsFungibleResourceBalanceable) {
	m.fungible_resource_balances = value
}

// SetLedgerHeaderSummary sets the ledger_header_summary property value. The ledger_header_summary property
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) SetLedgerHeaderSummary(value LedgerHeaderSummaryable) {
	m.ledger_header_summary = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LtsStateAccountAllFungibleResourceBalancesResponse) SetStateVersion(value *int64) {
	m.state_version = value
}

type LtsStateAccountAllFungibleResourceBalancesResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountAddress() *string
	GetFungibleResourceBalances() []LtsFungibleResourceBalanceable
	GetLedgerHeaderSummary() LedgerHeaderSummaryable
	GetStateVersion() *int64
	SetAccountAddress(value *string)
	SetFungibleResourceBalances(value []LtsFungibleResourceBalanceable)
	SetLedgerHeaderSummary(value LedgerHeaderSummaryable)
	SetStateVersion(value *int64)
}
