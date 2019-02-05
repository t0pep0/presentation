package main

import (
	"testing"
)

func BenchmarkNonBufChan(b *testing.B) {
	c := make(chan int)
	var d int
	go func() {
		for i := range c {
			d += intOp(i)
		}
	}()
	for i := 0; i < b.N; i++ {
		d = 0
		GenerateChan(c)
	}
	close(c)
}

func BenchmarkBuferizedChan(b *testing.B) {
	c := make(chan int, 10)
	var d int
	go func() {
		for i := range c {
			d += intOp(i)
		}
	}()
	for i := 0; i < b.N; i++ {
		d = 0
		GenerateChan(c)
	}
	close(c)
}

func BenchmarkGenerateCallback(b *testing.B) {
	d := 0
	callback := func(i int) {
		d += intOp(i)
	}
	for i := 0; i < b.N; i++ {
		d = 0
		GenerateCallback(callback)
	}
}

func BenchmarkGenerateAtomicChan(b *testing.B) {
	ac := NewAtomicChan()
	d := 0
	for i := 0; i < b.N; i++ {
		d = atomicChanWork(ac)
	}
	_ = d
}

func atomicChanWork(ac *AtomicChan) int {
	d := 0
	ok := true
	var a int64
	for ok {
		a, ok = ac.Next()
		d += intOp(int(a))
	}
	return d
}
