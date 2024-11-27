package models
type ObjectHook int

const (
    ONVIRTUALIZE_OBJECTHOOK ObjectHook = iota
    ONMOVE_OBJECTHOOK
    ONDROP_OBJECTHOOK
)

func (i ObjectHook) String() string {
    return []string{"OnVirtualize", "OnMove", "OnDrop"}[i]
}
func ParseObjectHook(v string) (any, error) {
    result := ONVIRTUALIZE_OBJECTHOOK
    switch v {
        case "OnVirtualize":
            result = ONVIRTUALIZE_OBJECTHOOK
        case "OnMove":
            result = ONMOVE_OBJECTHOOK
        case "OnDrop":
            result = ONDROP_OBJECTHOOK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeObjectHook(values []ObjectHook) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ObjectHook) isMultiValue() bool {
    return false
}
