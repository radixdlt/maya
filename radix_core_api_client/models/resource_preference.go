package models
type ResourcePreference int

const (
    ALLOWED_RESOURCEPREFERENCE ResourcePreference = iota
    DISALLOWED_RESOURCEPREFERENCE
)

func (i ResourcePreference) String() string {
    return []string{"Allowed", "Disallowed"}[i]
}
func ParseResourcePreference(v string) (any, error) {
    result := ALLOWED_RESOURCEPREFERENCE
    switch v {
        case "Allowed":
            result = ALLOWED_RESOURCEPREFERENCE
        case "Disallowed":
            result = DISALLOWED_RESOURCEPREFERENCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeResourcePreference(values []ResourcePreference) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ResourcePreference) isMultiValue() bool {
    return false
}
