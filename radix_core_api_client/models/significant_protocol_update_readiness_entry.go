package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SignificantProtocolUpdateReadinessEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The readiness_signal_name property
    readiness_signal_name *string
    // The signalled_stake property
    signalled_stake *string
}
// NewSignificantProtocolUpdateReadinessEntry instantiates a new SignificantProtocolUpdateReadinessEntry and sets the default values.
func NewSignificantProtocolUpdateReadinessEntry()(*SignificantProtocolUpdateReadinessEntry) {
    m := &SignificantProtocolUpdateReadinessEntry{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSignificantProtocolUpdateReadinessEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSignificantProtocolUpdateReadinessEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignificantProtocolUpdateReadinessEntry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SignificantProtocolUpdateReadinessEntry) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SignificantProtocolUpdateReadinessEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["readiness_signal_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadinessSignalName(val)
        }
        return nil
    }
    res["signalled_stake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignalledStake(val)
        }
        return nil
    }
    return res
}
// GetReadinessSignalName gets the readiness_signal_name property value. The readiness_signal_name property
// returns a *string when successful
func (m *SignificantProtocolUpdateReadinessEntry) GetReadinessSignalName()(*string) {
    return m.readiness_signal_name
}
// GetSignalledStake gets the signalled_stake property value. The signalled_stake property
// returns a *string when successful
func (m *SignificantProtocolUpdateReadinessEntry) GetSignalledStake()(*string) {
    return m.signalled_stake
}
// Serialize serializes information the current object
func (m *SignificantProtocolUpdateReadinessEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("readiness_signal_name", m.GetReadinessSignalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signalled_stake", m.GetSignalledStake())
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
func (m *SignificantProtocolUpdateReadinessEntry) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReadinessSignalName sets the readiness_signal_name property value. The readiness_signal_name property
func (m *SignificantProtocolUpdateReadinessEntry) SetReadinessSignalName(value *string)() {
    m.readiness_signal_name = value
}
// SetSignalledStake sets the signalled_stake property value. The signalled_stake property
func (m *SignificantProtocolUpdateReadinessEntry) SetSignalledStake(value *string)() {
    m.signalled_stake = value
}
type SignificantProtocolUpdateReadinessEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReadinessSignalName()(*string)
    GetSignalledStake()(*string)
    SetReadinessSignalName(value *string)()
    SetSignalledStake(value *string)()
}
