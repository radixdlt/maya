package models

import (
	"errors"
)

type SystemFieldKind int

const (
	BOOTLOADER_SYSTEMFIELDKIND SystemFieldKind = iota
	TYPEINFO_SYSTEMFIELDKIND
)

func (i SystemFieldKind) String() string {
	return []string{"BootLoader", "TypeInfo"}[i]
}

func ParseSystemFieldKind(v string) (any, error) {
	result := BOOTLOADER_SYSTEMFIELDKIND
	switch v {
	case "BootLoader":
		result = BOOTLOADER_SYSTEMFIELDKIND
	case "TypeInfo":
		result = TYPEINFO_SYSTEMFIELDKIND
	default:
		return 0, errors.New("Unknown SystemFieldKind value: " + v)
	}
	return &result, nil
}

func SerializeSystemFieldKind(values []SystemFieldKind) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i SystemFieldKind) isMultiValue() bool {
	return false
}
