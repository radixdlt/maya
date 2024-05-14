package models

import (
	"errors"
)

type AttachedModuleId int

const (
	METADATA_ATTACHEDMODULEID AttachedModuleId = iota
	ROYALTY_ATTACHEDMODULEID
	ROLEASSIGNMENT_ATTACHEDMODULEID
)

func (i AttachedModuleId) String() string {
	return []string{"Metadata", "Royalty", "RoleAssignment"}[i]
}

func ParseAttachedModuleId(v string) (any, error) {
	result := METADATA_ATTACHEDMODULEID
	switch v {
	case "Metadata":
		result = METADATA_ATTACHEDMODULEID
	case "Royalty":
		result = ROYALTY_ATTACHEDMODULEID
	case "RoleAssignment":
		result = ROLEASSIGNMENT_ATTACHEDMODULEID
	default:
		return 0, errors.New("Unknown AttachedModuleId value: " + v)
	}
	return &result, nil
}

func SerializeAttachedModuleId(values []AttachedModuleId) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i AttachedModuleId) isMultiValue() bool {
	return false
}
