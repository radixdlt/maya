package models

import (
	"errors"
)

type StreamProofsErrorDetailsType int

const (
	REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMPROOFSERRORDETAILSTYPE StreamProofsErrorDetailsType = iota
	REQUESTEDEPOCHOUTOFBOUNDS_STREAMPROOFSERRORDETAILSTYPE
)

func (i StreamProofsErrorDetailsType) String() string {
	return []string{"RequestedStateVersionOutOfBounds", "RequestedEpochOutOfBounds"}[i]
}

func ParseStreamProofsErrorDetailsType(v string) (any, error) {
	result := REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMPROOFSERRORDETAILSTYPE
	switch v {
	case "RequestedStateVersionOutOfBounds":
		result = REQUESTEDSTATEVERSIONOUTOFBOUNDS_STREAMPROOFSERRORDETAILSTYPE
	case "RequestedEpochOutOfBounds":
		result = REQUESTEDEPOCHOUTOFBOUNDS_STREAMPROOFSERRORDETAILSTYPE
	default:
		return 0, errors.New("Unknown StreamProofsErrorDetailsType value: " + v)
	}
	return &result, nil
}

func SerializeStreamProofsErrorDetailsType(values []StreamProofsErrorDetailsType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i StreamProofsErrorDetailsType) isMultiValue() bool {
	return false
}
