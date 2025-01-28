// use ferris_says::say;
// use std::io::{stdout, BufWriter};
use anyhow::{Context, Result};
use clap::Parser;

// Search for a pattern in a file and display the lines that contain it
#[derive(Parser)]
struct Cli {
    pattern: String,
    path: std::path::PathBuf,
}

fn main() -> Result<()> {
    let args = Cli::parse();
    let path = &args.path.into_os_string().into_string().unwrap();
    let content =
        std::fs::read_to_string(path).with_context(|| format!("Could not read file `{}`", path))?;
    for line in content.lines() {
        if line.contains(&args.pattern) {
            println!("{}", line);
        }
    }

    println!("pattern: {:?}, path: {:?}", args.pattern, path);
    Ok(())
}
