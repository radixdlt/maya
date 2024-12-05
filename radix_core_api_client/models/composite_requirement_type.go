package models
// `ProofRule` is now called `BasicRequirement`.
type CompositeRequirementType int

const (
    PROOFRULE_COMPOSITEREQUIREMENTTYPE CompositeRequirementType = iota
    ANYOF_COMPOSITEREQUIREMENTTYPE
    ALLOF_COMPOSITEREQUIREMENTTYPE
)

func (i CompositeRequirementType) String() string {
    return []string{"ProofRule", "AnyOf", "AllOf"}[i]
}
func ParseCompositeRequirementType(v string) (any, error) {
    result := PROOFRULE_COMPOSITEREQUIREMENTTYPE
    switch v {
        case "ProofRule":
            result = PROOFRULE_COMPOSITEREQUIREMENTTYPE
        case "AnyOf":
            result = ANYOF_COMPOSITEREQUIREMENTTYPE
        case "AllOf":
            result = ALLOF_COMPOSITEREQUIREMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCompositeRequirementType(values []CompositeRequirementType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CompositeRequirementType) isMultiValue() bool {
    return false
}
