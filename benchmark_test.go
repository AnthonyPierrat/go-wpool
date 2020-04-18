package wpool

import (
	"runtime"
	"testing"
	"time"
)

func BenchMarkWpool(b *testing.B) {
	wp := NewWorkerPool(runtime.NumCPU())
	wp.Start()
	defer wp.Stop()

	for i := 0; i < b.N; i++ {
		wp.Submit(func() {
			time.Sleep(1 * time.Second)
		})
	}
}

func BenchMarkNonWpool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second)
	}
}
