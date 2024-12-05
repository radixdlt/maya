package models
// This was added in Cuttlefish. Before that, this value was missing, but can be assumed to be V1.
type AlwaysVisibleGlobalNodesVersion int

const (
    V1_ALWAYSVISIBLEGLOBALNODESVERSION AlwaysVisibleGlobalNodesVersion = iota
    V2_ALWAYSVISIBLEGLOBALNODESVERSION
)

func (i AlwaysVisibleGlobalNodesVersion) String() string {
    return []string{"V1", "V2"}[i]
}
func ParseAlwaysVisibleGlobalNodesVersion(v string) (any, error) {
    result := V1_ALWAYSVISIBLEGLOBALNODESVERSION
    switch v {
        case "V1":
            result = V1_ALWAYSVISIBLEGLOBALNODESVERSION
        case "V2":
            result = V2_ALWAYSVISIBLEGLOBALNODESVERSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlwaysVisibleGlobalNodesVersion(values []AlwaysVisibleGlobalNodesVersion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlwaysVisibleGlobalNodesVersion) isMultiValue() bool {
    return false
}
