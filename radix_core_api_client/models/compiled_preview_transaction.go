package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompiledPreviewTransaction struct {
    PreviewTransaction
    // A hex-encoded, compiled `RawPreviewTransaction`.As of Cuttlefish, only `PreviewTransactionV2` is supported.A `PreviewTransactionV2` can be created with a v2 transaction builder:* If using Rust, it can be created with a `TransactionV2Builder` using `build_preview_transaction()`  and then converted to hex with `preview_transaction.to_raw().unwrap().to_hex()`* If using the toolkit, you can create this using the v2 transaction builder.Some subtleties:* Partial transactions can't be previewed. Instead, they must be wrapped inside a  transaction wrapper, so that the engine knows how to yield to them appropriately.* Currently the builder assumes that the signed partial transactions have real signatures.  This isn't strictly required, and we may create a builder in future which allows providing  public keys when building partial transactions for use in preview.* If you don't have signatures to hand, you can simply not sign the partial transactions,  and then use the `assume_all_signature_proofs` preview flag, although be advised that  this may result in the fee estimate being slightly lower during preview.* We may create more ergonomic builders for PreviewTransactions which allow use of  public keys to denote the signers of subintents. Let us know if this is important  for your use case.
    preview_transaction_hex *string
}
// NewCompiledPreviewTransaction instantiates a new CompiledPreviewTransaction and sets the default values.
func NewCompiledPreviewTransaction()(*CompiledPreviewTransaction) {
    m := &CompiledPreviewTransaction{
        PreviewTransaction: *NewPreviewTransaction(),
    }
    return m
}
// CreateCompiledPreviewTransactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompiledPreviewTransactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompiledPreviewTransaction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompiledPreviewTransaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PreviewTransaction.GetFieldDeserializers()
    res["preview_transaction_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewTransactionHex(val)
        }
        return nil
    }
    return res
}
// GetPreviewTransactionHex gets the preview_transaction_hex property value. A hex-encoded, compiled `RawPreviewTransaction`.As of Cuttlefish, only `PreviewTransactionV2` is supported.A `PreviewTransactionV2` can be created with a v2 transaction builder:* If using Rust, it can be created with a `TransactionV2Builder` using `build_preview_transaction()`  and then converted to hex with `preview_transaction.to_raw().unwrap().to_hex()`* If using the toolkit, you can create this using the v2 transaction builder.Some subtleties:* Partial transactions can't be previewed. Instead, they must be wrapped inside a  transaction wrapper, so that the engine knows how to yield to them appropriately.* Currently the builder assumes that the signed partial transactions have real signatures.  This isn't strictly required, and we may create a builder in future which allows providing  public keys when building partial transactions for use in preview.* If you don't have signatures to hand, you can simply not sign the partial transactions,  and then use the `assume_all_signature_proofs` preview flag, although be advised that  this may result in the fee estimate being slightly lower during preview.* We may create more ergonomic builders for PreviewTransactions which allow use of  public keys to denote the signers of subintents. Let us know if this is important  for your use case.
// returns a *string when successful
func (m *CompiledPreviewTransaction) GetPreviewTransactionHex()(*string) {
    return m.preview_transaction_hex
}
// Serialize serializes information the current object
func (m *CompiledPreviewTransaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PreviewTransaction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("preview_transaction_hex", m.GetPreviewTransactionHex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPreviewTransactionHex sets the preview_transaction_hex property value. A hex-encoded, compiled `RawPreviewTransaction`.As of Cuttlefish, only `PreviewTransactionV2` is supported.A `PreviewTransactionV2` can be created with a v2 transaction builder:* If using Rust, it can be created with a `TransactionV2Builder` using `build_preview_transaction()`  and then converted to hex with `preview_transaction.to_raw().unwrap().to_hex()`* If using the toolkit, you can create this using the v2 transaction builder.Some subtleties:* Partial transactions can't be previewed. Instead, they must be wrapped inside a  transaction wrapper, so that the engine knows how to yield to them appropriately.* Currently the builder assumes that the signed partial transactions have real signatures.  This isn't strictly required, and we may create a builder in future which allows providing  public keys when building partial transactions for use in preview.* If you don't have signatures to hand, you can simply not sign the partial transactions,  and then use the `assume_all_signature_proofs` preview flag, although be advised that  this may result in the fee estimate being slightly lower during preview.* We may create more ergonomic builders for PreviewTransactions which allow use of  public keys to denote the signers of subintents. Let us know if this is important  for your use case.
func (m *CompiledPreviewTransaction) SetPreviewTransactionHex(value *string)() {
    m.preview_transaction_hex = value
}
type CompiledPreviewTransactionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PreviewTransactionable
    GetPreviewTransactionHex()(*string)
    SetPreviewTransactionHex(value *string)()
}
