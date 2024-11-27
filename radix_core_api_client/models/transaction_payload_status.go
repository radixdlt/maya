package models
type TransactionPayloadStatus int

const (
    COMMITTEDSUCCESS_TRANSACTIONPAYLOADSTATUS TransactionPayloadStatus = iota
    COMMITTEDFAILURE_TRANSACTIONPAYLOADSTATUS
    PERMANENTLYREJECTED_TRANSACTIONPAYLOADSTATUS
    TRANSIENTLYREJECTED_TRANSACTIONPAYLOADSTATUS
    INMEMPOOL_TRANSACTIONPAYLOADSTATUS
    NOTINMEMPOOL_TRANSACTIONPAYLOADSTATUS
)

func (i TransactionPayloadStatus) String() string {
    return []string{"CommittedSuccess", "CommittedFailure", "PermanentlyRejected", "TransientlyRejected", "InMempool", "NotInMempool"}[i]
}
func ParseTransactionPayloadStatus(v string) (any, error) {
    result := COMMITTEDSUCCESS_TRANSACTIONPAYLOADSTATUS
    switch v {
        case "CommittedSuccess":
            result = COMMITTEDSUCCESS_TRANSACTIONPAYLOADSTATUS
        case "CommittedFailure":
            result = COMMITTEDFAILURE_TRANSACTIONPAYLOADSTATUS
        case "PermanentlyRejected":
            result = PERMANENTLYREJECTED_TRANSACTIONPAYLOADSTATUS
        case "TransientlyRejected":
            result = TRANSIENTLYREJECTED_TRANSACTIONPAYLOADSTATUS
        case "InMempool":
            result = INMEMPOOL_TRANSACTIONPAYLOADSTATUS
        case "NotInMempool":
            result = NOTINMEMPOOL_TRANSACTIONPAYLOADSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransactionPayloadStatus(values []TransactionPayloadStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransactionPayloadStatus) isMultiValue() bool {
    return false
}
