package models
type TargetIdentifierType int

const (
    METHOD_TARGETIDENTIFIERTYPE TargetIdentifierType = iota
    FUNCTION_TARGETIDENTIFIERTYPE
)

func (i TargetIdentifierType) String() string {
    return []string{"Method", "Function"}[i]
}
func ParseTargetIdentifierType(v string) (any, error) {
    result := METHOD_TARGETIDENTIFIERTYPE
    switch v {
        case "Method":
            result = METHOD_TARGETIDENTIFIERTYPE
        case "Function":
            result = FUNCTION_TARGETIDENTIFIERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTargetIdentifierType(values []TargetIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TargetIdentifierType) isMultiValue() bool {
    return false
}
