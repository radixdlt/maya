package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransactionSubmitRejectedErrorDetails indicates that the transaction was executed and resulted in a rejection,therefore the transaction is not being added into the mempool.
type TransactionSubmitRejectedErrorDetails struct {
	TransactionSubmitErrorDetails
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

// NewTransactionSubmitRejectedErrorDetails instantiates a new TransactionSubmitRejectedErrorDetails and sets the default values.
func NewTransactionSubmitRejectedErrorDetails() *TransactionSubmitRejectedErrorDetails {
	m := &TransactionSubmitRejectedErrorDetails{
		TransactionSubmitErrorDetails: *NewTransactionSubmitErrorDetails(),
	}
	return m
}

// CreateTransactionSubmitRejectedErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionSubmitRejectedErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionSubmitRejectedErrorDetails(), nil
}

// GetErrorMessage gets the error_message property value. An explanation of the error
// returns a *string when successful
func (m *TransactionSubmitRejectedErrorDetails) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionSubmitRejectedErrorDetails) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.TransactionSubmitErrorDetails.GetFieldDeserializers()
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
func (m *TransactionSubmitRejectedErrorDetails) GetInvalidFromEpoch() *int64 {
	return m.invalid_from_epoch
}

// GetIsFresh gets the is_fresh property value. Whether (true) this rejected status has just been calculated fresh, or (false) the status is from the pendingtransaction result cache.
// returns a *bool when successful
func (m *TransactionSubmitRejectedErrorDetails) GetIsFresh() *bool {
	return m.is_fresh
}

// GetIsIntentRejectionPermanent gets the is_intent_rejection_permanent property value. Whether the rejection of this intent is known to be permanent - this is a stronger statement than the payload rejectionbeing permanent, as it implies any payloads containing the intent will also be permanently rejected.
// returns a *bool when successful
func (m *TransactionSubmitRejectedErrorDetails) GetIsIntentRejectionPermanent() *bool {
	return m.is_intent_rejection_permanent
}

// GetIsPayloadRejectionPermanent gets the is_payload_rejection_permanent property value. Whether the rejection of this payload is known to be permanent.
// returns a *bool when successful
func (m *TransactionSubmitRejectedErrorDetails) GetIsPayloadRejectionPermanent() *bool {
	return m.is_payload_rejection_permanent
}

// GetRetryFromEpoch gets the retry_from_epoch property value. An integer between `0` and `10^10`, marking the epoch after which the node will consider recalculating the validity of the transaction.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
// returns a *int64 when successful
func (m *TransactionSubmitRejectedErrorDetails) GetRetryFromEpoch() *int64 {
	return m.retry_from_epoch
}

// GetRetryFromTimestamp gets the retry_from_timestamp property value. The retry_from_timestamp property
// returns a InstantMsable when successful
func (m *TransactionSubmitRejectedErrorDetails) GetRetryFromTimestamp() InstantMsable {
	return m.retry_from_timestamp
}

// Serialize serializes information the current object
func (m *TransactionSubmitRejectedErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.TransactionSubmitErrorDetails.Serialize(writer)
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
func (m *TransactionSubmitRejectedErrorDetails) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetInvalidFromEpoch sets the invalid_from_epoch property value. An integer between `0` and `10^10`, marking the epoch from which the transaction will no longer be valid, and be permanently rejected.Only present if the rejection isn't permanent.
func (m *TransactionSubmitRejectedErrorDetails) SetInvalidFromEpoch(value *int64) {
	m.invalid_from_epoch = value
}

// SetIsFresh sets the is_fresh property value. Whether (true) this rejected status has just been calculated fresh, or (false) the status is from the pendingtransaction result cache.
func (m *TransactionSubmitRejectedErrorDetails) SetIsFresh(value *bool) {
	m.is_fresh = value
}

// SetIsIntentRejectionPermanent sets the is_intent_rejection_permanent property value. Whether the rejection of this intent is known to be permanent - this is a stronger statement than the payload rejectionbeing permanent, as it implies any payloads containing the intent will also be permanently rejected.
func (m *TransactionSubmitRejectedErrorDetails) SetIsIntentRejectionPermanent(value *bool) {
	m.is_intent_rejection_permanent = value
}

// SetIsPayloadRejectionPermanent sets the is_payload_rejection_permanent property value. Whether the rejection of this payload is known to be permanent.
func (m *TransactionSubmitRejectedErrorDetails) SetIsPayloadRejectionPermanent(value *bool) {
	m.is_payload_rejection_permanent = value
}

// SetRetryFromEpoch sets the retry_from_epoch property value. An integer between `0` and `10^10`, marking the epoch after which the node will consider recalculating the validity of the transaction.Only present if the rejection is temporary due to a header specifying a "from epoch" in the future.
func (m *TransactionSubmitRejectedErrorDetails) SetRetryFromEpoch(value *int64) {
	m.retry_from_epoch = value
}

// SetRetryFromTimestamp sets the retry_from_timestamp property value. The retry_from_timestamp property
func (m *TransactionSubmitRejectedErrorDetails) SetRetryFromTimestamp(value InstantMsable) {
	m.retry_from_timestamp = value
}

type TransactionSubmitRejectedErrorDetailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TransactionSubmitErrorDetailsable
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
