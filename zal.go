//go:build !wasm
package main

import (
	"fmt"
	"os"

	"github.com/spenserblack/zal-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
