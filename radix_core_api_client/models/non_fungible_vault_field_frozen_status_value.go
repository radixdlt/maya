package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NonFungibleVaultFieldFrozenStatusValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The frozen_status property
    frozen_status FrozenStatusable
}
// NewNonFungibleVaultFieldFrozenStatusValue instantiates a new NonFungibleVaultFieldFrozenStatusValue and sets the default values.
func NewNonFungibleVaultFieldFrozenStatusValue()(*NonFungibleVaultFieldFrozenStatusValue) {
    m := &NonFungibleVaultFieldFrozenStatusValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNonFungibleVaultFieldFrozenStatusValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNonFungibleVaultFieldFrozenStatusValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNonFungibleVaultFieldFrozenStatusValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NonFungibleVaultFieldFrozenStatusValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NonFungibleVaultFieldFrozenStatusValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["frozen_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFrozenStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrozenStatus(val.(FrozenStatusable))
        }
        return nil
    }
    return res
}
// GetFrozenStatus gets the frozen_status property value. The frozen_status property
// returns a FrozenStatusable when successful
func (m *NonFungibleVaultFieldFrozenStatusValue) GetFrozenStatus()(FrozenStatusable) {
    return m.frozen_status
}
// Serialize serializes information the current object
func (m *NonFungibleVaultFieldFrozenStatusValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("frozen_status", m.GetFrozenStatus())
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
func (m *NonFungibleVaultFieldFrozenStatusValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFrozenStatus sets the frozen_status property value. The frozen_status property
func (m *NonFungibleVaultFieldFrozenStatusValue) SetFrozenStatus(value FrozenStatusable)() {
    m.frozen_status = value
}
type NonFungibleVaultFieldFrozenStatusValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFrozenStatus()(FrozenStatusable)
    SetFrozenStatus(value FrozenStatusable)()
}
