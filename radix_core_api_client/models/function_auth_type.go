package models

import (
	"errors"
)

type FunctionAuthType int

const (
	ALLOWALL_FUNCTIONAUTHTYPE FunctionAuthType = iota
	FUNCTIONACCESSRULES_FUNCTIONAUTHTYPE
	ROOTONLY_FUNCTIONAUTHTYPE
)

func (i FunctionAuthType) String() string {
	return []string{"AllowAll", "FunctionAccessRules", "RootOnly"}[i]
}

func ParseFunctionAuthType(v string) (any, error) {
	result := ALLOWALL_FUNCTIONAUTHTYPE
	switch v {
	case "AllowAll":
		result = ALLOWALL_FUNCTIONAUTHTYPE
	case "FunctionAccessRules":
		result = FUNCTIONACCESSRULES_FUNCTIONAUTHTYPE
	case "RootOnly":
		result = ROOTONLY_FUNCTIONAUTHTYPE
	default:
		return 0, errors.New("Unknown FunctionAuthType value: " + v)
	}
	return &result, nil
}

func SerializeFunctionAuthType(values []FunctionAuthType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i FunctionAuthType) isMultiValue() bool {
	return false
}
