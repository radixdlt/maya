package models
// The type of the ledger transaction
type LedgerTransactionType int

const (
    GENESIS_LEDGERTRANSACTIONTYPE LedgerTransactionType = iota
    USER_LEDGERTRANSACTIONTYPE
    USERV2_LEDGERTRANSACTIONTYPE
    ROUNDUPDATE_LEDGERTRANSACTIONTYPE
    FLASH_LEDGERTRANSACTIONTYPE
)

func (i LedgerTransactionType) String() string {
    return []string{"Genesis", "User", "UserV2", "RoundUpdate", "Flash"}[i]
}
func ParseLedgerTransactionType(v string) (any, error) {
    result := GENESIS_LEDGERTRANSACTIONTYPE
    switch v {
        case "Genesis":
            result = GENESIS_LEDGERTRANSACTIONTYPE
        case "User":
            result = USER_LEDGERTRANSACTIONTYPE
        case "UserV2":
            result = USERV2_LEDGERTRANSACTIONTYPE
        case "RoundUpdate":
            result = ROUNDUPDATE_LEDGERTRANSACTIONTYPE
        case "Flash":
            result = FLASH_LEDGERTRANSACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLedgerTransactionType(values []LedgerTransactionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LedgerTransactionType) isMultiValue() bool {
    return false
}
