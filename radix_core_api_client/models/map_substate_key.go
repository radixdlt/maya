package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MapSubstateKey struct {
    SubstateKey
    // The hex-encoded bytes of the substate key
    key_hex *string
}
// NewMapSubstateKey instantiates a new MapSubstateKey and sets the default values.
func NewMapSubstateKey()(*MapSubstateKey) {
    m := &MapSubstateKey{
        SubstateKey: *NewSubstateKey(),
    }
    return m
}
// CreateMapSubstateKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMapSubstateKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMapSubstateKey(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MapSubstateKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubstateKey.GetFieldDeserializers()
    res["key_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyHex(val)
        }
        return nil
    }
    return res
}
// GetKeyHex gets the key_hex property value. The hex-encoded bytes of the substate key
// returns a *string when successful
func (m *MapSubstateKey) GetKeyHex()(*string) {
    return m.key_hex
}
// Serialize serializes information the current object
func (m *MapSubstateKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubstateKey.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("key_hex", m.GetKeyHex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyHex sets the key_hex property value. The hex-encoded bytes of the substate key
func (m *MapSubstateKey) SetKeyHex(value *string)() {
    m.key_hex = value
}
type MapSubstateKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubstateKeyable
    GetKeyHex()(*string)
    SetKeyHex(value *string)()
}
