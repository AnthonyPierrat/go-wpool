package wpool

import (
	"runtime"
	"testing"
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

func TestShouldSubmitJob(t *testing.T) {

	// TODO

}
