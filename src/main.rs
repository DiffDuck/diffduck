use clap;
use std::{self, ops::Deref};

fn cli() -> clap::Command {
    clap::Command::new("diffduck")
        .about("diffduck: A commit message writing assistant")
        .subcommand_required(true)
        .arg_required_else_help(true)
        .allow_external_subcommands(true)
        .subcommand(clap::Command::new("commit").about("Start writing a commit message"))
}

fn main() {
    let matches = cli().get_matches();

    match matches.subcommand() {
        Some(("commit", _sub_matches)) => {
            println!("You ran diffduck commit.");
        }
        Some((external_command, sub_matches)) => {
            let args = sub_matches
                .get_many::<std::ffi::OsString>("")
                .into_iter()
                .flatten()
                .map(|x| x.to_string_lossy())
                .collect::<Vec<_>>();
            let args_joined = args.join(" ");
            println!("Unsupported diffduck command. Falling back to git.");
            println!("Running \"git {external_command} {args_joined}\"");
            std::process::Command::new("git")
                .arg(external_command)
                .args(args.iter().map(|x| x.deref()))
                .status()
                .unwrap();
        }
        _ => unreachable!(),
    }
}
