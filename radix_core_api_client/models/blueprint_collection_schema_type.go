package models
type BlueprintCollectionSchemaType int

const (
    KEYVALUE_BLUEPRINTCOLLECTIONSCHEMATYPE BlueprintCollectionSchemaType = iota
    INDEX_BLUEPRINTCOLLECTIONSCHEMATYPE
    SORTEDINDEX_BLUEPRINTCOLLECTIONSCHEMATYPE
)

func (i BlueprintCollectionSchemaType) String() string {
    return []string{"KeyValue", "Index", "SortedIndex"}[i]
}
func ParseBlueprintCollectionSchemaType(v string) (any, error) {
    result := KEYVALUE_BLUEPRINTCOLLECTIONSCHEMATYPE
    switch v {
        case "KeyValue":
            result = KEYVALUE_BLUEPRINTCOLLECTIONSCHEMATYPE
        case "Index":
            result = INDEX_BLUEPRINTCOLLECTIONSCHEMATYPE
        case "SortedIndex":
            result = SORTEDINDEX_BLUEPRINTCOLLECTIONSCHEMATYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBlueprintCollectionSchemaType(values []BlueprintCollectionSchemaType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BlueprintCollectionSchemaType) isMultiValue() bool {
    return false
}
