package models
type AccessRuleType int

const (
    PROTECTED_ACCESSRULETYPE AccessRuleType = iota
    ALLOWALL_ACCESSRULETYPE
    DENYALL_ACCESSRULETYPE
)

func (i AccessRuleType) String() string {
    return []string{"Protected", "AllowAll", "DenyAll"}[i]
}
func ParseAccessRuleType(v string) (any, error) {
    result := PROTECTED_ACCESSRULETYPE
    switch v {
        case "Protected":
            result = PROTECTED_ACCESSRULETYPE
        case "AllowAll":
            result = ALLOWALL_ACCESSRULETYPE
        case "DenyAll":
            result = DENYALL_ACCESSRULETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccessRuleType(values []AccessRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccessRuleType) isMultiValue() bool {
    return false
}
