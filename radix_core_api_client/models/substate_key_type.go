package models

import (
	"errors"
)

type SubstateKeyType int

const (
	FIELD_SUBSTATEKEYTYPE SubstateKeyType = iota
	MAP_SUBSTATEKEYTYPE
	SORTED_SUBSTATEKEYTYPE
)

func (i SubstateKeyType) String() string {
	return []string{"Field", "Map", "Sorted"}[i]
}

func ParseSubstateKeyType(v string) (any, error) {
	result := FIELD_SUBSTATEKEYTYPE
	switch v {
	case "Field":
		result = FIELD_SUBSTATEKEYTYPE
	case "Map":
		result = MAP_SUBSTATEKEYTYPE
	case "Sorted":
		result = SORTED_SUBSTATEKEYTYPE
	default:
		return 0, errors.New("Unknown SubstateKeyType value: " + v)
	}
	return &result, nil
}

func SerializeSubstateKeyType(values []SubstateKeyType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i SubstateKeyType) isMultiValue() bool {
	return false
}
