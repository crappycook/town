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

go build -o $PROJECT/target/bys-go-rpc
