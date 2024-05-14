package models

import (
	"errors"
)

type NonFungibleIdType int

const (
	STRING_NONFUNGIBLEIDTYPE NonFungibleIdType = iota
	INTEGER_NONFUNGIBLEIDTYPE
	BYTES_NONFUNGIBLEIDTYPE
	RUID_NONFUNGIBLEIDTYPE
)

func (i NonFungibleIdType) String() string {
	return []string{"String", "Integer", "Bytes", "RUID"}[i]
}

func ParseNonFungibleIdType(v string) (any, error) {
	result := STRING_NONFUNGIBLEIDTYPE
	switch v {
	case "String":
		result = STRING_NONFUNGIBLEIDTYPE
	case "Integer":
		result = INTEGER_NONFUNGIBLEIDTYPE
	case "Bytes":
		result = BYTES_NONFUNGIBLEIDTYPE
	case "RUID":
		result = RUID_NONFUNGIBLEIDTYPE
	default:
		return 0, errors.New("Unknown NonFungibleIdType value: " + v)
	}
	return &result, nil
}

func SerializeNonFungibleIdType(values []NonFungibleIdType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i NonFungibleIdType) isMultiValue() bool {
	return false
}
