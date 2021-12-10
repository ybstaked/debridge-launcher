package io

import (
	"io"
)

type (
	Reader          = io.Reader
	Writer          = io.Writer
	ReadWriter      = io.ReadWriter
	Closer          = io.Closer
	ReadCloser      = io.ReadCloser
	WriteCloser     = io.WriteCloser
	ReadWriteCloser = io.ReadWriteCloser
)

var (
	CopyBuffer = io.CopyBuffer
	TeeReader  = io.TeeReader
)

//

type readCloser struct {
	Reader
	Closer
}

//

func NewReadCloser(r Reader, c Closer) ReadCloser { return readCloser{r, c} }
