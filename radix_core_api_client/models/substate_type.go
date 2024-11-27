package models
type SubstateType int

const (
    BOOTLOADERMODULEFIELDSYSTEMBOOT_SUBSTATETYPE SubstateType = iota
    BOOTLOADERMODULEFIELDKERNELBOOT_SUBSTATETYPE
    BOOTLOADERMODULEFIELDVMBOOT_SUBSTATETYPE
    BOOTLOADERMODULEFIELDTRANSACTIONVALIDATIONCONFIGURATION_SUBSTATETYPE
    PROTOCOLUPDATESTATUSMODULEFIELDSUMMARY_SUBSTATETYPE
    TYPEINFOMODULEFIELDTYPEINFO_SUBSTATETYPE
    ROLEASSIGNMENTMODULEFIELDOWNERROLE_SUBSTATETYPE
    ROLEASSIGNMENTMODULERULEENTRY_SUBSTATETYPE
    ROLEASSIGNMENTMODULEMUTABILITYENTRY_SUBSTATETYPE
    ROYALTYMODULEFIELDSTATE_SUBSTATETYPE
    ROYALTYMODULEMETHODROYALTYENTRY_SUBSTATETYPE
    METADATAMODULEENTRY_SUBSTATETYPE
    PACKAGEFIELDROYALTYACCUMULATOR_SUBSTATETYPE
    PACKAGECODEVMTYPEENTRY_SUBSTATETYPE
    PACKAGECODEORIGINALCODEENTRY_SUBSTATETYPE
    PACKAGECODEINSTRUMENTEDCODEENTRY_SUBSTATETYPE
    SCHEMAENTRY_SUBSTATETYPE
    PACKAGEBLUEPRINTDEFINITIONENTRY_SUBSTATETYPE
    PACKAGEBLUEPRINTDEPENDENCIESENTRY_SUBSTATETYPE
    PACKAGEBLUEPRINTROYALTYENTRY_SUBSTATETYPE
    PACKAGEBLUEPRINTAUTHTEMPLATEENTRY_SUBSTATETYPE
    PACKAGEFIELDFUNCTIONACCESSRULES_SUBSTATETYPE
    FUNGIBLERESOURCEMANAGERFIELDDIVISIBILITY_SUBSTATETYPE
    FUNGIBLERESOURCEMANAGERFIELDTOTALSUPPLY_SUBSTATETYPE
    NONFUNGIBLERESOURCEMANAGERFIELDIDTYPE_SUBSTATETYPE
    NONFUNGIBLERESOURCEMANAGERFIELDTOTALSUPPLY_SUBSTATETYPE
    NONFUNGIBLERESOURCEMANAGERFIELDMUTABLEFIELDS_SUBSTATETYPE
    NONFUNGIBLERESOURCEMANAGERDATAENTRY_SUBSTATETYPE
    FUNGIBLEVAULTFIELDBALANCE_SUBSTATETYPE
    FUNGIBLEVAULTFIELDFROZENSTATUS_SUBSTATETYPE
    NONFUNGIBLEVAULTFIELDBALANCE_SUBSTATETYPE
    NONFUNGIBLEVAULTFIELDFROZENSTATUS_SUBSTATETYPE
    NONFUNGIBLEVAULTCONTENTSINDEXENTRY_SUBSTATETYPE
    CONSENSUSMANAGERFIELDCONFIG_SUBSTATETYPE
    CONSENSUSMANAGERFIELDSTATE_SUBSTATETYPE
    CONSENSUSMANAGERFIELDCURRENTVALIDATORSET_SUBSTATETYPE
    CONSENSUSMANAGERFIELDCURRENTPROPOSALSTATISTIC_SUBSTATETYPE
    CONSENSUSMANAGERFIELDCURRENTTIMEROUNDEDTOMINUTES_SUBSTATETYPE
    CONSENSUSMANAGERFIELDCURRENTTIME_SUBSTATETYPE
    CONSENSUSMANAGERFIELDVALIDATORREWARDS_SUBSTATETYPE
    CONSENSUSMANAGERREGISTEREDVALIDATORSBYSTAKEINDEXENTRY_SUBSTATETYPE
    VALIDATORFIELDSTATE_SUBSTATETYPE
    VALIDATORFIELDPROTOCOLUPDATEREADINESSSIGNAL_SUBSTATETYPE
    ACCOUNTFIELDSTATE_SUBSTATETYPE
    ACCOUNTVAULTENTRY_SUBSTATETYPE
    ACCOUNTRESOURCEPREFERENCEENTRY_SUBSTATETYPE
    ACCOUNTAUTHORIZEDDEPOSITORENTRY_SUBSTATETYPE
    ACCESSCONTROLLERFIELDSTATE_SUBSTATETYPE
    GENERICSCRYPTOCOMPONENTFIELDSTATE_SUBSTATETYPE
    GENERICKEYVALUESTOREENTRY_SUBSTATETYPE
    ONERESOURCEPOOLFIELDSTATE_SUBSTATETYPE
    TWORESOURCEPOOLFIELDSTATE_SUBSTATETYPE
    MULTIRESOURCEPOOLFIELDSTATE_SUBSTATETYPE
    TRANSACTIONTRACKERFIELDSTATE_SUBSTATETYPE
    TRANSACTIONTRACKERCOLLECTIONENTRY_SUBSTATETYPE
    ACCOUNTLOCKERACCOUNTCLAIMSENTRY_SUBSTATETYPE
)

func (i SubstateType) String() string {
    return []string{"BootLoaderModuleFieldSystemBoot", "BootLoaderModuleFieldKernelBoot", "BootLoaderModuleFieldVmBoot", "BootLoaderModuleFieldTransactionValidationConfiguration", "ProtocolUpdateStatusModuleFieldSummary", "TypeInfoModuleFieldTypeInfo", "RoleAssignmentModuleFieldOwnerRole", "RoleAssignmentModuleRuleEntry", "RoleAssignmentModuleMutabilityEntry", "RoyaltyModuleFieldState", "RoyaltyModuleMethodRoyaltyEntry", "MetadataModuleEntry", "PackageFieldRoyaltyAccumulator", "PackageCodeVmTypeEntry", "PackageCodeOriginalCodeEntry", "PackageCodeInstrumentedCodeEntry", "SchemaEntry", "PackageBlueprintDefinitionEntry", "PackageBlueprintDependenciesEntry", "PackageBlueprintRoyaltyEntry", "PackageBlueprintAuthTemplateEntry", "PackageFieldFunctionAccessRules", "FungibleResourceManagerFieldDivisibility", "FungibleResourceManagerFieldTotalSupply", "NonFungibleResourceManagerFieldIdType", "NonFungibleResourceManagerFieldTotalSupply", "NonFungibleResourceManagerFieldMutableFields", "NonFungibleResourceManagerDataEntry", "FungibleVaultFieldBalance", "FungibleVaultFieldFrozenStatus", "NonFungibleVaultFieldBalance", "NonFungibleVaultFieldFrozenStatus", "NonFungibleVaultContentsIndexEntry", "ConsensusManagerFieldConfig", "ConsensusManagerFieldState", "ConsensusManagerFieldCurrentValidatorSet", "ConsensusManagerFieldCurrentProposalStatistic", "ConsensusManagerFieldCurrentTimeRoundedToMinutes", "ConsensusManagerFieldCurrentTime", "ConsensusManagerFieldValidatorRewards", "ConsensusManagerRegisteredValidatorsByStakeIndexEntry", "ValidatorFieldState", "ValidatorFieldProtocolUpdateReadinessSignal", "AccountFieldState", "AccountVaultEntry", "AccountResourcePreferenceEntry", "AccountAuthorizedDepositorEntry", "AccessControllerFieldState", "GenericScryptoComponentFieldState", "GenericKeyValueStoreEntry", "OneResourcePoolFieldState", "TwoResourcePoolFieldState", "MultiResourcePoolFieldState", "TransactionTrackerFieldState", "TransactionTrackerCollectionEntry", "AccountLockerAccountClaimsEntry"}[i]
}
func ParseSubstateType(v string) (any, error) {
    result := BOOTLOADERMODULEFIELDSYSTEMBOOT_SUBSTATETYPE
    switch v {
        case "BootLoaderModuleFieldSystemBoot":
            result = BOOTLOADERMODULEFIELDSYSTEMBOOT_SUBSTATETYPE
        case "BootLoaderModuleFieldKernelBoot":
            result = BOOTLOADERMODULEFIELDKERNELBOOT_SUBSTATETYPE
        case "BootLoaderModuleFieldVmBoot":
            result = BOOTLOADERMODULEFIELDVMBOOT_SUBSTATETYPE
        case "BootLoaderModuleFieldTransactionValidationConfiguration":
            result = BOOTLOADERMODULEFIELDTRANSACTIONVALIDATIONCONFIGURATION_SUBSTATETYPE
        case "ProtocolUpdateStatusModuleFieldSummary":
            result = PROTOCOLUPDATESTATUSMODULEFIELDSUMMARY_SUBSTATETYPE
        case "TypeInfoModuleFieldTypeInfo":
            result = TYPEINFOMODULEFIELDTYPEINFO_SUBSTATETYPE
        case "RoleAssignmentModuleFieldOwnerRole":
            result = ROLEASSIGNMENTMODULEFIELDOWNERROLE_SUBSTATETYPE
        case "RoleAssignmentModuleRuleEntry":
            result = ROLEASSIGNMENTMODULERULEENTRY_SUBSTATETYPE
        case "RoleAssignmentModuleMutabilityEntry":
            result = ROLEASSIGNMENTMODULEMUTABILITYENTRY_SUBSTATETYPE
        case "RoyaltyModuleFieldState":
            result = ROYALTYMODULEFIELDSTATE_SUBSTATETYPE
        case "RoyaltyModuleMethodRoyaltyEntry":
            result = ROYALTYMODULEMETHODROYALTYENTRY_SUBSTATETYPE
        case "MetadataModuleEntry":
            result = METADATAMODULEENTRY_SUBSTATETYPE
        case "PackageFieldRoyaltyAccumulator":
            result = PACKAGEFIELDROYALTYACCUMULATOR_SUBSTATETYPE
        case "PackageCodeVmTypeEntry":
            result = PACKAGECODEVMTYPEENTRY_SUBSTATETYPE
        case "PackageCodeOriginalCodeEntry":
            result = PACKAGECODEORIGINALCODEENTRY_SUBSTATETYPE
        case "PackageCodeInstrumentedCodeEntry":
            result = PACKAGECODEINSTRUMENTEDCODEENTRY_SUBSTATETYPE
        case "SchemaEntry":
            result = SCHEMAENTRY_SUBSTATETYPE
        case "PackageBlueprintDefinitionEntry":
            result = PACKAGEBLUEPRINTDEFINITIONENTRY_SUBSTATETYPE
        case "PackageBlueprintDependenciesEntry":
            result = PACKAGEBLUEPRINTDEPENDENCIESENTRY_SUBSTATETYPE
        case "PackageBlueprintRoyaltyEntry":
            result = PACKAGEBLUEPRINTROYALTYENTRY_SUBSTATETYPE
        case "PackageBlueprintAuthTemplateEntry":
            result = PACKAGEBLUEPRINTAUTHTEMPLATEENTRY_SUBSTATETYPE
        case "PackageFieldFunctionAccessRules":
            result = PACKAGEFIELDFUNCTIONACCESSRULES_SUBSTATETYPE
        case "FungibleResourceManagerFieldDivisibility":
            result = FUNGIBLERESOURCEMANAGERFIELDDIVISIBILITY_SUBSTATETYPE
        case "FungibleResourceManagerFieldTotalSupply":
            result = FUNGIBLERESOURCEMANAGERFIELDTOTALSUPPLY_SUBSTATETYPE
        case "NonFungibleResourceManagerFieldIdType":
            result = NONFUNGIBLERESOURCEMANAGERFIELDIDTYPE_SUBSTATETYPE
        case "NonFungibleResourceManagerFieldTotalSupply":
            result = NONFUNGIBLERESOURCEMANAGERFIELDTOTALSUPPLY_SUBSTATETYPE
        case "NonFungibleResourceManagerFieldMutableFields":
            result = NONFUNGIBLERESOURCEMANAGERFIELDMUTABLEFIELDS_SUBSTATETYPE
        case "NonFungibleResourceManagerDataEntry":
            result = NONFUNGIBLERESOURCEMANAGERDATAENTRY_SUBSTATETYPE
        case "FungibleVaultFieldBalance":
            result = FUNGIBLEVAULTFIELDBALANCE_SUBSTATETYPE
        case "FungibleVaultFieldFrozenStatus":
            result = FUNGIBLEVAULTFIELDFROZENSTATUS_SUBSTATETYPE
        case "NonFungibleVaultFieldBalance":
            result = NONFUNGIBLEVAULTFIELDBALANCE_SUBSTATETYPE
        case "NonFungibleVaultFieldFrozenStatus":
            result = NONFUNGIBLEVAULTFIELDFROZENSTATUS_SUBSTATETYPE
        case "NonFungibleVaultContentsIndexEntry":
            result = NONFUNGIBLEVAULTCONTENTSINDEXENTRY_SUBSTATETYPE
        case "ConsensusManagerFieldConfig":
            result = CONSENSUSMANAGERFIELDCONFIG_SUBSTATETYPE
        case "ConsensusManagerFieldState":
            result = CONSENSUSMANAGERFIELDSTATE_SUBSTATETYPE
        case "ConsensusManagerFieldCurrentValidatorSet":
            result = CONSENSUSMANAGERFIELDCURRENTVALIDATORSET_SUBSTATETYPE
        case "ConsensusManagerFieldCurrentProposalStatistic":
            result = CONSENSUSMANAGERFIELDCURRENTPROPOSALSTATISTIC_SUBSTATETYPE
        case "ConsensusManagerFieldCurrentTimeRoundedToMinutes":
            result = CONSENSUSMANAGERFIELDCURRENTTIMEROUNDEDTOMINUTES_SUBSTATETYPE
        case "ConsensusManagerFieldCurrentTime":
            result = CONSENSUSMANAGERFIELDCURRENTTIME_SUBSTATETYPE
        case "ConsensusManagerFieldValidatorRewards":
            result = CONSENSUSMANAGERFIELDVALIDATORREWARDS_SUBSTATETYPE
        case "ConsensusManagerRegisteredValidatorsByStakeIndexEntry":
            result = CONSENSUSMANAGERREGISTEREDVALIDATORSBYSTAKEINDEXENTRY_SUBSTATETYPE
        case "ValidatorFieldState":
            result = VALIDATORFIELDSTATE_SUBSTATETYPE
        case "ValidatorFieldProtocolUpdateReadinessSignal":
            result = VALIDATORFIELDPROTOCOLUPDATEREADINESSSIGNAL_SUBSTATETYPE
        case "AccountFieldState":
            result = ACCOUNTFIELDSTATE_SUBSTATETYPE
        case "AccountVaultEntry":
            result = ACCOUNTVAULTENTRY_SUBSTATETYPE
        case "AccountResourcePreferenceEntry":
            result = ACCOUNTRESOURCEPREFERENCEENTRY_SUBSTATETYPE
        case "AccountAuthorizedDepositorEntry":
            result = ACCOUNTAUTHORIZEDDEPOSITORENTRY_SUBSTATETYPE
        case "AccessControllerFieldState":
            result = ACCESSCONTROLLERFIELDSTATE_SUBSTATETYPE
        case "GenericScryptoComponentFieldState":
            result = GENERICSCRYPTOCOMPONENTFIELDSTATE_SUBSTATETYPE
        case "GenericKeyValueStoreEntry":
            result = GENERICKEYVALUESTOREENTRY_SUBSTATETYPE
        case "OneResourcePoolFieldState":
            result = ONERESOURCEPOOLFIELDSTATE_SUBSTATETYPE
        case "TwoResourcePoolFieldState":
            result = TWORESOURCEPOOLFIELDSTATE_SUBSTATETYPE
        case "MultiResourcePoolFieldState":
            result = MULTIRESOURCEPOOLFIELDSTATE_SUBSTATETYPE
        case "TransactionTrackerFieldState":
            result = TRANSACTIONTRACKERFIELDSTATE_SUBSTATETYPE
        case "TransactionTrackerCollectionEntry":
            result = TRANSACTIONTRACKERCOLLECTIONENTRY_SUBSTATETYPE
        case "AccountLockerAccountClaimsEntry":
            result = ACCOUNTLOCKERACCOUNTCLAIMSENTRY_SUBSTATETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSubstateType(values []SubstateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SubstateType) isMultiValue() bool {
    return false
}
