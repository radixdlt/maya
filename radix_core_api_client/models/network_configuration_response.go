package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NetworkConfigurationResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address_types property
	address_types []AddressTypeable
	// The logical name of the network
	network *string
	// The network suffix used for Bech32m HRPs used for addressing.
	network_hrp_suffix *string
	// The logical id of the network
	network_id *int32
	// The current value of the protocol-based USD/XRD multiplier (i.e. an amount of XRDs to be paid for 1 USD).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
	usd_price_in_xrd *string
	// Different versions regarding the node, network and api.
	version NetworkConfigurationResponse_versionable
	// Key addresses for this network.
	well_known_addresses NetworkConfigurationResponse_well_known_addressesable
}

// NewNetworkConfigurationResponse instantiates a new NetworkConfigurationResponse and sets the default values.
func NewNetworkConfigurationResponse() *NetworkConfigurationResponse {
	m := &NetworkConfigurationResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNetworkConfigurationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkConfigurationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNetworkConfigurationResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkConfigurationResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddressTypes gets the address_types property value. The address_types property
// returns a []AddressTypeable when successful
func (m *NetworkConfigurationResponse) GetAddressTypes() []AddressTypeable {
	return m.address_types
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkConfigurationResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAddressTypeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AddressTypeable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AddressTypeable)
				}
			}
			m.SetAddressTypes(res)
		}
		return nil
	}
	res["network"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetwork(val)
		}
		return nil
	}
	res["network_hrp_suffix"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetworkHrpSuffix(val)
		}
		return nil
	}
	res["network_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNetworkId(val)
		}
		return nil
	}
	res["usd_price_in_xrd"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUsdPriceInXrd(val)
		}
		return nil
	}
	res["version"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNetworkConfigurationResponse_versionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVersion(val.(NetworkConfigurationResponse_versionable))
		}
		return nil
	}
	res["well_known_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNetworkConfigurationResponse_well_known_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWellKnownAddresses(val.(NetworkConfigurationResponse_well_known_addressesable))
		}
		return nil
	}
	return res
}

// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *NetworkConfigurationResponse) GetNetwork() *string {
	return m.network
}

// GetNetworkHrpSuffix gets the network_hrp_suffix property value. The network suffix used for Bech32m HRPs used for addressing.
// returns a *string when successful
func (m *NetworkConfigurationResponse) GetNetworkHrpSuffix() *string {
	return m.network_hrp_suffix
}

// GetNetworkId gets the network_id property value. The logical id of the network
// returns a *int32 when successful
func (m *NetworkConfigurationResponse) GetNetworkId() *int32 {
	return m.network_id
}

// GetUsdPriceInXrd gets the usd_price_in_xrd property value. The current value of the protocol-based USD/XRD multiplier (i.e. an amount of XRDs to be paid for 1 USD).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
// returns a *string when successful
func (m *NetworkConfigurationResponse) GetUsdPriceInXrd() *string {
	return m.usd_price_in_xrd
}

// GetVersion gets the version property value. Different versions regarding the node, network and api.
// returns a NetworkConfigurationResponse_versionable when successful
func (m *NetworkConfigurationResponse) GetVersion() NetworkConfigurationResponse_versionable {
	return m.version
}

// GetWellKnownAddresses gets the well_known_addresses property value. Key addresses for this network.
// returns a NetworkConfigurationResponse_well_known_addressesable when successful
func (m *NetworkConfigurationResponse) GetWellKnownAddresses() NetworkConfigurationResponse_well_known_addressesable {
	return m.well_known_addresses
}

// Serialize serializes information the current object
func (m *NetworkConfigurationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAddressTypes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddressTypes()))
		for i, v := range m.GetAddressTypes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("address_types", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("network", m.GetNetwork())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("network_hrp_suffix", m.GetNetworkHrpSuffix())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("network_id", m.GetNetworkId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("usd_price_in_xrd", m.GetUsdPriceInXrd())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("version", m.GetVersion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("well_known_addresses", m.GetWellKnownAddresses())
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
func (m *NetworkConfigurationResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddressTypes sets the address_types property value. The address_types property
func (m *NetworkConfigurationResponse) SetAddressTypes(value []AddressTypeable) {
	m.address_types = value
}

// SetNetwork sets the network property value. The logical name of the network
func (m *NetworkConfigurationResponse) SetNetwork(value *string) {
	m.network = value
}

// SetNetworkHrpSuffix sets the network_hrp_suffix property value. The network suffix used for Bech32m HRPs used for addressing.
func (m *NetworkConfigurationResponse) SetNetworkHrpSuffix(value *string) {
	m.network_hrp_suffix = value
}

// SetNetworkId sets the network_id property value. The logical id of the network
func (m *NetworkConfigurationResponse) SetNetworkId(value *int32) {
	m.network_id = value
}

// SetUsdPriceInXrd sets the usd_price_in_xrd property value. The current value of the protocol-based USD/XRD multiplier (i.e. an amount of XRDs to be paid for 1 USD).A decimal is formed of some signed integer `m` of attos (`10^(-18)`) units, where `-2^(192 - 1) <= m < 2^(192 - 1)`.
func (m *NetworkConfigurationResponse) SetUsdPriceInXrd(value *string) {
	m.usd_price_in_xrd = value
}

// SetVersion sets the version property value. Different versions regarding the node, network and api.
func (m *NetworkConfigurationResponse) SetVersion(value NetworkConfigurationResponse_versionable) {
	m.version = value
}

// SetWellKnownAddresses sets the well_known_addresses property value. Key addresses for this network.
func (m *NetworkConfigurationResponse) SetWellKnownAddresses(value NetworkConfigurationResponse_well_known_addressesable) {
	m.well_known_addresses = value
}

type NetworkConfigurationResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddressTypes() []AddressTypeable
	GetNetwork() *string
	GetNetworkHrpSuffix() *string
	GetNetworkId() *int32
	GetUsdPriceInXrd() *string
	GetVersion() NetworkConfigurationResponse_versionable
	GetWellKnownAddresses() NetworkConfigurationResponse_well_known_addressesable
	SetAddressTypes(value []AddressTypeable)
	SetNetwork(value *string)
	SetNetworkHrpSuffix(value *string)
	SetNetworkId(value *int32)
	SetUsdPriceInXrd(value *string)
	SetVersion(value NetworkConfigurationResponse_versionable)
	SetWellKnownAddresses(value NetworkConfigurationResponse_well_known_addressesable)
}
