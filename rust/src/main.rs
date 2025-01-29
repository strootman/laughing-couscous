// use ferris_says::say;
// use std::io::{stdout, BufWriter};
// use indicatif::ProgressBar;
use anyhow::{Context, Result};
use clap::Parser;

/* Notes: 
 * - Trying to put the `#[derive(Debug)]` macro on this type is interesting due to the `PathBuf`.
 *   I punted on this, because `PathBuf` does not implement the `Display` trait. You get into 
 *   ownership issues pretty quickly. Seems solvable, I just decided to move on.
 */
#[derive(Parser)]
struct Cli {
    pattern: String,
    path: std::path::PathBuf,
}

fn find_matches(content: &str, pattern: &str) {
    for line in content.lines() {
        if line.contains(pattern) {
            println!("{}", line);
        }
    }
}

// Search for a pattern in a file and display the lines that contain it
/*
 * https://rust-cli.github.io/book/index.html 
 * 
 */
fn main() -> Result<()> {
    // let pb = indicatif::ProgressBar::new(100);
    let args = Cli::parse();
    let path = &args.path.into_os_string().into_string().unwrap();
    let content =
        std::fs::read_to_string(path).with_context(|| format!("Could not read file `{}`", path))?;
    find_matches(&content, &args.pattern);

    println!("pattern: {}, path: {:?}", args.pattern, path);
    Ok(())
}
