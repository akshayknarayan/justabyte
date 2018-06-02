package justabyte

import (
	"bytes"
	"io"
)

type source struct {
	remaining int
}

func (b *source) Read(p []byte) (n int, err error) {
	if b.remaining == 0 {
		return 0, io.EOF
	}

	written := 0
	if l := len(p); l <= b.remaining {
		// fill in len(p) random bytes
		copy(p, bytes.Repeat([]byte{'b'}, l))
		b.remaining -= l
		written = l
	} else {
		copy(p, bytes.Repeat([]byte{'b'}, b.remaining))
		written = b.remaining
		b.remaining = 0
	}

	return written, nil
}

// New takes as input a length in MB and will yield that many bytes.
// calling Read will always fill the provided buffer, unless the predefined limit is reached
func New(size uint32) io.Reader {
	return &source{
		remaining: int(size) * 1024 * 1024,
	}
}
