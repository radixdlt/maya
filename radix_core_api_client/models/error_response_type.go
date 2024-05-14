package models

import (
	"errors"
)

type ErrorResponseType int

const (
	BASIC_ERRORRESPONSETYPE ErrorResponseType = iota
	TRANSACTIONSUBMIT_ERRORRESPONSETYPE
	LTSTRANSACTIONSUBMIT_ERRORRESPONSETYPE
	STREAMTRANSACTIONS_ERRORRESPONSETYPE
	STREAMPROOFS_ERRORRESPONSETYPE
)

func (i ErrorResponseType) String() string {
	return []string{"Basic", "TransactionSubmit", "LtsTransactionSubmit", "StreamTransactions", "StreamProofs"}[i]
}

func ParseErrorResponseType(v string) (any, error) {
	result := BASIC_ERRORRESPONSETYPE
	switch v {
	case "Basic":
		result = BASIC_ERRORRESPONSETYPE
	case "TransactionSubmit":
		result = TRANSACTIONSUBMIT_ERRORRESPONSETYPE
	case "LtsTransactionSubmit":
		result = LTSTRANSACTIONSUBMIT_ERRORRESPONSETYPE
	case "StreamTransactions":
		result = STREAMTRANSACTIONS_ERRORRESPONSETYPE
	case "StreamProofs":
		result = STREAMPROOFS_ERRORRESPONSETYPE
	default:
		return 0, errors.New("Unknown ErrorResponseType value: " + v)
	}
	return &result, nil
}

func SerializeErrorResponseType(values []ErrorResponseType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ErrorResponseType) isMultiValue() bool {
	return false
}
