package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionHeaderV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether the notary public key should be included in the transaction signers list
    notary_is_signatory *bool
    // The notary_public_key property
    notary_public_key PublicKeyable
    // An integer between `0` and `2^32 - 1 = 4294967295`, giving the validator tip as a basis points amount.That is, a value of `1` corresponds to an additional tip on 0.01% of the base fee.
    tip_basis_points *int64
}
// NewTransactionHeaderV2 instantiates a new TransactionHeaderV2 and sets the default values.
func NewTransactionHeaderV2()(*TransactionHeaderV2) {
    m := &TransactionHeaderV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionHeaderV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionHeaderV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionHeaderV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionHeaderV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionHeaderV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notary_is_signatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotaryIsSignatory(val)
        }
        return nil
    }
    res["notary_public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotaryPublicKey(val.(PublicKeyable))
        }
        return nil
    }
    res["tip_basis_points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTipBasisPoints(val)
        }
        return nil
    }
    return res
}
// GetNotaryIsSignatory gets the notary_is_signatory property value. Specifies whether the notary public key should be included in the transaction signers list
// returns a *bool when successful
func (m *TransactionHeaderV2) GetNotaryIsSignatory()(*bool) {
    return m.notary_is_signatory
}
// GetNotaryPublicKey gets the notary_public_key property value. The notary_public_key property
// returns a PublicKeyable when successful
func (m *TransactionHeaderV2) GetNotaryPublicKey()(PublicKeyable) {
    return m.notary_public_key
}
// GetTipBasisPoints gets the tip_basis_points property value. An integer between `0` and `2^32 - 1 = 4294967295`, giving the validator tip as a basis points amount.That is, a value of `1` corresponds to an additional tip on 0.01% of the base fee.
// returns a *int64 when successful
func (m *TransactionHeaderV2) GetTipBasisPoints()(*int64) {
    return m.tip_basis_points
}
// Serialize serializes information the current object
func (m *TransactionHeaderV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notary_is_signatory", m.GetNotaryIsSignatory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notary_public_key", m.GetNotaryPublicKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("tip_basis_points", m.GetTipBasisPoints())
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
func (m *TransactionHeaderV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotaryIsSignatory sets the notary_is_signatory property value. Specifies whether the notary public key should be included in the transaction signers list
func (m *TransactionHeaderV2) SetNotaryIsSignatory(value *bool)() {
    m.notary_is_signatory = value
}
// SetNotaryPublicKey sets the notary_public_key property value. The notary_public_key property
func (m *TransactionHeaderV2) SetNotaryPublicKey(value PublicKeyable)() {
    m.notary_public_key = value
}
// SetTipBasisPoints sets the tip_basis_points property value. An integer between `0` and `2^32 - 1 = 4294967295`, giving the validator tip as a basis points amount.That is, a value of `1` corresponds to an additional tip on 0.01% of the base fee.
func (m *TransactionHeaderV2) SetTipBasisPoints(value *int64)() {
    m.tip_basis_points = value
}
type TransactionHeaderV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNotaryIsSignatory()(*bool)
    GetNotaryPublicKey()(PublicKeyable)
    GetTipBasisPoints()(*int64)
    SetNotaryIsSignatory(value *bool)()
    SetNotaryPublicKey(value PublicKeyable)()
    SetTipBasisPoints(value *int64)()
}
