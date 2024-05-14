package models

import (
	"errors"
)

// The amount of information to return in the response."Basic" includes the type, validity information, and any relevant identifiers."Full" also includes the fully parsed information.If omitted, "Full" is used.
type TransactionParseRequest_response_mode int

const (
	BASIC_TRANSACTIONPARSEREQUEST_RESPONSE_MODE TransactionParseRequest_response_mode = iota
	FULL_TRANSACTIONPARSEREQUEST_RESPONSE_MODE
)

func (i TransactionParseRequest_response_mode) String() string {
	return []string{"Basic", "Full"}[i]
}

func ParseTransactionParseRequest_response_mode(v string) (any, error) {
	result := BASIC_TRANSACTIONPARSEREQUEST_RESPONSE_MODE
	switch v {
	case "Basic":
		result = BASIC_TRANSACTIONPARSEREQUEST_RESPONSE_MODE
	case "Full":
		result = FULL_TRANSACTIONPARSEREQUEST_RESPONSE_MODE
	default:
		return 0, errors.New("Unknown TransactionParseRequest_response_mode value: " + v)
	}
	return &result, nil
}

func SerializeTransactionParseRequest_response_mode(values []TransactionParseRequest_response_mode) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i TransactionParseRequest_response_mode) isMultiValue() bool {
	return false
}
