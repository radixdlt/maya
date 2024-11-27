package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionSubmitPriorityThresholdNotMetErrorDetails struct {
    TransactionSubmitErrorDetails
    // NOTE: This is kept for backwards compatibility, but we recommend using `min_tip_proportion_required` instead.A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
    // Deprecated: 
    min_tip_percentage_required *int32
    // A lower bound for tip proportion at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
    min_tip_proportion_required *string
    // NOTE: This is kept for backwards compatibility, but we recommend using `tip_proportion` instead.Tip percentage of the submitted (and rejected) transaction. For V2 transactions specifying basis point tips,the amount is rounded down.
    // Deprecated: 
    tip_percentage *int32
    // The string-encoded decimal tip proportion of the submitted (and rejected) transaction.This field will always be present on Cuttlefish nodes, but is marked as not-required for Cuttlefish launch,to avoid a dependency on clients to update after the node is updated.
    tip_proportion *string
}
// NewTransactionSubmitPriorityThresholdNotMetErrorDetails instantiates a new TransactionSubmitPriorityThresholdNotMetErrorDetails and sets the default values.
func NewTransactionSubmitPriorityThresholdNotMetErrorDetails()(*TransactionSubmitPriorityThresholdNotMetErrorDetails) {
    m := &TransactionSubmitPriorityThresholdNotMetErrorDetails{
        TransactionSubmitErrorDetails: *NewTransactionSubmitErrorDetails(),
    }
    return m
}
// CreateTransactionSubmitPriorityThresholdNotMetErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionSubmitPriorityThresholdNotMetErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionSubmitPriorityThresholdNotMetErrorDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TransactionSubmitErrorDetails.GetFieldDeserializers()
    res["min_tip_percentage_required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinTipPercentageRequired(val)
        }
        return nil
    }
    res["min_tip_proportion_required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinTipProportionRequired(val)
        }
        return nil
    }
    res["tip_percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTipPercentage(val)
        }
        return nil
    }
    res["tip_proportion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTipProportion(val)
        }
        return nil
    }
    return res
}
// GetMinTipPercentageRequired gets the min_tip_percentage_required property value. NOTE: This is kept for backwards compatibility, but we recommend using `min_tip_proportion_required` instead.A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
// Deprecated: 
// returns a *int32 when successful
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) GetMinTipPercentageRequired()(*int32) {
    return m.min_tip_percentage_required
}
// GetMinTipProportionRequired gets the min_tip_proportion_required property value. A lower bound for tip proportion at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
// returns a *string when successful
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) GetMinTipProportionRequired()(*string) {
    return m.min_tip_proportion_required
}
// GetTipPercentage gets the tip_percentage property value. NOTE: This is kept for backwards compatibility, but we recommend using `tip_proportion` instead.Tip percentage of the submitted (and rejected) transaction. For V2 transactions specifying basis point tips,the amount is rounded down.
// Deprecated: 
// returns a *int32 when successful
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) GetTipPercentage()(*int32) {
    return m.tip_percentage
}
// GetTipProportion gets the tip_proportion property value. The string-encoded decimal tip proportion of the submitted (and rejected) transaction.This field will always be present on Cuttlefish nodes, but is marked as not-required for Cuttlefish launch,to avoid a dependency on clients to update after the node is updated.
// returns a *string when successful
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) GetTipProportion()(*string) {
    return m.tip_proportion
}
// Serialize serializes information the current object
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TransactionSubmitErrorDetails.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("min_tip_percentage_required", m.GetMinTipPercentageRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("min_tip_proportion_required", m.GetMinTipProportionRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("tip_percentage", m.GetTipPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tip_proportion", m.GetTipProportion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMinTipPercentageRequired sets the min_tip_percentage_required property value. NOTE: This is kept for backwards compatibility, but we recommend using `min_tip_proportion_required` instead.A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
// Deprecated: 
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) SetMinTipPercentageRequired(value *int32)() {
    m.min_tip_percentage_required = value
}
// SetMinTipProportionRequired sets the min_tip_proportion_required property value. A lower bound for tip proportion at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) SetMinTipProportionRequired(value *string)() {
    m.min_tip_proportion_required = value
}
// SetTipPercentage sets the tip_percentage property value. NOTE: This is kept for backwards compatibility, but we recommend using `tip_proportion` instead.Tip percentage of the submitted (and rejected) transaction. For V2 transactions specifying basis point tips,the amount is rounded down.
// Deprecated: 
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) SetTipPercentage(value *int32)() {
    m.tip_percentage = value
}
// SetTipProportion sets the tip_proportion property value. The string-encoded decimal tip proportion of the submitted (and rejected) transaction.This field will always be present on Cuttlefish nodes, but is marked as not-required for Cuttlefish launch,to avoid a dependency on clients to update after the node is updated.
func (m *TransactionSubmitPriorityThresholdNotMetErrorDetails) SetTipProportion(value *string)() {
    m.tip_proportion = value
}
type TransactionSubmitPriorityThresholdNotMetErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TransactionSubmitErrorDetailsable
    GetMinTipPercentageRequired()(*int32)
    GetMinTipProportionRequired()(*string)
    GetTipPercentage()(*int32)
    GetTipProportion()(*string)
    SetMinTipPercentageRequired(value *int32)()
    SetMinTipProportionRequired(value *string)()
    SetTipPercentage(value *int32)()
    SetTipProportion(value *string)()
}
