package models
type PartitionDescriptionType int

const (
    LOGICAL_PARTITIONDESCRIPTIONTYPE PartitionDescriptionType = iota
    PHYSICAL_PARTITIONDESCRIPTIONTYPE
)

func (i PartitionDescriptionType) String() string {
    return []string{"Logical", "Physical"}[i]
}
func ParsePartitionDescriptionType(v string) (any, error) {
    result := LOGICAL_PARTITIONDESCRIPTIONTYPE
    switch v {
        case "Logical":
            result = LOGICAL_PARTITIONDESCRIPTIONTYPE
        case "Physical":
            result = PHYSICAL_PARTITIONDESCRIPTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePartitionDescriptionType(values []PartitionDescriptionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PartitionDescriptionType) isMultiValue() bool {
    return false
}
