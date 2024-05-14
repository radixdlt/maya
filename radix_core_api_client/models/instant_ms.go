package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InstantMs struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and include milliseconds (even if `000`).E.g.: `2023-01-26T18:30:09.453Z`.
	date_time *string
	// An integer between `0` and `10^14`, marking the unix timestamp in ms.
	unix_timestamp_ms *int64
}

// NewInstantMs instantiates a new InstantMs and sets the default values.
func NewInstantMs() *InstantMs {
	m := &InstantMs{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateInstantMsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstantMsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewInstantMs(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InstantMs) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDateTime gets the date_time property value. The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and include milliseconds (even if `000`).E.g.: `2023-01-26T18:30:09.453Z`.
// returns a *string when successful
func (m *InstantMs) GetDateTime() *string {
	return m.date_time
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InstantMs) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["date_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateTime(val)
		}
		return nil
	}
	res["unix_timestamp_ms"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUnixTimestampMs(val)
		}
		return nil
	}
	return res
}

// GetUnixTimestampMs gets the unix_timestamp_ms property value. An integer between `0` and `10^14`, marking the unix timestamp in ms.
// returns a *int64 when successful
func (m *InstantMs) GetUnixTimestampMs() *int64 {
	return m.unix_timestamp_ms
}

// Serialize serializes information the current object
func (m *InstantMs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("date_time", m.GetDateTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("unix_timestamp_ms", m.GetUnixTimestampMs())
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
func (m *InstantMs) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDateTime sets the date_time property value. The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and include milliseconds (even if `000`).E.g.: `2023-01-26T18:30:09.453Z`.
func (m *InstantMs) SetDateTime(value *string) {
	m.date_time = value
}

// SetUnixTimestampMs sets the unix_timestamp_ms property value. An integer between `0` and `10^14`, marking the unix timestamp in ms.
func (m *InstantMs) SetUnixTimestampMs(value *int64) {
	m.unix_timestamp_ms = value
}

type InstantMsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDateTime() *string
	GetUnixTimestampMs() *int64
	SetDateTime(value *string)
	SetUnixTimestampMs(value *int64)
}
