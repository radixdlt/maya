package models

import (
	"errors"
)

// The location against which to resolve this type reference.
type LocalTypeId_kind int

const (
	WELLKNOWN_LOCALTYPEID_KIND LocalTypeId_kind = iota
	SCHEMALOCAL_LOCALTYPEID_KIND
)

func (i LocalTypeId_kind) String() string {
	return []string{"WellKnown", "SchemaLocal"}[i]
}

func ParseLocalTypeId_kind(v string) (any, error) {
	result := WELLKNOWN_LOCALTYPEID_KIND
	switch v {
	case "WellKnown":
		result = WELLKNOWN_LOCALTYPEID_KIND
	case "SchemaLocal":
		result = SCHEMALOCAL_LOCALTYPEID_KIND
	default:
		return 0, errors.New("Unknown LocalTypeId_kind value: " + v)
	}
	return &result, nil
}

func SerializeLocalTypeId_kind(values []LocalTypeId_kind) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i LocalTypeId_kind) isMultiValue() bool {
	return false
}
