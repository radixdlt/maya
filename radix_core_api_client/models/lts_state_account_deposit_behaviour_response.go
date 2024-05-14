package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsStateAccountDepositBehaviourResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
	default_deposit_rule *DefaultDepositRule
	// Whether the input `badge` belongs to the account's set of authorized depositors.This field will only be present if any badge was passed in the request.
	is_badge_authorized_depositor *bool
	// The ledger_header_summary property
	ledger_header_summary LedgerHeaderSummaryable
	// A map from one of the input `resource_addresses` to its specific deposit behavior configured for this account.This field will only be present if an array of specific resource addresses was passed in the request (even if empty).
	resource_specific_behaviours LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable
	// The state_version property
	state_version *int64
}

// NewLtsStateAccountDepositBehaviourResponse instantiates a new LtsStateAccountDepositBehaviourResponse and sets the default values.
func NewLtsStateAccountDepositBehaviourResponse() *LtsStateAccountDepositBehaviourResponse {
	m := &LtsStateAccountDepositBehaviourResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsStateAccountDepositBehaviourResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsStateAccountDepositBehaviourResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsStateAccountDepositBehaviourResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultDepositRule gets the default_deposit_rule property value. This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
// returns a *DefaultDepositRule when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetDefaultDepositRule() *DefaultDepositRule {
	return m.default_deposit_rule
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_deposit_rule"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseDefaultDepositRule)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultDepositRule(val.(*DefaultDepositRule))
		}
		return nil
	}
	res["is_badge_authorized_depositor"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsBadgeAuthorizedDepositor(val)
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
	res["resource_specific_behaviours"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLtsStateAccountDepositBehaviourResponse_resource_specific_behavioursFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResourceSpecificBehaviours(val.(LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable))
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

// GetIsBadgeAuthorizedDepositor gets the is_badge_authorized_depositor property value. Whether the input `badge` belongs to the account's set of authorized depositors.This field will only be present if any badge was passed in the request.
// returns a *bool when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetIsBadgeAuthorizedDepositor() *bool {
	return m.is_badge_authorized_depositor
}

// GetLedgerHeaderSummary gets the ledger_header_summary property value. The ledger_header_summary property
// returns a LedgerHeaderSummaryable when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetLedgerHeaderSummary() LedgerHeaderSummaryable {
	return m.ledger_header_summary
}

// GetResourceSpecificBehaviours gets the resource_specific_behaviours property value. A map from one of the input `resource_addresses` to its specific deposit behavior configured for this account.This field will only be present if an array of specific resource addresses was passed in the request (even if empty).
// returns a LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetResourceSpecificBehaviours() LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable {
	return m.resource_specific_behaviours
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LtsStateAccountDepositBehaviourResponse) GetStateVersion() *int64 {
	return m.state_version
}

// Serialize serializes information the current object
func (m *LtsStateAccountDepositBehaviourResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetDefaultDepositRule() != nil {
		cast := (*m.GetDefaultDepositRule()).String()
		err := writer.WriteStringValue("default_deposit_rule", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_badge_authorized_depositor", m.GetIsBadgeAuthorizedDepositor())
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
		err := writer.WriteObjectValue("resource_specific_behaviours", m.GetResourceSpecificBehaviours())
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

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LtsStateAccountDepositBehaviourResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultDepositRule sets the default_deposit_rule property value. This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
func (m *LtsStateAccountDepositBehaviourResponse) SetDefaultDepositRule(value *DefaultDepositRule) {
	m.default_deposit_rule = value
}

// SetIsBadgeAuthorizedDepositor sets the is_badge_authorized_depositor property value. Whether the input `badge` belongs to the account's set of authorized depositors.This field will only be present if any badge was passed in the request.
func (m *LtsStateAccountDepositBehaviourResponse) SetIsBadgeAuthorizedDepositor(value *bool) {
	m.is_badge_authorized_depositor = value
}

// SetLedgerHeaderSummary sets the ledger_header_summary property value. The ledger_header_summary property
func (m *LtsStateAccountDepositBehaviourResponse) SetLedgerHeaderSummary(value LedgerHeaderSummaryable) {
	m.ledger_header_summary = value
}

// SetResourceSpecificBehaviours sets the resource_specific_behaviours property value. A map from one of the input `resource_addresses` to its specific deposit behavior configured for this account.This field will only be present if an array of specific resource addresses was passed in the request (even if empty).
func (m *LtsStateAccountDepositBehaviourResponse) SetResourceSpecificBehaviours(value LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable) {
	m.resource_specific_behaviours = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LtsStateAccountDepositBehaviourResponse) SetStateVersion(value *int64) {
	m.state_version = value
}

type LtsStateAccountDepositBehaviourResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultDepositRule() *DefaultDepositRule
	GetIsBadgeAuthorizedDepositor() *bool
	GetLedgerHeaderSummary() LedgerHeaderSummaryable
	GetResourceSpecificBehaviours() LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable
	GetStateVersion() *int64
	SetDefaultDepositRule(value *DefaultDepositRule)
	SetIsBadgeAuthorizedDepositor(value *bool)
	SetLedgerHeaderSummary(value LedgerHeaderSummaryable)
	SetResourceSpecificBehaviours(value LtsStateAccountDepositBehaviourResponse_resource_specific_behavioursable)
	SetStateVersion(value *int64)
}
