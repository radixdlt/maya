mod create_account;
mod deposit_and_withdraw;
mod publish_and_instantiate;

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
    DepositAndWithdraw(deposit_and_withdraw::DepositAndWithdraw),
}

impl Cli {
    pub fn run<O: std::io::Write>(self, out: &mut O) -> Result<(), Error> {
        match self {
            Self::PublishAndInstantiate(cmd) => cmd.run(out),
            Self::CreateAccount(cmd) => cmd.run(out),
            Self::DepositAndWithdraw(cmd) => cmd.run(out),
        }
    }
}
