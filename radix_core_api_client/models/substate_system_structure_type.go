package models

import (
	"errors"
)

type SubstateSystemStructureType int

const (
	SYSTEMFIELD_SUBSTATESYSTEMSTRUCTURETYPE SubstateSystemStructureType = iota
	SYSTEMSCHEMA_SUBSTATESYSTEMSTRUCTURETYPE
	KEYVALUESTOREENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	OBJECTFIELD_SUBSTATESYSTEMSTRUCTURETYPE
	OBJECTKEYVALUEPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	OBJECTINDEXPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	OBJECTSORTEDINDEXPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
)

func (i SubstateSystemStructureType) String() string {
	return []string{"SystemField", "SystemSchema", "KeyValueStoreEntry", "ObjectField", "ObjectKeyValuePartitionEntry", "ObjectIndexPartitionEntry", "ObjectSortedIndexPartitionEntry"}[i]
}

func ParseSubstateSystemStructureType(v string) (any, error) {
	result := SYSTEMFIELD_SUBSTATESYSTEMSTRUCTURETYPE
	switch v {
	case "SystemField":
		result = SYSTEMFIELD_SUBSTATESYSTEMSTRUCTURETYPE
	case "SystemSchema":
		result = SYSTEMSCHEMA_SUBSTATESYSTEMSTRUCTURETYPE
	case "KeyValueStoreEntry":
		result = KEYVALUESTOREENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	case "ObjectField":
		result = OBJECTFIELD_SUBSTATESYSTEMSTRUCTURETYPE
	case "ObjectKeyValuePartitionEntry":
		result = OBJECTKEYVALUEPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	case "ObjectIndexPartitionEntry":
		result = OBJECTINDEXPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	case "ObjectSortedIndexPartitionEntry":
		result = OBJECTSORTEDINDEXPARTITIONENTRY_SUBSTATESYSTEMSTRUCTURETYPE
	default:
		return 0, errors.New("Unknown SubstateSystemStructureType value: " + v)
	}
	return &result, nil
}

func SerializeSubstateSystemStructureType(values []SubstateSystemStructureType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i SubstateSystemStructureType) isMultiValue() bool {
	return false
}
