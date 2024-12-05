package models
type LtsTransactionSubmitErrorDetailsType int

const (
    PRIORITYTHRESHOLDNOTMET_LTSTRANSACTIONSUBMITERRORDETAILSTYPE LtsTransactionSubmitErrorDetailsType = iota
    INTENTALREADYCOMMITTED_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
    REJECTED_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
)

func (i LtsTransactionSubmitErrorDetailsType) String() string {
    return []string{"PriorityThresholdNotMet", "IntentAlreadyCommitted", "Rejected"}[i]
}
func ParseLtsTransactionSubmitErrorDetailsType(v string) (any, error) {
    result := PRIORITYTHRESHOLDNOTMET_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
    switch v {
        case "PriorityThresholdNotMet":
            result = PRIORITYTHRESHOLDNOTMET_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
        case "IntentAlreadyCommitted":
            result = INTENTALREADYCOMMITTED_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
        case "Rejected":
            result = REJECTED_LTSTRANSACTIONSUBMITERRORDETAILSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLtsTransactionSubmitErrorDetailsType(values []LtsTransactionSubmitErrorDetailsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LtsTransactionSubmitErrorDetailsType) isMultiValue() bool {
    return false
}
