package models

import (
	"errors"
)

type GenericSubstitutionType int

const (
	LOCAL_GENERICSUBSTITUTIONTYPE GenericSubstitutionType = iota
	REMOTE_GENERICSUBSTITUTIONTYPE
)

func (i GenericSubstitutionType) String() string {
	return []string{"Local", "Remote"}[i]
}

func ParseGenericSubstitutionType(v string) (any, error) {
	result := LOCAL_GENERICSUBSTITUTIONTYPE
	switch v {
	case "Local":
		result = LOCAL_GENERICSUBSTITUTIONTYPE
	case "Remote":
		result = REMOTE_GENERICSUBSTITUTIONTYPE
	default:
		return 0, errors.New("Unknown GenericSubstitutionType value: " + v)
	}
	return &result, nil
}

func SerializeGenericSubstitutionType(values []GenericSubstitutionType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i GenericSubstitutionType) isMultiValue() bool {
	return false
}
