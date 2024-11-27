package models
type PreviewTransactionType int

const (
    COMPILED_PREVIEWTRANSACTIONTYPE PreviewTransactionType = iota
)

func (i PreviewTransactionType) String() string {
    return []string{"Compiled"}[i]
}
func ParsePreviewTransactionType(v string) (any, error) {
    result := COMPILED_PREVIEWTRANSACTIONTYPE
    switch v {
        case "Compiled":
            result = COMPILED_PREVIEWTRANSACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePreviewTransactionType(values []PreviewTransactionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PreviewTransactionType) isMultiValue() bool {
    return false
}
