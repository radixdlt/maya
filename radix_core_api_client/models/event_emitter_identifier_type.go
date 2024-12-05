package models
type EventEmitterIdentifierType int

const (
    FUNCTION_EVENTEMITTERIDENTIFIERTYPE EventEmitterIdentifierType = iota
    METHOD_EVENTEMITTERIDENTIFIERTYPE
)

func (i EventEmitterIdentifierType) String() string {
    return []string{"Function", "Method"}[i]
}
func ParseEventEmitterIdentifierType(v string) (any, error) {
    result := FUNCTION_EVENTEMITTERIDENTIFIERTYPE
    switch v {
        case "Function":
            result = FUNCTION_EVENTEMITTERIDENTIFIERTYPE
        case "Method":
            result = METHOD_EVENTEMITTERIDENTIFIERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEventEmitterIdentifierType(values []EventEmitterIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EventEmitterIdentifierType) isMultiValue() bool {
    return false
}
