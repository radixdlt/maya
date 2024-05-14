package models

import (
	"errors"
)

type PresentedBadgeType int

const (
	RESOURCE_PRESENTEDBADGETYPE PresentedBadgeType = iota
	NONFUNGIBLE_PRESENTEDBADGETYPE
)

func (i PresentedBadgeType) String() string {
	return []string{"Resource", "NonFungible"}[i]
}

func ParsePresentedBadgeType(v string) (any, error) {
	result := RESOURCE_PRESENTEDBADGETYPE
	switch v {
	case "Resource":
		result = RESOURCE_PRESENTEDBADGETYPE
	case "NonFungible":
		result = NONFUNGIBLE_PRESENTEDBADGETYPE
	default:
		return 0, errors.New("Unknown PresentedBadgeType value: " + v)
	}
	return &result, nil
}

func SerializePresentedBadgeType(values []PresentedBadgeType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i PresentedBadgeType) isMultiValue() bool {
	return false
}
