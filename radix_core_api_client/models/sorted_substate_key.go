package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SortedSubstateKey struct {
	SubstateKey
	// The hex-encoded remaining bytes of the key
	key_hex *string
	// The hex-encoded bytes of the sorted part of the key
	sort_prefix_hex *string
}

// NewSortedSubstateKey instantiates a new SortedSubstateKey and sets the default values.
func NewSortedSubstateKey() *SortedSubstateKey {
	m := &SortedSubstateKey{
		SubstateKey: *NewSubstateKey(),
	}
	return m
}

// CreateSortedSubstateKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSortedSubstateKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSortedSubstateKey(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SortedSubstateKey) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.SubstateKey.GetFieldDeserializers()
	res["key_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeyHex(val)
		}
		return nil
	}
	res["sort_prefix_hex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSortPrefixHex(val)
		}
		return nil
	}
	return res
}

// GetKeyHex gets the key_hex property value. The hex-encoded remaining bytes of the key
// returns a *string when successful
func (m *SortedSubstateKey) GetKeyHex() *string {
	return m.key_hex
}

// GetSortPrefixHex gets the sort_prefix_hex property value. The hex-encoded bytes of the sorted part of the key
// returns a *string when successful
func (m *SortedSubstateKey) GetSortPrefixHex() *string {
	return m.sort_prefix_hex
}

// Serialize serializes information the current object
func (m *SortedSubstateKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.SubstateKey.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteStringValue("key_hex", m.GetKeyHex())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("sort_prefix_hex", m.GetSortPrefixHex())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKeyHex sets the key_hex property value. The hex-encoded remaining bytes of the key
func (m *SortedSubstateKey) SetKeyHex(value *string) {
	m.key_hex = value
}

// SetSortPrefixHex sets the sort_prefix_hex property value. The hex-encoded bytes of the sorted part of the key
func (m *SortedSubstateKey) SetSortPrefixHex(value *string) {
	m.sort_prefix_hex = value
}

type SortedSubstateKeyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubstateKeyable
	GetKeyHex() *string
	GetSortPrefixHex() *string
	SetKeyHex(value *string)
	SetSortPrefixHex(value *string)
}
