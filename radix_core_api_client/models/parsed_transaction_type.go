package models

import (
	"errors"
)

// The type of the parsed (partial) transaction payload
type ParsedTransactionType int

const (
	NOTARIZEDTRANSACTION_PARSEDTRANSACTIONTYPE ParsedTransactionType = iota
	SIGNEDTRANSACTIONINTENT_PARSEDTRANSACTIONTYPE
	TRANSACTIONINTENT_PARSEDTRANSACTIONTYPE
	LEDGERTRANSACTION_PARSEDTRANSACTIONTYPE
)

func (i ParsedTransactionType) String() string {
	return []string{"NotarizedTransaction", "SignedTransactionIntent", "TransactionIntent", "LedgerTransaction"}[i]
}

func ParseParsedTransactionType(v string) (any, error) {
	result := NOTARIZEDTRANSACTION_PARSEDTRANSACTIONTYPE
	switch v {
	case "NotarizedTransaction":
		result = NOTARIZEDTRANSACTION_PARSEDTRANSACTIONTYPE
	case "SignedTransactionIntent":
		result = SIGNEDTRANSACTIONINTENT_PARSEDTRANSACTIONTYPE
	case "TransactionIntent":
		result = TRANSACTIONINTENT_PARSEDTRANSACTIONTYPE
	case "LedgerTransaction":
		result = LEDGERTRANSACTION_PARSEDTRANSACTIONTYPE
	default:
		return 0, errors.New("Unknown ParsedTransactionType value: " + v)
	}
	return &result, nil
}

func SerializeParsedTransactionType(values []ParsedTransactionType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ParsedTransactionType) isMultiValue() bool {
	return false
}
