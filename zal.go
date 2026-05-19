package main

import (
	"io"
	"os"

	"github.com/spenserblack/zal-go/corrupter"
)

func main() {
	const text string = "Hello, world!\n"
	w := corrupter.New(os.Stdout)
	_, err := io.WriteString(w, text)
	if err != nil {
		panic(err)
	}
}
