package main

import (
	"testing"
)

func BenchmarkRegexpNotCompile(b *testing.B) {
	var length int
	for i := 0; i < b.N; i++ {
		length = len(RegexpNotCompile())
	}
	if length != 14 {
		b.Fail()
	}
}

func BenchmarkRegexpCompile(b *testing.B) {
	var length int
	for i := 0; i < b.N; i++ {
		length = len(RegexpCompile())
	}
	if length != 14 {
		println(length)
		b.Fail()
	}
}
