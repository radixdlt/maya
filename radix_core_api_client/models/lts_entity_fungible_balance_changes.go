package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsEntityFungibleBalanceChanges struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the entity's address
    entity_address *string
    // The fee_balance_change property
    fee_balance_change LtsFungibleResourceBalanceChangeable
    // If present, this field indicates fee-related balance changes, for example:- Payment of the fee (including tip and royalty)- Distribution of royalties- Distribution of the fee and tip to the consensus-manager, for distributing to the relevant  validator/s at end of epochSee https://www.radixdlt.com/blog/how-fees-work-in-babylon for further information on howfee payment works at Babylon.
    fee_balance_changes []LtsFeeFungibleResourceBalanceChangeable
    // The non_fee_balance_changes property
    non_fee_balance_changes []LtsFungibleResourceBalanceChangeable
}
// NewLtsEntityFungibleBalanceChanges instantiates a new LtsEntityFungibleBalanceChanges and sets the default values.
func NewLtsEntityFungibleBalanceChanges()(*LtsEntityFungibleBalanceChanges) {
    m := &LtsEntityFungibleBalanceChanges{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLtsEntityFungibleBalanceChangesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsEntityFungibleBalanceChangesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLtsEntityFungibleBalanceChanges(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsEntityFungibleBalanceChanges) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEntityAddress gets the entity_address property value. The Bech32m-encoded human readable version of the entity's address
// returns a *string when successful
func (m *LtsEntityFungibleBalanceChanges) GetEntityAddress()(*string) {
    return m.entity_address
}
// GetFeeBalanceChange gets the fee_balance_change property value. The fee_balance_change property
// returns a LtsFungibleResourceBalanceChangeable when successful
func (m *LtsEntityFungibleBalanceChanges) GetFeeBalanceChange()(LtsFungibleResourceBalanceChangeable) {
    return m.fee_balance_change
}
// GetFeeBalanceChanges gets the fee_balance_changes property value. If present, this field indicates fee-related balance changes, for example:- Payment of the fee (including tip and royalty)- Distribution of royalties- Distribution of the fee and tip to the consensus-manager, for distributing to the relevant  validator/s at end of epochSee https://www.radixdlt.com/blog/how-fees-work-in-babylon for further information on howfee payment works at Babylon.
// returns a []LtsFeeFungibleResourceBalanceChangeable when successful
func (m *LtsEntityFungibleBalanceChanges) GetFeeBalanceChanges()([]LtsFeeFungibleResourceBalanceChangeable) {
    return m.fee_balance_changes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsEntityFungibleBalanceChanges) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entity_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityAddress(val)
        }
        return nil
    }
    res["fee_balance_change"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLtsFungibleResourceBalanceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeeBalanceChange(val.(LtsFungibleResourceBalanceChangeable))
        }
        return nil
    }
    res["fee_balance_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLtsFeeFungibleResourceBalanceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LtsFeeFungibleResourceBalanceChangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LtsFeeFungibleResourceBalanceChangeable)
                }
            }
            m.SetFeeBalanceChanges(res)
        }
        return nil
    }
    res["non_fee_balance_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLtsFungibleResourceBalanceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LtsFungibleResourceBalanceChangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LtsFungibleResourceBalanceChangeable)
                }
            }
            m.SetNonFeeBalanceChanges(res)
        }
        return nil
    }
    return res
}
// GetNonFeeBalanceChanges gets the non_fee_balance_changes property value. The non_fee_balance_changes property
// returns a []LtsFungibleResourceBalanceChangeable when successful
func (m *LtsEntityFungibleBalanceChanges) GetNonFeeBalanceChanges()([]LtsFungibleResourceBalanceChangeable) {
    return m.non_fee_balance_changes
}
// Serialize serializes information the current object
func (m *LtsEntityFungibleBalanceChanges) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("entity_address", m.GetEntityAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fee_balance_change", m.GetFeeBalanceChange())
        if err != nil {
            return err
        }
    }
    if m.GetFeeBalanceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeeBalanceChanges()))
        for i, v := range m.GetFeeBalanceChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("fee_balance_changes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNonFeeBalanceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNonFeeBalanceChanges()))
        for i, v := range m.GetNonFeeBalanceChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("non_fee_balance_changes", cast)
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
func (m *LtsEntityFungibleBalanceChanges) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEntityAddress sets the entity_address property value. The Bech32m-encoded human readable version of the entity's address
func (m *LtsEntityFungibleBalanceChanges) SetEntityAddress(value *string)() {
    m.entity_address = value
}
// SetFeeBalanceChange sets the fee_balance_change property value. The fee_balance_change property
func (m *LtsEntityFungibleBalanceChanges) SetFeeBalanceChange(value LtsFungibleResourceBalanceChangeable)() {
    m.fee_balance_change = value
}
// SetFeeBalanceChanges sets the fee_balance_changes property value. If present, this field indicates fee-related balance changes, for example:- Payment of the fee (including tip and royalty)- Distribution of royalties- Distribution of the fee and tip to the consensus-manager, for distributing to the relevant  validator/s at end of epochSee https://www.radixdlt.com/blog/how-fees-work-in-babylon for further information on howfee payment works at Babylon.
func (m *LtsEntityFungibleBalanceChanges) SetFeeBalanceChanges(value []LtsFeeFungibleResourceBalanceChangeable)() {
    m.fee_balance_changes = value
}
// SetNonFeeBalanceChanges sets the non_fee_balance_changes property value. The non_fee_balance_changes property
func (m *LtsEntityFungibleBalanceChanges) SetNonFeeBalanceChanges(value []LtsFungibleResourceBalanceChangeable)() {
    m.non_fee_balance_changes = value
}
type LtsEntityFungibleBalanceChangesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntityAddress()(*string)
    GetFeeBalanceChange()(LtsFungibleResourceBalanceChangeable)
    GetFeeBalanceChanges()([]LtsFeeFungibleResourceBalanceChangeable)
    GetNonFeeBalanceChanges()([]LtsFungibleResourceBalanceChangeable)
    SetEntityAddress(value *string)()
    SetFeeBalanceChange(value LtsFungibleResourceBalanceChangeable)()
    SetFeeBalanceChanges(value []LtsFeeFungibleResourceBalanceChangeable)()
    SetNonFeeBalanceChanges(value []LtsFungibleResourceBalanceChangeable)()
}
