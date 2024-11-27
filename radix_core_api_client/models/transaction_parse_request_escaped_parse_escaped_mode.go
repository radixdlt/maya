package models
// The type of transaction payload that should be assumed. If omitted, "Any" is used - where the payload isattempted to be parsed as each of the following in turn: Notarized, Signed, Unsigned, Ledger.
type TransactionParseRequest_parse_mode int

const (
    ANY_TRANSACTIONPARSEREQUEST_PARSE_MODE TransactionParseRequest_parse_mode = iota
    NOTARIZED_TRANSACTIONPARSEREQUEST_PARSE_MODE
    SIGNED_TRANSACTIONPARSEREQUEST_PARSE_MODE
    UNSIGNED_TRANSACTIONPARSEREQUEST_PARSE_MODE
    LEDGER_TRANSACTIONPARSEREQUEST_PARSE_MODE
)

func (i TransactionParseRequest_parse_mode) String() string {
    return []string{"Any", "Notarized", "Signed", "Unsigned", "Ledger"}[i]
}
func ParseTransactionParseRequest_parse_mode(v string) (any, error) {
    result := ANY_TRANSACTIONPARSEREQUEST_PARSE_MODE
    switch v {
        case "Any":
            result = ANY_TRANSACTIONPARSEREQUEST_PARSE_MODE
        case "Notarized":
            result = NOTARIZED_TRANSACTIONPARSEREQUEST_PARSE_MODE
        case "Signed":
            result = SIGNED_TRANSACTIONPARSEREQUEST_PARSE_MODE
        case "Unsigned":
            result = UNSIGNED_TRANSACTIONPARSEREQUEST_PARSE_MODE
        case "Ledger":
            result = LEDGER_TRANSACTIONPARSEREQUEST_PARSE_MODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransactionParseRequest_parse_mode(values []TransactionParseRequest_parse_mode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransactionParseRequest_parse_mode) isMultiValue() bool {
    return false
}
