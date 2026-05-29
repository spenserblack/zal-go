//go:build wasm

package main

import (
	"bytes"
	"syscall/js"

	"github.com/spenserblack/zal-go/corrupter"
)

func main() {
	document := js.Global().Get("document")
	getElementById := func(id string) js.Value {
		return document.Call("getElementById", id)
	}
	inputTextarea := getElementById("zalgo-input")
	outputEl := getElementById("zalgo-output")

	buf := new(bytes.Buffer)
	corrupter := corrupter.New(buf)
	corrupter.Min = 5
	corrupter.Max = 10

	rawText := inputTextarea.Get("value").String()
	for _, r := range rawText {
		corrupter.WriteRune(r)
	}
	corruptedBytes := buf.Bytes()
	outputEl.Set("textContent", string(corruptedBytes))
}
