package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MempoolListResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contents property
    contents []MempoolTransactionHashesable
}
// NewMempoolListResponse instantiates a new MempoolListResponse and sets the default values.
func NewMempoolListResponse()(*MempoolListResponse) {
    m := &MempoolListResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMempoolListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMempoolListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMempoolListResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MempoolListResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContents gets the contents property value. The contents property
// returns a []MempoolTransactionHashesable when successful
func (m *MempoolListResponse) GetContents()([]MempoolTransactionHashesable) {
    return m.contents
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MempoolListResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMempoolTransactionHashesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MempoolTransactionHashesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MempoolTransactionHashesable)
                }
            }
            m.SetContents(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MempoolListResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContents()))
        for i, v := range m.GetContents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("contents", cast)
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
func (m *MempoolListResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContents sets the contents property value. The contents property
func (m *MempoolListResponse) SetContents(value []MempoolTransactionHashesable)() {
    m.contents = value
}
type MempoolListResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContents()([]MempoolTransactionHashesable)
    SetContents(value []MempoolTransactionHashesable)()
}
