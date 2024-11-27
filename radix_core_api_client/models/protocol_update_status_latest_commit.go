package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtocolUpdateStatusLatestCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The batch_group_index property
    batch_group_index *int64
    // The batch_group_name property
    batch_group_name *string
    // The batch_index property
    batch_index *int64
    // The batch_name property
    batch_name *string
}
// NewProtocolUpdateStatusLatestCommit instantiates a new ProtocolUpdateStatusLatestCommit and sets the default values.
func NewProtocolUpdateStatusLatestCommit()(*ProtocolUpdateStatusLatestCommit) {
    m := &ProtocolUpdateStatusLatestCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtocolUpdateStatusLatestCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtocolUpdateStatusLatestCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtocolUpdateStatusLatestCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtocolUpdateStatusLatestCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBatchGroupIndex gets the batch_group_index property value. The batch_group_index property
// returns a *int64 when successful
func (m *ProtocolUpdateStatusLatestCommit) GetBatchGroupIndex()(*int64) {
    return m.batch_group_index
}
// GetBatchGroupName gets the batch_group_name property value. The batch_group_name property
// returns a *string when successful
func (m *ProtocolUpdateStatusLatestCommit) GetBatchGroupName()(*string) {
    return m.batch_group_name
}
// GetBatchIndex gets the batch_index property value. The batch_index property
// returns a *int64 when successful
func (m *ProtocolUpdateStatusLatestCommit) GetBatchIndex()(*int64) {
    return m.batch_index
}
// GetBatchName gets the batch_name property value. The batch_name property
// returns a *string when successful
func (m *ProtocolUpdateStatusLatestCommit) GetBatchName()(*string) {
    return m.batch_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtocolUpdateStatusLatestCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batch_group_index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchGroupIndex(val)
        }
        return nil
    }
    res["batch_group_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchGroupName(val)
        }
        return nil
    }
    res["batch_index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchIndex(val)
        }
        return nil
    }
    res["batch_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtocolUpdateStatusLatestCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("batch_group_index", m.GetBatchGroupIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("batch_group_name", m.GetBatchGroupName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("batch_index", m.GetBatchIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("batch_name", m.GetBatchName())
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
func (m *ProtocolUpdateStatusLatestCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBatchGroupIndex sets the batch_group_index property value. The batch_group_index property
func (m *ProtocolUpdateStatusLatestCommit) SetBatchGroupIndex(value *int64)() {
    m.batch_group_index = value
}
// SetBatchGroupName sets the batch_group_name property value. The batch_group_name property
func (m *ProtocolUpdateStatusLatestCommit) SetBatchGroupName(value *string)() {
    m.batch_group_name = value
}
// SetBatchIndex sets the batch_index property value. The batch_index property
func (m *ProtocolUpdateStatusLatestCommit) SetBatchIndex(value *int64)() {
    m.batch_index = value
}
// SetBatchName sets the batch_name property value. The batch_name property
func (m *ProtocolUpdateStatusLatestCommit) SetBatchName(value *string)() {
    m.batch_name = value
}
type ProtocolUpdateStatusLatestCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatchGroupIndex()(*int64)
    GetBatchGroupName()(*string)
    GetBatchIndex()(*int64)
    GetBatchName()(*string)
    SetBatchGroupIndex(value *int64)()
    SetBatchGroupName(value *string)()
    SetBatchIndex(value *int64)()
    SetBatchName(value *string)()
}
