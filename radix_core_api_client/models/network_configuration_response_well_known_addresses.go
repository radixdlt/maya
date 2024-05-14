package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkConfigurationResponse_well_known_addresses key addresses for this network.
type NetworkConfigurationResponse_well_known_addresses struct {
	// The access_controller_package property
	access_controller_package *string
	// The account_owner_badge property
	account_owner_badge *string
	// The account_package property
	account_package *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The consensus_manager property
	consensus_manager *string
	// The consensus_manager_package property
	consensus_manager_package *string
	// The ed25519_signature_virtual_badge property
	ed25519_signature_virtual_badge *string
	// The faucet property
	faucet *string
	// The faucet_package property
	faucet_package *string
	// The genesis_helper property
	genesis_helper *string
	// The genesis_helper_package property
	genesis_helper_package *string
	// The global_caller_virtual_badge property
	global_caller_virtual_badge *string
	// The identity_owner_badge property
	identity_owner_badge *string
	// The identity_package property
	identity_package *string
	// The locker_package property
	locker_package *string
	// The metadata_module_package property
	metadata_module_package *string
	// The package_of_direct_caller_virtual_badge property
	package_of_direct_caller_virtual_badge *string
	// The package_owner_badge property
	package_owner_badge *string
	// The package_package property
	package_package *string
	// The pool_package property
	pool_package *string
	// The resource_package property
	resource_package *string
	// The role_assignment_module_package property
	role_assignment_module_package *string
	// The royalty_module_package property
	royalty_module_package *string
	// The secp256k1_signature_virtual_badge property
	secp256k1_signature_virtual_badge *string
	// The system_transaction_badge property
	system_transaction_badge *string
	// The transaction_processor_package property
	transaction_processor_package *string
	// The transaction_tracker property
	transaction_tracker *string
	// The validator_owner_badge property
	validator_owner_badge *string
	// The xrd property
	xrd *string
}

// NewNetworkConfigurationResponse_well_known_addresses instantiates a new NetworkConfigurationResponse_well_known_addresses and sets the default values.
func NewNetworkConfigurationResponse_well_known_addresses() *NetworkConfigurationResponse_well_known_addresses {
	m := &NetworkConfigurationResponse_well_known_addresses{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNetworkConfigurationResponse_well_known_addressesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNetworkConfigurationResponse_well_known_addressesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNetworkConfigurationResponse_well_known_addresses(), nil
}

// GetAccessControllerPackage gets the access_controller_package property value. The access_controller_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetAccessControllerPackage() *string {
	return m.access_controller_package
}

// GetAccountOwnerBadge gets the account_owner_badge property value. The account_owner_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetAccountOwnerBadge() *string {
	return m.account_owner_badge
}

// GetAccountPackage gets the account_package property value. The account_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetAccountPackage() *string {
	return m.account_package
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetConsensusManager gets the consensus_manager property value. The consensus_manager property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetConsensusManager() *string {
	return m.consensus_manager
}

// GetConsensusManagerPackage gets the consensus_manager_package property value. The consensus_manager_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetConsensusManagerPackage() *string {
	return m.consensus_manager_package
}

// GetEd25519SignatureVirtualBadge gets the ed25519_signature_virtual_badge property value. The ed25519_signature_virtual_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetEd25519SignatureVirtualBadge() *string {
	return m.ed25519_signature_virtual_badge
}

// GetFaucet gets the faucet property value. The faucet property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetFaucet() *string {
	return m.faucet
}

// GetFaucetPackage gets the faucet_package property value. The faucet_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetFaucetPackage() *string {
	return m.faucet_package
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["access_controller_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccessControllerPackage(val)
		}
		return nil
	}
	res["account_owner_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccountOwnerBadge(val)
		}
		return nil
	}
	res["account_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccountPackage(val)
		}
		return nil
	}
	res["consensus_manager"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetConsensusManager(val)
		}
		return nil
	}
	res["consensus_manager_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetConsensusManagerPackage(val)
		}
		return nil
	}
	res["ed25519_signature_virtual_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEd25519SignatureVirtualBadge(val)
		}
		return nil
	}
	res["faucet"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaucet(val)
		}
		return nil
	}
	res["faucet_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaucetPackage(val)
		}
		return nil
	}
	res["genesis_helper"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGenesisHelper(val)
		}
		return nil
	}
	res["genesis_helper_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGenesisHelperPackage(val)
		}
		return nil
	}
	res["global_caller_virtual_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGlobalCallerVirtualBadge(val)
		}
		return nil
	}
	res["identity_owner_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIdentityOwnerBadge(val)
		}
		return nil
	}
	res["identity_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIdentityPackage(val)
		}
		return nil
	}
	res["locker_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLockerPackage(val)
		}
		return nil
	}
	res["metadata_module_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMetadataModulePackage(val)
		}
		return nil
	}
	res["package_of_direct_caller_virtual_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPackageOfDirectCallerVirtualBadge(val)
		}
		return nil
	}
	res["package_owner_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPackageOwnerBadge(val)
		}
		return nil
	}
	res["package_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPackagePackage(val)
		}
		return nil
	}
	res["pool_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPoolPackage(val)
		}
		return nil
	}
	res["resource_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResourcePackage(val)
		}
		return nil
	}
	res["role_assignment_module_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoleAssignmentModulePackage(val)
		}
		return nil
	}
	res["royalty_module_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRoyaltyModulePackage(val)
		}
		return nil
	}
	res["secp256k1_signature_virtual_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSecp256k1SignatureVirtualBadge(val)
		}
		return nil
	}
	res["system_transaction_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSystemTransactionBadge(val)
		}
		return nil
	}
	res["transaction_processor_package"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionProcessorPackage(val)
		}
		return nil
	}
	res["transaction_tracker"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionTracker(val)
		}
		return nil
	}
	res["validator_owner_badge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValidatorOwnerBadge(val)
		}
		return nil
	}
	res["xrd"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetXrd(val)
		}
		return nil
	}
	return res
}

// GetGenesisHelper gets the genesis_helper property value. The genesis_helper property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetGenesisHelper() *string {
	return m.genesis_helper
}

// GetGenesisHelperPackage gets the genesis_helper_package property value. The genesis_helper_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetGenesisHelperPackage() *string {
	return m.genesis_helper_package
}

// GetGlobalCallerVirtualBadge gets the global_caller_virtual_badge property value. The global_caller_virtual_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetGlobalCallerVirtualBadge() *string {
	return m.global_caller_virtual_badge
}

// GetIdentityOwnerBadge gets the identity_owner_badge property value. The identity_owner_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetIdentityOwnerBadge() *string {
	return m.identity_owner_badge
}

// GetIdentityPackage gets the identity_package property value. The identity_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetIdentityPackage() *string {
	return m.identity_package
}

// GetLockerPackage gets the locker_package property value. The locker_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetLockerPackage() *string {
	return m.locker_package
}

// GetMetadataModulePackage gets the metadata_module_package property value. The metadata_module_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetMetadataModulePackage() *string {
	return m.metadata_module_package
}

// GetPackageOfDirectCallerVirtualBadge gets the package_of_direct_caller_virtual_badge property value. The package_of_direct_caller_virtual_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetPackageOfDirectCallerVirtualBadge() *string {
	return m.package_of_direct_caller_virtual_badge
}

// GetPackageOwnerBadge gets the package_owner_badge property value. The package_owner_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetPackageOwnerBadge() *string {
	return m.package_owner_badge
}

// GetPackagePackage gets the package_package property value. The package_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetPackagePackage() *string {
	return m.package_package
}

// GetPoolPackage gets the pool_package property value. The pool_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetPoolPackage() *string {
	return m.pool_package
}

// GetResourcePackage gets the resource_package property value. The resource_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetResourcePackage() *string {
	return m.resource_package
}

// GetRoleAssignmentModulePackage gets the role_assignment_module_package property value. The role_assignment_module_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetRoleAssignmentModulePackage() *string {
	return m.role_assignment_module_package
}

// GetRoyaltyModulePackage gets the royalty_module_package property value. The royalty_module_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetRoyaltyModulePackage() *string {
	return m.royalty_module_package
}

// GetSecp256k1SignatureVirtualBadge gets the secp256k1_signature_virtual_badge property value. The secp256k1_signature_virtual_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetSecp256k1SignatureVirtualBadge() *string {
	return m.secp256k1_signature_virtual_badge
}

// GetSystemTransactionBadge gets the system_transaction_badge property value. The system_transaction_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetSystemTransactionBadge() *string {
	return m.system_transaction_badge
}

// GetTransactionProcessorPackage gets the transaction_processor_package property value. The transaction_processor_package property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetTransactionProcessorPackage() *string {
	return m.transaction_processor_package
}

// GetTransactionTracker gets the transaction_tracker property value. The transaction_tracker property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetTransactionTracker() *string {
	return m.transaction_tracker
}

// GetValidatorOwnerBadge gets the validator_owner_badge property value. The validator_owner_badge property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetValidatorOwnerBadge() *string {
	return m.validator_owner_badge
}

// GetXrd gets the xrd property value. The xrd property
// returns a *string when successful
func (m *NetworkConfigurationResponse_well_known_addresses) GetXrd() *string {
	return m.xrd
}

// Serialize serializes information the current object
func (m *NetworkConfigurationResponse_well_known_addresses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("access_controller_package", m.GetAccessControllerPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("account_owner_badge", m.GetAccountOwnerBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("account_package", m.GetAccountPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("consensus_manager", m.GetConsensusManager())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("consensus_manager_package", m.GetConsensusManagerPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("ed25519_signature_virtual_badge", m.GetEd25519SignatureVirtualBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("faucet", m.GetFaucet())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("faucet_package", m.GetFaucetPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("genesis_helper", m.GetGenesisHelper())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("genesis_helper_package", m.GetGenesisHelperPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("global_caller_virtual_badge", m.GetGlobalCallerVirtualBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("identity_owner_badge", m.GetIdentityOwnerBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("identity_package", m.GetIdentityPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("locker_package", m.GetLockerPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("metadata_module_package", m.GetMetadataModulePackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("package_of_direct_caller_virtual_badge", m.GetPackageOfDirectCallerVirtualBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("package_owner_badge", m.GetPackageOwnerBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("package_package", m.GetPackagePackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("pool_package", m.GetPoolPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("resource_package", m.GetResourcePackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("role_assignment_module_package", m.GetRoleAssignmentModulePackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("royalty_module_package", m.GetRoyaltyModulePackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("secp256k1_signature_virtual_badge", m.GetSecp256k1SignatureVirtualBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("system_transaction_badge", m.GetSystemTransactionBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("transaction_processor_package", m.GetTransactionProcessorPackage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("transaction_tracker", m.GetTransactionTracker())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("validator_owner_badge", m.GetValidatorOwnerBadge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("xrd", m.GetXrd())
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

// SetAccessControllerPackage sets the access_controller_package property value. The access_controller_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetAccessControllerPackage(value *string) {
	m.access_controller_package = value
}

// SetAccountOwnerBadge sets the account_owner_badge property value. The account_owner_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetAccountOwnerBadge(value *string) {
	m.account_owner_badge = value
}

// SetAccountPackage sets the account_package property value. The account_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetAccountPackage(value *string) {
	m.account_package = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkConfigurationResponse_well_known_addresses) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetConsensusManager sets the consensus_manager property value. The consensus_manager property
func (m *NetworkConfigurationResponse_well_known_addresses) SetConsensusManager(value *string) {
	m.consensus_manager = value
}

// SetConsensusManagerPackage sets the consensus_manager_package property value. The consensus_manager_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetConsensusManagerPackage(value *string) {
	m.consensus_manager_package = value
}

// SetEd25519SignatureVirtualBadge sets the ed25519_signature_virtual_badge property value. The ed25519_signature_virtual_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetEd25519SignatureVirtualBadge(value *string) {
	m.ed25519_signature_virtual_badge = value
}

// SetFaucet sets the faucet property value. The faucet property
func (m *NetworkConfigurationResponse_well_known_addresses) SetFaucet(value *string) {
	m.faucet = value
}

// SetFaucetPackage sets the faucet_package property value. The faucet_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetFaucetPackage(value *string) {
	m.faucet_package = value
}

// SetGenesisHelper sets the genesis_helper property value. The genesis_helper property
func (m *NetworkConfigurationResponse_well_known_addresses) SetGenesisHelper(value *string) {
	m.genesis_helper = value
}

// SetGenesisHelperPackage sets the genesis_helper_package property value. The genesis_helper_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetGenesisHelperPackage(value *string) {
	m.genesis_helper_package = value
}

// SetGlobalCallerVirtualBadge sets the global_caller_virtual_badge property value. The global_caller_virtual_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetGlobalCallerVirtualBadge(value *string) {
	m.global_caller_virtual_badge = value
}

// SetIdentityOwnerBadge sets the identity_owner_badge property value. The identity_owner_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetIdentityOwnerBadge(value *string) {
	m.identity_owner_badge = value
}

// SetIdentityPackage sets the identity_package property value. The identity_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetIdentityPackage(value *string) {
	m.identity_package = value
}

// SetLockerPackage sets the locker_package property value. The locker_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetLockerPackage(value *string) {
	m.locker_package = value
}

// SetMetadataModulePackage sets the metadata_module_package property value. The metadata_module_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetMetadataModulePackage(value *string) {
	m.metadata_module_package = value
}

// SetPackageOfDirectCallerVirtualBadge sets the package_of_direct_caller_virtual_badge property value. The package_of_direct_caller_virtual_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetPackageOfDirectCallerVirtualBadge(value *string) {
	m.package_of_direct_caller_virtual_badge = value
}

// SetPackageOwnerBadge sets the package_owner_badge property value. The package_owner_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetPackageOwnerBadge(value *string) {
	m.package_owner_badge = value
}

// SetPackagePackage sets the package_package property value. The package_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetPackagePackage(value *string) {
	m.package_package = value
}

// SetPoolPackage sets the pool_package property value. The pool_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetPoolPackage(value *string) {
	m.pool_package = value
}

// SetResourcePackage sets the resource_package property value. The resource_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetResourcePackage(value *string) {
	m.resource_package = value
}

// SetRoleAssignmentModulePackage sets the role_assignment_module_package property value. The role_assignment_module_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetRoleAssignmentModulePackage(value *string) {
	m.role_assignment_module_package = value
}

// SetRoyaltyModulePackage sets the royalty_module_package property value. The royalty_module_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetRoyaltyModulePackage(value *string) {
	m.royalty_module_package = value
}

// SetSecp256k1SignatureVirtualBadge sets the secp256k1_signature_virtual_badge property value. The secp256k1_signature_virtual_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetSecp256k1SignatureVirtualBadge(value *string) {
	m.secp256k1_signature_virtual_badge = value
}

// SetSystemTransactionBadge sets the system_transaction_badge property value. The system_transaction_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetSystemTransactionBadge(value *string) {
	m.system_transaction_badge = value
}

// SetTransactionProcessorPackage sets the transaction_processor_package property value. The transaction_processor_package property
func (m *NetworkConfigurationResponse_well_known_addresses) SetTransactionProcessorPackage(value *string) {
	m.transaction_processor_package = value
}

// SetTransactionTracker sets the transaction_tracker property value. The transaction_tracker property
func (m *NetworkConfigurationResponse_well_known_addresses) SetTransactionTracker(value *string) {
	m.transaction_tracker = value
}

// SetValidatorOwnerBadge sets the validator_owner_badge property value. The validator_owner_badge property
func (m *NetworkConfigurationResponse_well_known_addresses) SetValidatorOwnerBadge(value *string) {
	m.validator_owner_badge = value
}

// SetXrd sets the xrd property value. The xrd property
func (m *NetworkConfigurationResponse_well_known_addresses) SetXrd(value *string) {
	m.xrd = value
}

type NetworkConfigurationResponse_well_known_addressesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccessControllerPackage() *string
	GetAccountOwnerBadge() *string
	GetAccountPackage() *string
	GetConsensusManager() *string
	GetConsensusManagerPackage() *string
	GetEd25519SignatureVirtualBadge() *string
	GetFaucet() *string
	GetFaucetPackage() *string
	GetGenesisHelper() *string
	GetGenesisHelperPackage() *string
	GetGlobalCallerVirtualBadge() *string
	GetIdentityOwnerBadge() *string
	GetIdentityPackage() *string
	GetLockerPackage() *string
	GetMetadataModulePackage() *string
	GetPackageOfDirectCallerVirtualBadge() *string
	GetPackageOwnerBadge() *string
	GetPackagePackage() *string
	GetPoolPackage() *string
	GetResourcePackage() *string
	GetRoleAssignmentModulePackage() *string
	GetRoyaltyModulePackage() *string
	GetSecp256k1SignatureVirtualBadge() *string
	GetSystemTransactionBadge() *string
	GetTransactionProcessorPackage() *string
	GetTransactionTracker() *string
	GetValidatorOwnerBadge() *string
	GetXrd() *string
	SetAccessControllerPackage(value *string)
	SetAccountOwnerBadge(value *string)
	SetAccountPackage(value *string)
	SetConsensusManager(value *string)
	SetConsensusManagerPackage(value *string)
	SetEd25519SignatureVirtualBadge(value *string)
	SetFaucet(value *string)
	SetFaucetPackage(value *string)
	SetGenesisHelper(value *string)
	SetGenesisHelperPackage(value *string)
	SetGlobalCallerVirtualBadge(value *string)
	SetIdentityOwnerBadge(value *string)
	SetIdentityPackage(value *string)
	SetLockerPackage(value *string)
	SetMetadataModulePackage(value *string)
	SetPackageOfDirectCallerVirtualBadge(value *string)
	SetPackageOwnerBadge(value *string)
	SetPackagePackage(value *string)
	SetPoolPackage(value *string)
	SetResourcePackage(value *string)
	SetRoleAssignmentModulePackage(value *string)
	SetRoyaltyModulePackage(value *string)
	SetSecp256k1SignatureVirtualBadge(value *string)
	SetSystemTransactionBadge(value *string)
	SetTransactionProcessorPackage(value *string)
	SetTransactionTracker(value *string)
	SetValidatorOwnerBadge(value *string)
	SetXrd(value *string)
}
