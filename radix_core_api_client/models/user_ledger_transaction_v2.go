package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserLedgerTransactionV2 struct {
    LedgerTransaction
    // The notarized_transaction property
    notarized_transaction NotarizedTransactionV2able
}
// NewUserLedgerTransactionV2 instantiates a new UserLedgerTransactionV2 and sets the default values.
func NewUserLedgerTransactionV2()(*UserLedgerTransactionV2) {
    m := &UserLedgerTransactionV2{
        LedgerTransaction: *NewLedgerTransaction(),
    }
    return m
}
// CreateUserLedgerTransactionV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserLedgerTransactionV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserLedgerTransactionV2(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserLedgerTransactionV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LedgerTransaction.GetFieldDeserializers()
    res["notarized_transaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotarizedTransactionV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotarizedTransaction(val.(NotarizedTransactionV2able))
        }
        return nil
    }
    return res
}
// GetNotarizedTransaction gets the notarized_transaction property value. The notarized_transaction property
// returns a NotarizedTransactionV2able when successful
func (m *UserLedgerTransactionV2) GetNotarizedTransaction()(NotarizedTransactionV2able) {
    return m.notarized_transaction
}
// Serialize serializes information the current object
func (m *UserLedgerTransactionV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LedgerTransaction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("notarized_transaction", m.GetNotarizedTransaction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNotarizedTransaction sets the notarized_transaction property value. The notarized_transaction property
func (m *UserLedgerTransactionV2) SetNotarizedTransaction(value NotarizedTransactionV2able)() {
    m.notarized_transaction = value
}
type UserLedgerTransactionV2able interface {
    LedgerTransactionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNotarizedTransaction()(NotarizedTransactionV2able)
    SetNotarizedTransaction(value NotarizedTransactionV2able)()
}
