package linenumberreader

import (
	"io"
)

type LineNumberReader struct {
	LineNumber int

	skipLF bool
	r      io.Reader
}

func NewLineNumberReader(r io.Reader) *LineNumberReader {
	return &LineNumberReader{
		LineNumber: 0,
		r:          r,
	}
}

func (r *LineNumberReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n == 0 {
		return
	}

	var c byte
	for i := 0; i < n; i++ {
		c = p[i]
		if r.skipLF {
			r.skipLF = false
			if c == '\n' {
				continue
			}
		}

		switch c {
		case '\r':
			r.skipLF = true
			fallthrough
		case '\n':
			r.LineNumber++
		}
	}

	if c == '\r' || c == '\n' {
		r.LineNumber--
	}

	return
}
