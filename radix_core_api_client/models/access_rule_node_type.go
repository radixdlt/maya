package models

import (
	"errors"
)

type AccessRuleNodeType int

const (
	PROOFRULE_ACCESSRULENODETYPE AccessRuleNodeType = iota
	ANYOF_ACCESSRULENODETYPE
	ALLOF_ACCESSRULENODETYPE
)

func (i AccessRuleNodeType) String() string {
	return []string{"ProofRule", "AnyOf", "AllOf"}[i]
}

func ParseAccessRuleNodeType(v string) (any, error) {
	result := PROOFRULE_ACCESSRULENODETYPE
	switch v {
	case "ProofRule":
		result = PROOFRULE_ACCESSRULENODETYPE
	case "AnyOf":
		result = ANYOF_ACCESSRULENODETYPE
	case "AllOf":
		result = ALLOF_ACCESSRULENODETYPE
	default:
		return 0, errors.New("Unknown AccessRuleNodeType value: " + v)
	}
	return &result, nil
}

func SerializeAccessRuleNodeType(values []AccessRuleNodeType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i AccessRuleNodeType) isMultiValue() bool {
	return false
}
