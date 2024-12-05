package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleLocalId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id_type property
    id_type *NonFungibleIdType
    // The hex-encoded SBOR-encoded bytes of its non-fungible id
    sbor_hex *string
    // The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
    simple_rep *string
}
// NewNonFungibleLocalId instantiates a new NonFungibleLocalId and sets the default values.
func NewNonFungibleLocalId()(*NonFungibleLocalId) {
    m := &NonFungibleLocalId{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNonFungibleLocalIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleLocalIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleLocalId(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleLocalId) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleLocalId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonFungibleIdType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdType(val.(*NonFungibleIdType))
        }
        return nil
    }
    res["sbor_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSborHex(val)
        }
        return nil
    }
    res["simple_rep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimpleRep(val)
        }
        return nil
    }
    return res
}
// GetIdType gets the id_type property value. The id_type property
// returns a *NonFungibleIdType when successful
func (m *NonFungibleLocalId) GetIdType()(*NonFungibleIdType) {
    return m.id_type
}
// GetSborHex gets the sbor_hex property value. The hex-encoded SBOR-encoded bytes of its non-fungible id
// returns a *string when successful
func (m *NonFungibleLocalId) GetSborHex()(*string) {
    return m.sbor_hex
}
// GetSimpleRep gets the simple_rep property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
// returns a *string when successful
func (m *NonFungibleLocalId) GetSimpleRep()(*string) {
    return m.simple_rep
}
// Serialize serializes information the current object
func (m *NonFungibleLocalId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIdType() != nil {
        cast := (*m.GetIdType()).String()
        err := writer.WriteStringValue("id_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sbor_hex", m.GetSborHex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("simple_rep", m.GetSimpleRep())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NonFungibleLocalId) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdType sets the id_type property value. The id_type property
func (m *NonFungibleLocalId) SetIdType(value *NonFungibleIdType)() {
    m.id_type = value
}
// SetSborHex sets the sbor_hex property value. The hex-encoded SBOR-encoded bytes of its non-fungible id
func (m *NonFungibleLocalId) SetSborHex(value *string)() {
    m.sbor_hex = value
}
// SetSimpleRep sets the simple_rep property value. The simple string representation of the non-fungible id.* For string ids, this is `<the-string-id>`* For integer ids, this is `#the-integer-id#`* For bytes ids, this is `[the-lower-case-hex-representation]`* For RUID ids, this is `{...-...-...-...}` where `...` are each 16 hex characters.A given non-fungible resource has a fixed `NonFungibleIdType`, so this representation uniquely identifies this non-fungibleunder the given resource address.
func (m *NonFungibleLocalId) SetSimpleRep(value *string)() {
    m.simple_rep = value
}
type NonFungibleLocalIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdType()(*NonFungibleIdType)
    GetSborHex()(*string)
    GetSimpleRep()(*string)
    SetIdType(value *NonFungibleIdType)()
    SetSborHex(value *string)()
    SetSimpleRep(value *string)()
}
