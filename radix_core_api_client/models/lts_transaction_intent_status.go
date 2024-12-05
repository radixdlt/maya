package models
// The status of the transaction intent, as determined by the node.FateUncertain or FateUncertainButLikelyRejection mean that it's still possible that a payload containing the transaction
type LtsTransactionIntentStatus int

const (
    COMMITTEDSUCCESS_LTSTRANSACTIONINTENTSTATUS LtsTransactionIntentStatus = iota
    COMMITTEDFAILURE_LTSTRANSACTIONINTENTSTATUS
    PERMANENTREJECTION_LTSTRANSACTIONINTENTSTATUS
    INMEMPOOL_LTSTRANSACTIONINTENTSTATUS
    NOTSEEN_LTSTRANSACTIONINTENTSTATUS
    FATEUNCERTAIN_LTSTRANSACTIONINTENTSTATUS
    FATEUNCERTAINBUTLIKELYREJECTION_LTSTRANSACTIONINTENTSTATUS
)

func (i LtsTransactionIntentStatus) String() string {
    return []string{"CommittedSuccess", "CommittedFailure", "PermanentRejection", "InMempool", "NotSeen", "FateUncertain", "FateUncertainButLikelyRejection"}[i]
}
func ParseLtsTransactionIntentStatus(v string) (any, error) {
    result := COMMITTEDSUCCESS_LTSTRANSACTIONINTENTSTATUS
    switch v {
        case "CommittedSuccess":
            result = COMMITTEDSUCCESS_LTSTRANSACTIONINTENTSTATUS
        case "CommittedFailure":
            result = COMMITTEDFAILURE_LTSTRANSACTIONINTENTSTATUS
        case "PermanentRejection":
            result = PERMANENTREJECTION_LTSTRANSACTIONINTENTSTATUS
        case "InMempool":
            result = INMEMPOOL_LTSTRANSACTIONINTENTSTATUS
        case "NotSeen":
            result = NOTSEEN_LTSTRANSACTIONINTENTSTATUS
        case "FateUncertain":
            result = FATEUNCERTAIN_LTSTRANSACTIONINTENTSTATUS
        case "FateUncertainButLikelyRejection":
            result = FATEUNCERTAINBUTLIKELYREJECTION_LTSTRANSACTIONINTENTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLtsTransactionIntentStatus(values []LtsTransactionIntentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LtsTransactionIntentStatus) isMultiValue() bool {
    return false
}
