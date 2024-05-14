package models

import (
	"errors"
)

type PlaintextMessageContentType int

const (
	STRING_PLAINTEXTMESSAGECONTENTTYPE PlaintextMessageContentType = iota
	BINARY_PLAINTEXTMESSAGECONTENTTYPE
)

func (i PlaintextMessageContentType) String() string {
	return []string{"String", "Binary"}[i]
}

func ParsePlaintextMessageContentType(v string) (any, error) {
	result := STRING_PLAINTEXTMESSAGECONTENTTYPE
	switch v {
	case "String":
		result = STRING_PLAINTEXTMESSAGECONTENTTYPE
	case "Binary":
		result = BINARY_PLAINTEXTMESSAGECONTENTTYPE
	default:
		return 0, errors.New("Unknown PlaintextMessageContentType value: " + v)
	}
	return &result, nil
}

func SerializePlaintextMessageContentType(values []PlaintextMessageContentType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i PlaintextMessageContentType) isMultiValue() bool {
	return false
}
