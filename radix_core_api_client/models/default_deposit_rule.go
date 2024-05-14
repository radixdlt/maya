package models

import (
	"errors"
)

// This setting has the following interpretations:- Allow: Allows the deposit of all resources - the deny list is honored in this state.- Reject: Disallows the deposit of all resources - the allow list is honored in this state.- AllowExisting: Only deposits of existing resources *or* XRD are accepted - both allow and deny lists are honored in this mode.
type DefaultDepositRule int

const (
	ACCEPT_DEFAULTDEPOSITRULE DefaultDepositRule = iota
	REJECT_DEFAULTDEPOSITRULE
	ALLOWEXISTING_DEFAULTDEPOSITRULE
)

func (i DefaultDepositRule) String() string {
	return []string{"Accept", "Reject", "AllowExisting"}[i]
}

func ParseDefaultDepositRule(v string) (any, error) {
	result := ACCEPT_DEFAULTDEPOSITRULE
	switch v {
	case "Accept":
		result = ACCEPT_DEFAULTDEPOSITRULE
	case "Reject":
		result = REJECT_DEFAULTDEPOSITRULE
	case "AllowExisting":
		result = ALLOWEXISTING_DEFAULTDEPOSITRULE
	default:
		return 0, errors.New("Unknown DefaultDepositRule value: " + v)
	}
	return &result, nil
}

func SerializeDefaultDepositRule(values []DefaultDepositRule) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i DefaultDepositRule) isMultiValue() bool {
	return false
}
