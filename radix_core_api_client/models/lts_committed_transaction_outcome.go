package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LtsCommittedTransactionOutcome for the given transaction, contains the status, total fee summary and individual entity resource balance changes.The balance changes accounts for the fee payments as well.
type LtsCommittedTransactionOutcome struct {
	// The hex-encoded transaction accumulator hash. This hash captures the order of all transactions on ledger.This hash is `ACC_{N+1} = combine(ACC_N, LEDGER_HASH_{N}))` (where `combine()` is an arbitrary deterministic function we use).
	accumulator_hash *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
	fungible_entity_balance_changes []LtsEntityFungibleBalanceChangesable
	// Non fungible changes per entity and resource
	non_fungible_entity_balance_changes []LtsEntityNonFungibleBalanceChangesable
	// An integer between `0` and `10^14`, marking the proposer timestamp in ms.
	proposer_timestamp_ms *int64
	// A list of the resultant fungible account balances for any balances which changed in this transaction.Only balances for accounts are returned, not any other kind of entity.
	resultant_account_fungible_balances []LtsResultantAccountFungibleBalancesable
	// The state_version property
	state_version *int64
	// The status of the transaction
	status *LtsCommittedTransactionStatus
	// The string-encoded decimal representing the total amount of XRD paid as fee (execution, validator tip and royalties).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	total_fee *string
	// The user_transaction_identifiers property
	user_transaction_identifiers TransactionIdentifiersable
}

// NewLtsCommittedTransactionOutcome instantiates a new LtsCommittedTransactionOutcome and sets the default values.
func NewLtsCommittedTransactionOutcome() *LtsCommittedTransactionOutcome {
	m := &LtsCommittedTransactionOutcome{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsCommittedTransactionOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsCommittedTransactionOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsCommittedTransactionOutcome(), nil
}

// GetAccumulatorHash gets the accumulator_hash property value. The hex-encoded transaction accumulator hash. This hash captures the order of all transactions on ledger.This hash is `ACC_{N+1} = combine(ACC_N, LEDGER_HASH_{N}))` (where `combine()` is an arbitrary deterministic function we use).
// returns a *string when successful
func (m *LtsCommittedTransactionOutcome) GetAccumulatorHash() *string {
	return m.accumulator_hash
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsCommittedTransactionOutcome) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsCommittedTransactionOutcome) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["accumulator_hash"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccumulatorHash(val)
		}
		return nil
	}
	res["fungible_entity_balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsEntityFungibleBalanceChangesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsEntityFungibleBalanceChangesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsEntityFungibleBalanceChangesable)
				}
			}
			m.SetFungibleEntityBalanceChanges(res)
		}
		return nil
	}
	res["non_fungible_entity_balance_changes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsEntityNonFungibleBalanceChangesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsEntityNonFungibleBalanceChangesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsEntityNonFungibleBalanceChangesable)
				}
			}
			m.SetNonFungibleEntityBalanceChanges(res)
		}
		return nil
	}
	res["proposer_timestamp_ms"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProposerTimestampMs(val)
		}
		return nil
	}
	res["resultant_account_fungible_balances"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLtsResultantAccountFungibleBalancesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LtsResultantAccountFungibleBalancesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LtsResultantAccountFungibleBalancesable)
				}
			}
			m.SetResultantAccountFungibleBalances(res)
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
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseLtsCommittedTransactionStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val.(*LtsCommittedTransactionStatus))
		}
		return nil
	}
	res["total_fee"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotalFee(val)
		}
		return nil
	}
	res["user_transaction_identifiers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTransactionIdentifiersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserTransactionIdentifiers(val.(TransactionIdentifiersable))
		}
		return nil
	}
	return res
}

// GetFungibleEntityBalanceChanges gets the fungible_entity_balance_changes property value. A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
// returns a []LtsEntityFungibleBalanceChangesable when successful
func (m *LtsCommittedTransactionOutcome) GetFungibleEntityBalanceChanges() []LtsEntityFungibleBalanceChangesable {
	return m.fungible_entity_balance_changes
}

// GetNonFungibleEntityBalanceChanges gets the non_fungible_entity_balance_changes property value. Non fungible changes per entity and resource
// returns a []LtsEntityNonFungibleBalanceChangesable when successful
func (m *LtsCommittedTransactionOutcome) GetNonFungibleEntityBalanceChanges() []LtsEntityNonFungibleBalanceChangesable {
	return m.non_fungible_entity_balance_changes
}

// GetProposerTimestampMs gets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
// returns a *int64 when successful
func (m *LtsCommittedTransactionOutcome) GetProposerTimestampMs() *int64 {
	return m.proposer_timestamp_ms
}

// GetResultantAccountFungibleBalances gets the resultant_account_fungible_balances property value. A list of the resultant fungible account balances for any balances which changed in this transaction.Only balances for accounts are returned, not any other kind of entity.
// returns a []LtsResultantAccountFungibleBalancesable when successful
func (m *LtsCommittedTransactionOutcome) GetResultantAccountFungibleBalances() []LtsResultantAccountFungibleBalancesable {
	return m.resultant_account_fungible_balances
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LtsCommittedTransactionOutcome) GetStateVersion() *int64 {
	return m.state_version
}

// GetStatus gets the status property value. The status of the transaction
// returns a *LtsCommittedTransactionStatus when successful
func (m *LtsCommittedTransactionOutcome) GetStatus() *LtsCommittedTransactionStatus {
	return m.status
}

// GetTotalFee gets the total_fee property value. The string-encoded decimal representing the total amount of XRD paid as fee (execution, validator tip and royalties).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *LtsCommittedTransactionOutcome) GetTotalFee() *string {
	return m.total_fee
}

// GetUserTransactionIdentifiers gets the user_transaction_identifiers property value. The user_transaction_identifiers property
// returns a TransactionIdentifiersable when successful
func (m *LtsCommittedTransactionOutcome) GetUserTransactionIdentifiers() TransactionIdentifiersable {
	return m.user_transaction_identifiers
}

// Serialize serializes information the current object
func (m *LtsCommittedTransactionOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("accumulator_hash", m.GetAccumulatorHash())
		if err != nil {
			return err
		}
	}
	if m.GetFungibleEntityBalanceChanges() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFungibleEntityBalanceChanges()))
		for i, v := range m.GetFungibleEntityBalanceChanges() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("fungible_entity_balance_changes", cast)
		if err != nil {
			return err
		}
	}
	if m.GetNonFungibleEntityBalanceChanges() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonFungibleEntityBalanceChanges()))
		for i, v := range m.GetNonFungibleEntityBalanceChanges() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("non_fungible_entity_balance_changes", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("proposer_timestamp_ms", m.GetProposerTimestampMs())
		if err != nil {
			return err
		}
	}
	if m.GetResultantAccountFungibleBalances() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResultantAccountFungibleBalances()))
		for i, v := range m.GetResultantAccountFungibleBalances() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("resultant_account_fungible_balances", cast)
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
	if m.GetStatus() != nil {
		cast := (*m.GetStatus()).String()
		err := writer.WriteStringValue("status", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("total_fee", m.GetTotalFee())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("user_transaction_identifiers", m.GetUserTransactionIdentifiers())
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

// SetAccumulatorHash sets the accumulator_hash property value. The hex-encoded transaction accumulator hash. This hash captures the order of all transactions on ledger.This hash is `ACC_{N+1} = combine(ACC_N, LEDGER_HASH_{N}))` (where `combine()` is an arbitrary deterministic function we use).
func (m *LtsCommittedTransactionOutcome) SetAccumulatorHash(value *string) {
	m.accumulator_hash = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsCommittedTransactionOutcome) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFungibleEntityBalanceChanges sets the fungible_entity_balance_changes property value. A list of all fungible balance updates which occurred in this transaction, aggregated by the global entity (such as account)which owns the vaults which were updated.
func (m *LtsCommittedTransactionOutcome) SetFungibleEntityBalanceChanges(value []LtsEntityFungibleBalanceChangesable) {
	m.fungible_entity_balance_changes = value
}

// SetNonFungibleEntityBalanceChanges sets the non_fungible_entity_balance_changes property value. Non fungible changes per entity and resource
func (m *LtsCommittedTransactionOutcome) SetNonFungibleEntityBalanceChanges(value []LtsEntityNonFungibleBalanceChangesable) {
	m.non_fungible_entity_balance_changes = value
}

// SetProposerTimestampMs sets the proposer_timestamp_ms property value. An integer between `0` and `10^14`, marking the proposer timestamp in ms.
func (m *LtsCommittedTransactionOutcome) SetProposerTimestampMs(value *int64) {
	m.proposer_timestamp_ms = value
}

// SetResultantAccountFungibleBalances sets the resultant_account_fungible_balances property value. A list of the resultant fungible account balances for any balances which changed in this transaction.Only balances for accounts are returned, not any other kind of entity.
func (m *LtsCommittedTransactionOutcome) SetResultantAccountFungibleBalances(value []LtsResultantAccountFungibleBalancesable) {
	m.resultant_account_fungible_balances = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LtsCommittedTransactionOutcome) SetStateVersion(value *int64) {
	m.state_version = value
}

// SetStatus sets the status property value. The status of the transaction
func (m *LtsCommittedTransactionOutcome) SetStatus(value *LtsCommittedTransactionStatus) {
	m.status = value
}

// SetTotalFee sets the total_fee property value. The string-encoded decimal representing the total amount of XRD paid as fee (execution, validator tip and royalties).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *LtsCommittedTransactionOutcome) SetTotalFee(value *string) {
	m.total_fee = value
}

// SetUserTransactionIdentifiers sets the user_transaction_identifiers property value. The user_transaction_identifiers property
func (m *LtsCommittedTransactionOutcome) SetUserTransactionIdentifiers(value TransactionIdentifiersable) {
	m.user_transaction_identifiers = value
}

type LtsCommittedTransactionOutcomeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccumulatorHash() *string
	GetFungibleEntityBalanceChanges() []LtsEntityFungibleBalanceChangesable
	GetNonFungibleEntityBalanceChanges() []LtsEntityNonFungibleBalanceChangesable
	GetProposerTimestampMs() *int64
	GetResultantAccountFungibleBalances() []LtsResultantAccountFungibleBalancesable
	GetStateVersion() *int64
	GetStatus() *LtsCommittedTransactionStatus
	GetTotalFee() *string
	GetUserTransactionIdentifiers() TransactionIdentifiersable
	SetAccumulatorHash(value *string)
	SetFungibleEntityBalanceChanges(value []LtsEntityFungibleBalanceChangesable)
	SetNonFungibleEntityBalanceChanges(value []LtsEntityNonFungibleBalanceChangesable)
	SetProposerTimestampMs(value *int64)
	SetResultantAccountFungibleBalances(value []LtsResultantAccountFungibleBalancesable)
	SetStateVersion(value *int64)
	SetStatus(value *LtsCommittedTransactionStatus)
	SetTotalFee(value *string)
	SetUserTransactionIdentifiers(value TransactionIdentifiersable)
}
