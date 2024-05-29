use crate::network_connection_provider::*;

#[derive(Debug)]
pub enum Error {
    PrivateKeyError,
    GatewayExecutorError(ExecutionServiceError<GatewayExecutorError>),
    SimulatorExecutorError(String),
    IoError(std::io::Error),
}

impl From<ExecutionServiceError<GatewayExecutorError>> for Error {
    fn from(value: ExecutionServiceError<GatewayExecutorError>) -> Self {
        Self::GatewayExecutorError(value)
    }
}
