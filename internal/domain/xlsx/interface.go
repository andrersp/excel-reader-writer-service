package xlsx

import (
	"bytes"
	"io"
)

type Reader interface {
	Read(io.Reader) ([][]string, error)
}

type Writer interface {
	Write([]Sheet) (bytes.Buffer, error)
}
