package main

import (
	"math"
	"sync/atomic"
)

const max = 1000

func GenerateChan(c chan int) {
	for i := 0; i < max; i++ {
		c <- i
	}
}

func GenerateCallback(callback func(int)) {
	for i := 0; i < max; i++ {
		callback(i)
	}
}

type AtomicChan struct {
	Value *int64
}

func (ac *AtomicChan) Next() (n int64, ok bool) {
	n = atomic.AddInt64(ac.Value, 1)
	return n, n < max
}

func NewAtomicChan() *AtomicChan {
	ac := new(AtomicChan)
	ac.Value = new(int64)
	atomic.StoreInt64(ac.Value, 0)
	return ac
}

func intOp(i int) int {
	return int(math.Exp(float64(i)))
}
