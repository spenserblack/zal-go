// Package corrupter provides utilities for corrupting text.
package corrupter

import (
	"bytes"
	"io"
	"math/rand"
)

// Corrupter corrupts text.
type Corrupter struct {
	// Min is the minimum corruption of each character. Undefined behavior if it less
	// than 0.
	Min int
	// Max is the maximum corruption of each character. Undefined behavior if it is less
	// than 0 or less than Min.
	Max int
	// w is the underlying writer.
	w io.Writer
}

// New creates a new corrupter. It sets reasonable defaults for Min and Max, but these
// can be modified.
func New(w io.Writer) *Corrupter {
	return &Corrupter{
		Min: 1,
		Max: 10,
		w:   w,
	}
}

// Write implements io.Writer.
func (c *Corrupter) Write(p []byte) (n int, err error) {
	// TODO UTF-8 encode the bytes and call WriteRune in a loop, instead?
	buf := new(bytes.Buffer)
	var written int
	for _, b := range p {
		// Just write the byte directly if it can't be corrupted.
		if !isCorruptable(b) {
			written, err = c.w.Write([]byte{b})
			n += written
			if err != nil {
				return
			}
			continue
		}
		err = buf.WriteByte(b)
		if err != nil {
			return
		}
		// NOTE We make the maximum *inclusive* with +1
		corruptionTimes := c.Min + rand.Intn(c.Max-c.Min+1)
		for i := 0; i < corruptionTimes; i++ {
			_, err = buf.WriteRune(selectCorruption())
			if err != nil {
				return
			}
		}
		written, err = c.w.Write(buf.Bytes())
		n += written
		if err != nil {
			return
		}
		buf.Reset()
	}
	return
}

// writeRuneBuffer is a bytes buffer for writing runes.
var writeRuneBuffer = new(bytes.Buffer)

// WriteRune writes a single rune, writing it as UTF-8 bytes.
func (c *Corrupter) WriteRune(r rune) (n int, err error) {
	writeRuneBuffer.Reset()
	_, err = writeRuneBuffer.WriteRune(r)
	if err != nil {
		return
	}
	if isCorruptableRune(r) {
		// NOTE We make the maximum *inclusive* with +1
		times := c.Min + rand.Intn(c.Max-c.Min+1)
		for i := 0; i < times; i++ {
			_, err = writeRuneBuffer.WriteRune(selectCorruption())
			if err != nil {
				return
			}
		}
	}
	bytes := writeRuneBuffer.Bytes()
	return c.w.Write(bytes)
}

// isCorruptable checks if a rune can be corrupted.
func isCorruptable(b byte) bool {
	return b != '\n' && b != '\r'
}

// isCorruptableRune checks if a rune can be corrupted.
func isCorruptableRune(r rune) bool {
	return r != '\n' && r != '\r'
}
