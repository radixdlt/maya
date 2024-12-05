package models
type TypeInfoType int

const (
    OBJECT_TYPEINFOTYPE TypeInfoType = iota
    KEYVALUESTORE_TYPEINFOTYPE
)

func (i TypeInfoType) String() string {
    return []string{"Object", "KeyValueStore"}[i]
}
func ParseTypeInfoType(v string) (any, error) {
    result := OBJECT_TYPEINFOTYPE
    switch v {
        case "Object":
            result = OBJECT_TYPEINFOTYPE
        case "KeyValueStore":
            result = KEYVALUESTORE_TYPEINFOTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTypeInfoType(values []TypeInfoType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TypeInfoType) isMultiValue() bool {
    return false
}
