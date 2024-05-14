package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LedgerStateSummary struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The header_summary property
	header_summary LedgerHeaderSummaryable
	// The state_version property
	state_version *int64
}

// NewLedgerStateSummary instantiates a new LedgerStateSummary and sets the default values.
func NewLedgerStateSummary() *LedgerStateSummary {
	m := &LedgerStateSummary{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLedgerStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLedgerStateSummary(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerStateSummary) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerStateSummary) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["header_summary"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLedgerHeaderSummaryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHeaderSummary(val.(LedgerHeaderSummaryable))
		}
		return nil
	}
	res["state_version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStateVersion(val)
		}
		return nil
	}
	return res
}

// GetHeaderSummary gets the header_summary property value. The header_summary property
// returns a LedgerHeaderSummaryable when successful
func (m *LedgerStateSummary) GetHeaderSummary() LedgerHeaderSummaryable {
	return m.header_summary
}

// GetStateVersion gets the state_version property value. The state_version property
// returns a *int64 when successful
func (m *LedgerStateSummary) GetStateVersion() *int64 {
	return m.state_version
}

// Serialize serializes information the current object
func (m *LedgerStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("header_summary", m.GetHeaderSummary())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("state_version", m.GetStateVersion())
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
func (m *LedgerStateSummary) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetHeaderSummary sets the header_summary property value. The header_summary property
func (m *LedgerStateSummary) SetHeaderSummary(value LedgerHeaderSummaryable) {
	m.header_summary = value
}

// SetStateVersion sets the state_version property value. The state_version property
func (m *LedgerStateSummary) SetStateVersion(value *int64) {
	m.state_version = value
}

type LedgerStateSummaryable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetHeaderSummary() LedgerHeaderSummaryable
	GetStateVersion() *int64
	SetHeaderSummary(value LedgerHeaderSummaryable)
	SetStateVersion(value *int64)
}
