package models

import (
	"errors"
)

// The status of the transaction payload, as per this node.A NotInMempool status means that it wasn't rejected at last execution attempt, but it's not currently in the mempool either.
type LtsTransactionPayloadStatus int

const (
	COMMITTEDSUCCESS_LTSTRANSACTIONPAYLOADSTATUS LtsTransactionPayloadStatus = iota
	COMMITTEDFAILURE_LTSTRANSACTIONPAYLOADSTATUS
	PERMANENTLYREJECTED_LTSTRANSACTIONPAYLOADSTATUS
	TRANSIENTLYREJECTED_LTSTRANSACTIONPAYLOADSTATUS
	INMEMPOOL_LTSTRANSACTIONPAYLOADSTATUS
	NOTINMEMPOOL_LTSTRANSACTIONPAYLOADSTATUS
)

func (i LtsTransactionPayloadStatus) String() string {
	return []string{"CommittedSuccess", "CommittedFailure", "PermanentlyRejected", "TransientlyRejected", "InMempool", "NotInMempool"}[i]
}

func ParseLtsTransactionPayloadStatus(v string) (any, error) {
	result := COMMITTEDSUCCESS_LTSTRANSACTIONPAYLOADSTATUS
	switch v {
	case "CommittedSuccess":
		result = COMMITTEDSUCCESS_LTSTRANSACTIONPAYLOADSTATUS
	case "CommittedFailure":
		result = COMMITTEDFAILURE_LTSTRANSACTIONPAYLOADSTATUS
	case "PermanentlyRejected":
		result = PERMANENTLYREJECTED_LTSTRANSACTIONPAYLOADSTATUS
	case "TransientlyRejected":
		result = TRANSIENTLYREJECTED_LTSTRANSACTIONPAYLOADSTATUS
	case "InMempool":
		result = INMEMPOOL_LTSTRANSACTIONPAYLOADSTATUS
	case "NotInMempool":
		result = NOTINMEMPOOL_LTSTRANSACTIONPAYLOADSTATUS
	default:
		return 0, errors.New("Unknown LtsTransactionPayloadStatus value: " + v)
	}
	return &result, nil
}

func SerializeLtsTransactionPayloadStatus(values []LtsTransactionPayloadStatus) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i LtsTransactionPayloadStatus) isMultiValue() bool {
	return false
}
