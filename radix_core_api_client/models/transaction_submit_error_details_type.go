package models

import (
	"errors"
)

type TransactionSubmitErrorDetailsType int

const (
	PRIORITYTHRESHOLDNOTMET_TRANSACTIONSUBMITERRORDETAILSTYPE TransactionSubmitErrorDetailsType = iota
	INTENTALREADYCOMMITTED_TRANSACTIONSUBMITERRORDETAILSTYPE
	REJECTED_TRANSACTIONSUBMITERRORDETAILSTYPE
)

func (i TransactionSubmitErrorDetailsType) String() string {
	return []string{"PriorityThresholdNotMet", "IntentAlreadyCommitted", "Rejected"}[i]
}

func ParseTransactionSubmitErrorDetailsType(v string) (any, error) {
	result := PRIORITYTHRESHOLDNOTMET_TRANSACTIONSUBMITERRORDETAILSTYPE
	switch v {
	case "PriorityThresholdNotMet":
		result = PRIORITYTHRESHOLDNOTMET_TRANSACTIONSUBMITERRORDETAILSTYPE
	case "IntentAlreadyCommitted":
		result = INTENTALREADYCOMMITTED_TRANSACTIONSUBMITERRORDETAILSTYPE
	case "Rejected":
		result = REJECTED_TRANSACTIONSUBMITERRORDETAILSTYPE
	default:
		return 0, errors.New("Unknown TransactionSubmitErrorDetailsType value: " + v)
	}
	return &result, nil
}

func SerializeTransactionSubmitErrorDetailsType(values []TransactionSubmitErrorDetailsType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i TransactionSubmitErrorDetailsType) isMultiValue() bool {
	return false
}
