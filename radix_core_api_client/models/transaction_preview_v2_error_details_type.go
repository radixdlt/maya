package models
type TransactionPreviewV2ErrorDetailsType int

const (
    INVALID_TRANSACTIONPREVIEWV2ERRORDETAILSTYPE TransactionPreviewV2ErrorDetailsType = iota
)

func (i TransactionPreviewV2ErrorDetailsType) String() string {
    return []string{"Invalid"}[i]
}
func ParseTransactionPreviewV2ErrorDetailsType(v string) (any, error) {
    result := INVALID_TRANSACTIONPREVIEWV2ERRORDETAILSTYPE
    switch v {
        case "Invalid":
            result = INVALID_TRANSACTIONPREVIEWV2ERRORDETAILSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransactionPreviewV2ErrorDetailsType(values []TransactionPreviewV2ErrorDetailsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransactionPreviewV2ErrorDetailsType) isMultiValue() bool {
    return false
}
