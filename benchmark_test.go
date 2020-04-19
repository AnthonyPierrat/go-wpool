package wpool

import (
	"runtime"
	"testing"
)

func BenchmarkWpool(b *testing.B) {
	wp := NewWorkerPool(runtime.NumCPU())
	wp.Start()
	defer wp.Stop()

	for i := 0; i < b.N; i++ {
		wp.Submit(func() {
			Fib(10)
		})
	}
}

func BenchmarkNonWpool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
