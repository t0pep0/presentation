package main

import (
	"os"
	"testing"
)

func BenchmarkGetIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			GetIface(os.Stdin)
		}
	}
}

func BenchmarkGetFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			GetFile(os.Stdin)
		}
	}
}
