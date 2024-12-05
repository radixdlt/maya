package models
type TransactionTrackerTransactionStatus int

const (
    COMMITTEDSUCCESS_TRANSACTIONTRACKERTRANSACTIONSTATUS TransactionTrackerTransactionStatus = iota
    COMMITTEDFAILURE_TRANSACTIONTRACKERTRANSACTIONSTATUS
    CANCELLED_TRANSACTIONTRACKERTRANSACTIONSTATUS
)

func (i TransactionTrackerTransactionStatus) String() string {
    return []string{"CommittedSuccess", "CommittedFailure", "Cancelled"}[i]
}
func ParseTransactionTrackerTransactionStatus(v string) (any, error) {
    result := COMMITTEDSUCCESS_TRANSACTIONTRACKERTRANSACTIONSTATUS
    switch v {
        case "CommittedSuccess":
            result = COMMITTEDSUCCESS_TRANSACTIONTRACKERTRANSACTIONSTATUS
        case "CommittedFailure":
            result = COMMITTEDFAILURE_TRANSACTIONTRACKERTRANSACTIONSTATUS
        case "Cancelled":
            result = CANCELLED_TRANSACTIONTRACKERTRANSACTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransactionTrackerTransactionStatus(values []TransactionTrackerTransactionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransactionTrackerTransactionStatus) isMultiValue() bool {
    return false
}
