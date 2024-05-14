package models

import (
	"errors"
)

// The status of the transaction
type TransactionStatus int

const (
	SUCCEEDED_TRANSACTIONSTATUS TransactionStatus = iota
	FAILED_TRANSACTIONSTATUS
	REJECTED_TRANSACTIONSTATUS
)

func (i TransactionStatus) String() string {
	return []string{"Succeeded", "Failed", "Rejected"}[i]
}

func ParseTransactionStatus(v string) (any, error) {
	result := SUCCEEDED_TRANSACTIONSTATUS
	switch v {
	case "Succeeded":
		result = SUCCEEDED_TRANSACTIONSTATUS
	case "Failed":
		result = FAILED_TRANSACTIONSTATUS
	case "Rejected":
		result = REJECTED_TRANSACTIONSTATUS
	default:
		return 0, errors.New("Unknown TransactionStatus value: " + v)
	}
	return &result, nil
}

func SerializeTransactionStatus(values []TransactionStatus) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i TransactionStatus) isMultiValue() bool {
	return false
}
