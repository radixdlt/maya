package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AddressType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address_byte_length property
    address_byte_length *int32
    // The address_byte_prefix property
    address_byte_prefix *int32
    // To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
    entity_type *EntityType
    // The hrp_prefix property
    hrp_prefix *string
}
// NewAddressType instantiates a new AddressType and sets the default values.
func NewAddressType()(*AddressType) {
    m := &AddressType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAddressTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAddressTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddressType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AddressType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddressByteLength gets the address_byte_length property value. The address_byte_length property
// returns a *int32 when successful
func (m *AddressType) GetAddressByteLength()(*int32) {
    return m.address_byte_length
}
// GetAddressBytePrefix gets the address_byte_prefix property value. The address_byte_prefix property
// returns a *int32 when successful
func (m *AddressType) GetAddressBytePrefix()(*int32) {
    return m.address_byte_prefix
}
// GetEntityType gets the entity_type property value. To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
// returns a *EntityType when successful
func (m *AddressType) GetEntityType()(*EntityType) {
    return m.entity_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AddressType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address_byte_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressByteLength(val)
        }
        return nil
    }
    res["address_byte_prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressBytePrefix(val)
        }
        return nil
    }
    res["entity_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val.(*EntityType))
        }
        return nil
    }
    res["hrp_prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHrpPrefix(val)
        }
        return nil
    }
    return res
}
// GetHrpPrefix gets the hrp_prefix property value. The hrp_prefix property
// returns a *string when successful
func (m *AddressType) GetHrpPrefix()(*string) {
    return m.hrp_prefix
}
// Serialize serializes information the current object
func (m *AddressType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("address_byte_length", m.GetAddressByteLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("address_byte_prefix", m.GetAddressBytePrefix())
        if err != nil {
            return err
        }
    }
    if m.GetEntityType() != nil {
        cast := (*m.GetEntityType()).String()
        err := writer.WriteStringValue("entity_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hrp_prefix", m.GetHrpPrefix())
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
func (m *AddressType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddressByteLength sets the address_byte_length property value. The address_byte_length property
func (m *AddressType) SetAddressByteLength(value *int32)() {
    m.address_byte_length = value
}
// SetAddressBytePrefix sets the address_byte_prefix property value. The address_byte_prefix property
func (m *AddressType) SetAddressBytePrefix(value *int32)() {
    m.address_byte_prefix = value
}
// SetEntityType sets the entity_type property value. To improve clarity, some `EntityType` names have been updated at Cuttlefish.But, due to backwards compatibility, the old names continue to be used here.Notably:* `GlobalVirtualSecp256k1Account` is now `GlobalPreallocatedSecp256k1Account` elsewhere* `GlobalVirtualSecp256k1Identity` is now `GlobalPreallocatedSecp256k1Identity` elsewhere* `GlobalVirtualEd25519Account` is now `GlobalPreallocatedEd25519Account` elsewhere* `GlobalVirtualEd25519Identity` is now `GlobalPreallocatedEd25519Identity` elsewhere
func (m *AddressType) SetEntityType(value *EntityType)() {
    m.entity_type = value
}
// SetHrpPrefix sets the hrp_prefix property value. The hrp_prefix property
func (m *AddressType) SetHrpPrefix(value *string)() {
    m.hrp_prefix = value
}
type AddressTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressByteLength()(*int32)
    GetAddressBytePrefix()(*int32)
    GetEntityType()(*EntityType)
    GetHrpPrefix()(*string)
    SetAddressByteLength(value *int32)()
    SetAddressBytePrefix(value *int32)()
    SetEntityType(value *EntityType)()
    SetHrpPrefix(value *string)()
}
