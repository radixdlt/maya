package models

import (
	"errors"
)

// The status of the transaction intent, as determined by the node.FateUncertain or FateUncertainButLikelyRejection mean that it's still possible that a payload containing the transaction
type TransactionIntentStatus int

const (
	COMMITTEDSUCCESS_TRANSACTIONINTENTSTATUS TransactionIntentStatus = iota
	COMMITTEDFAILURE_TRANSACTIONINTENTSTATUS
	PERMANENTREJECTION_TRANSACTIONINTENTSTATUS
	INMEMPOOL_TRANSACTIONINTENTSTATUS
	NOTSEEN_TRANSACTIONINTENTSTATUS
	FATEUNCERTAIN_TRANSACTIONINTENTSTATUS
	FATEUNCERTAINBUTLIKELYREJECTION_TRANSACTIONINTENTSTATUS
)

func (i TransactionIntentStatus) String() string {
	return []string{"CommittedSuccess", "CommittedFailure", "PermanentRejection", "InMempool", "NotSeen", "FateUncertain", "FateUncertainButLikelyRejection"}[i]
}

func ParseTransactionIntentStatus(v string) (any, error) {
	result := COMMITTEDSUCCESS_TRANSACTIONINTENTSTATUS
	switch v {
	case "CommittedSuccess":
		result = COMMITTEDSUCCESS_TRANSACTIONINTENTSTATUS
	case "CommittedFailure":
		result = COMMITTEDFAILURE_TRANSACTIONINTENTSTATUS
	case "PermanentRejection":
		result = PERMANENTREJECTION_TRANSACTIONINTENTSTATUS
	case "InMempool":
		result = INMEMPOOL_TRANSACTIONINTENTSTATUS
	case "NotSeen":
		result = NOTSEEN_TRANSACTIONINTENTSTATUS
	case "FateUncertain":
		result = FATEUNCERTAIN_TRANSACTIONINTENTSTATUS
	case "FateUncertainButLikelyRejection":
		result = FATEUNCERTAINBUTLIKELYREJECTION_TRANSACTIONINTENTSTATUS
	default:
		return 0, errors.New("Unknown TransactionIntentStatus value: " + v)
	}
	return &result, nil
}

func SerializeTransactionIntentStatus(values []TransactionIntentStatus) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i TransactionIntentStatus) isMultiValue() bool {
	return false
}
