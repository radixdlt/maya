package models

import (
	"errors"
)

type RoleSpecification int

const (
	NORMAL_ROLESPECIFICATION RoleSpecification = iota
	USEOUTER_ROLESPECIFICATION
)

func (i RoleSpecification) String() string {
	return []string{"Normal", "UseOuter"}[i]
}

func ParseRoleSpecification(v string) (any, error) {
	result := NORMAL_ROLESPECIFICATION
	switch v {
	case "Normal":
		result = NORMAL_ROLESPECIFICATION
	case "UseOuter":
		result = USEOUTER_ROLESPECIFICATION
	default:
		return 0, errors.New("Unknown RoleSpecification value: " + v)
	}
	return &result, nil
}

func SerializeRoleSpecification(values []RoleSpecification) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i RoleSpecification) isMultiValue() bool {
	return false
}
