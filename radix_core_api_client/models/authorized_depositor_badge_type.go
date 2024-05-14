package models

import (
	"errors"
)

type AuthorizedDepositorBadgeType int

const (
	RESOURCE_AUTHORIZEDDEPOSITORBADGETYPE AuthorizedDepositorBadgeType = iota
	NONFUNGIBLE_AUTHORIZEDDEPOSITORBADGETYPE
)

func (i AuthorizedDepositorBadgeType) String() string {
	return []string{"Resource", "NonFungible"}[i]
}

func ParseAuthorizedDepositorBadgeType(v string) (any, error) {
	result := RESOURCE_AUTHORIZEDDEPOSITORBADGETYPE
	switch v {
	case "Resource":
		result = RESOURCE_AUTHORIZEDDEPOSITORBADGETYPE
	case "NonFungible":
		result = NONFUNGIBLE_AUTHORIZEDDEPOSITORBADGETYPE
	default:
		return 0, errors.New("Unknown AuthorizedDepositorBadgeType value: " + v)
	}
	return &result, nil
}

func SerializeAuthorizedDepositorBadgeType(values []AuthorizedDepositorBadgeType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i AuthorizedDepositorBadgeType) isMultiValue() bool {
	return false
}
