package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountResourcePreferenceEntrySubstate struct {
	Substate
	// The key property
	key ResourceKeyable
	// The value property
	value AccountResourcePreferenceEntryValueable
}

// NewAccountResourcePreferenceEntrySubstate instantiates a new AccountResourcePreferenceEntrySubstate and sets the default values.
func NewAccountResourcePreferenceEntrySubstate() *AccountResourcePreferenceEntrySubstate {
	m := &AccountResourcePreferenceEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateAccountResourcePreferenceEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountResourcePreferenceEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountResourcePreferenceEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountResourcePreferenceEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateResourceKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(ResourceKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccountResourcePreferenceEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(AccountResourcePreferenceEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a ResourceKeyable when successful
func (m *AccountResourcePreferenceEntrySubstate) GetKey() ResourceKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a AccountResourcePreferenceEntryValueable when successful
func (m *AccountResourcePreferenceEntrySubstate) GetValue() AccountResourcePreferenceEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *AccountResourcePreferenceEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Substate.Serialize(writer)
	if err != nil {
		return err
	}
	{
		err = writer.WriteObjectValue("key", m.GetKey())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("value", m.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetKey sets the key property value. The key property
func (m *AccountResourcePreferenceEntrySubstate) SetKey(value ResourceKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *AccountResourcePreferenceEntrySubstate) SetValue(value AccountResourcePreferenceEntryValueable) {
	m.value = value
}

type AccountResourcePreferenceEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() ResourceKeyable
	GetValue() AccountResourcePreferenceEntryValueable
	SetKey(value ResourceKeyable)
	SetValue(value AccountResourcePreferenceEntryValueable)
}
