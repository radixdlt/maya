package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LtsTransactionSubmitPriorityThresholdNotMetErrorDetails struct {
	LtsTransactionSubmitErrorDetails
	// A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
	min_tip_percentage_required *int32
	// Tip percentage of the submitted (and rejected) transaction.
	tip_percentage *int32
}

// NewLtsTransactionSubmitPriorityThresholdNotMetErrorDetails instantiates a new LtsTransactionSubmitPriorityThresholdNotMetErrorDetails and sets the default values.
func NewLtsTransactionSubmitPriorityThresholdNotMetErrorDetails() *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails {
	m := &LtsTransactionSubmitPriorityThresholdNotMetErrorDetails{
		LtsTransactionSubmitErrorDetails: *NewLtsTransactionSubmitErrorDetails(),
	}
	return m
}

// CreateLtsTransactionSubmitPriorityThresholdNotMetErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionSubmitPriorityThresholdNotMetErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionSubmitPriorityThresholdNotMetErrorDetails(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LtsTransactionSubmitErrorDetails.GetFieldDeserializers()
	res["min_tip_percentage_required"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMinTipPercentageRequired(val)
		}
		return nil
	}
	res["tip_percentage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTipPercentage(val)
		}
		return nil
	}
	return res
}

// GetMinTipPercentageRequired gets the min_tip_percentage_required property value. A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
// returns a *int32 when successful
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) GetMinTipPercentageRequired() *int32 {
	return m.min_tip_percentage_required
}

// GetTipPercentage gets the tip_percentage property value. Tip percentage of the submitted (and rejected) transaction.
// returns a *int32 when successful
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) GetTipPercentage() *int32 {
	return m.tip_percentage
}

// Serialize serializes information the current object
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LtsTransactionSubmitErrorDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteInt32Value("min_tip_percentage_required", m.GetMinTipPercentageRequired())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteInt32Value("tip_percentage", m.GetTipPercentage())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetMinTipPercentageRequired sets the min_tip_percentage_required property value. A lower bound for tip percentage at current mempool state. Anything lower than this will very likely result in a mempool rejection.A missing value means there is no tip that can guarantee submission.
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) SetMinTipPercentageRequired(value *int32) {
	m.min_tip_percentage_required = value
}

// SetTipPercentage sets the tip_percentage property value. Tip percentage of the submitted (and rejected) transaction.
func (m *LtsTransactionSubmitPriorityThresholdNotMetErrorDetails) SetTipPercentage(value *int32) {
	m.tip_percentage = value
}

type LtsTransactionSubmitPriorityThresholdNotMetErrorDetailsable interface {
	LtsTransactionSubmitErrorDetailsable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMinTipPercentageRequired() *int32
	GetTipPercentage() *int32
	SetMinTipPercentageRequired(value *int32)
	SetTipPercentage(value *int32)
}
