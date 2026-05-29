//go:build wasm

package main

import (
	"bytes"
	"strconv"
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
	minCorruption := must(intValue(getElementById("minimum-corruption")))
	maxCorruption := must(intValue(getElementById("maximum-corruption")))

	buf := new(bytes.Buffer)
	corrupter := corrupter.New(buf)
	corrupter.Min = minCorruption
	corrupter.Max = maxCorruption

	rawText := inputTextarea.Get("value").String()
	for _, r := range rawText {
		corrupter.WriteRune(r)
	}
	corruptedBytes := buf.Bytes()
	outputEl.Set("textContent", string(corruptedBytes))
}

func intValue(el js.Value) (int, error) {
	value := el.Get("value").String()
	return strconv.Atoi(value)
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
