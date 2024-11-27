package models
// The SystemVersion was added at Cuttlefish. Before that it can be assumed to be V1.
type SystemVersion int

const (
    V1_SYSTEMVERSION SystemVersion = iota
    V2_SYSTEMVERSION
    V3_SYSTEMVERSION
)

func (i SystemVersion) String() string {
    return []string{"V1", "V2", "V3"}[i]
}
func ParseSystemVersion(v string) (any, error) {
    result := V1_SYSTEMVERSION
    switch v {
        case "V1":
            result = V1_SYSTEMVERSION
        case "V2":
            result = V2_SYSTEMVERSION
        case "V3":
            result = V3_SYSTEMVERSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemVersion(values []SystemVersion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemVersion) isMultiValue() bool {
    return false
}
