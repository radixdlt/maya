package models

import (
	"errors"
)

// The status of the transaction
type LtsCommittedTransactionStatus int

const (
	SUCCESS_LTSCOMMITTEDTRANSACTIONSTATUS LtsCommittedTransactionStatus = iota
	FAILURE_LTSCOMMITTEDTRANSACTIONSTATUS
)

func (i LtsCommittedTransactionStatus) String() string {
	return []string{"Success", "Failure"}[i]
}

func ParseLtsCommittedTransactionStatus(v string) (any, error) {
	result := SUCCESS_LTSCOMMITTEDTRANSACTIONSTATUS
	switch v {
	case "Success":
		result = SUCCESS_LTSCOMMITTEDTRANSACTIONSTATUS
	case "Failure":
		result = FAILURE_LTSCOMMITTEDTRANSACTIONSTATUS
	default:
		return 0, errors.New("Unknown LtsCommittedTransactionStatus value: " + v)
	}
	return &result, nil
}

func SerializeLtsCommittedTransactionStatus(values []LtsCommittedTransactionStatus) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i LtsCommittedTransactionStatus) isMultiValue() bool {
	return false
}
