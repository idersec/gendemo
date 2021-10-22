#!/bin/bash

export DB_CONNECT_MODE="DSN"

PROJECT_DIR=$(dirname "$0")
GENERATE_DIR="$PROJECT_DIR/cmd/generate"

cd "$GENERATE_DIR" || exit

echo "Start Generating"
go run .
