package models

import (
	"errors"
)

// The constraints on the concrete type provided to fill the generic type parameter.Note: currently, we only support the wildcard (i.e. `Any`) generic type.
type GenericTypeParameterConstraints int

const (
	ANY_GENERICTYPEPARAMETERCONSTRAINTS GenericTypeParameterConstraints = iota
)

func (i GenericTypeParameterConstraints) String() string {
	return []string{"Any"}[i]
}

func ParseGenericTypeParameterConstraints(v string) (any, error) {
	result := ANY_GENERICTYPEPARAMETERCONSTRAINTS
	switch v {
	case "Any":
		result = ANY_GENERICTYPEPARAMETERCONSTRAINTS
	default:
		return 0, errors.New("Unknown GenericTypeParameterConstraints value: " + v)
	}
	return &result, nil
}

func SerializeGenericTypeParameterConstraints(values []GenericTypeParameterConstraints) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i GenericTypeParameterConstraints) isMultiValue() bool {
	return false
}
