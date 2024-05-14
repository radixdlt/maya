package models

import (
	"errors"
)

type MethodAuthType int

const (
	ALLOWALL_METHODAUTHTYPE MethodAuthType = iota
	STATICROLEDEFINITION_METHODAUTHTYPE
)

func (i MethodAuthType) String() string {
	return []string{"AllowAll", "StaticRoleDefinition"}[i]
}

func ParseMethodAuthType(v string) (any, error) {
	result := ALLOWALL_METHODAUTHTYPE
	switch v {
	case "AllowAll":
		result = ALLOWALL_METHODAUTHTYPE
	case "StaticRoleDefinition":
		result = STATICROLEDEFINITION_METHODAUTHTYPE
	default:
		return 0, errors.New("Unknown MethodAuthType value: " + v)
	}
	return &result, nil
}

func SerializeMethodAuthType(values []MethodAuthType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i MethodAuthType) isMultiValue() bool {
	return false
}
