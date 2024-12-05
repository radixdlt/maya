package models
type BlueprintPayloadDefType int

const (
    STATIC_BLUEPRINTPAYLOADDEFTYPE BlueprintPayloadDefType = iota
    GENERIC_BLUEPRINTPAYLOADDEFTYPE
)

func (i BlueprintPayloadDefType) String() string {
    return []string{"Static", "Generic"}[i]
}
func ParseBlueprintPayloadDefType(v string) (any, error) {
    result := STATIC_BLUEPRINTPAYLOADDEFTYPE
    switch v {
        case "Static":
            result = STATIC_BLUEPRINTPAYLOADDEFTYPE
        case "Generic":
            result = GENERIC_BLUEPRINTPAYLOADDEFTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBlueprintPayloadDefType(values []BlueprintPayloadDefType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BlueprintPayloadDefType) isMultiValue() bool {
    return false
}
