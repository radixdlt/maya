package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActiveValidator struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Bech32m-encoded human readable version of the component address
    address *string
    // The key property
    key EcdsaSecp256k1PublicKeyable
    // A string-encoded decimal representing the validator's voting power for this epoch.This is a snapshot of the amount of XRD staked to the validator at the start of the epoch.
    stake *string
}
// NewActiveValidator instantiates a new ActiveValidator and sets the default values.
func NewActiveValidator()(*ActiveValidator) {
    m := &ActiveValidator{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActiveValidatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActiveValidatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActiveValidator(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActiveValidator) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The Bech32m-encoded human readable version of the component address
// returns a *string when successful
func (m *ActiveValidator) GetAddress()(*string) {
    return m.address
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActiveValidator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEcdsaSecp256k1PublicKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(EcdsaSecp256k1PublicKeyable))
        }
        return nil
    }
    res["stake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStake(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a EcdsaSecp256k1PublicKeyable when successful
func (m *ActiveValidator) GetKey()(EcdsaSecp256k1PublicKeyable) {
    return m.key
}
// GetStake gets the stake property value. A string-encoded decimal representing the validator's voting power for this epoch.This is a snapshot of the amount of XRD staked to the validator at the start of the epoch.
// returns a *string when successful
func (m *ActiveValidator) GetStake()(*string) {
    return m.stake
}
// Serialize serializes information the current object
func (m *ActiveValidator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stake", m.GetStake())
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
func (m *ActiveValidator) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The Bech32m-encoded human readable version of the component address
func (m *ActiveValidator) SetAddress(value *string)() {
    m.address = value
}
// SetKey sets the key property value. The key property
func (m *ActiveValidator) SetKey(value EcdsaSecp256k1PublicKeyable)() {
    m.key = value
}
// SetStake sets the stake property value. A string-encoded decimal representing the validator's voting power for this epoch.This is a snapshot of the amount of XRD staked to the validator at the start of the epoch.
func (m *ActiveValidator) SetStake(value *string)() {
    m.stake = value
}
type ActiveValidatorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetKey()(EcdsaSecp256k1PublicKeyable)
    GetStake()(*string)
    SetAddress(value *string)()
    SetKey(value EcdsaSecp256k1PublicKeyable)()
    SetStake(value *string)()
}
