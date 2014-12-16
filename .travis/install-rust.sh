#!/bin/bash
set -ev

export RUST_URL='https://static.rust-lang.org/dist/rust-nightly-x86_64-unknown-linux-gnu.tar.gz';
export CARGO_URL='https://static.rust-lang.org/cargo-dist/cargo-nightly-x86_64-unknown-linux-gnu.tar.gz';

if [ ! -d "$HOME/vendor/rust" ]; then
    mkdir -p $HOME/vendor/rust

    echo "Installing Rust and Cargo..."
    curl -sL $RUST_URL | tar --strip-components=1 -C $HOME/vendor/rust -xzf -
    curl -sL $CARGO_URL | tar --strip-components=1 -C $HOME/vendor/rust -xzf -
fi

export PATH="$HOME/vendor/rust/bin:$PATH"
export LD_LIBRARY_PATH="$HOME/vendor/rust/lib:$LD_LIBRARY_PATH"

rustc --version
cargo --version
