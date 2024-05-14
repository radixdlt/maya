package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountLockerAccountClaimsEntrySubstate struct {
	Substate
	// The key property
	key AccountAddressKeyable
	// The value property
	value AccountLockerAccountClaimsEntryValueable
}

// NewAccountLockerAccountClaimsEntrySubstate instantiates a new AccountLockerAccountClaimsEntrySubstate and sets the default values.
func NewAccountLockerAccountClaimsEntrySubstate() *AccountLockerAccountClaimsEntrySubstate {
	m := &AccountLockerAccountClaimsEntrySubstate{
		Substate: *NewSubstate(),
	}
	return m
}

// CreateAccountLockerAccountClaimsEntrySubstateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountLockerAccountClaimsEntrySubstateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountLockerAccountClaimsEntrySubstate(), nil
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountLockerAccountClaimsEntrySubstate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Substate.GetFieldDeserializers()
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccountAddressKeyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val.(AccountAddressKeyable))
		}
		return nil
	}
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccountLockerAccountClaimsEntryValueFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val.(AccountLockerAccountClaimsEntryValueable))
		}
		return nil
	}
	return res
}

// GetKey gets the key property value. The key property
// returns a AccountAddressKeyable when successful
func (m *AccountLockerAccountClaimsEntrySubstate) GetKey() AccountAddressKeyable {
	return m.key
}

// GetValue gets the value property value. The value property
// returns a AccountLockerAccountClaimsEntryValueable when successful
func (m *AccountLockerAccountClaimsEntrySubstate) GetValue() AccountLockerAccountClaimsEntryValueable {
	return m.value
}

// Serialize serializes information the current object
func (m *AccountLockerAccountClaimsEntrySubstate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AccountLockerAccountClaimsEntrySubstate) SetKey(value AccountAddressKeyable) {
	m.key = value
}

// SetValue sets the value property value. The value property
func (m *AccountLockerAccountClaimsEntrySubstate) SetValue(value AccountLockerAccountClaimsEntryValueable) {
	m.value = value
}

type AccountLockerAccountClaimsEntrySubstateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	Substateable
	GetKey() AccountAddressKeyable
	GetValue() AccountLockerAccountClaimsEntryValueable
	SetKey(value AccountAddressKeyable)
	SetValue(value AccountLockerAccountClaimsEntryValueable)
}
