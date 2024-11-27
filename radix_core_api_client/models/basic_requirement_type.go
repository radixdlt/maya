package models
type BasicRequirementType int

const (
    REQUIRE_BASICREQUIREMENTTYPE BasicRequirementType = iota
    AMOUNTOF_BASICREQUIREMENTTYPE
    ALLOF_BASICREQUIREMENTTYPE
    ANYOF_BASICREQUIREMENTTYPE
    COUNTOF_BASICREQUIREMENTTYPE
)

func (i BasicRequirementType) String() string {
    return []string{"Require", "AmountOf", "AllOf", "AnyOf", "CountOf"}[i]
}
func ParseBasicRequirementType(v string) (any, error) {
    result := REQUIRE_BASICREQUIREMENTTYPE
    switch v {
        case "Require":
            result = REQUIRE_BASICREQUIREMENTTYPE
        case "AmountOf":
            result = AMOUNTOF_BASICREQUIREMENTTYPE
        case "AllOf":
            result = ALLOF_BASICREQUIREMENTTYPE
        case "AnyOf":
            result = ANYOF_BASICREQUIREMENTTYPE
        case "CountOf":
            result = COUNTOF_BASICREQUIREMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBasicRequirementType(values []BasicRequirementType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BasicRequirementType) isMultiValue() bool {
    return false
}
