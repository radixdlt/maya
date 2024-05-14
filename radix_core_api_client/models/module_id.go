package models

import (
	"errors"
)

type ModuleId int

const (
	MAIN_MODULEID ModuleId = iota
	METADATA_MODULEID
	ROYALTY_MODULEID
	ROLEASSIGNMENT_MODULEID
)

func (i ModuleId) String() string {
	return []string{"Main", "Metadata", "Royalty", "RoleAssignment"}[i]
}

func ParseModuleId(v string) (any, error) {
	result := MAIN_MODULEID
	switch v {
	case "Main":
		result = MAIN_MODULEID
	case "Metadata":
		result = METADATA_MODULEID
	case "Royalty":
		result = ROYALTY_MODULEID
	case "RoleAssignment":
		result = ROLEASSIGNMENT_MODULEID
	default:
		return 0, errors.New("Unknown ModuleId value: " + v)
	}
	return &result, nil
}

func SerializeModuleId(values []ModuleId) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ModuleId) isMultiValue() bool {
	return false
}
