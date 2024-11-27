package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionValidationConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The manifest_validation property
    manifest_validation *ManifestValidationRuleset
    // The max_epoch_range property
    max_epoch_range *string
    // The max_instructions property
    max_instructions *string
    // The max_references_per_intent property
    max_references_per_intent *string
    // The max_signer_signatures_per_intent property
    max_signer_signatures_per_intent *string
    // The max_subintent_depth property
    max_subintent_depth *string
    // The max_tip_basis_points property
    max_tip_basis_points *int64
    // Only applies to V1 transactions
    max_tip_percentage *int32
    // The max_total_references property
    max_total_references *string
    // The max_total_signature_validations property
    max_total_signature_validations *string
    // The message_validation property
    message_validation MessageValidationConfigable
    // The min_tip_basis_points property
    min_tip_basis_points *int64
    // Only applies to V1 transactions
    min_tip_percentage *int32
    // The preparation_settings property
    preparation_settings PreparationSettingsable
    // The v1_transactions_allow_notary_to_duplicate_signer property
    v1_transactions_allow_notary_to_duplicate_signer *bool
    // The v2_transactions_allowed property
    v2_transactions_allowed *bool
}
// NewTransactionValidationConfig instantiates a new TransactionValidationConfig and sets the default values.
func NewTransactionValidationConfig()(*TransactionValidationConfig) {
    m := &TransactionValidationConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransactionValidationConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionValidationConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransactionValidationConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionValidationConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionValidationConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["manifest_validation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManifestValidationRuleset)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifestValidation(val.(*ManifestValidationRuleset))
        }
        return nil
    }
    res["max_epoch_range"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxEpochRange(val)
        }
        return nil
    }
    res["max_instructions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxInstructions(val)
        }
        return nil
    }
    res["max_references_per_intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxReferencesPerIntent(val)
        }
        return nil
    }
    res["max_signer_signatures_per_intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSignerSignaturesPerIntent(val)
        }
        return nil
    }
    res["max_subintent_depth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSubintentDepth(val)
        }
        return nil
    }
    res["max_tip_basis_points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTipBasisPoints(val)
        }
        return nil
    }
    res["max_tip_percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTipPercentage(val)
        }
        return nil
    }
    res["max_total_references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTotalReferences(val)
        }
        return nil
    }
    res["max_total_signature_validations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTotalSignatureValidations(val)
        }
        return nil
    }
    res["message_validation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageValidationConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageValidation(val.(MessageValidationConfigable))
        }
        return nil
    }
    res["min_tip_basis_points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinTipBasisPoints(val)
        }
        return nil
    }
    res["min_tip_percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinTipPercentage(val)
        }
        return nil
    }
    res["preparation_settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreparationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreparationSettings(val.(PreparationSettingsable))
        }
        return nil
    }
    res["v1_transactions_allow_notary_to_duplicate_signer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1TransactionsAllowNotaryToDuplicateSigner(val)
        }
        return nil
    }
    res["v2_transactions_allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV2TransactionsAllowed(val)
        }
        return nil
    }
    return res
}
// GetManifestValidation gets the manifest_validation property value. The manifest_validation property
// returns a *ManifestValidationRuleset when successful
func (m *TransactionValidationConfig) GetManifestValidation()(*ManifestValidationRuleset) {
    return m.manifest_validation
}
// GetMaxEpochRange gets the max_epoch_range property value. The max_epoch_range property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxEpochRange()(*string) {
    return m.max_epoch_range
}
// GetMaxInstructions gets the max_instructions property value. The max_instructions property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxInstructions()(*string) {
    return m.max_instructions
}
// GetMaxReferencesPerIntent gets the max_references_per_intent property value. The max_references_per_intent property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxReferencesPerIntent()(*string) {
    return m.max_references_per_intent
}
// GetMaxSignerSignaturesPerIntent gets the max_signer_signatures_per_intent property value. The max_signer_signatures_per_intent property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxSignerSignaturesPerIntent()(*string) {
    return m.max_signer_signatures_per_intent
}
// GetMaxSubintentDepth gets the max_subintent_depth property value. The max_subintent_depth property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxSubintentDepth()(*string) {
    return m.max_subintent_depth
}
// GetMaxTipBasisPoints gets the max_tip_basis_points property value. The max_tip_basis_points property
// returns a *int64 when successful
func (m *TransactionValidationConfig) GetMaxTipBasisPoints()(*int64) {
    return m.max_tip_basis_points
}
// GetMaxTipPercentage gets the max_tip_percentage property value. Only applies to V1 transactions
// returns a *int32 when successful
func (m *TransactionValidationConfig) GetMaxTipPercentage()(*int32) {
    return m.max_tip_percentage
}
// GetMaxTotalReferences gets the max_total_references property value. The max_total_references property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxTotalReferences()(*string) {
    return m.max_total_references
}
// GetMaxTotalSignatureValidations gets the max_total_signature_validations property value. The max_total_signature_validations property
// returns a *string when successful
func (m *TransactionValidationConfig) GetMaxTotalSignatureValidations()(*string) {
    return m.max_total_signature_validations
}
// GetMessageValidation gets the message_validation property value. The message_validation property
// returns a MessageValidationConfigable when successful
func (m *TransactionValidationConfig) GetMessageValidation()(MessageValidationConfigable) {
    return m.message_validation
}
// GetMinTipBasisPoints gets the min_tip_basis_points property value. The min_tip_basis_points property
// returns a *int64 when successful
func (m *TransactionValidationConfig) GetMinTipBasisPoints()(*int64) {
    return m.min_tip_basis_points
}
// GetMinTipPercentage gets the min_tip_percentage property value. Only applies to V1 transactions
// returns a *int32 when successful
func (m *TransactionValidationConfig) GetMinTipPercentage()(*int32) {
    return m.min_tip_percentage
}
// GetPreparationSettings gets the preparation_settings property value. The preparation_settings property
// returns a PreparationSettingsable when successful
func (m *TransactionValidationConfig) GetPreparationSettings()(PreparationSettingsable) {
    return m.preparation_settings
}
// GetV1TransactionsAllowNotaryToDuplicateSigner gets the v1_transactions_allow_notary_to_duplicate_signer property value. The v1_transactions_allow_notary_to_duplicate_signer property
// returns a *bool when successful
func (m *TransactionValidationConfig) GetV1TransactionsAllowNotaryToDuplicateSigner()(*bool) {
    return m.v1_transactions_allow_notary_to_duplicate_signer
}
// GetV2TransactionsAllowed gets the v2_transactions_allowed property value. The v2_transactions_allowed property
// returns a *bool when successful
func (m *TransactionValidationConfig) GetV2TransactionsAllowed()(*bool) {
    return m.v2_transactions_allowed
}
// Serialize serializes information the current object
func (m *TransactionValidationConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManifestValidation() != nil {
        cast := (*m.GetManifestValidation()).String()
        err := writer.WriteStringValue("manifest_validation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_epoch_range", m.GetMaxEpochRange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_instructions", m.GetMaxInstructions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_references_per_intent", m.GetMaxReferencesPerIntent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_signer_signatures_per_intent", m.GetMaxSignerSignaturesPerIntent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_subintent_depth", m.GetMaxSubintentDepth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("max_tip_basis_points", m.GetMaxTipBasisPoints())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("max_tip_percentage", m.GetMaxTipPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_total_references", m.GetMaxTotalReferences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("max_total_signature_validations", m.GetMaxTotalSignatureValidations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message_validation", m.GetMessageValidation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("min_tip_basis_points", m.GetMinTipBasisPoints())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("min_tip_percentage", m.GetMinTipPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("preparation_settings", m.GetPreparationSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v1_transactions_allow_notary_to_duplicate_signer", m.GetV1TransactionsAllowNotaryToDuplicateSigner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v2_transactions_allowed", m.GetV2TransactionsAllowed())
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
func (m *TransactionValidationConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManifestValidation sets the manifest_validation property value. The manifest_validation property
func (m *TransactionValidationConfig) SetManifestValidation(value *ManifestValidationRuleset)() {
    m.manifest_validation = value
}
// SetMaxEpochRange sets the max_epoch_range property value. The max_epoch_range property
func (m *TransactionValidationConfig) SetMaxEpochRange(value *string)() {
    m.max_epoch_range = value
}
// SetMaxInstructions sets the max_instructions property value. The max_instructions property
func (m *TransactionValidationConfig) SetMaxInstructions(value *string)() {
    m.max_instructions = value
}
// SetMaxReferencesPerIntent sets the max_references_per_intent property value. The max_references_per_intent property
func (m *TransactionValidationConfig) SetMaxReferencesPerIntent(value *string)() {
    m.max_references_per_intent = value
}
// SetMaxSignerSignaturesPerIntent sets the max_signer_signatures_per_intent property value. The max_signer_signatures_per_intent property
func (m *TransactionValidationConfig) SetMaxSignerSignaturesPerIntent(value *string)() {
    m.max_signer_signatures_per_intent = value
}
// SetMaxSubintentDepth sets the max_subintent_depth property value. The max_subintent_depth property
func (m *TransactionValidationConfig) SetMaxSubintentDepth(value *string)() {
    m.max_subintent_depth = value
}
// SetMaxTipBasisPoints sets the max_tip_basis_points property value. The max_tip_basis_points property
func (m *TransactionValidationConfig) SetMaxTipBasisPoints(value *int64)() {
    m.max_tip_basis_points = value
}
// SetMaxTipPercentage sets the max_tip_percentage property value. Only applies to V1 transactions
func (m *TransactionValidationConfig) SetMaxTipPercentage(value *int32)() {
    m.max_tip_percentage = value
}
// SetMaxTotalReferences sets the max_total_references property value. The max_total_references property
func (m *TransactionValidationConfig) SetMaxTotalReferences(value *string)() {
    m.max_total_references = value
}
// SetMaxTotalSignatureValidations sets the max_total_signature_validations property value. The max_total_signature_validations property
func (m *TransactionValidationConfig) SetMaxTotalSignatureValidations(value *string)() {
    m.max_total_signature_validations = value
}
// SetMessageValidation sets the message_validation property value. The message_validation property
func (m *TransactionValidationConfig) SetMessageValidation(value MessageValidationConfigable)() {
    m.message_validation = value
}
// SetMinTipBasisPoints sets the min_tip_basis_points property value. The min_tip_basis_points property
func (m *TransactionValidationConfig) SetMinTipBasisPoints(value *int64)() {
    m.min_tip_basis_points = value
}
// SetMinTipPercentage sets the min_tip_percentage property value. Only applies to V1 transactions
func (m *TransactionValidationConfig) SetMinTipPercentage(value *int32)() {
    m.min_tip_percentage = value
}
// SetPreparationSettings sets the preparation_settings property value. The preparation_settings property
func (m *TransactionValidationConfig) SetPreparationSettings(value PreparationSettingsable)() {
    m.preparation_settings = value
}
// SetV1TransactionsAllowNotaryToDuplicateSigner sets the v1_transactions_allow_notary_to_duplicate_signer property value. The v1_transactions_allow_notary_to_duplicate_signer property
func (m *TransactionValidationConfig) SetV1TransactionsAllowNotaryToDuplicateSigner(value *bool)() {
    m.v1_transactions_allow_notary_to_duplicate_signer = value
}
// SetV2TransactionsAllowed sets the v2_transactions_allowed property value. The v2_transactions_allowed property
func (m *TransactionValidationConfig) SetV2TransactionsAllowed(value *bool)() {
    m.v2_transactions_allowed = value
}
type TransactionValidationConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifestValidation()(*ManifestValidationRuleset)
    GetMaxEpochRange()(*string)
    GetMaxInstructions()(*string)
    GetMaxReferencesPerIntent()(*string)
    GetMaxSignerSignaturesPerIntent()(*string)
    GetMaxSubintentDepth()(*string)
    GetMaxTipBasisPoints()(*int64)
    GetMaxTipPercentage()(*int32)
    GetMaxTotalReferences()(*string)
    GetMaxTotalSignatureValidations()(*string)
    GetMessageValidation()(MessageValidationConfigable)
    GetMinTipBasisPoints()(*int64)
    GetMinTipPercentage()(*int32)
    GetPreparationSettings()(PreparationSettingsable)
    GetV1TransactionsAllowNotaryToDuplicateSigner()(*bool)
    GetV2TransactionsAllowed()(*bool)
    SetManifestValidation(value *ManifestValidationRuleset)()
    SetMaxEpochRange(value *string)()
    SetMaxInstructions(value *string)()
    SetMaxReferencesPerIntent(value *string)()
    SetMaxSignerSignaturesPerIntent(value *string)()
    SetMaxSubintentDepth(value *string)()
    SetMaxTipBasisPoints(value *int64)()
    SetMaxTipPercentage(value *int32)()
    SetMaxTotalReferences(value *string)()
    SetMaxTotalSignatureValidations(value *string)()
    SetMessageValidation(value MessageValidationConfigable)()
    SetMinTipBasisPoints(value *int64)()
    SetMinTipPercentage(value *int32)()
    SetPreparationSettings(value PreparationSettingsable)()
    SetV1TransactionsAllowNotaryToDuplicateSigner(value *bool)()
    SetV2TransactionsAllowed(value *bool)()
}
