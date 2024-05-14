package models

import (
	"errors"
)

type LedgerStateSelectorType int

const (
	BYSTATEVERSION_LEDGERSTATESELECTORTYPE LedgerStateSelectorType = iota
)

func (i LedgerStateSelectorType) String() string {
	return []string{"ByStateVersion"}[i]
}

func ParseLedgerStateSelectorType(v string) (any, error) {
	result := BYSTATEVERSION_LEDGERSTATESELECTORTYPE
	switch v {
	case "ByStateVersion":
		result = BYSTATEVERSION_LEDGERSTATESELECTORTYPE
	default:
		return 0, errors.New("Unknown LedgerStateSelectorType value: " + v)
	}
	return &result, nil
}

func SerializeLedgerStateSelectorType(values []LedgerStateSelectorType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i LedgerStateSelectorType) isMultiValue() bool {
	return false
}
