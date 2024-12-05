package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActiveValidatorIndex struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An index of a validator within a specific active validator set (ordered by stake descending). This is the same ordering as used by other parts of the API which list active validators.
    index *int32
}
// NewActiveValidatorIndex instantiates a new ActiveValidatorIndex and sets the default values.
func NewActiveValidatorIndex()(*ActiveValidatorIndex) {
    m := &ActiveValidatorIndex{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActiveValidatorIndexFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActiveValidatorIndexFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActiveValidatorIndex(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActiveValidatorIndex) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActiveValidatorIndex) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. An index of a validator within a specific active validator set (ordered by stake descending). This is the same ordering as used by other parts of the API which list active validators.
// returns a *int32 when successful
func (m *ActiveValidatorIndex) GetIndex()(*int32) {
    return m.index
}
// Serialize serializes information the current object
func (m *ActiveValidatorIndex) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
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
func (m *ActiveValidatorIndex) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndex sets the index property value. An index of a validator within a specific active validator set (ordered by stake descending). This is the same ordering as used by other parts of the API which list active validators.
func (m *ActiveValidatorIndex) SetIndex(value *int32)() {
    m.index = value
}
type ActiveValidatorIndexable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndex()(*int32)
    SetIndex(value *int32)()
}
