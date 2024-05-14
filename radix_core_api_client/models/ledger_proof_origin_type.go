package models

import (
	"errors"
)

type LedgerProofOriginType int

const (
	GENESIS_LEDGERPROOFORIGINTYPE LedgerProofOriginType = iota
	CONSENSUS_LEDGERPROOFORIGINTYPE
	PROTOCOLUPDATE_LEDGERPROOFORIGINTYPE
)

func (i LedgerProofOriginType) String() string {
	return []string{"Genesis", "Consensus", "ProtocolUpdate"}[i]
}

func ParseLedgerProofOriginType(v string) (any, error) {
	result := GENESIS_LEDGERPROOFORIGINTYPE
	switch v {
	case "Genesis":
		result = GENESIS_LEDGERPROOFORIGINTYPE
	case "Consensus":
		result = CONSENSUS_LEDGERPROOFORIGINTYPE
	case "ProtocolUpdate":
		result = PROTOCOLUPDATE_LEDGERPROOFORIGINTYPE
	default:
		return 0, errors.New("Unknown LedgerProofOriginType value: " + v)
	}
	return &result, nil
}

func SerializeLedgerProofOriginType(values []LedgerProofOriginType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i LedgerProofOriginType) isMultiValue() bool {
	return false
}
