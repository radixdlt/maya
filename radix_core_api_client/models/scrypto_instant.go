package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScryptoInstant struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and a second-precision (i.e. *skipping* the `.000` milliseconds part).E.g.: `2023-01-26T18:30:09Z`.Note: This field will *not* be present if the actual on-ledger `unix_timestamp_seconds`value is outside the basic range supported by the RFC 3339 / ISO 8601 standard, whichstarts at year 1583 (i.e. the beginning of the Gregorian calendar) and ends at year9999 (inclusive).
	date_time *string
	// A decimal string-encoded 64-bit signed integer, marking the unix timestamp in seconds.Note: this field accurately represents the full range of possible on-ledger values (i.e.`-2^63 <= seconds < 2^63`). This is contrary to the `InstantMs` type used in otherplaces of this API.
	unix_timestamp_seconds *string
}

// NewScryptoInstant instantiates a new ScryptoInstant and sets the default values.
func NewScryptoInstant() *ScryptoInstant {
	m := &ScryptoInstant{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateScryptoInstantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScryptoInstantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewScryptoInstant(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScryptoInstant) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDateTime gets the date_time property value. The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and a second-precision (i.e. *skipping* the `.000` milliseconds part).E.g.: `2023-01-26T18:30:09Z`.Note: This field will *not* be present if the actual on-ledger `unix_timestamp_seconds`value is outside the basic range supported by the RFC 3339 / ISO 8601 standard, whichstarts at year 1583 (i.e. the beginning of the Gregorian calendar) and ends at year9999 (inclusive).
// returns a *string when successful
func (m *ScryptoInstant) GetDateTime() *string {
	return m.date_time
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScryptoInstant) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["unix_timestamp_seconds"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUnixTimestampSeconds(val)
		}
		return nil
	}
	return res
}

// GetUnixTimestampSeconds gets the unix_timestamp_seconds property value. A decimal string-encoded 64-bit signed integer, marking the unix timestamp in seconds.Note: this field accurately represents the full range of possible on-ledger values (i.e.`-2^63 <= seconds < 2^63`). This is contrary to the `InstantMs` type used in otherplaces of this API.
// returns a *string when successful
func (m *ScryptoInstant) GetUnixTimestampSeconds() *string {
	return m.unix_timestamp_seconds
}

// Serialize serializes information the current object
func (m *ScryptoInstant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("date_time", m.GetDateTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("unix_timestamp_seconds", m.GetUnixTimestampSeconds())
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
func (m *ScryptoInstant) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDateTime sets the date_time property value. The RFC 3339 / ISO 8601 string representation of the timestamp. Will always use "Z"(denoting UTC) and a second-precision (i.e. *skipping* the `.000` milliseconds part).E.g.: `2023-01-26T18:30:09Z`.Note: This field will *not* be present if the actual on-ledger `unix_timestamp_seconds`value is outside the basic range supported by the RFC 3339 / ISO 8601 standard, whichstarts at year 1583 (i.e. the beginning of the Gregorian calendar) and ends at year9999 (inclusive).
func (m *ScryptoInstant) SetDateTime(value *string) {
	m.date_time = value
}

// SetUnixTimestampSeconds sets the unix_timestamp_seconds property value. A decimal string-encoded 64-bit signed integer, marking the unix timestamp in seconds.Note: this field accurately represents the full range of possible on-ledger values (i.e.`-2^63 <= seconds < 2^63`). This is contrary to the `InstantMs` type used in otherplaces of this API.
func (m *ScryptoInstant) SetUnixTimestampSeconds(value *string) {
	m.unix_timestamp_seconds = value
}

type ScryptoInstantable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDateTime() *string
	GetUnixTimestampSeconds() *string
	SetDateTime(value *string)
	SetUnixTimestampSeconds(value *string)
}
