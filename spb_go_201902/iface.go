package main

import (
	"os"
)

import (
	"io"
)

type Iface struct {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt
}

func GetIface(rwc io.ReadWriteCloser) {
	_ = rwc
}

func GetFile(f *os.File) {
	_ = f
}
