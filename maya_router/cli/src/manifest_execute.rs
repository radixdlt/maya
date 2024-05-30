use crate::network_connection_provider::*;
use crate::network_selector::*;
use radix_common::prelude::*;
use radix_transactions::{model::TransactionManifestV1, signing::PrivateKey};

pub fn execute(
    network: NetworkConfig,
    private_key: PrivateKey,
    manifest: TransactionManifestV1,
) -> Result<ExecutionReceiptSuccessContents, ExecutionServiceError<GatewayExecutorError>> {
    let mut gateway_network_provider = GatewayNetworkConnector::new(
        network.gateway_base_url(),
        network.network_definition(),
        PollingConfiguration {
            interval_in_seconds: 10,
            retries: 10,
        },
    );

    let signer_keys = vec![&private_key];
    let address = ComponentAddress::virtual_account_from_public_key(&private_key.public_key());

    let execution_service = ExecutionService::new(
        &mut gateway_network_provider,
        address,
        &private_key,
        &signer_keys,
    );
    execution_service.execute_manifest(manifest, false)
}

pub fn execute_2(
    network_connector: &GatewayNetworkConnector,
    private_key: PrivateKey,
    manifest: TransactionManifestV1,
) -> Result<ExecutionReceiptSuccessContents, ExecutionServiceError<GatewayExecutorError>> {
    let signer_keys = vec![&private_key];
    let address = ComponentAddress::virtual_account_from_public_key(&private_key.public_key());

    let execution_service =
        ExecutionService::new(network_connector, address, &private_key, &signer_keys);
    execution_service.execute_manifest(manifest, false)
}
