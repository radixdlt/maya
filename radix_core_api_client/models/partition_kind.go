package models
// The type of the partition:- Field  - Has key: TupleKey(u8) also known as an offset  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- KeyValue ("ConcurrentMap")  - Has key: MapKey(Vec<u8>)  - No iteration exposed to engine  - Is versioned / locked substate-by-substate- Index  - Has key: MapKey(Vec<u8>)  - Iteration exposed to engine via the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety- SortedIndex  - Has key: SortedU16Key(U16, Vec<u8>)  - Iteration exposed to engine via the user-controlled U16 prefix and then the MapKey's database key (ie hash of the key)  - Is versioned / locked in its entirety
type PartitionKind int

const (
    FIELD_PARTITIONKIND PartitionKind = iota
    KEYVALUE_PARTITIONKIND
    INDEX_PARTITIONKIND
    SORTEDINDEX_PARTITIONKIND
)

func (i PartitionKind) String() string {
    return []string{"Field", "KeyValue", "Index", "SortedIndex"}[i]
}
func ParsePartitionKind(v string) (any, error) {
    result := FIELD_PARTITIONKIND
    switch v {
        case "Field":
            result = FIELD_PARTITIONKIND
        case "KeyValue":
            result = KEYVALUE_PARTITIONKIND
        case "Index":
            result = INDEX_PARTITIONKIND
        case "SortedIndex":
            result = SORTEDINDEX_PARTITIONKIND
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePartitionKind(values []PartitionKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PartitionKind) isMultiValue() bool {
    return false
}
