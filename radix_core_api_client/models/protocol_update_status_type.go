package models
type ProtocolUpdateStatusType int

const (
    INPROGRESS_PROTOCOLUPDATESTATUSTYPE ProtocolUpdateStatusType = iota
    COMPLETE_PROTOCOLUPDATESTATUSTYPE
)

func (i ProtocolUpdateStatusType) String() string {
    return []string{"InProgress", "Complete"}[i]
}
func ParseProtocolUpdateStatusType(v string) (any, error) {
    result := INPROGRESS_PROTOCOLUPDATESTATUSTYPE
    switch v {
        case "InProgress":
            result = INPROGRESS_PROTOCOLUPDATESTATUSTYPE
        case "Complete":
            result = COMPLETE_PROTOCOLUPDATESTATUSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProtocolUpdateStatusType(values []ProtocolUpdateStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProtocolUpdateStatusType) isMultiValue() bool {
    return false
}
