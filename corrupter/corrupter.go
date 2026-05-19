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
		w: w,
	}
}

// Write implements io.Writer.
func (c *Corrupter) Write(p []byte) (n int, err error) {
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
		corruptionTimes := c.Min + rand.Intn(c.Max - c.Min)
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

// isCorruptable checks if a rune can be corrupted.
func isCorruptable(b byte) bool {
	return b != '\n' && b != '\r'
}
