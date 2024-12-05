package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtocolUpdateLedgerProofOrigin represents a proof from the execution of a non-genesis protocol update.The execution of a protocol update is organised into batch groups, andthen these batch groups are organised into batches, with each batch committedatomically.NOTE: Some of these values may be placeholder values for protocol updates pre-Cuttlefishon nodes which haven't resynced since Cuttlefish. In particular, the following values might beinvalid on such nodes:* `config_hash` (placeholder of all zeros)* `batch_group_idx` (placeholder of 0)* `batch_group_name` (placeholder of "")* `batch_idx` (placeholder of 0)* `batch_name` (placeholder of "")* `is_end_of_update` (placeholder of false)
type ProtocolUpdateLedgerProofOrigin struct {
    LedgerProofOrigin
    // The batch_group_idx property
    batch_group_idx *int64
    // The batch_group_name property
    batch_group_name *string
    // The batch_idx property
    batch_idx *int64
    // The batch_name property
    batch_name *string
    // The config_hash property
    config_hash *string
    // The is_end_of_update property
    is_end_of_update *bool
    // The protocol_version_name property
    protocol_version_name *string
}
// NewProtocolUpdateLedgerProofOrigin instantiates a new ProtocolUpdateLedgerProofOrigin and sets the default values.
func NewProtocolUpdateLedgerProofOrigin()(*ProtocolUpdateLedgerProofOrigin) {
    m := &ProtocolUpdateLedgerProofOrigin{
        LedgerProofOrigin: *NewLedgerProofOrigin(),
    }
    return m
}
// CreateProtocolUpdateLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtocolUpdateLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtocolUpdateLedgerProofOrigin(), nil
}
// GetBatchGroupIdx gets the batch_group_idx property value. The batch_group_idx property
// returns a *int64 when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetBatchGroupIdx()(*int64) {
    return m.batch_group_idx
}
// GetBatchGroupName gets the batch_group_name property value. The batch_group_name property
// returns a *string when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetBatchGroupName()(*string) {
    return m.batch_group_name
}
// GetBatchIdx gets the batch_idx property value. The batch_idx property
// returns a *int64 when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetBatchIdx()(*int64) {
    return m.batch_idx
}
// GetBatchName gets the batch_name property value. The batch_name property
// returns a *string when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetBatchName()(*string) {
    return m.batch_name
}
// GetConfigHash gets the config_hash property value. The config_hash property
// returns a *string when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetConfigHash()(*string) {
    return m.config_hash
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LedgerProofOrigin.GetFieldDeserializers()
    res["batch_group_idx"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchGroupIdx(val)
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
    res["batch_idx"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatchIdx(val)
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
    res["config_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigHash(val)
        }
        return nil
    }
    res["is_end_of_update"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndOfUpdate(val)
        }
        return nil
    }
    res["protocol_version_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocolVersionName(val)
        }
        return nil
    }
    return res
}
// GetIsEndOfUpdate gets the is_end_of_update property value. The is_end_of_update property
// returns a *bool when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetIsEndOfUpdate()(*bool) {
    return m.is_end_of_update
}
// GetProtocolVersionName gets the protocol_version_name property value. The protocol_version_name property
// returns a *string when successful
func (m *ProtocolUpdateLedgerProofOrigin) GetProtocolVersionName()(*string) {
    return m.protocol_version_name
}
// Serialize serializes information the current object
func (m *ProtocolUpdateLedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LedgerProofOrigin.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("batch_group_idx", m.GetBatchGroupIdx())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("batch_group_name", m.GetBatchGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("batch_idx", m.GetBatchIdx())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("batch_name", m.GetBatchName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("config_hash", m.GetConfigHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("is_end_of_update", m.GetIsEndOfUpdate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("protocol_version_name", m.GetProtocolVersionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBatchGroupIdx sets the batch_group_idx property value. The batch_group_idx property
func (m *ProtocolUpdateLedgerProofOrigin) SetBatchGroupIdx(value *int64)() {
    m.batch_group_idx = value
}
// SetBatchGroupName sets the batch_group_name property value. The batch_group_name property
func (m *ProtocolUpdateLedgerProofOrigin) SetBatchGroupName(value *string)() {
    m.batch_group_name = value
}
// SetBatchIdx sets the batch_idx property value. The batch_idx property
func (m *ProtocolUpdateLedgerProofOrigin) SetBatchIdx(value *int64)() {
    m.batch_idx = value
}
// SetBatchName sets the batch_name property value. The batch_name property
func (m *ProtocolUpdateLedgerProofOrigin) SetBatchName(value *string)() {
    m.batch_name = value
}
// SetConfigHash sets the config_hash property value. The config_hash property
func (m *ProtocolUpdateLedgerProofOrigin) SetConfigHash(value *string)() {
    m.config_hash = value
}
// SetIsEndOfUpdate sets the is_end_of_update property value. The is_end_of_update property
func (m *ProtocolUpdateLedgerProofOrigin) SetIsEndOfUpdate(value *bool)() {
    m.is_end_of_update = value
}
// SetProtocolVersionName sets the protocol_version_name property value. The protocol_version_name property
func (m *ProtocolUpdateLedgerProofOrigin) SetProtocolVersionName(value *string)() {
    m.protocol_version_name = value
}
type ProtocolUpdateLedgerProofOriginable interface {
    LedgerProofOriginable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatchGroupIdx()(*int64)
    GetBatchGroupName()(*string)
    GetBatchIdx()(*int64)
    GetBatchName()(*string)
    GetConfigHash()(*string)
    GetIsEndOfUpdate()(*bool)
    GetProtocolVersionName()(*string)
    SetBatchGroupIdx(value *int64)()
    SetBatchGroupName(value *string)()
    SetBatchIdx(value *int64)()
    SetBatchName(value *string)()
    SetConfigHash(value *string)()
    SetIsEndOfUpdate(value *bool)()
    SetProtocolVersionName(value *string)()
}
