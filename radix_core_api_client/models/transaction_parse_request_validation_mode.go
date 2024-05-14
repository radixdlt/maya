package models

import (
	"errors"
)

// The type of validation that should be performed, if the payload correctly decompiles as a Notarized Transaction.This is only relevant for Notarized payloads. If omitted, "Static" is used.
type TransactionParseRequest_validation_mode int

const (
	NONE_TRANSACTIONPARSEREQUEST_VALIDATION_MODE TransactionParseRequest_validation_mode = iota
	STATIC_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
	FULL_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
)

func (i TransactionParseRequest_validation_mode) String() string {
	return []string{"None", "Static", "Full"}[i]
}

func ParseTransactionParseRequest_validation_mode(v string) (any, error) {
	result := NONE_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
	switch v {
	case "None":
		result = NONE_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
	case "Static":
		result = STATIC_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
	case "Full":
		result = FULL_TRANSACTIONPARSEREQUEST_VALIDATION_MODE
	default:
		return 0, errors.New("Unknown TransactionParseRequest_validation_mode value: " + v)
	}
	return &result, nil
}

func SerializeTransactionParseRequest_validation_mode(values []TransactionParseRequest_validation_mode) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i TransactionParseRequest_validation_mode) isMultiValue() bool {
	return false
}
