mod create_account;
mod publish;

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
    Publish(publish::Publish),
    CreateAccount(create_account::CreateAccount),
}

impl Cli {
    pub fn run<O: std::io::Write>(self, out: &mut O) -> Result<(), Error> {
        match self {
            Self::Publish(cmd) => cmd.run(out),
            Self::CreateAccount(cmd) => cmd.run(out),
        }
    }
}
