package models

import (
	"errors"
)

type ProofRuleType int

const (
	REQUIRE_PROOFRULETYPE ProofRuleType = iota
	AMOUNTOF_PROOFRULETYPE
	ALLOF_PROOFRULETYPE
	ANYOF_PROOFRULETYPE
	COUNTOF_PROOFRULETYPE
)

func (i ProofRuleType) String() string {
	return []string{"Require", "AmountOf", "AllOf", "AnyOf", "CountOf"}[i]
}

func ParseProofRuleType(v string) (any, error) {
	result := REQUIRE_PROOFRULETYPE
	switch v {
	case "Require":
		result = REQUIRE_PROOFRULETYPE
	case "AmountOf":
		result = AMOUNTOF_PROOFRULETYPE
	case "AllOf":
		result = ALLOF_PROOFRULETYPE
	case "AnyOf":
		result = ANYOF_PROOFRULETYPE
	case "CountOf":
		result = COUNTOF_PROOFRULETYPE
	default:
		return 0, errors.New("Unknown ProofRuleType value: " + v)
	}
	return &result, nil
}

func SerializeProofRuleType(values []ProofRuleType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i ProofRuleType) isMultiValue() bool {
	return false
}
