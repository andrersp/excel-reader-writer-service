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

type Excel interface {
	Read(io.Reader) ([][]string, error)
	Write([]XlsxRequest) (bytes.Buffer, error)
}
