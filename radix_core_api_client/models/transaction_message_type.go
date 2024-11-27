package models
type TransactionMessageType int

const (
    PLAINTEXT_TRANSACTIONMESSAGETYPE TransactionMessageType = iota
    ENCRYPTED_TRANSACTIONMESSAGETYPE
)

func (i TransactionMessageType) String() string {
    return []string{"Plaintext", "Encrypted"}[i]
}
func ParseTransactionMessageType(v string) (any, error) {
    result := PLAINTEXT_TRANSACTIONMESSAGETYPE
    switch v {
        case "Plaintext":
            result = PLAINTEXT_TRANSACTIONMESSAGETYPE
        case "Encrypted":
            result = ENCRYPTED_TRANSACTIONMESSAGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransactionMessageType(values []TransactionMessageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransactionMessageType) isMultiValue() bool {
    return false
}
