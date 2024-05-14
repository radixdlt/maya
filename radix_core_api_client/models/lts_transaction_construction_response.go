package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionConstructionResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An integer between `0` and `10^10`, marking the current epoch
	current_epoch *int64
	// The ledger_clock property
	ledger_clock InstantMsable
}

// NewLtsTransactionConstructionResponse instantiates a new LtsTransactionConstructionResponse and sets the default values.
func NewLtsTransactionConstructionResponse() *LtsTransactionConstructionResponse {
	m := &LtsTransactionConstructionResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLtsTransactionConstructionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionConstructionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionConstructionResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LtsTransactionConstructionResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCurrentEpoch gets the current_epoch property value. An integer between `0` and `10^10`, marking the current epoch
// returns a *int64 when successful
func (m *LtsTransactionConstructionResponse) GetCurrentEpoch() *int64 {
	return m.current_epoch
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionConstructionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["current_epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrentEpoch(val)
		}
		return nil
	}
	res["ledger_clock"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLedgerClock(val.(InstantMsable))
		}
		return nil
	}
	return res
}

// GetLedgerClock gets the ledger_clock property value. The ledger_clock property
// returns a InstantMsable when successful
func (m *LtsTransactionConstructionResponse) GetLedgerClock() InstantMsable {
	return m.ledger_clock
}

// Serialize serializes information the current object
func (m *LtsTransactionConstructionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("current_epoch", m.GetCurrentEpoch())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("ledger_clock", m.GetLedgerClock())
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
func (m *LtsTransactionConstructionResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCurrentEpoch sets the current_epoch property value. An integer between `0` and `10^10`, marking the current epoch
func (m *LtsTransactionConstructionResponse) SetCurrentEpoch(value *int64) {
	m.current_epoch = value
}

// SetLedgerClock sets the ledger_clock property value. The ledger_clock property
func (m *LtsTransactionConstructionResponse) SetLedgerClock(value InstantMsable) {
	m.ledger_clock = value
}

type LtsTransactionConstructionResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCurrentEpoch() *int64
	GetLedgerClock() InstantMsable
	SetCurrentEpoch(value *int64)
	SetLedgerClock(value InstantMsable)
}
