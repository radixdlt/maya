package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FeeSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction execution.
    execution_cost_units_consumed *int64
    // An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction finalization.
    finalization_cost_units_consumed *int64
    // The string-encoded decimal representing the total amount of XRD burned in the transaction as part of execution costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_total_execution_cost *string
    // The string-encoded decimal representing the total amount of XRD burned in the transaction as part of finalization costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_total_finalization_cost *string
    // The string-encoded decimal representing the total amount of XRD paid in royalties as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_total_royalty_cost *string
    // The string-encoded decimal representing the total amount of XRD paid in state expansion costs as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_total_storage_cost *string
    // The string-encoded decimal representing the total amount of XRD tipped to validators in the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_total_tipping_cost *string
}
// NewFeeSummary instantiates a new FeeSummary and sets the default values.
func NewFeeSummary()(*FeeSummary) {
    m := &FeeSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFeeSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFeeSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeeSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FeeSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExecutionCostUnitsConsumed gets the execution_cost_units_consumed property value. An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction execution.
// returns a *int64 when successful
func (m *FeeSummary) GetExecutionCostUnitsConsumed()(*int64) {
    return m.execution_cost_units_consumed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FeeSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["execution_cost_units_consumed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionCostUnitsConsumed(val)
        }
        return nil
    }
    res["finalization_cost_units_consumed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalizationCostUnitsConsumed(val)
        }
        return nil
    }
    res["xrd_total_execution_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdTotalExecutionCost(val)
        }
        return nil
    }
    res["xrd_total_finalization_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdTotalFinalizationCost(val)
        }
        return nil
    }
    res["xrd_total_royalty_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdTotalRoyaltyCost(val)
        }
        return nil
    }
    res["xrd_total_storage_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdTotalStorageCost(val)
        }
        return nil
    }
    res["xrd_total_tipping_cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdTotalTippingCost(val)
        }
        return nil
    }
    return res
}
// GetFinalizationCostUnitsConsumed gets the finalization_cost_units_consumed property value. An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction finalization.
// returns a *int64 when successful
func (m *FeeSummary) GetFinalizationCostUnitsConsumed()(*int64) {
    return m.finalization_cost_units_consumed
}
// GetXrdTotalExecutionCost gets the xrd_total_execution_cost property value. The string-encoded decimal representing the total amount of XRD burned in the transaction as part of execution costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeSummary) GetXrdTotalExecutionCost()(*string) {
    return m.xrd_total_execution_cost
}
// GetXrdTotalFinalizationCost gets the xrd_total_finalization_cost property value. The string-encoded decimal representing the total amount of XRD burned in the transaction as part of finalization costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeSummary) GetXrdTotalFinalizationCost()(*string) {
    return m.xrd_total_finalization_cost
}
// GetXrdTotalRoyaltyCost gets the xrd_total_royalty_cost property value. The string-encoded decimal representing the total amount of XRD paid in royalties as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeSummary) GetXrdTotalRoyaltyCost()(*string) {
    return m.xrd_total_royalty_cost
}
// GetXrdTotalStorageCost gets the xrd_total_storage_cost property value. The string-encoded decimal representing the total amount of XRD paid in state expansion costs as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeSummary) GetXrdTotalStorageCost()(*string) {
    return m.xrd_total_storage_cost
}
// GetXrdTotalTippingCost gets the xrd_total_tipping_cost property value. The string-encoded decimal representing the total amount of XRD tipped to validators in the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *FeeSummary) GetXrdTotalTippingCost()(*string) {
    return m.xrd_total_tipping_cost
}
// Serialize serializes information the current object
func (m *FeeSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("execution_cost_units_consumed", m.GetExecutionCostUnitsConsumed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("finalization_cost_units_consumed", m.GetFinalizationCostUnitsConsumed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_total_execution_cost", m.GetXrdTotalExecutionCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_total_finalization_cost", m.GetXrdTotalFinalizationCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_total_royalty_cost", m.GetXrdTotalRoyaltyCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_total_storage_cost", m.GetXrdTotalStorageCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_total_tipping_cost", m.GetXrdTotalTippingCost())
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
func (m *FeeSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExecutionCostUnitsConsumed sets the execution_cost_units_consumed property value. An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction execution.
func (m *FeeSummary) SetExecutionCostUnitsConsumed(value *int64)() {
    m.execution_cost_units_consumed = value
}
// SetFinalizationCostUnitsConsumed sets the finalization_cost_units_consumed property value. An integer between `0` and `2^32 - 1`, representing the amount of cost units consumed by the transaction finalization.
func (m *FeeSummary) SetFinalizationCostUnitsConsumed(value *int64)() {
    m.finalization_cost_units_consumed = value
}
// SetXrdTotalExecutionCost sets the xrd_total_execution_cost property value. The string-encoded decimal representing the total amount of XRD burned in the transaction as part of execution costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeSummary) SetXrdTotalExecutionCost(value *string)() {
    m.xrd_total_execution_cost = value
}
// SetXrdTotalFinalizationCost sets the xrd_total_finalization_cost property value. The string-encoded decimal representing the total amount of XRD burned in the transaction as part of finalization costs.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeSummary) SetXrdTotalFinalizationCost(value *string)() {
    m.xrd_total_finalization_cost = value
}
// SetXrdTotalRoyaltyCost sets the xrd_total_royalty_cost property value. The string-encoded decimal representing the total amount of XRD paid in royalties as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeSummary) SetXrdTotalRoyaltyCost(value *string)() {
    m.xrd_total_royalty_cost = value
}
// SetXrdTotalStorageCost sets the xrd_total_storage_cost property value. The string-encoded decimal representing the total amount of XRD paid in state expansion costs as part of the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeSummary) SetXrdTotalStorageCost(value *string)() {
    m.xrd_total_storage_cost = value
}
// SetXrdTotalTippingCost sets the xrd_total_tipping_cost property value. The string-encoded decimal representing the total amount of XRD tipped to validators in the transaction.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *FeeSummary) SetXrdTotalTippingCost(value *string)() {
    m.xrd_total_tipping_cost = value
}
type FeeSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExecutionCostUnitsConsumed()(*int64)
    GetFinalizationCostUnitsConsumed()(*int64)
    GetXrdTotalExecutionCost()(*string)
    GetXrdTotalFinalizationCost()(*string)
    GetXrdTotalRoyaltyCost()(*string)
    GetXrdTotalStorageCost()(*string)
    GetXrdTotalTippingCost()(*string)
    SetExecutionCostUnitsConsumed(value *int64)()
    SetFinalizationCostUnitsConsumed(value *int64)()
    SetXrdTotalExecutionCost(value *string)()
    SetXrdTotalFinalizationCost(value *string)()
    SetXrdTotalRoyaltyCost(value *string)()
    SetXrdTotalStorageCost(value *string)()
    SetXrdTotalTippingCost(value *string)()
}
