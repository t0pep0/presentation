package main

import (
	"io"
	"os"
)

func GetIface(rwc io.ReadWriteCloser) {
	rwc.Read([]byte{})
}

func GetFile(f *os.File) {
	f.Read([]byte{})
}
