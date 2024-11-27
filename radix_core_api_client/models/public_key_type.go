package models
type PublicKeyType int

const (
    ECDSASECP256K1_PUBLICKEYTYPE PublicKeyType = iota
    EDDSAED25519_PUBLICKEYTYPE
)

func (i PublicKeyType) String() string {
    return []string{"EcdsaSecp256k1", "EddsaEd25519"}[i]
}
func ParsePublicKeyType(v string) (any, error) {
    result := ECDSASECP256K1_PUBLICKEYTYPE
    switch v {
        case "EcdsaSecp256k1":
            result = ECDSASECP256K1_PUBLICKEYTYPE
        case "EddsaEd25519":
            result = EDDSAED25519_PUBLICKEYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePublicKeyType(values []PublicKeyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PublicKeyType) isMultiValue() bool {
    return false
}
