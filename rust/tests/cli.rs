use assert_cmd::prelude::*;
use predicates::prelude::*;
use std::process::Command;


#[test]
fn file_doesnt_exist() -> Result<(), Box<dyn std::error::Error>> {
    let mut cmd = Command::cargo_bin("hello-rust")?;

    cmd.arg("foobar").arg("test/file/doesnt/exist");
    cmd.assert()
        .failure()
        .stderr(predicate::str::contains("Could not read file"));
    Ok(())
}