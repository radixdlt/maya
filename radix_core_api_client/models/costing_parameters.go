package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CostingParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction execution.
    execution_cost_unit_limit *int64
    // An integer between `0` and `2^32 - 1`, representing the maximum number of cost units which can be used before fee is locked from a vault.
    execution_cost_unit_loan *int64
    // The string-encoded decimal representing the XRD price of a single cost unit of transaction execution.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    execution_cost_unit_price *string
    // An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction finalization.
    finalization_cost_unit_limit *int64
    // The string-encoded decimal representing the XRD price of a single cost unit of transaction finalization.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    finalization_cost_unit_price *string
    // NOTE: V2 transactions specify the tip in basis points, which gets rounded down for this `tip_percentage` field.It is recommended to instead use the `tip_proportion` field to get a fully accurate value.An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
    // Deprecated: 
    tip_percentage *int32
    // A string-encoded decimal, giving the validator tip as a proportional amount.A value of `"0.01"` corresponds to 1% of the fee being paid as a tip.NOTE: This field is not marked as required for Cuttlefish launch, to permit cuttlefish clients to talk to pre-cuttlefish nodes.This can be changed after Cuttlefish enactment once all nodes are on Cuttlefish.
    tip_proportion *string
    // The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_archive_storage_price *string
    // The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_storage_price *string
    // The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_usd_price *string
}
// NewCostingParameters instantiates a new CostingParameters and sets the default values.
func NewCostingParameters()(*CostingParameters) {
    m := &CostingParameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCostingParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCostingParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCostingParameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CostingParameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExecutionCostUnitLimit gets the execution_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction execution.
// returns a *int64 when successful
func (m *CostingParameters) GetExecutionCostUnitLimit()(*int64) {
    return m.execution_cost_unit_limit
}
// GetExecutionCostUnitLoan gets the execution_cost_unit_loan property value. An integer between `0` and `2^32 - 1`, representing the maximum number of cost units which can be used before fee is locked from a vault.
// returns a *int64 when successful
func (m *CostingParameters) GetExecutionCostUnitLoan()(*int64) {
    return m.execution_cost_unit_loan
}
// GetExecutionCostUnitPrice gets the execution_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction execution.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingParameters) GetExecutionCostUnitPrice()(*string) {
    return m.execution_cost_unit_price
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CostingParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["execution_cost_unit_limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionCostUnitLimit(val)
        }
        return nil
    }
    res["execution_cost_unit_loan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionCostUnitLoan(val)
        }
        return nil
    }
    res["execution_cost_unit_price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionCostUnitPrice(val)
        }
        return nil
    }
    res["finalization_cost_unit_limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalizationCostUnitLimit(val)
        }
        return nil
    }
    res["finalization_cost_unit_price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalizationCostUnitPrice(val)
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
    res["xrd_archive_storage_price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdArchiveStoragePrice(val)
        }
        return nil
    }
    res["xrd_storage_price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdStoragePrice(val)
        }
        return nil
    }
    res["xrd_usd_price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXrdUsdPrice(val)
        }
        return nil
    }
    return res
}
// GetFinalizationCostUnitLimit gets the finalization_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction finalization.
// returns a *int64 when successful
func (m *CostingParameters) GetFinalizationCostUnitLimit()(*int64) {
    return m.finalization_cost_unit_limit
}
// GetFinalizationCostUnitPrice gets the finalization_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction finalization.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingParameters) GetFinalizationCostUnitPrice()(*string) {
    return m.finalization_cost_unit_price
}
// GetTipPercentage gets the tip_percentage property value. NOTE: V2 transactions specify the tip in basis points, which gets rounded down for this `tip_percentage` field.It is recommended to instead use the `tip_proportion` field to get a fully accurate value.An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
// Deprecated: 
// returns a *int32 when successful
func (m *CostingParameters) GetTipPercentage()(*int32) {
    return m.tip_percentage
}
// GetTipProportion gets the tip_proportion property value. A string-encoded decimal, giving the validator tip as a proportional amount.A value of `"0.01"` corresponds to 1% of the fee being paid as a tip.NOTE: This field is not marked as required for Cuttlefish launch, to permit cuttlefish clients to talk to pre-cuttlefish nodes.This can be changed after Cuttlefish enactment once all nodes are on Cuttlefish.
// returns a *string when successful
func (m *CostingParameters) GetTipProportion()(*string) {
    return m.tip_proportion
}
// GetXrdArchiveStoragePrice gets the xrd_archive_storage_price property value. The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingParameters) GetXrdArchiveStoragePrice()(*string) {
    return m.xrd_archive_storage_price
}
// GetXrdStoragePrice gets the xrd_storage_price property value. The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingParameters) GetXrdStoragePrice()(*string) {
    return m.xrd_storage_price
}
// GetXrdUsdPrice gets the xrd_usd_price property value. The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *CostingParameters) GetXrdUsdPrice()(*string) {
    return m.xrd_usd_price
}
// Serialize serializes information the current object
func (m *CostingParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("execution_cost_unit_limit", m.GetExecutionCostUnitLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("execution_cost_unit_loan", m.GetExecutionCostUnitLoan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("execution_cost_unit_price", m.GetExecutionCostUnitPrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("finalization_cost_unit_limit", m.GetFinalizationCostUnitLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("finalization_cost_unit_price", m.GetFinalizationCostUnitPrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("tip_percentage", m.GetTipPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tip_proportion", m.GetTipProportion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_archive_storage_price", m.GetXrdArchiveStoragePrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_storage_price", m.GetXrdStoragePrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("xrd_usd_price", m.GetXrdUsdPrice())
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
func (m *CostingParameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExecutionCostUnitLimit sets the execution_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction execution.
func (m *CostingParameters) SetExecutionCostUnitLimit(value *int64)() {
    m.execution_cost_unit_limit = value
}
// SetExecutionCostUnitLoan sets the execution_cost_unit_loan property value. An integer between `0` and `2^32 - 1`, representing the maximum number of cost units which can be used before fee is locked from a vault.
func (m *CostingParameters) SetExecutionCostUnitLoan(value *int64)() {
    m.execution_cost_unit_loan = value
}
// SetExecutionCostUnitPrice sets the execution_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction execution.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingParameters) SetExecutionCostUnitPrice(value *string)() {
    m.execution_cost_unit_price = value
}
// SetFinalizationCostUnitLimit sets the finalization_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction finalization.
func (m *CostingParameters) SetFinalizationCostUnitLimit(value *int64)() {
    m.finalization_cost_unit_limit = value
}
// SetFinalizationCostUnitPrice sets the finalization_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction finalization.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingParameters) SetFinalizationCostUnitPrice(value *string)() {
    m.finalization_cost_unit_price = value
}
// SetTipPercentage sets the tip_percentage property value. NOTE: V2 transactions specify the tip in basis points, which gets rounded down for this `tip_percentage` field.It is recommended to instead use the `tip_proportion` field to get a fully accurate value.An integer between `0` and `65535`, giving the validator tip as a percentage amount. A value of `1` corresponds to 1% of the fee.
// Deprecated: 
func (m *CostingParameters) SetTipPercentage(value *int32)() {
    m.tip_percentage = value
}
// SetTipProportion sets the tip_proportion property value. A string-encoded decimal, giving the validator tip as a proportional amount.A value of `"0.01"` corresponds to 1% of the fee being paid as a tip.NOTE: This field is not marked as required for Cuttlefish launch, to permit cuttlefish clients to talk to pre-cuttlefish nodes.This can be changed after Cuttlefish enactment once all nodes are on Cuttlefish.
func (m *CostingParameters) SetTipProportion(value *string)() {
    m.tip_proportion = value
}
// SetXrdArchiveStoragePrice sets the xrd_archive_storage_price property value. The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingParameters) SetXrdArchiveStoragePrice(value *string)() {
    m.xrd_archive_storage_price = value
}
// SetXrdStoragePrice sets the xrd_storage_price property value. The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingParameters) SetXrdStoragePrice(value *string)() {
    m.xrd_storage_price = value
}
// SetXrdUsdPrice sets the xrd_usd_price property value. The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *CostingParameters) SetXrdUsdPrice(value *string)() {
    m.xrd_usd_price = value
}
type CostingParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExecutionCostUnitLimit()(*int64)
    GetExecutionCostUnitLoan()(*int64)
    GetExecutionCostUnitPrice()(*string)
    GetFinalizationCostUnitLimit()(*int64)
    GetFinalizationCostUnitPrice()(*string)
    GetTipPercentage()(*int32)
    GetTipProportion()(*string)
    GetXrdArchiveStoragePrice()(*string)
    GetXrdStoragePrice()(*string)
    GetXrdUsdPrice()(*string)
    SetExecutionCostUnitLimit(value *int64)()
    SetExecutionCostUnitLoan(value *int64)()
    SetExecutionCostUnitPrice(value *string)()
    SetFinalizationCostUnitLimit(value *int64)()
    SetFinalizationCostUnitPrice(value *string)()
    SetTipPercentage(value *int32)()
    SetTipProportion(value *string)()
    SetXrdArchiveStoragePrice(value *string)()
    SetXrdStoragePrice(value *string)()
    SetXrdUsdPrice(value *string)()
}
