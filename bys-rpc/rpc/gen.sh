#!/usr/bin/env bash
set -e

# Script that generates .pb.go files from the .proto files.

# Install or update gRPC and the protoc plugin for Golong.
# $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Get the directory of this script file.
DIR=$(dirname "$0")

protoc -I="$DIR"/../../proto --go_out="$DIR"/proto --go_opt=paths=source_relative \
    --go-grpc_out="$DIR"/proto --go-grpc_opt=paths=source_relative \
    "$DIR"/../../proto/host_status_service.proto
