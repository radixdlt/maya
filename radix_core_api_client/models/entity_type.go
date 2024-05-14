package models

import (
	"errors"
)

type EntityType int

const (
	GLOBALPACKAGE_ENTITYTYPE EntityType = iota
	GLOBALCONSENSUSMANAGER_ENTITYTYPE
	GLOBALVALIDATOR_ENTITYTYPE
	GLOBALGENERICCOMPONENT_ENTITYTYPE
	GLOBALACCOUNT_ENTITYTYPE
	GLOBALACCOUNTLOCKER_ENTITYTYPE
	GLOBALIDENTITY_ENTITYTYPE
	GLOBALACCESSCONTROLLER_ENTITYTYPE
	GLOBALVIRTUALSECP256K1ACCOUNT_ENTITYTYPE
	GLOBALVIRTUALSECP256K1IDENTITY_ENTITYTYPE
	GLOBALVIRTUALED25519ACCOUNT_ENTITYTYPE
	GLOBALVIRTUALED25519IDENTITY_ENTITYTYPE
	GLOBALFUNGIBLERESOURCE_ENTITYTYPE
	INTERNALFUNGIBLEVAULT_ENTITYTYPE
	GLOBALNONFUNGIBLERESOURCE_ENTITYTYPE
	INTERNALNONFUNGIBLEVAULT_ENTITYTYPE
	INTERNALGENERICCOMPONENT_ENTITYTYPE
	INTERNALKEYVALUESTORE_ENTITYTYPE
	GLOBALONERESOURCEPOOL_ENTITYTYPE
	GLOBALTWORESOURCEPOOL_ENTITYTYPE
	GLOBALMULTIRESOURCEPOOL_ENTITYTYPE
	GLOBALTRANSACTIONTRACKER_ENTITYTYPE
)

func (i EntityType) String() string {
	return []string{"GlobalPackage", "GlobalConsensusManager", "GlobalValidator", "GlobalGenericComponent", "GlobalAccount", "GlobalAccountLocker", "GlobalIdentity", "GlobalAccessController", "GlobalVirtualSecp256k1Account", "GlobalVirtualSecp256k1Identity", "GlobalVirtualEd25519Account", "GlobalVirtualEd25519Identity", "GlobalFungibleResource", "InternalFungibleVault", "GlobalNonFungibleResource", "InternalNonFungibleVault", "InternalGenericComponent", "InternalKeyValueStore", "GlobalOneResourcePool", "GlobalTwoResourcePool", "GlobalMultiResourcePool", "GlobalTransactionTracker"}[i]
}

func ParseEntityType(v string) (any, error) {
	result := GLOBALPACKAGE_ENTITYTYPE
	switch v {
	case "GlobalPackage":
		result = GLOBALPACKAGE_ENTITYTYPE
	case "GlobalConsensusManager":
		result = GLOBALCONSENSUSMANAGER_ENTITYTYPE
	case "GlobalValidator":
		result = GLOBALVALIDATOR_ENTITYTYPE
	case "GlobalGenericComponent":
		result = GLOBALGENERICCOMPONENT_ENTITYTYPE
	case "GlobalAccount":
		result = GLOBALACCOUNT_ENTITYTYPE
	case "GlobalAccountLocker":
		result = GLOBALACCOUNTLOCKER_ENTITYTYPE
	case "GlobalIdentity":
		result = GLOBALIDENTITY_ENTITYTYPE
	case "GlobalAccessController":
		result = GLOBALACCESSCONTROLLER_ENTITYTYPE
	case "GlobalVirtualSecp256k1Account":
		result = GLOBALVIRTUALSECP256K1ACCOUNT_ENTITYTYPE
	case "GlobalVirtualSecp256k1Identity":
		result = GLOBALVIRTUALSECP256K1IDENTITY_ENTITYTYPE
	case "GlobalVirtualEd25519Account":
		result = GLOBALVIRTUALED25519ACCOUNT_ENTITYTYPE
	case "GlobalVirtualEd25519Identity":
		result = GLOBALVIRTUALED25519IDENTITY_ENTITYTYPE
	case "GlobalFungibleResource":
		result = GLOBALFUNGIBLERESOURCE_ENTITYTYPE
	case "InternalFungibleVault":
		result = INTERNALFUNGIBLEVAULT_ENTITYTYPE
	case "GlobalNonFungibleResource":
		result = GLOBALNONFUNGIBLERESOURCE_ENTITYTYPE
	case "InternalNonFungibleVault":
		result = INTERNALNONFUNGIBLEVAULT_ENTITYTYPE
	case "InternalGenericComponent":
		result = INTERNALGENERICCOMPONENT_ENTITYTYPE
	case "InternalKeyValueStore":
		result = INTERNALKEYVALUESTORE_ENTITYTYPE
	case "GlobalOneResourcePool":
		result = GLOBALONERESOURCEPOOL_ENTITYTYPE
	case "GlobalTwoResourcePool":
		result = GLOBALTWORESOURCEPOOL_ENTITYTYPE
	case "GlobalMultiResourcePool":
		result = GLOBALMULTIRESOURCEPOOL_ENTITYTYPE
	case "GlobalTransactionTracker":
		result = GLOBALTRANSACTIONTRACKER_ENTITYTYPE
	default:
		return 0, errors.New("Unknown EntityType value: " + v)
	}
	return &result, nil
}

func SerializeEntityType(values []EntityType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i EntityType) isMultiValue() bool {
	return false
}
