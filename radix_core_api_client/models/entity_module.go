package models

import (
	"errors"
)

type EntityModule int

const (
	TYPEINFO_ENTITYMODULE EntityModule = iota
	METADATA_ENTITYMODULE
	ROLEASSIGNMENT_ENTITYMODULE
	ROYALTY_ENTITYMODULE
	MAIN_ENTITYMODULE
	SCHEMA_ENTITYMODULE
	BOOTLOADER_ENTITYMODULE
)

func (i EntityModule) String() string {
	return []string{"TypeInfo", "Metadata", "RoleAssignment", "Royalty", "Main", "Schema", "BootLoader"}[i]
}

func ParseEntityModule(v string) (any, error) {
	result := TYPEINFO_ENTITYMODULE
	switch v {
	case "TypeInfo":
		result = TYPEINFO_ENTITYMODULE
	case "Metadata":
		result = METADATA_ENTITYMODULE
	case "RoleAssignment":
		result = ROLEASSIGNMENT_ENTITYMODULE
	case "Royalty":
		result = ROYALTY_ENTITYMODULE
	case "Main":
		result = MAIN_ENTITYMODULE
	case "Schema":
		result = SCHEMA_ENTITYMODULE
	case "BootLoader":
		result = BOOTLOADER_ENTITYMODULE
	default:
		return 0, errors.New("Unknown EntityModule value: " + v)
	}
	return &result, nil
}

func SerializeEntityModule(values []EntityModule) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i EntityModule) isMultiValue() bool {
	return false
}
