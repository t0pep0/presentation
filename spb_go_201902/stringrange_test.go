package main

import (
	"testing"
)

func BenchmarkFindEnglishWords(b *testing.B) {
	var length int
	for i := 0; i < b.N; i++ {
		length = len(FindEnglishWords())
	}
	if length != 14 {
		b.Fail()
	}
}

func BenchmarkReplaceGo(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = ReplaceGo()
	}
	if len(res) != len(data) {
		b.Fail()
	}
}

func BenchmarkStringReplaceGo(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = StringReplaceGo()
	}
	if len(res) != len(data) {
		b.Fail()
	}
}

func BenchmarkDumbReplaceGo(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = DumbReplaceGo()
	}
	if len(res) != len(data) {
		b.Log(len(res), len(data))
		b.Fail()
	}
}
