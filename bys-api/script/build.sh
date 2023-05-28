#!/usr/bin/env bash
set -e

# Get the directory of this script file.
DIR=$(dirname "$0")

cd "$DIR"/..

# Project root
PROJECT=$(pwd)

if [ ! -d "$PROJECT/target" ]; then
    mkdir $PROJECT/target
fi

if [ -d "$PROJECT/target/logs" ]; then
    # Remove logs
    rm -rf $PROJECT/target/logs
fi

go build -o $PROJECT/target/bys-go-api $PROJECT/cmd/main.go
