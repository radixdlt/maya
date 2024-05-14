package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionPreviewRequest_flags struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The assume_all_signature_proofs property
	assume_all_signature_proofs *bool
	// Can be used to skip performing auth checks during runtime.This could be used to e.g.:* Preview protocol update style transactions* Mint resources for previewing trades with resources you don't own. If doing this, be warned:  * Only resources which were potentially mintable/burnable at creation time  will be mintable/burnable, due to feature flags on the resource.  * Please see the below warning about unexpected results if using this approach.Please be warned - this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution* This mode can subtly break invariants some dApp code might rely on, or result in unexpectedbehaviour, so the resulting execution result might not be valid for your needs. For example,if I used this flag to mint pool units to preview a redemption (or some dApp interaction whichbehind the scenes redeemed them), they'd redeem for less than they're currently worth,because the blueprint code relies on the total supply of the pool units to calculate theirredemption worth, and you've just inflated the total supply through the mint operation.
	disable_auth_checks *bool
	// The skip_epoch_check property
	skip_epoch_check *bool
	// The use_free_credit property
	use_free_credit *bool
}

// NewTransactionPreviewRequest_flags instantiates a new TransactionPreviewRequest_flags and sets the default values.
func NewTransactionPreviewRequest_flags() *TransactionPreviewRequest_flags {
	m := &TransactionPreviewRequest_flags{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionPreviewRequest_flagsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionPreviewRequest_flagsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionPreviewRequest_flags(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionPreviewRequest_flags) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAssumeAllSignatureProofs gets the assume_all_signature_proofs property value. The assume_all_signature_proofs property
// returns a *bool when successful
func (m *TransactionPreviewRequest_flags) GetAssumeAllSignatureProofs() *bool {
	return m.assume_all_signature_proofs
}

// GetDisableAuthChecks gets the disable_auth_checks property value. Can be used to skip performing auth checks during runtime.This could be used to e.g.:* Preview protocol update style transactions* Mint resources for previewing trades with resources you don't own. If doing this, be warned:  * Only resources which were potentially mintable/burnable at creation time  will be mintable/burnable, due to feature flags on the resource.  * Please see the below warning about unexpected results if using this approach.Please be warned - this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution* This mode can subtly break invariants some dApp code might rely on, or result in unexpectedbehaviour, so the resulting execution result might not be valid for your needs. For example,if I used this flag to mint pool units to preview a redemption (or some dApp interaction whichbehind the scenes redeemed them), they'd redeem for less than they're currently worth,because the blueprint code relies on the total supply of the pool units to calculate theirredemption worth, and you've just inflated the total supply through the mint operation.
// returns a *bool when successful
func (m *TransactionPreviewRequest_flags) GetDisableAuthChecks() *bool {
	return m.disable_auth_checks
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionPreviewRequest_flags) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["assume_all_signature_proofs"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAssumeAllSignatureProofs(val)
		}
		return nil
	}
	res["disable_auth_checks"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisableAuthChecks(val)
		}
		return nil
	}
	res["skip_epoch_check"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSkipEpochCheck(val)
		}
		return nil
	}
	res["use_free_credit"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUseFreeCredit(val)
		}
		return nil
	}
	return res
}

// GetSkipEpochCheck gets the skip_epoch_check property value. The skip_epoch_check property
// returns a *bool when successful
func (m *TransactionPreviewRequest_flags) GetSkipEpochCheck() *bool {
	return m.skip_epoch_check
}

// GetUseFreeCredit gets the use_free_credit property value. The use_free_credit property
// returns a *bool when successful
func (m *TransactionPreviewRequest_flags) GetUseFreeCredit() *bool {
	return m.use_free_credit
}

// Serialize serializes information the current object
func (m *TransactionPreviewRequest_flags) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("assume_all_signature_proofs", m.GetAssumeAllSignatureProofs())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("disable_auth_checks", m.GetDisableAuthChecks())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("skip_epoch_check", m.GetSkipEpochCheck())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("use_free_credit", m.GetUseFreeCredit())
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
func (m *TransactionPreviewRequest_flags) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAssumeAllSignatureProofs sets the assume_all_signature_proofs property value. The assume_all_signature_proofs property
func (m *TransactionPreviewRequest_flags) SetAssumeAllSignatureProofs(value *bool) {
	m.assume_all_signature_proofs = value
}

// SetDisableAuthChecks sets the disable_auth_checks property value. Can be used to skip performing auth checks during runtime.This could be used to e.g.:* Preview protocol update style transactions* Mint resources for previewing trades with resources you don't own. If doing this, be warned:  * Only resources which were potentially mintable/burnable at creation time  will be mintable/burnable, due to feature flags on the resource.  * Please see the below warning about unexpected results if using this approach.Please be warned - this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution* This mode can subtly break invariants some dApp code might rely on, or result in unexpectedbehaviour, so the resulting execution result might not be valid for your needs. For example,if I used this flag to mint pool units to preview a redemption (or some dApp interaction whichbehind the scenes redeemed them), they'd redeem for less than they're currently worth,because the blueprint code relies on the total supply of the pool units to calculate theirredemption worth, and you've just inflated the total supply through the mint operation.
func (m *TransactionPreviewRequest_flags) SetDisableAuthChecks(value *bool) {
	m.disable_auth_checks = value
}

// SetSkipEpochCheck sets the skip_epoch_check property value. The skip_epoch_check property
func (m *TransactionPreviewRequest_flags) SetSkipEpochCheck(value *bool) {
	m.skip_epoch_check = value
}

// SetUseFreeCredit sets the use_free_credit property value. The use_free_credit property
func (m *TransactionPreviewRequest_flags) SetUseFreeCredit(value *bool) {
	m.use_free_credit = value
}

type TransactionPreviewRequest_flagsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAssumeAllSignatureProofs() *bool
	GetDisableAuthChecks() *bool
	GetSkipEpochCheck() *bool
	GetUseFreeCredit() *bool
	SetAssumeAllSignatureProofs(value *bool)
	SetDisableAuthChecks(value *bool)
	SetSkipEpochCheck(value *bool)
	SetUseFreeCredit(value *bool)
}
