package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LtsTransactionSubmitRejectedErrorDetails indicates that the transaction was executed and resulted in a rejection,therefore the transaction is not being added into the mempool.
type LtsTransactionSubmitRejectedErrorDetails struct {
	LtsTransactionSubmitErrorDetails
	// An explanation of the error
	error_message *string
	// An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the rejection isn't permanent.
	invalid_from_epoch *int64
	// Whether (true) this rejected status has just been calculated fresh, or (false) the status is from the pendingtransaction result cache.
	is_fresh *bool
	// Whether the rejection of this intent is known to be permanent - this is a stronger statement than the payload rejectionbeing permanent, as it implies any payloads containing the intent will also be permanently rejected.
	is_intent_rejection_permanent *bool
	// Whether the rejection of this payload is known to be permanent.
	is_payload_rejection_permanent *bool
	// An integer between `0` and `10^10`, marking the epoch after which the node will consider recalculating the validity of the transaction.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
	retry_from_epoch *int64
	// The retry_from_timestamp property
	retry_from_timestamp InstantMsable
}

// NewLtsTransactionSubmitRejectedErrorDetails instantiates a new LtsTransactionSubmitRejectedErrorDetails and sets the default values.
func NewLtsTransactionSubmitRejectedErrorDetails() *LtsTransactionSubmitRejectedErrorDetails {
	m := &LtsTransactionSubmitRejectedErrorDetails{
		LtsTransactionSubmitErrorDetails: *NewLtsTransactionSubmitErrorDetails(),
	}
	return m
}

// CreateLtsTransactionSubmitRejectedErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLtsTransactionSubmitRejectedErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLtsTransactionSubmitRejectedErrorDetails(), nil
}

// GetErrorMessage gets the error_message property value. An explanation of the error
// returns a *string when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.LtsTransactionSubmitErrorDetails.GetFieldDeserializers()
	res["error_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrorMessage(val)
		}
		return nil
	}
	res["invalid_from_epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInvalidFromEpoch(val)
		}
		return nil
	}
	res["is_fresh"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsFresh(val)
		}
		return nil
	}
	res["is_intent_rejection_permanent"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsIntentRejectionPermanent(val)
		}
		return nil
	}
	res["is_payload_rejection_permanent"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsPayloadRejectionPermanent(val)
		}
		return nil
	}
	res["retry_from_epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRetryFromEpoch(val)
		}
		return nil
	}
	res["retry_from_timestamp"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateInstantMsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRetryFromTimestamp(val.(InstantMsable))
		}
		return nil
	}
	return res
}

// GetInvalidFromEpoch gets the invalid_from_epoch property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the rejection isn't permanent.
// returns a *int64 when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetInvalidFromEpoch() *int64 {
	return m.invalid_from_epoch
}

// GetIsFresh gets the is_fresh property value. Whether (true) this rejected status has just been calculated fresh, or (false) the status is from the pendingtransaction result cache.
// returns a *bool when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetIsFresh() *bool {
	return m.is_fresh
}

// GetIsIntentRejectionPermanent gets the is_intent_rejection_permanent property value. Whether the rejection of this intent is known to be permanent - this is a stronger statement than the payload rejectionbeing permanent, as it implies any payloads containing the intent will also be permanently rejected.
// returns a *bool when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetIsIntentRejectionPermanent() *bool {
	return m.is_intent_rejection_permanent
}

// GetIsPayloadRejectionPermanent gets the is_payload_rejection_permanent property value. Whether the rejection of this payload is known to be permanent.
// returns a *bool when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetIsPayloadRejectionPermanent() *bool {
	return m.is_payload_rejection_permanent
}

// GetRetryFromEpoch gets the retry_from_epoch property value. An integer between `0` and `10^10`, marking the epoch after which the node will consider recalculating the validity of the transaction.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
// returns a *int64 when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetRetryFromEpoch() *int64 {
	return m.retry_from_epoch
}

// GetRetryFromTimestamp gets the retry_from_timestamp property value. The retry_from_timestamp property
// returns a InstantMsable when successful
func (m *LtsTransactionSubmitRejectedErrorDetails) GetRetryFromTimestamp() InstantMsable {
	return m.retry_from_timestamp
}

// Serialize serializes information the current object
func (m *LtsTransactionSubmitRejectedErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.LtsTransactionSubmitErrorDetails.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("error_message", m.GetErrorMessage())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteInt64Value("invalid_from_epoch", m.GetInvalidFromEpoch())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("is_fresh", m.GetIsFresh())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("is_intent_rejection_permanent", m.GetIsIntentRejectionPermanent())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("is_payload_rejection_permanent", m.GetIsPayloadRejectionPermanent())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteInt64Value("retry_from_epoch", m.GetRetryFromEpoch())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("retry_from_timestamp", m.GetRetryFromTimestamp())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetErrorMessage sets the error_message property value. An explanation of the error
func (m *LtsTransactionSubmitRejectedErrorDetails) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetInvalidFromEpoch sets the invalid_from_epoch property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the rejection isn't permanent.
func (m *LtsTransactionSubmitRejectedErrorDetails) SetInvalidFromEpoch(value *int64) {
	m.invalid_from_epoch = value
}

// SetIsFresh sets the is_fresh property value. Whether (true) this rejected status has just been calculated fresh, or (false) the status is from the pendingtransaction result cache.
func (m *LtsTransactionSubmitRejectedErrorDetails) SetIsFresh(value *bool) {
	m.is_fresh = value
}

// SetIsIntentRejectionPermanent sets the is_intent_rejection_permanent property value. Whether the rejection of this intent is known to be permanent - this is a stronger statement than the payload rejectionbeing permanent, as it implies any payloads containing the intent will also be permanently rejected.
func (m *LtsTransactionSubmitRejectedErrorDetails) SetIsIntentRejectionPermanent(value *bool) {
	m.is_intent_rejection_permanent = value
}

// SetIsPayloadRejectionPermanent sets the is_payload_rejection_permanent property value. Whether the rejection of this payload is known to be permanent.
func (m *LtsTransactionSubmitRejectedErrorDetails) SetIsPayloadRejectionPermanent(value *bool) {
	m.is_payload_rejection_permanent = value
}

// SetRetryFromEpoch sets the retry_from_epoch property value. An integer between `0` and `10^10`, marking the epoch after which the node will consider recalculating the validity of the transaction.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
func (m *LtsTransactionSubmitRejectedErrorDetails) SetRetryFromEpoch(value *int64) {
	m.retry_from_epoch = value
}

// SetRetryFromTimestamp sets the retry_from_timestamp property value. The retry_from_timestamp property
func (m *LtsTransactionSubmitRejectedErrorDetails) SetRetryFromTimestamp(value InstantMsable) {
	m.retry_from_timestamp = value
}

type LtsTransactionSubmitRejectedErrorDetailsable interface {
	LtsTransactionSubmitErrorDetailsable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetErrorMessage() *string
	GetInvalidFromEpoch() *int64
	GetIsFresh() *bool
	GetIsIntentRejectionPermanent() *bool
	GetIsPayloadRejectionPermanent() *bool
	GetRetryFromEpoch() *int64
	GetRetryFromTimestamp() InstantMsable
	SetErrorMessage(value *string)
	SetInvalidFromEpoch(value *int64)
	SetIsFresh(value *bool)
	SetIsIntentRejectionPermanent(value *bool)
	SetIsPayloadRejectionPermanent(value *bool)
	SetRetryFromEpoch(value *int64)
	SetRetryFromTimestamp(value InstantMsable)
}
