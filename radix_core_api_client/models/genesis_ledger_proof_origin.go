package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GenesisLedgerProofOrigin represents a proof from the execution of the babylon genesis protocol update, which startsthe babylon-node ledger.Behind-the-scenes, this is now the same as a `ProtocolUpdateLedgerProofOrigin`,but is kept separate for backwards-compatibility.NOTE: Some of these values may be placeholder values on nodes which haven't resynced sinceCuttlefish. In particular, the following values might be invalid on such nodes:* `batch_group_idx` (placeholder of 0)* `batch_group_name` (placeholder of "")* `batch_idx` (placeholder of 0)* `batch_name` (placeholder of "")* `is_end_of_update` (placeholder of false)
type GenesisLedgerProofOrigin struct {
    LedgerProofOrigin
    // The batch_group_idx property
    batch_group_idx *int64
    // The batch_group_name property
    batch_group_name *string
    // The batch_idx property
    batch_idx *int64
    // The batch_name property
    batch_name *string
    // The genesis_opaque_hash property
    genesis_opaque_hash *string
    // The is_end_of_update property
    is_end_of_update *bool
    // The protocol_version_name property
    protocol_version_name *string
}
// NewGenesisLedgerProofOrigin instantiates a new GenesisLedgerProofOrigin and sets the default values.
func NewGenesisLedgerProofOrigin()(*GenesisLedgerProofOrigin) {
    m := &GenesisLedgerProofOrigin{
        LedgerProofOrigin: *NewLedgerProofOrigin(),
    }
    return m
}
// CreateGenesisLedgerProofOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenesisLedgerProofOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenesisLedgerProofOrigin(), nil
}
// GetBatchGroupIdx gets the batch_group_idx property value. The batch_group_idx property
// returns a *int64 when successful
func (m *GenesisLedgerProofOrigin) GetBatchGroupIdx()(*int64) {
    return m.batch_group_idx
}
// GetBatchGroupName gets the batch_group_name property value. The batch_group_name property
// returns a *string when successful
func (m *GenesisLedgerProofOrigin) GetBatchGroupName()(*string) {
    return m.batch_group_name
}
// GetBatchIdx gets the batch_idx property value. The batch_idx property
// returns a *int64 when successful
func (m *GenesisLedgerProofOrigin) GetBatchIdx()(*int64) {
    return m.batch_idx
}
// GetBatchName gets the batch_name property value. The batch_name property
// returns a *string when successful
func (m *GenesisLedgerProofOrigin) GetBatchName()(*string) {
    return m.batch_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenesisLedgerProofOrigin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["genesis_opaque_hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGenesisOpaqueHash(val)
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
// GetGenesisOpaqueHash gets the genesis_opaque_hash property value. The genesis_opaque_hash property
// returns a *string when successful
func (m *GenesisLedgerProofOrigin) GetGenesisOpaqueHash()(*string) {
    return m.genesis_opaque_hash
}
// GetIsEndOfUpdate gets the is_end_of_update property value. The is_end_of_update property
// returns a *bool when successful
func (m *GenesisLedgerProofOrigin) GetIsEndOfUpdate()(*bool) {
    return m.is_end_of_update
}
// GetProtocolVersionName gets the protocol_version_name property value. The protocol_version_name property
// returns a *string when successful
func (m *GenesisLedgerProofOrigin) GetProtocolVersionName()(*string) {
    return m.protocol_version_name
}
// Serialize serializes information the current object
func (m *GenesisLedgerProofOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("genesis_opaque_hash", m.GetGenesisOpaqueHash())
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
func (m *GenesisLedgerProofOrigin) SetBatchGroupIdx(value *int64)() {
    m.batch_group_idx = value
}
// SetBatchGroupName sets the batch_group_name property value. The batch_group_name property
func (m *GenesisLedgerProofOrigin) SetBatchGroupName(value *string)() {
    m.batch_group_name = value
}
// SetBatchIdx sets the batch_idx property value. The batch_idx property
func (m *GenesisLedgerProofOrigin) SetBatchIdx(value *int64)() {
    m.batch_idx = value
}
// SetBatchName sets the batch_name property value. The batch_name property
func (m *GenesisLedgerProofOrigin) SetBatchName(value *string)() {
    m.batch_name = value
}
// SetGenesisOpaqueHash sets the genesis_opaque_hash property value. The genesis_opaque_hash property
func (m *GenesisLedgerProofOrigin) SetGenesisOpaqueHash(value *string)() {
    m.genesis_opaque_hash = value
}
// SetIsEndOfUpdate sets the is_end_of_update property value. The is_end_of_update property
func (m *GenesisLedgerProofOrigin) SetIsEndOfUpdate(value *bool)() {
    m.is_end_of_update = value
}
// SetProtocolVersionName sets the protocol_version_name property value. The protocol_version_name property
func (m *GenesisLedgerProofOrigin) SetProtocolVersionName(value *string)() {
    m.protocol_version_name = value
}
type GenesisLedgerProofOriginable interface {
    LedgerProofOriginable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatchGroupIdx()(*int64)
    GetBatchGroupName()(*string)
    GetBatchIdx()(*int64)
    GetBatchName()(*string)
    GetGenesisOpaqueHash()(*string)
    GetIsEndOfUpdate()(*bool)
    GetProtocolVersionName()(*string)
    SetBatchGroupIdx(value *int64)()
    SetBatchGroupName(value *string)()
    SetBatchIdx(value *int64)()
    SetBatchName(value *string)()
    SetGenesisOpaqueHash(value *string)()
    SetIsEndOfUpdate(value *bool)()
    SetProtocolVersionName(value *string)()
}
