package models

import (
	"errors"
)

type FieldSchemaFeatureConditionType int

const (
	ALWAYS_FIELDSCHEMAFEATURECONDITIONTYPE FieldSchemaFeatureConditionType = iota
	IFOWNFEATURE_FIELDSCHEMAFEATURECONDITIONTYPE
	IFOUTEROBJECTFEATURE_FIELDSCHEMAFEATURECONDITIONTYPE
)

func (i FieldSchemaFeatureConditionType) String() string {
	return []string{"Always", "IfOwnFeature", "IfOuterObjectFeature"}[i]
}

func ParseFieldSchemaFeatureConditionType(v string) (any, error) {
	result := ALWAYS_FIELDSCHEMAFEATURECONDITIONTYPE
	switch v {
	case "Always":
		result = ALWAYS_FIELDSCHEMAFEATURECONDITIONTYPE
	case "IfOwnFeature":
		result = IFOWNFEATURE_FIELDSCHEMAFEATURECONDITIONTYPE
	case "IfOuterObjectFeature":
		result = IFOUTEROBJECTFEATURE_FIELDSCHEMAFEATURECONDITIONTYPE
	default:
		return 0, errors.New("Unknown FieldSchemaFeatureConditionType value: " + v)
	}
	return &result, nil
}

func SerializeFieldSchemaFeatureConditionType(values []FieldSchemaFeatureConditionType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i FieldSchemaFeatureConditionType) isMultiValue() bool {
	return false
}
