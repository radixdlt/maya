package models

import (
	"errors"
)

type VmType int

const (
	NATIVE_VMTYPE VmType = iota
	SCRYPTOV1_VMTYPE
)

func (i VmType) String() string {
	return []string{"Native", "ScryptoV1"}[i]
}

func ParseVmType(v string) (any, error) {
	result := NATIVE_VMTYPE
	switch v {
	case "Native":
		result = NATIVE_VMTYPE
	case "ScryptoV1":
		result = SCRYPTOV1_VMTYPE
	default:
		return 0, errors.New("Unknown VmType value: " + v)
	}
	return &result, nil
}

func SerializeVmType(values []VmType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i VmType) isMultiValue() bool {
	return false
}
