package models
type ManifestValidationRuleset int

const (
    BASIC_MANIFESTVALIDATIONRULESET ManifestValidationRuleset = iota
    CUTTLEFISH_MANIFESTVALIDATIONRULESET
)

func (i ManifestValidationRuleset) String() string {
    return []string{"Basic", "Cuttlefish"}[i]
}
func ParseManifestValidationRuleset(v string) (any, error) {
    result := BASIC_MANIFESTVALIDATIONRULESET
    switch v {
        case "Basic":
            result = BASIC_MANIFESTVALIDATIONRULESET
        case "Cuttlefish":
            result = CUTTLEFISH_MANIFESTVALIDATIONRULESET
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManifestValidationRuleset(values []ManifestValidationRuleset) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManifestValidationRuleset) isMultiValue() bool {
    return false
}
