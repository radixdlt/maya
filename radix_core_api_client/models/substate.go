package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Substate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The is_locked property
    is_locked *bool
    // The substate_type property
    substate_type *SubstateType
}
// NewSubstate instantiates a new Substate and sets the default values.
func NewSubstate()(*Substate) {
    m := &Substate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("substate_type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "AccessControllerFieldState":
                        return NewAccessControllerFieldStateSubstate(), nil
                    case "AccountAuthorizedDepositorEntry":
                        return NewAccountAuthorizedDepositorEntrySubstate(), nil
                    case "AccountFieldState":
                        return NewAccountFieldStateSubstate(), nil
                    case "AccountLockerAccountClaimsEntry":
                        return NewAccountLockerAccountClaimsEntrySubstate(), nil
                    case "AccountResourcePreferenceEntry":
                        return NewAccountResourcePreferenceEntrySubstate(), nil
                    case "AccountVaultEntry":
                        return NewAccountVaultEntrySubstate(), nil
                    case "BootLoaderModuleFieldKernelBoot":
                        return NewBootLoaderModuleFieldKernelBootSubstate(), nil
                    case "BootLoaderModuleFieldSystemBoot":
                        return NewBootLoaderModuleFieldSystemBootSubstate(), nil
                    case "BootLoaderModuleFieldTransactionValidationConfiguration":
                        return NewBootLoaderModuleFieldTransactionValidationConfigurationSubstate(), nil
                    case "BootLoaderModuleFieldVmBoot":
                        return NewBootLoaderModuleFieldVmBootSubstate(), nil
                    case "ConsensusManagerFieldConfig":
                        return NewConsensusManagerFieldConfigSubstate(), nil
                    case "ConsensusManagerFieldCurrentProposalStatistic":
                        return NewConsensusManagerFieldCurrentProposalStatisticSubstate(), nil
                    case "ConsensusManagerFieldCurrentTime":
                        return NewConsensusManagerFieldCurrentTimeSubstate(), nil
                    case "ConsensusManagerFieldCurrentTimeRoundedToMinutes":
                        return NewConsensusManagerFieldCurrentTimeRoundedToMinutesSubstate(), nil
                    case "ConsensusManagerFieldCurrentValidatorSet":
                        return NewConsensusManagerFieldCurrentValidatorSetSubstate(), nil
                    case "ConsensusManagerFieldState":
                        return NewConsensusManagerFieldStateSubstate(), nil
                    case "ConsensusManagerFieldValidatorRewards":
                        return NewConsensusManagerFieldValidatorRewardsSubstate(), nil
                    case "ConsensusManagerRegisteredValidatorsByStakeIndexEntry":
                        return NewConsensusManagerRegisteredValidatorsByStakeIndexEntrySubstate(), nil
                    case "FungibleResourceManagerFieldDivisibility":
                        return NewFungibleResourceManagerFieldDivisibilitySubstate(), nil
                    case "FungibleResourceManagerFieldTotalSupply":
                        return NewFungibleResourceManagerFieldTotalSupplySubstate(), nil
                    case "FungibleVaultFieldBalance":
                        return NewFungibleVaultFieldBalanceSubstate(), nil
                    case "FungibleVaultFieldFrozenStatus":
                        return NewFungibleVaultFieldFrozenStatusSubstate(), nil
                    case "GenericKeyValueStoreEntry":
                        return NewGenericKeyValueStoreEntrySubstate(), nil
                    case "GenericScryptoComponentFieldState":
                        return NewGenericScryptoComponentFieldStateSubstate(), nil
                    case "MetadataModuleEntry":
                        return NewMetadataModuleEntrySubstate(), nil
                    case "MultiResourcePoolFieldState":
                        return NewMultiResourcePoolFieldStateSubstate(), nil
                    case "NonFungibleResourceManagerDataEntry":
                        return NewNonFungibleResourceManagerDataEntrySubstate(), nil
                    case "NonFungibleResourceManagerFieldIdType":
                        return NewNonFungibleResourceManagerFieldIdTypeSubstate(), nil
                    case "NonFungibleResourceManagerFieldMutableFields":
                        return NewNonFungibleResourceManagerFieldMutableFieldsSubstate(), nil
                    case "NonFungibleResourceManagerFieldTotalSupply":
                        return NewNonFungibleResourceManagerFieldTotalSupplySubstate(), nil
                    case "NonFungibleVaultContentsIndexEntry":
                        return NewNonFungibleVaultContentsIndexEntrySubstate(), nil
                    case "NonFungibleVaultFieldBalance":
                        return NewNonFungibleVaultFieldBalanceSubstate(), nil
                    case "NonFungibleVaultFieldFrozenStatus":
                        return NewNonFungibleVaultFieldFrozenStatusSubstate(), nil
                    case "OneResourcePoolFieldState":
                        return NewOneResourcePoolFieldStateSubstate(), nil
                    case "PackageBlueprintAuthTemplateEntry":
                        return NewPackageBlueprintAuthTemplateEntrySubstate(), nil
                    case "PackageBlueprintDefinitionEntry":
                        return NewPackageBlueprintDefinitionEntrySubstate(), nil
                    case "PackageBlueprintDependenciesEntry":
                        return NewPackageBlueprintDependenciesEntrySubstate(), nil
                    case "PackageBlueprintRoyaltyEntry":
                        return NewPackageBlueprintRoyaltyEntrySubstate(), nil
                    case "PackageCodeInstrumentedCodeEntry":
                        return NewPackageCodeInstrumentedCodeEntrySubstate(), nil
                    case "PackageCodeOriginalCodeEntry":
                        return NewPackageCodeOriginalCodeEntrySubstate(), nil
                    case "PackageCodeVmTypeEntry":
                        return NewPackageCodeVmTypeEntrySubstate(), nil
                    case "PackageFieldRoyaltyAccumulator":
                        return NewPackageFieldRoyaltyAccumulatorSubstate(), nil
                    case "ProtocolUpdateStatusModuleFieldSummary":
                        return NewProtocolUpdateStatusModuleFieldSummarySubstate(), nil
                    case "RoleAssignmentModuleFieldOwnerRole":
                        return NewRoleAssignmentModuleFieldOwnerRoleSubstate(), nil
                    case "RoleAssignmentModuleRuleEntry":
                        return NewRoleAssignmentModuleRuleEntrySubstate(), nil
                    case "RoyaltyModuleFieldState":
                        return NewRoyaltyModuleFieldStateSubstate(), nil
                    case "RoyaltyModuleMethodRoyaltyEntry":
                        return NewRoyaltyModuleMethodRoyaltyEntrySubstate(), nil
                    case "SchemaEntry":
                        return NewSchemaEntrySubstate(), nil
                    case "TransactionTrackerCollectionEntry":
                        return NewTransactionTrackerCollectionEntrySubstate(), nil
                    case "TransactionTrackerFieldState":
                        return NewTransactionTrackerFieldStateSubstate(), nil
                    case "TwoResourcePoolFieldState":
                        return NewTwoResourcePoolFieldStateSubstate(), nil
                    case "TypeInfoModuleFieldTypeInfo":
                        return NewTypeInfoModuleFieldTypeInfoSubstate(), nil
                    case "ValidatorFieldProtocolUpdateReadinessSignal":
                        return NewValidatorFieldProtocolUpdateReadinessSignalSubstate(), nil
                    case "ValidatorFieldState":
                        return NewValidatorFieldStateSubstate(), nil
                }
            }
        }
    }
    return NewSubstate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Substate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Substate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_locked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocked(val)
        }
        return nil
    }
    res["substate_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubstateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubstateType(val.(*SubstateType))
        }
        return nil
    }
    return res
}
// GetIsLocked gets the is_locked property value. The is_locked property
// returns a *bool when successful
func (m *Substate) GetIsLocked()(*bool) {
    return m.is_locked
}
// GetSubstateType gets the substate_type property value. The substate_type property
// returns a *SubstateType when successful
func (m *Substate) GetSubstateType()(*SubstateType) {
    return m.substate_type
}
// Serialize serializes information the current object
func (m *Substate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_locked", m.GetIsLocked())
        if err != nil {
            return err
        }
    }
    if m.GetSubstateType() != nil {
        cast := (*m.GetSubstateType()).String()
        err := writer.WriteStringValue("substate_type", &cast)
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
func (m *Substate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsLocked sets the is_locked property value. The is_locked property
func (m *Substate) SetIsLocked(value *bool)() {
    m.is_locked = value
}
// SetSubstateType sets the substate_type property value. The substate_type property
func (m *Substate) SetSubstateType(value *SubstateType)() {
    m.substate_type = value
}
type Substateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsLocked()(*bool)
    GetSubstateType()(*SubstateType)
    SetIsLocked(value *bool)()
    SetSubstateType(value *SubstateType)()
}
