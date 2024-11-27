package models
type RequirementType int

const (
    RESOURCE_REQUIREMENTTYPE RequirementType = iota
    NONFUNGIBLE_REQUIREMENTTYPE
)

func (i RequirementType) String() string {
    return []string{"Resource", "NonFungible"}[i]
}
func ParseRequirementType(v string) (any, error) {
    result := RESOURCE_REQUIREMENTTYPE
    switch v {
        case "Resource":
            result = RESOURCE_REQUIREMENTTYPE
        case "NonFungible":
            result = NONFUNGIBLE_REQUIREMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRequirementType(values []RequirementType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequirementType) isMultiValue() bool {
    return false
}
