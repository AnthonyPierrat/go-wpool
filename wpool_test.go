package wpool

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestShouldCreateWorkerPool(t *testing.T) {
	workerPool := NewWorkerPool(3)

	if workerPool == nil {
		t.Fatalf("Could not create worker")
	}

	if workerPool.maxWorkers != 3 {
		t.Fatalf("Error assigning max workers")
	}
}

func TestShouldCreateWorkerPoolWithCPU(t *testing.T) {
	workerPool := NewWorkerPool(-1)

	if workerPool.maxWorkers != runtime.NumCPU() {
		t.Fatalf("Error assigning max workers")
	}
}

func TestShouldSubmitAndStop(t *testing.T) {
	workerPool := NewWorkerPool(-1)
	workerPool.Start()
	defer workerPool.Stop()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workerPool.Submit(func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		})
	}

	wg.Wait()
}
