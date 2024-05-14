package models

import (
	"errors"
)

type StreamProofsFilterType int

const (
	ANY_STREAMPROOFSFILTERTYPE StreamProofsFilterType = iota
	NEWEPOCHS_STREAMPROOFSFILTERTYPE
	PROTOCOLUPDATEINITIALIZATIONS_STREAMPROOFSFILTERTYPE
	PROTOCOLUPDATEEXECUTION_STREAMPROOFSFILTERTYPE
)

func (i StreamProofsFilterType) String() string {
	return []string{"Any", "NewEpochs", "ProtocolUpdateInitializations", "ProtocolUpdateExecution"}[i]
}

func ParseStreamProofsFilterType(v string) (any, error) {
	result := ANY_STREAMPROOFSFILTERTYPE
	switch v {
	case "Any":
		result = ANY_STREAMPROOFSFILTERTYPE
	case "NewEpochs":
		result = NEWEPOCHS_STREAMPROOFSFILTERTYPE
	case "ProtocolUpdateInitializations":
		result = PROTOCOLUPDATEINITIALIZATIONS_STREAMPROOFSFILTERTYPE
	case "ProtocolUpdateExecution":
		result = PROTOCOLUPDATEEXECUTION_STREAMPROOFSFILTERTYPE
	default:
		return 0, errors.New("Unknown StreamProofsFilterType value: " + v)
	}
	return &result, nil
}

func SerializeStreamProofsFilterType(values []StreamProofsFilterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i StreamProofsFilterType) isMultiValue() bool {
	return false
}
