package models
type LtsFeeFungibleResourceBalanceChangeType int

const (
    FEEPAYMENT_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE LtsFeeFungibleResourceBalanceChangeType = iota
    FEEDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
    TIPDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
    ROYALTYDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
)

func (i LtsFeeFungibleResourceBalanceChangeType) String() string {
    return []string{"FeePayment", "FeeDistributed", "TipDistributed", "RoyaltyDistributed"}[i]
}
func ParseLtsFeeFungibleResourceBalanceChangeType(v string) (any, error) {
    result := FEEPAYMENT_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
    switch v {
        case "FeePayment":
            result = FEEPAYMENT_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
        case "FeeDistributed":
            result = FEEDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
        case "TipDistributed":
            result = TIPDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
        case "RoyaltyDistributed":
            result = ROYALTYDISTRIBUTED_LTSFEEFUNGIBLERESOURCEBALANCECHANGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLtsFeeFungibleResourceBalanceChangeType(values []LtsFeeFungibleResourceBalanceChangeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LtsFeeFungibleResourceBalanceChangeType) isMultiValue() bool {
    return false
}
