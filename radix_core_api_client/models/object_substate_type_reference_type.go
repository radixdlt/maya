package models

import (
	"errors"
)

type ObjectSubstateTypeReferenceType int

const (
	PACKAGE_OBJECTSUBSTATETYPEREFERENCETYPE ObjectSubstateTypeReferenceType = iota
	OBJECTINSTANCE_OBJECTSUBSTATETYPEREFERENCETYPE
)

func (i ObjectSubstateTypeReferenceType) String() string {
	return []string{"Package", "ObjectInstance"}[i]
}

func ParseObjectSubstateTypeReferenceType(v string) (any, error) {
	result := PACKAGE_OBJECTSUBSTATETYPEREFERENCETYPE
	switch v {
	case "Package":
		result = PACKAGE_OBJECTSUBSTATETYPEREFERENCETYPE
	case "ObjectInstance":
		result = OBJECTINSTANCE_OBJECTSUBSTATETYPEREFERENCETYPE
	default:
		return 0, errors.New("Unknown ObjectSubstateTypeReferenceType value: " + v)
	}
	return &result, nil
}

func SerializeObjectSubstateTypeReferenceType(values []ObjectSubstateTypeReferenceType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ObjectSubstateTypeReferenceType) isMultiValue() bool {
	return false
}
