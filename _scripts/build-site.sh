#!/bin/sh
SCRIPT="$(readlink -f "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT")"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
SITE_DIR="$ROOT_DIR/_site"
ASSETS_DIR="$SITE_DIR/assets"
cd "$ROOT_DIR"
mkdir -p "$ASSETS_DIR"
cp -f "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$ASSETS_DIR/wasm_exec.js"
cp -f "$(go env GOROOT)/LICENSE" "$ASSETS_DIR/LICENSE-wasm_exec"
GOOS=js GOARCH=wasm go build -o "$ASSETS_DIR/main.wasm" ./wasm.go
