package models
type SystemFieldKind int

const (
    VMBOOT_SYSTEMFIELDKIND SystemFieldKind = iota
    TYPEINFO_SYSTEMFIELDKIND
    SYSTEMBOOT_SYSTEMFIELDKIND
    KERNELBOOT_SYSTEMFIELDKIND
    TRANSACTIONVALIDATIONCONFIGURATION_SYSTEMFIELDKIND
    PROTOCOLUPDATESTATUSSUMMARY_SYSTEMFIELDKIND
)

func (i SystemFieldKind) String() string {
    return []string{"VmBoot", "TypeInfo", "SystemBoot", "KernelBoot", "TransactionValidationConfiguration", "ProtocolUpdateStatusSummary"}[i]
}
func ParseSystemFieldKind(v string) (any, error) {
    result := VMBOOT_SYSTEMFIELDKIND
    switch v {
        case "VmBoot":
            result = VMBOOT_SYSTEMFIELDKIND
        case "TypeInfo":
            result = TYPEINFO_SYSTEMFIELDKIND
        case "SystemBoot":
            result = SYSTEMBOOT_SYSTEMFIELDKIND
        case "KernelBoot":
            result = KERNELBOOT_SYSTEMFIELDKIND
        case "TransactionValidationConfiguration":
            result = TRANSACTIONVALIDATIONCONFIGURATION_SYSTEMFIELDKIND
        case "ProtocolUpdateStatusSummary":
            result = PROTOCOLUPDATESTATUSSUMMARY_SYSTEMFIELDKIND
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemFieldKind(values []SystemFieldKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemFieldKind) isMultiValue() bool {
    return false
}
