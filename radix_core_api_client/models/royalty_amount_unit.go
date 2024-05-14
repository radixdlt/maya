package models

import (
	"errors"
)

type RoyaltyAmount_unit int

const (
	XRD_ROYALTYAMOUNT_UNIT RoyaltyAmount_unit = iota
	USD_ROYALTYAMOUNT_UNIT
)

func (i RoyaltyAmount_unit) String() string {
	return []string{"XRD", "USD"}[i]
}

func ParseRoyaltyAmount_unit(v string) (any, error) {
	result := XRD_ROYALTYAMOUNT_UNIT
	switch v {
	case "XRD":
		result = XRD_ROYALTYAMOUNT_UNIT
	case "USD":
		result = USD_ROYALTYAMOUNT_UNIT
	default:
		return 0, errors.New("Unknown RoyaltyAmount_unit value: " + v)
	}
	return &result, nil
}

func SerializeRoyaltyAmount_unit(values []RoyaltyAmount_unit) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i RoyaltyAmount_unit) isMultiValue() bool {
	return false
}
