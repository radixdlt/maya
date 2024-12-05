package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreviewFlags struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If enabled, each manifest processor's auth zone will be given a simulated proof ofevery signature, which can be used to pass signature access rules.This can be used to preview transactions even if the required signatures are notknown ahead of time.See the documentation on[advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules#signature-requirements)for more information.
    assume_all_signature_proofs *bool
    // If enabled, all authorization checks are skipped during execution.This could be used to e.g.:* Preview protocol update style transactions.* Mint resources for previewing trades with resources you don't own.  If doing this, be warned: only resources which were potentially mintable/burnable  at creation time will be mintable/burnable, due to feature flags on the resource.Warning: this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution.* This mode can subtly break invariants some dApp code might rely on, or result in unexpected  behaviour, so the execution result might not be valid for your needs. For example,  if this flag was used to mint pool units to preview a redemption (or some dApp interaction which  behind the scenes redeemed them), they'd redeem for less than they're currently worth,  because the blueprint code relies on the total supply of the pool units to calculate their  redemption worth, and you've just inflated the total supply through the mint operation.
    disable_auth_checks *bool
    // If enabled, the various runtime epoch-related verifications are skipped:- The `start_epoch_inclusive` and `end_epoch_exclusive` parameters, if specified, are ignored.- The duplicate intent checks (which rely on the expiry epoch) are also ignored.However, if the start and end epoch are provided, they must still be statically valid.We recommend using a value of `start_epoch_inclusive = 1` and `end_epoch_exclusive = 2` in thiscase.
    skip_epoch_check *bool
    // If enabled, a large simulated pool of XRD is marked as locked.This mode can be used to estimate fees. To get a reliable estimate, we recommend that yourtransaction is as close as possible to the real transaction. For example:- You should still use a lock fee command, but you can set it to lock a fee of 0.- You should include the public keys that will sign the transaction, so the cost of  signature verification and payload size can be accounted for.
    use_free_credit *bool
}
// NewPreviewFlags instantiates a new PreviewFlags and sets the default values.
func NewPreviewFlags()(*PreviewFlags) {
    m := &PreviewFlags{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreviewFlagsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreviewFlagsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreviewFlags(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreviewFlags) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssumeAllSignatureProofs gets the assume_all_signature_proofs property value. If enabled, each manifest processor's auth zone will be given a simulated proof ofevery signature, which can be used to pass signature access rules.This can be used to preview transactions even if the required signatures are notknown ahead of time.See the documentation on[advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules#signature-requirements)for more information.
// returns a *bool when successful
func (m *PreviewFlags) GetAssumeAllSignatureProofs()(*bool) {
    return m.assume_all_signature_proofs
}
// GetDisableAuthChecks gets the disable_auth_checks property value. If enabled, all authorization checks are skipped during execution.This could be used to e.g.:* Preview protocol update style transactions.* Mint resources for previewing trades with resources you don't own.  If doing this, be warned: only resources which were potentially mintable/burnable  at creation time will be mintable/burnable, due to feature flags on the resource.Warning: this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution.* This mode can subtly break invariants some dApp code might rely on, or result in unexpected  behaviour, so the execution result might not be valid for your needs. For example,  if this flag was used to mint pool units to preview a redemption (or some dApp interaction which  behind the scenes redeemed them), they'd redeem for less than they're currently worth,  because the blueprint code relies on the total supply of the pool units to calculate their  redemption worth, and you've just inflated the total supply through the mint operation.
// returns a *bool when successful
func (m *PreviewFlags) GetDisableAuthChecks()(*bool) {
    return m.disable_auth_checks
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreviewFlags) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assume_all_signature_proofs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssumeAllSignatureProofs(val)
        }
        return nil
    }
    res["disable_auth_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAuthChecks(val)
        }
        return nil
    }
    res["skip_epoch_check"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipEpochCheck(val)
        }
        return nil
    }
    res["use_free_credit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetSkipEpochCheck gets the skip_epoch_check property value. If enabled, the various runtime epoch-related verifications are skipped:- The `start_epoch_inclusive` and `end_epoch_exclusive` parameters, if specified, are ignored.- The duplicate intent checks (which rely on the expiry epoch) are also ignored.However, if the start and end epoch are provided, they must still be statically valid.We recommend using a value of `start_epoch_inclusive = 1` and `end_epoch_exclusive = 2` in thiscase.
// returns a *bool when successful
func (m *PreviewFlags) GetSkipEpochCheck()(*bool) {
    return m.skip_epoch_check
}
// GetUseFreeCredit gets the use_free_credit property value. If enabled, a large simulated pool of XRD is marked as locked.This mode can be used to estimate fees. To get a reliable estimate, we recommend that yourtransaction is as close as possible to the real transaction. For example:- You should still use a lock fee command, but you can set it to lock a fee of 0.- You should include the public keys that will sign the transaction, so the cost of  signature verification and payload size can be accounted for.
// returns a *bool when successful
func (m *PreviewFlags) GetUseFreeCredit()(*bool) {
    return m.use_free_credit
}
// Serialize serializes information the current object
func (m *PreviewFlags) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PreviewFlags) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssumeAllSignatureProofs sets the assume_all_signature_proofs property value. If enabled, each manifest processor's auth zone will be given a simulated proof ofevery signature, which can be used to pass signature access rules.This can be used to preview transactions even if the required signatures are notknown ahead of time.See the documentation on[advanced access rules](https://docs.radixdlt.com/docs/advanced-accessrules#signature-requirements)for more information.
func (m *PreviewFlags) SetAssumeAllSignatureProofs(value *bool)() {
    m.assume_all_signature_proofs = value
}
// SetDisableAuthChecks sets the disable_auth_checks property value. If enabled, all authorization checks are skipped during execution.This could be used to e.g.:* Preview protocol update style transactions.* Mint resources for previewing trades with resources you don't own.  If doing this, be warned: only resources which were potentially mintable/burnable  at creation time will be mintable/burnable, due to feature flags on the resource.Warning: this mode of operation is quite a departure from normal operation:* Calculated fees will likely be lower than a standard execution.* This mode can subtly break invariants some dApp code might rely on, or result in unexpected  behaviour, so the execution result might not be valid for your needs. For example,  if this flag was used to mint pool units to preview a redemption (or some dApp interaction which  behind the scenes redeemed them), they'd redeem for less than they're currently worth,  because the blueprint code relies on the total supply of the pool units to calculate their  redemption worth, and you've just inflated the total supply through the mint operation.
func (m *PreviewFlags) SetDisableAuthChecks(value *bool)() {
    m.disable_auth_checks = value
}
// SetSkipEpochCheck sets the skip_epoch_check property value. If enabled, the various runtime epoch-related verifications are skipped:- The `start_epoch_inclusive` and `end_epoch_exclusive` parameters, if specified, are ignored.- The duplicate intent checks (which rely on the expiry epoch) are also ignored.However, if the start and end epoch are provided, they must still be statically valid.We recommend using a value of `start_epoch_inclusive = 1` and `end_epoch_exclusive = 2` in thiscase.
func (m *PreviewFlags) SetSkipEpochCheck(value *bool)() {
    m.skip_epoch_check = value
}
// SetUseFreeCredit sets the use_free_credit property value. If enabled, a large simulated pool of XRD is marked as locked.This mode can be used to estimate fees. To get a reliable estimate, we recommend that yourtransaction is as close as possible to the real transaction. For example:- You should still use a lock fee command, but you can set it to lock a fee of 0.- You should include the public keys that will sign the transaction, so the cost of  signature verification and payload size can be accounted for.
func (m *PreviewFlags) SetUseFreeCredit(value *bool)() {
    m.use_free_credit = value
}
type PreviewFlagsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssumeAllSignatureProofs()(*bool)
    GetDisableAuthChecks()(*bool)
    GetSkipEpochCheck()(*bool)
    GetUseFreeCredit()(*bool)
    SetAssumeAllSignatureProofs(value *bool)()
    SetDisableAuthChecks(value *bool)()
    SetSkipEpochCheck(value *bool)()
    SetUseFreeCredit(value *bool)()
}
