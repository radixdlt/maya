mod create_account;
mod deposit;
mod info;
mod publish_and_instantiate;
mod transfer_between_vaults;
mod withdraw;

use clap::Parser;
use maya_router_cli::error::*;

fn main() -> Result<(), Error> {
    env_logger::init();
    let mut out = std::io::stdout();
    let cli = <Cli as clap::Parser>::parse();
    cli.run(&mut out)
}

#[derive(Parser, Debug)]
pub enum Cli {
    PublishAndInstantiate(publish_and_instantiate::PublishAndInstantiate),
    CreateAccount(create_account::CreateAccount),
    Deposit(deposit::Deposit),
    Withdraw(withdraw::Withdraw),
    TransferBetweenVaults(transfer_between_vaults::TransferBetweenVaults),
    Info(info::Info),
}

impl Cli {
    pub fn run<O: std::io::Write>(self, out: &mut O) -> Result<(), Error> {
        match self {
            Self::PublishAndInstantiate(cmd) => cmd.run(out),
            Self::CreateAccount(cmd) => cmd.run(out),
            Self::Deposit(cmd) => cmd.run(out),
            Self::Withdraw(cmd) => cmd.run(out),
            Self::TransferBetweenVaults(cmd) => cmd.run(out),
            Self::Info(cmd) => cmd.run(out),
        }
    }
}
