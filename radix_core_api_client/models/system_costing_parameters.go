package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SystemCostingParameters struct {
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
    // The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_archive_storage_price *string
    // The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_storage_price *string
    // The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
    xrd_usd_price *string
}
// NewSystemCostingParameters instantiates a new SystemCostingParameters and sets the default values.
func NewSystemCostingParameters()(*SystemCostingParameters) {
    m := &SystemCostingParameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSystemCostingParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSystemCostingParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSystemCostingParameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SystemCostingParameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExecutionCostUnitLimit gets the execution_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction execution.
// returns a *int64 when successful
func (m *SystemCostingParameters) GetExecutionCostUnitLimit()(*int64) {
    return m.execution_cost_unit_limit
}
// GetExecutionCostUnitLoan gets the execution_cost_unit_loan property value. An integer between `0` and `2^32 - 1`, representing the maximum number of cost units which can be used before fee is locked from a vault.
// returns a *int64 when successful
func (m *SystemCostingParameters) GetExecutionCostUnitLoan()(*int64) {
    return m.execution_cost_unit_loan
}
// GetExecutionCostUnitPrice gets the execution_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction execution.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SystemCostingParameters) GetExecutionCostUnitPrice()(*string) {
    return m.execution_cost_unit_price
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SystemCostingParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *SystemCostingParameters) GetFinalizationCostUnitLimit()(*int64) {
    return m.finalization_cost_unit_limit
}
// GetFinalizationCostUnitPrice gets the finalization_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction finalization.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SystemCostingParameters) GetFinalizationCostUnitPrice()(*string) {
    return m.finalization_cost_unit_price
}
// GetXrdArchiveStoragePrice gets the xrd_archive_storage_price property value. The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SystemCostingParameters) GetXrdArchiveStoragePrice()(*string) {
    return m.xrd_archive_storage_price
}
// GetXrdStoragePrice gets the xrd_storage_price property value. The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SystemCostingParameters) GetXrdStoragePrice()(*string) {
    return m.xrd_storage_price
}
// GetXrdUsdPrice gets the xrd_usd_price property value. The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *SystemCostingParameters) GetXrdUsdPrice()(*string) {
    return m.xrd_usd_price
}
// Serialize serializes information the current object
func (m *SystemCostingParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SystemCostingParameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExecutionCostUnitLimit sets the execution_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction execution.
func (m *SystemCostingParameters) SetExecutionCostUnitLimit(value *int64)() {
    m.execution_cost_unit_limit = value
}
// SetExecutionCostUnitLoan sets the execution_cost_unit_loan property value. An integer between `0` and `2^32 - 1`, representing the maximum number of cost units which can be used before fee is locked from a vault.
func (m *SystemCostingParameters) SetExecutionCostUnitLoan(value *int64)() {
    m.execution_cost_unit_loan = value
}
// SetExecutionCostUnitPrice sets the execution_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction execution.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SystemCostingParameters) SetExecutionCostUnitPrice(value *string)() {
    m.execution_cost_unit_price = value
}
// SetFinalizationCostUnitLimit sets the finalization_cost_unit_limit property value. An integer between `0` and `2^32 - 1`, representing the maximum amount of cost units available for the transaction finalization.
func (m *SystemCostingParameters) SetFinalizationCostUnitLimit(value *int64)() {
    m.finalization_cost_unit_limit = value
}
// SetFinalizationCostUnitPrice sets the finalization_cost_unit_price property value. The string-encoded decimal representing the XRD price of a single cost unit of transaction finalization.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SystemCostingParameters) SetFinalizationCostUnitPrice(value *string)() {
    m.finalization_cost_unit_price = value
}
// SetXrdArchiveStoragePrice sets the xrd_archive_storage_price property value. The string-encoded decimal representing the price of 1 byte of archive storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SystemCostingParameters) SetXrdArchiveStoragePrice(value *string)() {
    m.xrd_archive_storage_price = value
}
// SetXrdStoragePrice sets the xrd_storage_price property value. The string-encoded decimal representing the price of 1 byte of state storage, expressed in XRD.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SystemCostingParameters) SetXrdStoragePrice(value *string)() {
    m.xrd_storage_price = value
}
// SetXrdUsdPrice sets the xrd_usd_price property value. The string-encoded decimal representing what amount of XRD is consumed by a Royalty of 1 USD.This is fixed for a given protocol version, so is not an accurate representation of the XRD price.A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *SystemCostingParameters) SetXrdUsdPrice(value *string)() {
    m.xrd_usd_price = value
}
type SystemCostingParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExecutionCostUnitLimit()(*int64)
    GetExecutionCostUnitLoan()(*int64)
    GetExecutionCostUnitPrice()(*string)
    GetFinalizationCostUnitLimit()(*int64)
    GetFinalizationCostUnitPrice()(*string)
    GetXrdArchiveStoragePrice()(*string)
    GetXrdStoragePrice()(*string)
    GetXrdUsdPrice()(*string)
    SetExecutionCostUnitLimit(value *int64)()
    SetExecutionCostUnitLoan(value *int64)()
    SetExecutionCostUnitPrice(value *string)()
    SetFinalizationCostUnitLimit(value *int64)()
    SetFinalizationCostUnitPrice(value *string)()
    SetXrdArchiveStoragePrice(value *string)()
    SetXrdStoragePrice(value *string)()
    SetXrdUsdPrice(value *string)()
}
