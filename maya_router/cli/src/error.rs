#[derive(Debug)]
pub enum Error {
    PrivateKeyError,
    GatewayExecutorError(String),
    SimulatorExecutorError(String),
    IoError(std::io::Error),
}
