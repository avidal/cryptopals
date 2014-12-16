#!/bin/bash
set -ev

export GO_URL="https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz"

if [ ! -d "$HOME/vendor/go" ]; then
    mkdir -p $HOME/vendor/go

    echo "Installing Go..."
    curl -sL $GO_URL | tar --strip-components=1 -C $HOME/vendor/go -xzf -
fi

export PATH="$HOME/vendor/go/bin:$PATH"
export GOROOT="$HOME/vendor/go"
unset GOPATH

go version
