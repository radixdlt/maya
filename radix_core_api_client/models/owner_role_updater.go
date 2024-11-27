package models
type OwnerRoleUpdater int

const (
    NONE_OWNERROLEUPDATER OwnerRoleUpdater = iota
    OWNER_OWNERROLEUPDATER
    OBJECT_OWNERROLEUPDATER
)

func (i OwnerRoleUpdater) String() string {
    return []string{"None", "Owner", "Object"}[i]
}
func ParseOwnerRoleUpdater(v string) (any, error) {
    result := NONE_OWNERROLEUPDATER
    switch v {
        case "None":
            result = NONE_OWNERROLEUPDATER
        case "Owner":
            result = OWNER_OWNERROLEUPDATER
        case "Object":
            result = OBJECT_OWNERROLEUPDATER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOwnerRoleUpdater(values []OwnerRoleUpdater) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnerRoleUpdater) isMultiValue() bool {
    return false
}
