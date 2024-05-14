package models

import (
	"errors"
)

type ResourceType int

const (
	FUNGIBLE_RESOURCETYPE ResourceType = iota
	NONFUNGIBLE_RESOURCETYPE
)

func (i ResourceType) String() string {
	return []string{"Fungible", "NonFungible"}[i]
}

func ParseResourceType(v string) (any, error) {
	result := FUNGIBLE_RESOURCETYPE
	switch v {
	case "Fungible":
		result = FUNGIBLE_RESOURCETYPE
	case "NonFungible":
		result = NONFUNGIBLE_RESOURCETYPE
	default:
		return 0, errors.New("Unknown ResourceType value: " + v)
	}
	return &result, nil
}

func SerializeResourceType(values []ResourceType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ResourceType) isMultiValue() bool {
	return false
}
