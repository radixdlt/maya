package models

import (
	"errors"
)

type StreamTransactionsErrorDetailsType int

const (
	REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMTRANSACTIONSERRORDETAILSTYPE StreamTransactionsErrorDetailsType = iota
)

func (i StreamTransactionsErrorDetailsType) String() string {
	return []string{"RequestedStateVersionOutOfBounds"}[i]
}

func ParseStreamTransactionsErrorDetailsType(v string) (any, error) {
	result := REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMTRANSACTIONSERRORDETAILSTYPE
	switch v {
	case "RequestedStateVersionOutOfBounds":
		result = REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMTRANSACTIONSERRORDETAILSTYPE
	default:
		return 0, errors.New("Unknown StreamTransactionsErrorDetailsType value: " + v)
	}
	return &result, nil
}

func SerializeStreamTransactionsErrorDetailsType(values []StreamTransactionsErrorDetailsType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i StreamTransactionsErrorDetailsType) isMultiValue() bool {
	return false
}
